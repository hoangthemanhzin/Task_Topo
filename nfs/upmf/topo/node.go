package topo

import (
	"etrib5gc/logctx"
	"etrib5gc/mesh/httpdp"
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/pfcp/pfcptypes"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"etrib5gc/sbi/upf/upf2upmf"
	"etrib5gc/util/fsm"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync/atomic"
	"time"
)

type PfcpInfo struct {
	Ip   net.IP
	Port int
}

type UpfNode struct {
	logctx.LogWriter
	done   chan struct{} //for closing heartbeat loop
	upfcli sbi.ConsumerClient
	//Other attributes
	//for statemachine
	fsm.State
	// worker    common.Executer
	// hbtimer   common.UeTimer //heartbeat timer
	// assotimer common.UeTimer //association timer
	// rectimer  common.UeTimer //recovery timer
	Heartbeat int // heartbeat interval
	Hbtry     int //num of times trying to send a heartbeat
	//attributes
	Id        string              // unique in a topo
	Infs      map[string][]NetInf //inf identities as keys
	Slices    []models.Snssai
	Pfcpinfo  PfcpInfo
	Static    bool //is the node from config?
	Isactive  uint32
	Issession uint32 //Added a field to check if UPF is handle any user sessions.
}

func (node *UpfNode) HasSbiIp() bool {
	return len(node.Pfcpinfo.Ip) > 0
}

func (topoNode *UpfNode) GetTopoNodeId() string {
	return topoNode.Id
}

func (node *UpfNode) hasPfcpIp() bool {
	return len(node.Pfcpinfo.Ip) > 0
}

func (pfcpinfo *PfcpInfo) NodeId() string {
	return fmt.Sprintf("%s:%d", pfcpinfo.Ip.String(), pfcpinfo.Port)
}

func (node *UpfNode) IsActive() bool {
	return atomic.LoadUint32(&node.Isactive) == 1
}

// does the node serve the snssai?
func (node *UpfNode) Serve(snssai models.Snssai) bool {
	for _, s := range node.Slices {
		if s.Sst == snssai.Sst && strings.Compare(s.Sd, snssai.Sd) == 0 {
			return true
		}
	}
	return false
}

func NewUpfNode(t *Topo, id string, heartbeat int, pfcpinfo *PfcpConfig, static bool) (upf *UpfNode) {

	upfaddr := fmt.Sprintf("%s:%d", pfcpinfo.Ip, pfcpinfo.Port)
	upf = &UpfNode{
		LogWriter: t.WithFields(logctx.Fields{
			"upfid": "dummyid", //add identity of the UPF here for logging, for example its IP
		}),
		done: make(chan struct{}),

		//create upfcli from UPF's Ip address and Sbi Port
		upfcli: httpdp.NewClientWithAddr(upfaddr),
		//initilize other attributes
		//State:     fsm.NewState(UPF_NONASSOCIATED),
		Id:        id,
		Heartbeat: heartbeat,
		Static:    static,
		Infs:      make(map[string][]NetInf),
		Isactive:  1,
		Issession: 1,
	}
	if pfcpinfo != nil {
		upf.Pfcpinfo.Ip = net.ParseIP(pfcpinfo.Ip)
		upf.Pfcpinfo.Port = pfcpinfo.Port
	}
	//finally create a gorouting for sending heartbeats
	go upf.hbLoop(t)
	return
}

// get a NetInf from the node
func (node *UpfNode) getInf(network string, index int) (inf *NetInf, err error) {
	if infs, ok := node.Infs[network]; ok {
		if index >= len(infs) {
			err = fmt.Errorf("Get interface at network '%s' of node '%s' error: index out of range[len=%d;index=%d]", network, node.Id, len(infs), index)
		} else {
			inf = &infs[index]
		}
	} else {
		err = fmt.Errorf("node '%s' is not in network '%s'", node.Id, network)
	}
	return
}

func (upf *UpfNode) hbLoop(topo *Topo) {
	//upf.LogWriter.Info(fmt.Sprintf("Node starts"))
	t := time.NewTicker(time.Duration(10) * time.Second) //fire every 10 second
	for {
		select {
		case <-t.C:
			//TODO: send heartbeat here
			msg := n42.HeartbeatRequest{
				Nonce: rand.Int63(),
				Msg: pfcpmsg.HeartbeatRequest{
					RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
						RecoveryTimeStamp: time.Now(),
					},
				},
			}
			if /*hbrsp*/ _, err := upf2upmf.Heartbeat(upf.upfcli, msg); err != nil {
				//TODO; terminate the loop and remove the UpfNode from topology :
				upf.Infof(fmt.Sprintf("Heartbeat does not respond"))
				upf.RemoteUpfNode(upf.Id, topo)
				upf.Infof("Topo 's data when remove UpfNode : ", topo.Nodes)
				upf.terminate()
			} else {
				//fine, receive a heartbeat response, handle it
				upf.Infof(fmt.Sprintf("Heartbeat responded"))
			}

		case <-upf.done:
			return
		}
	}
}

func (upf *UpfNode) terminate() {
	upf.done <- struct{}{} //send a signal to exit hbLoop
	upf.Info("Node terminated")
}
