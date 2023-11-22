package uepc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

// sbi producer handler for DeleteIndividualUEPolicyAssociation
func OnDeleteIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	polAssoId := ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		//polAssoId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "polAssoId is required",
		})
		return
	}

	if prob := prod.UEPC_HandleDeleteIndividualUEPolicyAssociation(polAssoId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, nil)
	}
	return
}

// sbi producer handler for ReadIndividualUEPolicyAssociation
func OnReadIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	polAssoId := ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		//polAssoId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "polAssoId is required",
		})

		return
	}

	if rsp, prob := prod.UEPC_HandleReadIndividualUEPolicyAssociation(polAssoId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, rsp)
	}
	return
}

// sbi producer handler for ReportObservedEventTriggersForIndividualUEPolicyAssociation
func OnReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	polAssoId := ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		//polAssoId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "polAssoId is required",
		})
		return
	}

	var input models.PolicyAssociationUpdateRequest
	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEPC_HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(polAssoId, input); prob != nil {
			response.SetProblem(prob)
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

// sbi producer handler for CreateIndividualUEPolicyAssociation
func OnCreateIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input models.PolicyAssociationRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEPC_HandleCreateIndividualUEPolicyAssociation(input); prob != nil {
			response.SetProblem(prob)
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

type Producer interface {
	UEPC_HandleDeleteIndividualUEPolicyAssociation(polAssoId string) *models.ProblemDetails
	UEPC_HandleReadIndividualUEPolicyAssociation(polAssoId string) (*models.PolicyAssociation, *models.ProblemDetails)
	UEPC_HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(polAssoId string, body models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)
	UEPC_HandleCreateIndividualUEPolicyAssociation(body models.PolicyAssociationRequest) (*models.PolicyAssociation, *models.ProblemDetails)
}
