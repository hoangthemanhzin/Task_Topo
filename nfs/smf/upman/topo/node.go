package topo

import (
	"etrib5gc/common"
	"etrib5gc/pfcp"
	"etrib5gc/sbi/models"
	"etrib5gc/util/fsm"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
	"time"
)

const (
	MAX_HEARTBEAT_TRY    int = 3
	ASSOCIATION_INTERVAL int = 10000 //ms
	RECOVERY_INTERVAL    int = 10000 //ms
)

type PfcpInfo struct {
	Ip   net.IP
	Port int
}

func (pfcpinfo *PfcpInfo) NodeId() string {
	return fmt.Sprintf("%s:%d", pfcpinfo.Ip.String(), pfcpinfo.Port)
}

type topoNode struct {
	//for statemachine
	fsm.State
	worker    common.Executer
	hbtimer   common.UeTimer //heartbeat timer
	assotimer common.UeTimer //association timer
	rectimer  common.UeTimer //recovery timer

	heartbeat int // heartbeat interval
	hbtry     int //num of times trying to send a heartbeat

	//attributes
	id     string              // unique in a topo
	infs   map[string][]NetInf //inf identities as keys
	slices []models.Snssai

	pfcpinfo PfcpInfo
	static   bool //is the node from config?
	isactive uint32
}

func newNode(id string, heartbeat int, pfcpinfo *PfcpConfig, static bool) (node *topoNode) {
	node = &topoNode{
		State:     fsm.NewState(UPF_NONASSOCIATED),
		id:        id,
		heartbeat: heartbeat,
		static:    static,
		infs:      make(map[string][]NetInf),
	}
	if pfcpinfo != nil {
		node.pfcpinfo.Ip = net.ParseIP(pfcpinfo.Ip)
		node.pfcpinfo.Port = pfcpinfo.Port
		log.Tracef("%s has pfcpinfo %s[%d]", node.id, pfcpinfo.Ip, len(node.pfcpinfo.Ip))
	}
	return
}

func (node *topoNode) UdpAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   node.pfcpinfo.Ip,
		Port: node.pfcpinfo.Port,
	}
}

func (node *topoNode) setActive(state bool) {
	if state {
		atomic.StoreUint32(&node.isactive, 1)
	} else {
		atomic.StoreUint32(&node.isactive, 0)
	}
}

func (node *topoNode) isActive() bool {
	return atomic.LoadUint32(&node.isactive) == 1
}
func (node *topoNode) hasPfcpIp() bool {
	return len(node.pfcpinfo.Ip) > 0
}

func (node *topoNode) start(pfcp *pfcp.Pfcp) {
	log.Tracef("node %s starting", node.id)
	node.worker = common.NewExecuter(1024)
	onheartbeat := func() {
		node.sendEvent(SendHbEvent, pfcp)
	}
	onassociation := func() {
		node.sendEvent(StartEvent, pfcp)
	}

	node.hbtimer = common.NewTimer(time.Duration(node.heartbeat)*time.Millisecond, onheartbeat /*node.worker*/, nil)
	node.rectimer = common.NewTimer(time.Duration(RECOVERY_INTERVAL)*time.Millisecond, onheartbeat /*node.worker*/, nil)
	node.assotimer = common.NewTimer(time.Duration(ASSOCIATION_INTERVAL)*time.Millisecond, onassociation /*node.worker*/, nil)
	if node.hasPfcpIp() {
		node.sendEvent(StartEvent, pfcp)
	}
}

// send an event to the state machine for handling
func (node *topoNode) sendEvent(ev fsm.EventType, args interface{}) (err error) {
	return _sm.SendEvent(node.worker, node, ev, args)
}

func (node *topoNode) stop() {
	log.Tracef("node %s stopping", node.id)
	node.worker.Terminate()
}

// does the node serve the snssai?
func (node *topoNode) serve(snssai models.Snssai) bool {
	for _, s := range node.slices {
		if s.Sst == snssai.Sst && strings.Compare(s.Sd, snssai.Sd) == 0 {
			return true
		}
	}
	return false
}

// get a NetInf from the node
func (node *topoNode) getInf(network string, index int) (inf *NetInf, err error) {
	if infs, ok := node.infs[network]; ok {
		if index >= len(infs) {
			err = fmt.Errorf("Get interface at network '%s' of node '%s' error: index out of range[len=%d;index=%d]", network, node.id, len(infs), index)
		} else {
			inf = &infs[index]
		}
	} else {
		err = fmt.Errorf("node '%s' is not in network '%s'", node.id, network)
	}
	return
}

func (node *topoNode) sendAssociationRequest(pfcp *pfcp.Pfcp) (err error) {
	log.Debugf("%s send association to %s", node.id, node.pfcpinfo.NodeId())
	_, err = pfcp.SendPfcpAssociationSetupRequest(node)
	return
}

func (node *topoNode) sendHeartbeat(pfcp *pfcp.Pfcp) (err error) {
	node.hbtry++
	log.Debugf("%s send heartbeat", node.id)

	if _, err = pfcp.SendPfcpHeartbeatRequest(node); err == nil {
		node.hbtry = 0
	}
	return
}
