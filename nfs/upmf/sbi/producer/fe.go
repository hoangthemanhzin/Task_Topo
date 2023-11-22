package producer

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
)

func (prod *Producer) GetTopo() (rsp *n43.TopoUpf, prob *models.ProblemDetails) {
	prod.Infof("Receive GetTopo")
	//var err error
	rsp = &n43.TopoUpf{
		SlicesTopo: prod.topology.Slices,
		Nets:       prod.topology.Nets,
		Nodes:      make(map[string]*n43.UpfNode),
	}
	var nodes2Fe *n43.UpfNode
	for key, value := range prod.topology.Nodes {
		nodes2Fe = &n43.UpfNode{
			Id:      value.Id,
			Slices:  value.Slices,
			SbiIp:   value.Pfcpinfo.Ip,
			SbiPort: value.Pfcpinfo.Port,
			Infs:    make(map[string][]n43.NetInf),
		}
		//var infs2Fe []*n43.NetInf
		for keyInfs, valInfs := range value.Infs {
			var inf2Fe *n43.NetInf
			for _, valueNetInf := range valInfs {
				inf2Fe = &n43.NetInf{
					Id:      valueNetInf.Id,
					Netname: valueNetInf.Netname,
					Nettype: valueNetInf.Nettype,
					Addr:    valueNetInf.Addr.GetIpAddr(),
				}
				//infs2Fe[indexNetInf] = inf2Fe
				nodes2Fe.Infs[keyInfs] = append(nodes2Fe.Infs[keyInfs], *inf2Fe)
			}
		}
		rsp.Nodes[key] = nodes2Fe
	}
	return
}
