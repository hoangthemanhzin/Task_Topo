package smpc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

//sbi producer handler for DeleteSMPolicy
func OnDeleteSMPolicy(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	smPolicyId := ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		//smPolicyId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "smPolicyId is required",
		})
		return
	}

	var input models.SmPolicyDeleteData
	if err := ctx.DecodeRequest(&input); err == nil {
		if prob := prod.SMPC_HandleDeleteSMPolicy(smPolicyId, input); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, nil)
		}
	} else {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}
	return
}

//sbi producer handler for GetSMPolicy
func OnGetSMPolicy(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	smPolicyId := ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		//smPolicyId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "smPolicyId is required",
		})
		return
	}

	if result, prob := prod.SMPC_HandleGetSMPolicy(smPolicyId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, &result)
	}

	return
}

//sbi producer handler for UpdateSMPolicy
func OnUpdateSMPolicy(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	smPolicyId := ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		//smPolicyId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "smPolicyId is required",
		})
		return
	}

	var input models.SmPolicyUpdateContextData

	if err := ctx.DecodeRequest(&input); err == nil {
		if result, prob := prod.SMPC_HandleUpdateSMPolicy(smPolicyId, input); prob != nil {
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

//sbi producer handler for CreateSMPolicy
func OnCreateSMPolicy(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input models.SmPolicyContextData
	if err := ctx.DecodeRequest(&input); err == nil {
		if result, prob := prod.SMPC_HandleCreateSMPolicy(input); prob != nil {
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

type Producer interface {
	SMPC_HandleDeleteSMPolicy(smPolicyId string, body models.SmPolicyDeleteData) *models.ProblemDetails
	SMPC_HandleGetSMPolicy(smPolicyId string) (models.SmPolicyControl, *models.ProblemDetails)
	SMPC_HandleUpdateSMPolicy(smPolicyId string, body models.SmPolicyUpdateContextData) (models.SmPolicyDecision, *models.ProblemDetails)
	SMPC_HandleCreateSMPolicy(body models.SmPolicyContextData) (models.SmPolicyDecision, *models.ProblemDetails)
}
