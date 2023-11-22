package upmf2activate

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"net/http"
)

// sbi producer handler for activate, deactivate Upfs :
func UpfsActivate(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
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

func UpfsDeActivate(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	var query *n42.UpfDeactivateQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleDeactivate(query); prob == nil {
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
	//HandleGetUpfPath(query *n43.UpfPathQuery) (*n43.UpfPath, *models.ProblemDetails)
	HandleActivate(query *n42.UpfActivateQuery) (*n42.UpfActivate, *models.ProblemDetails)
	HandleDeactivate(query *n42.UpfDeactivateQuery) (*n42.UpfDeactivate, *models.ProblemDetails)
}
