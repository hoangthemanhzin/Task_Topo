package registry

import (
	"etrib5gc/mesh/models"
	"fmt"
	"sync"
)

type ServiceManager struct {
	services      map[models.ServiceName]*Service
	endpoints     map[string]*Endpoint //for searching endpoint by its identity
	addr2endpoint map[string]*Endpoint //for searching endpoint by its address
	rwmutex       sync.RWMutex
}

func newServiceManager() (sman *ServiceManager) {
	sman = &ServiceManager{
		services:      make(map[models.ServiceName]*Service),
		endpoints:     make(map[string]*Endpoint),
		addr2endpoint: make(map[string]*Endpoint),
	}
	return
}

func (sman *ServiceManager) initialize(msg *models.SubscribeResponse) {
	sman.rwmutex.Lock()
	defer sman.rwmutex.Unlock()
	//TODO: check if received services are the requested ones
	//create and add services
	for _, info := range msg.Services {
		if len(info.Id) == 0 {
			log.Warnf("A service with empty name")
			continue
		}
		if len(info.Selectors) == 0 {
			log.Warnf("Service %s has no selector", info.Id)
			continue
		}

		log.Infof("Add service %s", info.Id)
		service := newService(info)
		sman.services[info.Id] = service
	}

	for _, info := range msg.Endpoints {
		ep := newEndpoint(info)
		for _, service := range sman.services {
			//can endpoint be added to the service?
			if service.addEndpoint(ep) {
				log.Infof("Add endpoint %s [%s]", ep.id, service.Id)
				//then add to the general lists too
				sman.endpoints[ep.id] = ep
				sman.addr2endpoint[ep.Addr()] = ep
				break
			}
		}
	}

}

// update endpoint list
func (sman *ServiceManager) update(msg *models.EndpointUpdates) {
	sman.rwmutex.RLock()
	defer sman.rwmutex.RUnlock()

	log.Tracef("update endpoint list")
	for _, item := range msg.Left {
		for _, s := range sman.services {
			if ep := s.removeEndpoint(item); ep != nil {
				log.Infof("Delete endpoint %s [%s]", item, s.Id)
				delete(sman.endpoints, item)
				delete(sman.addr2endpoint, ep.Addr())
				break
			}
		}
	}
	for _, item := range msg.Join {
		ep := newEndpoint(item)
		for _, service := range sman.services {
			//can endpoint be added to the service?
			if service.addEndpoint(ep) {
				log.Infof("Add endpoint %s [%s]", ep.id, service.Id)
				//then add to the general lists too
				sman.endpoints[ep.id] = ep
				sman.addr2endpoint[ep.Addr()] = ep
				break
			}
		}
	}
}

func (sman *ServiceManager) search(id models.ServiceName, match models.RouteMatch) (m MatchedGroup, err error) {
	sman.rwmutex.RLock()
	defer sman.rwmutex.RUnlock()
	log.Tracef("Search an instance of service %s", id)
	if service, ok := sman.services[id]; !ok {
		err = fmt.Errorf("service %s not exist", string(id))
		return
	} else {
		if match != nil {
			if id, ok := match["instanceid"]; ok {
				if ep, ok := sman.endpoints[id]; ok {
					m = &SelectedEndpoint{ep}
				} else {
					err = fmt.Errorf("endpoint %s not found", id)
				}
				return
			}
		}
		m, err = service.match(match)
	}
	return
}

func (sman *ServiceManager) findEndpointByAddr(addr string) (ep *Endpoint) {
	sman.rwmutex.RLock()
	defer sman.rwmutex.RUnlock()
	ep, _ = sman.addr2endpoint[addr]
	return
}

func (sman *ServiceManager) createEndpointWithAddr(addr string) (ep *Endpoint) {
	ep = newEndpointWithAddr(addr)
	sman.rwmutex.Lock()
	defer sman.rwmutex.Unlock()
	sman.addr2endpoint[addr] = ep
	return
}
