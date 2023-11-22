package controller

import (
	"bytes"
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh/models"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Endpoint struct {
	logctx.LogWriter
	k8s       bool
	id        string //unique identity (either uuid or pod id)
	ip        net.IP
	name      string //hostname or pod name
	sbiPort   int
	agentPort int
	labels    map[string]string
	quit      chan bool
	dead      func(*Endpoint)                 //callback when ping fails
	cli       *http.Client                    //for sending request to the endpoint
	services  map[models.ServiceName]*Service //list of subscribed services
}

func newEndpoint(msg *models.RegistrationRequest, dead func(*Endpoint)) (ep *Endpoint, err error) {
	ep = &Endpoint{
		k8s:       msg.K8s,
		id:        strings.Clone(msg.Id),
		name:      strings.Clone(msg.Name),
		labels:    msg.Labels,
		sbiPort:   msg.SbiPort,
		agentPort: msg.AgentPort,
		services:  make(map[models.ServiceName]*Service),
		dead:      dead,
		quit:      make(chan bool),
		cli:       &http.Client{}, //TODO: add parameters
	}
	if !msg.K8s {
		//check if UUID is valid
		if _, err = uuid.Parse(ep.id); err != nil {
			log.Errorf("Registration with invalid UUID: %s", err.Error())
			return
		}

		//parse IP address if the endpoint is not in K8s
		if ep.ip = net.ParseIP(msg.Ip); ep.ip == nil {
			err = fmt.Errorf("Failed to parse IP address for %s", msg.Ip)
		}
	}
	ep.LogWriter = logctx.WithFields(logctx.Fields{
		"endpoint-id": ep.id,
		"endpoint-ip": ep.ip.String(),
	})
	return
}

func (ep *Endpoint) agentUri() string {
	return fmt.Sprintf("%s:%d", ep.ip.String(), ep.agentPort)
}

func (ep *Endpoint) sbiUri() string {
	return fmt.Sprintf("%s:%d", ep.ip.String(), ep.sbiPort)
}

func (ep *Endpoint) loop(wg *sync.WaitGroup) {
	defer wg.Done()
	//defer log.Infof("%s is closed", ep.id)
	defer ep.dead(ep)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			if !ep.ping() {
				return
			}
		case <-ep.quit:
			return
		}
	}
}
func (ep *Endpoint) close() {
	//log.Infof("Close endpoint %s", ep.idr
	close(ep.quit)
}
func (ep *Endpoint) ping() bool {
	url := fmt.Sprintf("http://%s/ping", ep.agentUri())
	if _, err := http.Get(url); err != nil {
		ep.Errorf(err.Error())
		return false
	}
	ep.Infof("ping %s", url)
	return true
}

// pull pod information from K8s control plane
func (ep *Endpoint) k8sUpdate(ctrl *Controller) (err error) {
	return nil
}

// check if an endpoint is serving a service?
func (ep *Endpoint) isServing(s *Service) bool {
	//check for labels matching to selectors
	for k, v := range s.selectors {
		if v1, ok := ep.labels[k]; !ok {
			return false
		} else if strings.Compare(v1, v) != 0 {
			return false
		}
	}

	return true
}

func (ep *Endpoint) notifyEpLeft(lefter *Endpoint) {
	if ep == lefter {
		return
	}
	var err error
	msgbody, _ := json.Marshal(&models.EndpointUpdates{
		Left: []string{lefter.id},
	})
	body := bytes.NewBuffer(msgbody)
	url := fmt.Sprintf("http://%s/endpoint", ep.agentUri())
	req, _ := http.NewRequest(http.MethodPost, url, body)
	var rsp *http.Response
	if rsp, err = ep.cli.Do(req); err == nil {
		var rspbody []byte
		defer rsp.Body.Close()
		if rspbody, err = ioutil.ReadAll(rsp.Body); err == nil {
			var rspmsg models.EndpointUpdatesConfirm
			if err = json.Unmarshal(rspbody, &rspmsg); err == nil {
				ep.Infof("ok=%v", rspmsg.Ok)
			}
		}
	}

}

func (ep *Endpoint) notifyEpJoin(joiner *Endpoint) {
	if ep == joiner {
		return
	}
	var err error
	msgbody, _ := json.Marshal(&models.EndpointUpdates{
		Join: []models.Endpoint{
			models.Endpoint{
				Id:      joiner.id,
				Labels:  joiner.labels,
				Ip:      joiner.ip.String(),
				SbiPort: joiner.sbiPort,
			},
		},
	})
	body := bytes.NewBuffer(msgbody)
	url := fmt.Sprintf("http://%s/endpoint", ep.agentUri())
	req, _ := http.NewRequest(http.MethodPost, url, body)
	var rsp *http.Response
	if rsp, err = ep.cli.Do(req); err == nil {
		var rspbody []byte
		defer rsp.Body.Close()
		if rspbody, err = ioutil.ReadAll(rsp.Body); err == nil {
			var rspmsg models.EndpointUpdatesConfirm
			if err = json.Unmarshal(rspbody, &rspmsg); err == nil {
				ep.Infof("ok=%v", rspmsg.Ok)
			}
		}
	}

}

// remove the endpoint from subsciber lists from services
func (ep *Endpoint) unsubscribeServices() {
	for _, service := range ep.services {
		service.removeSub(ep)
	}
}
