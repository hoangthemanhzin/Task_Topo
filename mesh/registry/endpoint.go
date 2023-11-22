package registry

import (
	"etrib5gc/mesh/models"
	"fmt"
	"net/http"
	"time"
)

type Endpoints []*Endpoint //endpoint list
type Endpoint struct {
	service *Service
	id      string
	addr    string
	labels  models.Labels
	//	models.Endpoint
	cli *http.Client
	//statistics
	numReqs int
	created time.Time
}

func newEndpoint(def models.Endpoint) (ep *Endpoint) {
	ep = &Endpoint{
		id:      def.Id,
		labels:  def.Labels,
		addr:    fmt.Sprintf("%s:%d", def.Ip, def.SbiPort),
		created: time.Now(),
	}
	//TODO: create http client with transport (TLS)
	ep.cli = &http.Client{}
	//TODO:initialize needed statistics
	return
}

func newEndpointWithAddr(addr string) (ep *Endpoint) {
	ep = &Endpoint{
		addr:    addr,
		created: time.Now(),
	}
	//TODO: create http client with transport (TLS)
	ep.cli = &http.Client{}
	//TODO:initialize needed statistics
	return
}
func (ep *Endpoint) Client() *http.Client {
	return ep.cli
}

func (ep *Endpoint) Traffic() uint64 {
	return uint64(ep.numReqs)
}

func (ep *Endpoint) Addr() string {
	return ep.addr
	//return fmt.Sprintf("%s:%d", ep.Endpoint.Sbi.Ip, ep.Endpoint.Sbi.Port)
}

func (eps Endpoints) Traffic() (t uint64) {
	for _, ep := range eps {
		t += ep.Traffic()
	}
	return
}
