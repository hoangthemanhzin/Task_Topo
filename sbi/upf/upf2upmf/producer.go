package upf2upmf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"net/http"
)

func OnHeartbeat(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	var req n42.HeartbeatRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleHeartbeat(&req); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		})

	}
	return
}

func OnActivate(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	var query *n42.UpfActivateQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleActivate(query); prob == nil {
			response.SetBody(200, rsp.Msg)
		} else {
			response.SetProblem(prob)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		})
	}
	return
}

type Producer interface {
	HandleHeartbeat(req *n42.HeartbeatRequest) (*n42.HeartbeatResponse, *models.ProblemDetails)
	HandleActivate(req *n42.UpfActivateQuery) (*n42.UpfActivate, *models.ProblemDetails)
}
