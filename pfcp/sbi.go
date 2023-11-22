package pfcp

import (
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/pfcp/pfcptypes"
	"etrib5gc/sbi/models/n41"
	"etrib5gc/sbi/upf/upf2smf"
	"fmt"
)

type sbiPfcp struct {
	id *pfcptypes.NodeID
}

func NewSbiPfcp(id *pfcptypes.NodeID) PfcpSender {
	return &sbiPfcp{
		id: id,
	}
}

func (proto *sbiPfcp) SendPfcpSessionEstablishmentRequest(session PFCPSession) (rsp *pfcpmsg.PFCPSessionEstablishmentResponse, err error) {
	body := &pfcpmsg.PFCPSessionEstablishmentRequest{
		NodeID: proto.id,
	}

	session.FillEstablishmentRequest(body)

	msg := n41.SessionEstablishmentRequest{
		Msg: body,
	}
	var sbirsp *n41.SessionEstablishmentResponse
	if sbirsp, err = upf2smf.SessionEstablishment(session.UpfCli(), msg); err != nil {
		return
	}

	if sbirsp.Seid == session.LocalSeid() {
		rsp = sbirsp.Msg
	} else {
		err = fmt.Errorf("mismatched SEID")
	}
	return
}

func (proto *sbiPfcp) SendPfcpSessionModificationRequest(session PFCPSession) (rsp *pfcpmsg.PFCPSessionModificationResponse, err error) {
	body := &pfcpmsg.PFCPSessionModificationRequest{}
	session.FillModificationRequest(body)
	msg := n41.SessionModificationRequest{
		Msg: body,
	}

	var sbirsp *n41.SessionModificationResponse
	if sbirsp, err = upf2smf.SessionModification(session.UpfCli(), session.RemoteSeid(), msg); err != nil {
		return
	}
	if sbirsp.Seid == session.LocalSeid() {
		rsp = sbirsp.Msg
	} else {
		err = fmt.Errorf("mismatched SEID")
	}

	return
}

func (proto *sbiPfcp) SendPfcpSessionDeletionRequest(session PFCPSession) (rsp *pfcpmsg.PFCPSessionDeletionResponse, err error) {
	body := &pfcpmsg.PFCPSessionDeletionRequest{}
	session.FillDeletionRequest(body)
	msg := n41.SessionDeletionRequest{
		Msg: body,
	}
	var sbirsp *n41.SessionDeletionResponse
	if sbirsp, err = upf2smf.SessionDeletion(session.UpfCli(), session.RemoteSeid(), msg); err != nil {
		return
	}
	if sbirsp.Seid == session.LocalSeid() {
		rsp = sbirsp.Msg
	} else {
		err = fmt.Errorf("mismatched SEID")
	}

	return
}
