package producer

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"fmt"
	"net/http"
)

func (prod *Producer) HandleGetUpfPath(query *n43.UpfPathQuery) (rsp *n43.UpfPath, prob *models.ProblemDetails) {
	prod.Infof("Receive GetUpfPath from SMF")
	var err error
	// @ManhHT Sua lai o day :
	/*
		var topo = prod.topo //tungtq: Wrong, you need an existing topo
		if rsp, err = prod.ctx.GetPath(topo, query); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: fmt.Sprintf("Path not found: %+v", err),
			}
		}
	*/

	//tungtq: should be like this. Note the `topology` is existing in the
	//Producer and it refers to existing Topology
	if rsp, err = prod.topology.GetPath(query); err != nil {
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: fmt.Sprintf("Path not found: %+v", err),
		}
	}

	return
}


