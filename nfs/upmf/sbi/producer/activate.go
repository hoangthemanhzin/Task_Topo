package producer

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"net/http"
)

func (prod *Producer) HandleActivate(query *n42.UpfActivateQuery)(rsp *n42.UpfActivate, prob *models.ProblemDetails){
	prod.Infof("Activate Upfs ")
	rsp = &n42.UpfActivate{
		Msg: 	make(map[string]string),
	}
	rsp = prod.topology.ActivateUpf(query)
	// for _, upfIdQuery := range query.UpfIds{
	// 	for upf, upfNode := range prod.topology.Nodes{
	// 		if(upfIdQuery == upfNode.Id){
	// 			if(upfNode.Isactive == 1){
	// 				rsp.Msg[upfNode.Id] = "Upf is in activate state"
	// 			}else{
	// 				upf2upmf.UpfActivate(prod.topology.Nodes[upf])
	// 				prod.topology.Nodes[upf].Isactive = 1
	// 				rsp.Msg[upfNode.Id] = "Upf status is activated"
	// 			}
	// 		}
	// 	} 
	// }
	if(rsp == nil){
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
		}
	}
	return
}

func (prod Producer) HandleDeactivate(query *n42.UpfDeactivateQuery) (rsp *n42.UpfDeactivate,prob *models.ProblemDetails){
	prod.Infof("Deactivate Upfs ")
	rsp = &n42.UpfDeactivate{
		Msg: 	make(map[string]string),
	}
	for _, upfIdQuery := range query.UpfIds{
		for upf, upfNode := range prod.topology.Nodes{
			if(upfIdQuery == upfNode.Id){
				if(upfNode.Isactive == 1 && upfNode.Issession == 1){
					prod.topology.Nodes[upf].Issession = 0
					prod.topology.Nodes[upf].Isactive = 0
					rsp.Msg[upfNode.Id] = "Upf status is deactivated"
				}else{
					rsp.Msg[upfNode.Id] = "Upf status is in deactivate"
				}
			}
		}
	} 
	if(rsp == nil){
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
		}
	}
	return
}
