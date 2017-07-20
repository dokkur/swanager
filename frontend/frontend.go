package frontend

import (
	"sync"

	"github.com/dokkur/swanager/command"
	"github.com/dokkur/swanager/config"
	"github.com/dokkur/swanager/core/entities"
	vampRouter "github.com/dokkur/swanager/frontend/vamp_router"
)

// Updatable notifies some component with service
type Updatable interface {
	Update([]entities.Service, []entities.Node)
}

var frontends = make([]Updatable, 0)
var lock sync.Mutex

func init() {
	if config.VampRouterURL != "" {
		frontends = append(frontends, &vampRouter.VampRouter{
			URL: config.VampRouterURL,
		})
	}
}

// Update updates frontend config
//   - Only one update can be run at a time, others will be blocked, consider to run it in coroutine
func Update() {
	if len(frontends) == 0 {
		return
	}

	// Only one update at a time
	lock.Lock()
	defer lock.Unlock()

	list, listRespChan, listErrChan := command.NewServiceListCommand(command.ServiceList{WithStatuses: true})

	nodeList, nodesRespChan, nodesErrChan := command.NewNodeListCommand(command.NodeList{OnlyAvailable: true})

	command.RunAsync(list)
	command.RunAsync(nodeList)

	runningServices := make([]entities.Service, 0)
	var nodes []entities.Node

	select {
	case services := <-listRespChan:
		for _, service := range services {
			if len(service.Status) > 0 {
				runningServices = append(runningServices, service)
			}
		}
	case <-listErrChan:
		return
	}

	select {
	case nodes = <-nodesRespChan:
	case <-nodesErrChan:
		return
	}

	var wg sync.WaitGroup
	for _, frontend := range frontends {
		wg.Add(1)
		go func(front Updatable, runServs []entities.Service, wg *sync.WaitGroup) {
			defer wg.Done()
			front.Update(runServs, nodes)
		}(frontend, runningServices, &wg)
	}
	wg.Wait()
}
