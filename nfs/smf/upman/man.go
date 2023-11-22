package upman

import (
	"etrib5gc/common"
	"etrib5gc/pfcp"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models/n43"
	"etrib5gc/sbi/upmf/upmf2smf"
)

type UpManager struct {
	nodes  map[string]*UpNode //index nodes with UPF identity (IP+Port)
	sender pfcp.PfcpSender
}

func NewUpManager(sender pfcp.PfcpSender) (upman *UpManager, err error) {
	_initLog()
	upman = &UpManager{
		sender: sender,
		nodes:  make(map[string]*UpNode),
	}
	return
}

/*
func (upman *UpManager) findNode(id string) (node *UpNode) {
	node, _ = upman.nodes[id]
	return
}
*/

func (upman *UpManager) CreateTunnel(query *n43.UpfPathQuery, upmfcli sbi.ConsumerClient) (tunnel *UpTunnel, err error) {
	var path *n43.UpfPath
	if path, err = upmf2smf.GetUpfPath(upmfcli, *query); err != nil {
		err = common.WrapError("GetUpfPat failed", err)
	} else {
		log.Infof("Create tunnel")
		tunnel = newTunnel(upman.sender, upmfcli, path)
	}
	for _, node := range tunnel.path {
		log.Infof("set node : %s", node.upf.Id())
		upman.nodes[node.upf.Id()] = node
	}
	return
}

//func (upman *UpManager) CreateTunnel(upmfcli sbi.ConsumerClient) (tunnel *UpTunnel, err error)
