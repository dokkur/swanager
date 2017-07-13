package frontend

import (
	"sync"

	"github.com/dokkur/swanager/command"
	"github.com/dokkur/swanager/core/entities"
	vampRouter "github.com/dokkur/swanager/frontend/vamp_router"
)

// Updatable notifies some component with service
type Updatable interface {
	Update([]entities.Service)
}

var frontends = make([]Updatable, 0)

func init() {
	frontends = append(frontends, vampRouter.VampRouter{URL: "http://localhost:345345"})
}

// Update updates frontend config
func Update() {
	cmd, respChan, errChan := command.NewServiceListCommand(command.ServiceList{WithStatuses: true})

	command.RunAsync(cmd)

	runningServices := make([]entities.Service, 0)

	select {
	case services := <-respChan:
		for _, service := range services {
			if len(service.Status) > 0 {
				runningServices = append(runningServices, service)
			}
		}
	case <-errChan:
		return
	}

	var wg sync.WaitGroup
	for _, frontend := range frontends {
		wg.Add(1)
		go func(front Updatable, runServs []entities.Service, wg *sync.WaitGroup) {
			defer wg.Done()
			front.Update(runServs)
		}(frontend, runningServices, &wg)
	}
	wg.Wait()
}
