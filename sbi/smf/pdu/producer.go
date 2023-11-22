package pdu

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
)

// sbi producer handler for ReleaseSmContext
func OnReleaseSmContext(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	smContextRef := ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		//smContextRef is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "smContextRef is required",
		})
		return
	}

	var input models.ReleaseSmContextRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.PDU_HandleReleaseSmContext(smContextRef, &input); prob != nil {
			response.SetBody(400, prob)
		} else {
			response.SetBody(200, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

// sbi producer handler for UpdateSmContext
func OnUpdateSmContext(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	smContextRef := ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		//smContextRef is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "smContextRef is required",
		})
		return
	}

	var input models.UpdateSmContextRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, ersp := prod.PDU_HandleUpdateSmContext(smContextRef, input); ersp != nil {
			response.SetBody(400, ersp)
		} else {
			response.SetBody(200, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}
	return
}

// sbi producer handler for PostSmContexts
func OnPostSmContexts(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	callbackstr := ctx.Header("Callback")
	if len(callbackstr) == 0 {
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "Callback is not set in the request header",
		})
		return
	}
	callback := models.Callback(callbackstr)
	var input models.PostSmContextsRequest
	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, ersp := prod.PDU_HandlePostSmContexts(input, callback); ersp != nil {
			response.SetBody(400, ersp)
		} else {
			response.SetBody(201, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: fmt.Sprintf("Failed to decode input: %s", err.Error()),
		})
	}

	return
}

type Producer interface {
	PDU_HandlePostSmContexts(body models.PostSmContextsRequest, callback models.Callback) (*models.PostSmContextsResponse, *models.PostSmContextsErrorResponse)
	PDU_HandleUpdateSmContext(smContextRef string, body models.UpdateSmContextRequest) (*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse)
	PDU_HandleReleaseSmContext(smContextRef string, body *models.ReleaseSmContextRequest) (*models.SmContextReleasedData, *models.ProblemDetails)
}
