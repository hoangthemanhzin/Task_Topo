package upmf2upf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"net/http"
)

func OnRegister(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	var req n42.RegistrationRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleRegistration(&req); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(201, rsp)
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
	HandleRegistration(req *n42.RegistrationRequest) (*n42.RegistrationResponse, *models.ProblemDetails)
}
