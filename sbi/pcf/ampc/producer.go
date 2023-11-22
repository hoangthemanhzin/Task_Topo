package ampc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

//sbi producer handler for CreateIndividualAMPolicyAssociation
func OnCreateIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input models.PolicyAssociationRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if result, prob := prod.AMPC_HandleCreateIndividualAMPolicyAssociation(input); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, &result)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}
	return
}

//sbi producer handler for DeleteIndividualAMPolicyAssociation
func OnDeleteIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
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

	if prob := prod.AMPC_HandleDeleteIndividualAMPolicyAssociation(polAssoId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, nil)
	}
	return
}

//sbi producer handler for ReadIndividualAMPolicyAssociation
func OnReadIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
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

	if result, prob := prod.AMPC_HandleReadIndividualAMPolicyAssociation(polAssoId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, &result)
	}

	return
}

//sbi producer handler for ReportObservedEventTriggersForIndividualAMPolicyAssociation
func OnReportObservedEventTriggersForIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
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

	if prob := ctx.DecodeRequest(&input); prob == nil {
		if result, prob := prod.AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId, input); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, &result)
		}
	}
	return
}

type Producer interface {
	AMPC_HandleCreateIndividualAMPolicyAssociation(body models.PolicyAssociationRequest) (models.PolicyAssociation, *models.ProblemDetails)
	AMPC_HandleDeleteIndividualAMPolicyAssociation(polAssoId string) *models.ProblemDetails
	AMPC_HandleReadIndividualAMPolicyAssociation(polAssoId string) (models.PolicyAssociation, *models.ProblemDetails)
	AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId string, body models.PolicyAssociationUpdateRequest) (models.PolicyUpdate, *models.ProblemDetails)
}
