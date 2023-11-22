package controller

import (
	"etrib5gc/mesh/models"
	"sync"
)

type EndpointManager struct {
	endpoints map[string]*Endpoint
	rwmutex   sync.RWMutex
	wg        sync.WaitGroup //wait for all endpoint's loop
	ctrl      *Controller
}

func newEndpointManager(c *Controller) (epman *EndpointManager) {
	epman = &EndpointManager{
		ctrl:      c,
		endpoints: make(map[string]*Endpoint),
	}
	return
}
func (epman *EndpointManager) findById(id string) (ep *Endpoint) {
	epman.rwmutex.RLock()
	defer epman.rwmutex.RUnlock()
	ep, _ = epman.endpoints[id]
	return
}

func (epman *EndpointManager) registerEndpoint(msg *models.RegistrationRequest) (ep *Endpoint, err error) {
	log.Infof("Registration from %s-%s-%s:%d", msg.Id, msg.Name, msg.Ip, msg.SbiPort)
	if ep = epman.findById(msg.Id); ep != nil {
		return
	}

	if ep, err = newEndpoint(msg, epman.onEndpointDead); err != nil {
		return
	}
	if msg.K8s {
		//pull Pod information from K8s control plane
		if err = ep.k8sUpdate(epman.ctrl); err != nil {
			return
		}
	}
	epman.add(ep)
	return
}
func (epman *EndpointManager) onEndpointDead(ep *Endpoint) {
	epman.remove(ep)
}

func (epman *EndpointManager) add(ep *Endpoint) {
	epman.rwmutex.Lock()
	defer epman.rwmutex.Unlock()

	//add the endpoint and index it by sbi address
	epman.endpoints[ep.id] = ep
	//epman.addr2endpoint[ep.sbiUri()] = ep

	//start a goroutine to send pings
	epman.wg.Add(1)
	go ep.loop(&epman.wg)

	//send notification to subscribers
	epman.ctrl.sendEpEvent(EP_JOIN, ep)
	ep.Infof("endpoint is registered")
}
func (epman *EndpointManager) remove(ep *Endpoint) {
	epman.rwmutex.Lock()
	defer epman.rwmutex.Unlock()
	defer ep.Infof("endpoint is removed")
	log.Infof("remove %s", ep.id)
	delete(epman.endpoints, ep.id)
	//delete(epman.addr2endpoint, ep.sbiUri())
	//send notification to subscribers
	epman.ctrl.sendEpEvent(EP_LEFT, ep)
}
func (epman *EndpointManager) close() {
	//get endpoint list
	epman.rwmutex.Lock()
	endpoints := []*Endpoint{}
	for _, ep := range epman.endpoints {
		endpoints = append(endpoints, ep)
	}
	epman.rwmutex.Unlock()

	//close (remove) each of them
	for _, ep := range endpoints {
		ep.close()
	}
	epman.wg.Wait()
}

// get endpoints of subscribed services for an endpoint
func (epman *EndpointManager) getEndpoints(sub *Endpoint) (endpoints []*Endpoint) {
	epman.rwmutex.RLock()
	defer epman.rwmutex.RUnlock()
	for _, ep := range epman.endpoints {
		for _, s := range sub.services {
			if ep.isServing(s) {
				endpoints = append(endpoints, ep)
				break
			}
		}
	}
	return
}
