package controller

import (
	"etrib5gc/mesh/models"
	"sync"
)

type Service struct {
	id          models.ServiceName
	selectors   models.Selectors
	groups      map[string]models.EndpointGroup
	routes      []models.RouteRule
	subscribers map[string]*Endpoint
	mutex       sync.Mutex
}

func newService(info *models.Service) (s *Service) {
	s = &Service{
		id:          info.Id,
		selectors:   info.Selectors,
		groups:      info.Groups,
		routes:      info.Routes,
		subscribers: make(map[string]*Endpoint),
	}
	return
}

func (s *Service) addSub(ep *Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Infof("add %s as a subscriber to %s", ep.id, s.id)
	s.subscribers[ep.id] = ep
}

func (s *Service) removeSub(ep *Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.subscribers, ep.id)
	log.Infof("remove subscriber %s from %s", ep.id, s.id)
}

func (s *Service) getSubscribers() (endpoints []*Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, subs := range s.subscribers {
		endpoints = append(endpoints, subs)
	}
	return
}
