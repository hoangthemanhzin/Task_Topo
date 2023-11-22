package topo

import (
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/pfcp/pfcptypes"
	"fmt"
)

func (topo *UpfTopo) findNodeByPfcpId(id string) (node *topoNode) {
	node, _ = topo.pfcpid2node[id]
	return
}

func (topo *UpfTopo) HandleAssociationSetupRequest(upfid string, req *pfcpmsg.PFCPAssociationSetupRequest) (rsp *pfcpmsg.PFCPAssociationSetupResponse, err error) {
	//TODO: set cause value
	cause := pfcptypes.Cause{
		CauseValue: pfcptypes.CauseRequestAccepted,
	}
	rsp = &pfcpmsg.PFCPAssociationSetupResponse{
		//	NodeID: proto.ctx.NodeId(),
		Cause: &cause,
		//RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
		//	RecoveryTimeStamp: proto.fwd.When(),
		//},
		CPFunctionFeatures: &pfcptypes.CPFunctionFeatures{
			SupportedFeatures: 0,
		},
	}

	var node *topoNode
	if node = topo.findNodeByPfcpId(upfid); node != nil {
		node.sendEvent(AssociatedEvent, nil)
	} else {
		//NOTE: This is for future architectural changes if the Upf is in the
		//configured topopogy and the request has its identity we find the Upf
		//in the current topology.  Otherwise, the upf is none-configurede,
		//create a new node then add to TopologyManager (if the request has
		//enough information)
		cause.CauseValue = pfcptypes.CauseRequestRejected
	}
	return
}
func (topo *UpfTopo) HandleAssociationReleaseRequest(upfid string, req *pfcpmsg.PFCPAssociationReleaseRequest) (rsp *pfcpmsg.PFCPAssociationReleaseResponse, err error) {
	cause := pfcptypes.Cause{
		CauseValue: pfcptypes.CauseRequestAccepted,
	}
	rsp = &pfcpmsg.PFCPAssociationReleaseResponse{
		//NodeID: proto.ctx.NodeId(),
		Cause: &cause,
	}

	var node *topoNode
	if node = topo.findNodeByPfcpId(upfid); node == nil {
		cause.CauseValue = pfcptypes.CauseRequestRejected
	} else {
		node.sendEvent(DisassociatedEvent, nil)
	}
	return
}
func (topo *UpfTopo) HandleHeartbeatRequest(upfid string, req *pfcpmsg.HeartbeatRequest) (rsp *pfcpmsg.HeartbeatResponse, err error) {
	var node *topoNode
	if node = topo.findNodeByPfcpId(upfid); node == nil {
		err = fmt.Errorf("Upf not found")
		return
	}
	rsp = &pfcpmsg.HeartbeatResponse{
		//	RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
		//		RecoveryTimeStamp: proto.fwd.When(),
		//	},
	}
	node.sendEvent(ConnectedEvent, nil)
	return
}
