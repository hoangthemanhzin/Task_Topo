package controller

import (
	"etrib5gc/mesh/models"
	"sync"
)

type ServiceManager struct {
	epman    *EndpointManager
	services map[models.ServiceName]*Service
	rwmutex  sync.RWMutex
}

func newServiceManager(epman *EndpointManager, infoitems []models.Service) (sman *ServiceManager) {
	sman = &ServiceManager{
		epman:    epman,
		services: make(map[models.ServiceName]*Service),
	}
	for _, item := range infoitems {
		log.Infof("add service %v", item)
		sman.services[item.Id] = newService(&item)
	}
	return
}

// an endpoint subscribes a list of services
func (sman *ServiceManager) addSubscriber(ep *Endpoint, services []models.ServiceName) {
	sman.rwmutex.Lock()
	defer sman.rwmutex.Unlock()
	//make unique service list
	uservices := make(map[models.ServiceName]bool)
	for _, s := range services {
		uservices[s] = true
	}
	//only subscribe to existing services
	for sid, _ := range uservices {
		if service, ok := sman.services[sid]; ok {
			//subscribe the endpoint
			service.addSub(ep)
			//add the subscibed service
			ep.services[sid] = service
		} else {
			log.Warnf("Service %s not exists", sid)
		}
	}
}

func (sman *ServiceManager) getSubscribers(ep *Endpoint) (endpoints []*Endpoint) {
	var service *Service
	if service = sman.getService(ep); service == nil {
		return
	}
	endpoints = service.getSubscribers()
	return
}

func (sman *ServiceManager) getService(ep *Endpoint) *Service {
	sman.rwmutex.RLock()
	defer sman.rwmutex.RUnlock()
	for _, service := range sman.services {
		if ep.isServing(service) {
			log.Infof("endpoint %s with labels=%v is serving service %s", ep.id, ep.labels, service.id)
			return service
		}
	}
	return nil
}
