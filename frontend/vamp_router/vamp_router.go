package vampRouter

import (
	"sync"

	"github.com/magneticio/vamp-router/haproxy"

	"github.com/dokkur/swanager/core/entities"
)

// VampRouter represenst vamp-router integration
type VampRouter struct {
	URL string

	frontendsMutex sync.Mutex
	frontends      []haproxy.Frontend

	backendsMutex sync.Mutex
	backends      []haproxy.Backend
}

// Update updates vamp-router configuration
func (vr *VampRouter) Update(services []entities.Service) {
	vr.cleanup()

	if len(services) == 0 {
		return
	}

	var wg sync.WaitGroup
	for _, service := range services {
		if len(service.FrontendEndpoints) == 0 {
			continue
		}

		wg.Add(1)
		go func(serv entities.Service, wg *sync.WaitGroup) {
			defer wg.Done()
			vr.parseService(serv)
		}(service, &wg)
	}
	wg.Wait()
}

func (vr *VampRouter) parseService(service entities.Service) {
	for _, endpoint := range service.FrontendEndpoints {
		vr.parseEndpoint(service, endpoint)
	}
}

func (vr *VampRouter) parseEndpoint(service entities.Service,
	endpoint entities.FrontendEndpoint) {

}

func (vr *VampRouter) addFrontend(front haproxy.Frontend) {
	vr.frontendsMutex.Lock()
	defer vr.frontendsMutex.Unlock()

	vr.frontends = append(vr.frontends, front)
}

func (vr *VampRouter) addBackend(back haproxy.Backend) {
	vr.backendsMutex.Lock()
	defer vr.backendsMutex.Unlock()

	vr.backends = append(vr.backends, back)
}

func (vr *VampRouter) cleanup() {
	vr.frontendsMutex.Lock()
	defer vr.frontendsMutex.Unlock()
	vr.backendsMutex.Lock()
	defer vr.backendsMutex.Unlock()

	vr.frontends = make([]haproxy.Frontend, 0)
	vr.backends = make([]haproxy.Backend, 0)
}
