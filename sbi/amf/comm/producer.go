package comm

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

// sbi producer handler for CreateUEContext
func OnCreateUEContext(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		})
		return
	}

	var input models.CreateUEContextRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, ersp := prod.COMM_HandleCreateUEContext(ueContextId, input); ersp != nil {
			response.SetBody(403, ersp)
		} else {
			response.SetBody(201, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Title:  "Internal Error",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}
	return
}

// sbi producer handler for UEContextTransfer
func OnUEContextTransfer(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		})

		return
	}

	var input models.UEContextTransferRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.COMM_HandleUEContextTransfer(ueContextId, input); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Title:  "Internal Error",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})

	}
	return
}

// sbi producer handler for N1N2MessageTransfer
func OnN1N2MessageTransfer(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		})

		return
	}

	var input models.N1N2MessageTransferRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, ersp := prod.COMM_HandleN1N2MessageTransfer(ueContextId, input); ersp != nil {
			response.SetBody(409, ersp)
		} else {
			response.SetBody(200, &rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Title:  "Internal Error",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})

	}

	return
}

type Producer interface {
	COMM_HandleCreateUEContext(string, models.CreateUEContextRequest) (*models.CreateUEContextResponse, *models.CreateUEContextErrorResponse)
	COMM_HandleUEContextTransfer(string, models.UEContextTransferRequest) (*models.UEContextTransferResponse, *models.ProblemDetails)
	COMM_HandleN1N2MessageTransfer(string, models.N1N2MessageTransferRequest) (*models.N1N2MessageTransferRspData, *models.N1N2MessageTransferError)
}
