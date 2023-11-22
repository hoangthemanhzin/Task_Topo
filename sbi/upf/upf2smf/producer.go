package upf2smf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n41"
	"fmt"
	"net/http"
	"strconv"
)

func OnSessionEstablishment(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	var req n41.SessionEstablishmentRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleSessionEstablishment(&req); prob != nil {
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

func OnSessionModification(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	var err error
	var seid uint64
	seidstr := ctx.Param("seid")
	if seid, err = strconv.ParseUint(seidstr, 10, 64); err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: fmt.Sprintf("Convert SEID failed: %+v", err),
		})
		return
	}

	prod := handler.(Producer)
	var req n41.SessionModificationRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleSessionModification(seid, &req); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(201, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: fmt.Sprintf("Decode Request failed: %+v", err),
		})

	}
	return
}

func OnSessionDeletion(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	var err error
	var seid uint64
	seidstr := ctx.Param("seid")
	if seid, err = strconv.ParseUint(seidstr, 10, 64); err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: fmt.Sprintf("Convert SEID failed: %+v", err),
		})
		return
	}

	prod := handler.(Producer)
	var req n41.SessionDeletionRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleSessionDeletion(seid, &req); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(201, rsp)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: fmt.Sprintf("Decode Request failed: %+v", err),
		})

	}
	return
}

type Producer interface {
	HandleSessionEstablishment(req *n41.SessionEstablishmentRequest) (*n41.SessionEstablishmentResponse, *models.ProblemDetails)
	HandleSessionModification(seid uint64, req *n41.SessionModificationRequest) (*n41.SessionModificationResponse, *models.ProblemDetails)
	HandleSessionDeletion(seid uint64, req *n41.SessionDeletionRequest) (*n41.SessionDeletionResponse, *models.ProblemDetails)
}
