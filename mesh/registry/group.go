package registry

import (
	"fmt"
	"math/rand"
)

type Destination struct {
	group *EndpointGroup
	lb    uint8 //load balancing algorithm
}

func (d *Destination) Select() (ep *Endpoint, err error) {
	return d.group.selectEp(d.lb)
}

type EndpointGroup struct {
	endpoints map[string]*Endpoint
}

func (g *EndpointGroup) add(ep *Endpoint) {
	g.endpoints[ep.id] = ep
}

func (g *EndpointGroup) remove(ep *Endpoint) {
	delete(g.endpoints, ep.id)
}

//select one endpoint from the destination
func (g *EndpointGroup) selectEp(lb uint8) (ep *Endpoint, err error) {
	num := len(g.endpoints)
	if num == 0 {
		err = fmt.Errorf("No endpoint to select")
		return
	}
	endpoints := []*Endpoint{}
	for _, ep := range g.endpoints {
		endpoints = append(endpoints, ep)
	}
	//only support a simple random load balancer for now
	switch lb {
	case LB_RANDOM:
		fallthrough
	case LB_ROUND_ROBIN:
		fallthrough
	case LB_LEAST_REQUEST:
		log.Tracef("Pick a random instance from %d instance(s)", num)
		//draw a randome index
		id := rand.Intn(num)
		//return the selected endpoint
		ep = endpoints[id]
	default:
		err = fmt.Errorf("Unknown load balancer %d", lb)
	}
	return
}

type DestinationSet struct {
	destinations []Destination
	wmarkers     []float64 //weight markers //increasing values from 0 to 1
}

//apply the traffic spliting logic to choose one destination from the set then
//select a endpoint
func (s *DestinationSet) Select() (ep *Endpoint, err error) {
	//get the index of an destination
	index := s.sample()
	//select an endpoint from the group
	ep, err = s.destinations[index].Select()
	return
}

//select a destination index from the set with randomizationg using traffic ratios
func (s *DestinationSet) sample() int {
	tmp := rand.Float64()
	for i := 0; i < len(s.wmarkers); i++ {
		if s.wmarkers[i] > tmp {
			return i
		}
	}
	return 0
}

type SelectedEndpoint struct {
	*Endpoint
}

func (s *SelectedEndpoint) Select() (*Endpoint, error) {
	return s.Endpoint, nil
}
