package upmf2smf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"net/http"
)

// sbi producer handler for GetNfGroupIDs
func OnGetUpfPath(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	var query *n43.UpfPathQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleGetUpfPath(query); prob == nil {
			response.SetBody(200, rsp)
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
	HandleGetUpfPath(query *n43.UpfPathQuery) (*n43.UpfPath, *models.ProblemDetails)
}
