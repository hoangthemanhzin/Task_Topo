package registry

import (
	"etrib5gc/mesh/models"
	"fmt"
)

type Service struct {
	models.Service                           //service definition
	endpoints      map[string]*Endpoint      //current endpoint list
	groups         map[string]*EndpointGroup //group of endpoints for routing
	defaultGroup   Destination               //default destination including all endpoints
}

func newService(def models.Service) (s *Service) {
	s = &Service{
		Service:   def,
		endpoints: make(map[string]*Endpoint),
		groups:    make(map[string]*EndpointGroup),
	}

	s.defaultGroup = Destination{
		group: &EndpointGroup{
			endpoints: s.endpoints,
		},
		lb: LB_RANDOM,
	}
	//create routing groups
	for id, _ := range def.Groups {
		s.groups[id] = &EndpointGroup{
			endpoints: make(map[string]*Endpoint),
		}
	}
	return
}

// check to include an endpoint
func (s *Service) addEndpoint(ep *Endpoint) (b bool) {
	if b = s.Selectors.Match(ep.labels); b {
		//attach service to endpoint
		ep.service = s
		//add endpoint to service's endpoint list
		s.endpoints[ep.id] = ep
		log.Tracef("Endpoint %s with labels %v belong to service %s", ep.id, ep.labels, s.Id)
		//add endpoint to subgroups
		for gid, gdef := range s.Groups {
			if gdef.Selectors.Match(ep.labels) {
				s.groups[gid].add(ep)
			}
		}
	}
	return
}

func (s *Service) removeEndpoint(id string) *Endpoint {
	if ep, ok := s.endpoints[id]; ok {
		ep.service = nil
		delete(s.endpoints, id)
		//remove from subgroups
		for _, group := range s.groups {
			group.remove(ep)
		}
		return ep
	}
	return nil
}

// find a routing destination
func (s *Service) match(match models.RouteMatch) (m MatchedGroup, err error) {
	matched := false
	for _, route := range s.Routes {
		if route.IsMatched(match) {
			matched = true
			//compose destination set
			log.Tracef("Match to destination %v", route.Destinations)
			m, err = s.createDestinations(route.Destinations)
			break
		}
	}
	if !matched {
		//nothing match, return the default destination
		log.Tracef("Match to default group")
		m = &s.defaultGroup
	}
	return
}

func (s *Service) createDestinations(info []models.Destination) (m MatchedGroup, err error) {
	num := len(info)
	if num == 0 {
		err = fmt.Errorf("Empty destination set")
		return
	}
	if num == 1 {
		if g, ok := s.groups[info[0].GroupId]; !ok {
			err = fmt.Errorf("Destination %s not found", info[0].GroupId)
		} else {
			//return the single group
			m = &Destination{
				group: g,
				lb:    info[0].Lb,
			}
		}
		return
	}
	//create a DestinationSet
	destinationset := DestinationSet{
		destinations: make([]Destination, num),
		wmarkers:     make([]float64, num),
	}
	//set the groups
	for i := 0; i < num; i++ {
		if g, ok := s.groups[info[i].GroupId]; !ok {
			err = fmt.Errorf("Destination %s not found", info[i].GroupId)
			return
		} else {
			destinationset.destinations[i] = Destination{
				group: g,
				lb:    info[i].Lb,
			}
		}
	}

	//set weight markers for traffic ratios
	destinationset.wmarkers[0] = float64(info[0].Weight)
	for i := 1; i < num; i++ {
		destinationset.wmarkers[i] = destinationset.wmarkers[i-1] + float64(info[i].Weight)
	}
	//ranging  (0,1]
	for i := 0; i < num; i++ {
		destinationset.wmarkers[i] /= destinationset.wmarkers[num-1]
	}
	m = &destinationset
	return
}
