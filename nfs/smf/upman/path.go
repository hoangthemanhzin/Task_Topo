package upman

import (
	"etrib5gc/nfs/smf/upman/up"
	"etrib5gc/pfcp"
	"etrib5gc/sbi/models/n43"
)

type UpPath []*UpNode

func newUpPath(topopath *n43.UpfPath, sender pfcp.PfcpSender) (uppath UpPath) {
	plen := len(topopath.Path)
	uppath = make([]*UpNode, plen)
	for i, pathnode := range topopath.Path {
		upnode := &UpNode{
			upf: up.NewUpf(sender, pathnode.PfcpIp, pathnode.PfcpPort),
		}
		upnode.ul.local = upnode
		upnode.ul.localip = pathnode.UlIp
		upnode.dl.local = upnode
		upnode.dl.localip = pathnode.DlIp

		if i > 0 {
			upnode.dl.remote = uppath[i-1]
		}
		if i < plen-1 {
			upnode.ul.remote = uppath[i+1]
		}

		uppath[i] = upnode
	}
	return
}
