package upman

import (
	"etrib5gc/nfs/smf/upman/up"
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/pfcpmsg"
	"net"
)

type LinkInfo struct {
	local, remote *UpNode
	localip       net.IP
	teid          uint32
	pdr           *up.PDR
}

// one upf node in a data path of a tunnel
type UpNode struct {
	upf     *up.Upf         //hold node interfaces information
	ul      LinkInfo        //link toward uplink UPF in the current path
	dl      LinkInfo        //link toward downlink UPF (or RAN) in the current path
	session *up.PfcpSession //Pfcp session belong to an SmContext at this UPF
}

func (node *UpNode) UlPdr() *up.PDR {
	return node.ul.pdr
}
func (node *UpNode) DlPdr() *up.PDR {
	return node.dl.pdr
}

// create Pfcp sessions and PDRs for uplink and downlink tunnels attached to
// them
func (node *UpNode) createSessionAndPdrs(seid uint64) {
	node.session = node.upf.CreateSession(seid)
	//	if node.ul.remote != nil { //if not the anchor upf
	node.ul.pdr = node.session.CreatePdr()
	node.ul.teid = node.upf.GenerateTeid()
	//	}
	node.dl.pdr = node.session.CreatePdr()
	node.dl.teid = node.upf.GenerateTeid()
}

func (node *UpNode) deactivate() (err error) {
	node.deactivateLink(&node.ul)
	node.deactivateLink(&node.dl)
	node.session.RemovePdrs()
	return
}
func (node *UpNode) deactivateLink(link *LinkInfo) {
	node.upf.FreeTeid(link.teid)
	link.pdr = nil
	link.remote = nil
}

func (node *UpNode) releaseFn(sender pfcp.PfcpSender) func() error {
	return func() (err error) {
		_, err = sender.SendPfcpSessionDeletionRequest(node.session)
		//TODO: handle response
		return nil
	}
}

func (node *UpNode) updateFn(sender pfcp.PfcpSender) func() error {
	return func() (err error) {
		log.Trace("send pfcp session update")
		if node.session.RemoteSeid() == 0 {
			log.Infof("Send pfcp SessionEstablishment to UPF[%s]", node.upf.Id())
			var rsp *pfcpmsg.PFCPSessionEstablishmentResponse
			rsp, err = sender.SendPfcpSessionEstablishmentRequest(node.session)
			if err != nil {
				log.Errorf("Send PfcpSessionEstablishment failed: %+v", err)
			} else {
				node.session.SetRemoteSeid(rsp.UPFSEID.Seid)
			}
		} else {
			log.Infof("Send pfcp SessionModification to UPF[%s]", node.upf.Id())
			if _, err = sender.SendPfcpSessionModificationRequest(node.session); err != nil {
				log.Errorf("Send PfcpSessionModification failed: %+v", err)
			}
		}
		//TODO: handle response
		return nil
	}
}
