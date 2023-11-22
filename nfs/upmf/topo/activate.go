package topo

import (
	"etrib5gc/sbi/models/n42"
	"etrib5gc/sbi/upf/upf2upmf"
)

func (t *Topo) ActivateUpf(req *n42.UpfActivateQuery) (rsp *n42.UpfActivate) {
	for _, upfIdQuery := range req.UpfIds {
		for upf, upfNode := range t.Nodes {
			if upfIdQuery == upfNode.Id {
				if upfNode.Isactive == 1 {
					rsp.Msg[upfNode.Id] = "Upf is in activate state"
				} else {
					upf2upmf.UpfActivate(t.Nodes[upf].upfcli, *req)
					t.Nodes[upf].Isactive = 1
					//rsp.Msg[upfNode.Id] = "Upf status is activated"
				}
			}
		}
	}
	return
}
