package uea

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

// sbi producer handler for Delete5gAkaAuthenticationResult
func OnDelete5gAkaAuthenticationResult(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		})
		return
	}

	if prob := prod.UEA_HandleDelete5gAkaAuthenticationResult(authCtxId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, nil)
	}

	return
}

// sbi producer handler for DeleteEapAuthenticationResult
func OnDeleteEapAuthenticationResult(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		})

		return
	}

	if prob := prod.UEA_HandleDeleteEapAuthenticationResult(authCtxId); prob != nil {
		response.SetProblem(prob)
	} else {
		response.SetBody(200, nil)
	}
	return
}

// sbi producer handler for EapAuthMethod
func OnEapAuthMethod(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		})

		return
	}

	var input models.EapSession
	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEA_HandleEapAuthMethod(authCtxId, &input); prob != nil {
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

// sbi producer handler for RgAuthenticationsPost
func OnRgAuthenticationsPost(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input models.RgAuthenticationInfo

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEA_HandleRgAuthenticationsPost(input); prob != nil {
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

// sbi producer handler for UeAuthenticationsAuthCtxId5gAkaConfirmationPut
func OnUeAuthenticationsAuthCtxId5gAkaConfirmationPut(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		})
		return
	}

	var input models.ConfirmationData
	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId, &input); prob != nil {
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

// sbi producer handler for UeAuthenticationsDeregisterPost
func OnUeAuthenticationsDeregisterPost(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input models.DeregistrationInfo

	if err := ctx.DecodeRequest(&input); err == nil {
		if prob := prod.UEA_HandleUeAuthenticationsDeregisterPost(input); prob != nil {
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

// sbi producer handler for UeAuthenticationsPost
func OnUeAuthenticationsPost(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input models.AuthenticationInfo
	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEA_HandleUeAuthenticationsPost(input); prob != nil {
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
	UEA_HandleDelete5gAkaAuthenticationResult(authCtxId string) *models.ProblemDetails
	UEA_HandleDeleteEapAuthenticationResult(authCtxId string) *models.ProblemDetails
	UEA_HandleEapAuthMethod(authCtxId string, body *models.EapSession) (*models.EapSession, *models.ProblemDetails)
	UEA_HandleRgAuthenticationsPost(body models.RgAuthenticationInfo) (*models.RgAuthCtx, *models.ProblemDetails)
	UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId string, body *models.ConfirmationData) (*models.ConfirmationDataResponse, *models.ProblemDetails)
	UEA_HandleUeAuthenticationsDeregisterPost(body models.DeregistrationInfo) *models.ProblemDetails
	UEA_HandleUeAuthenticationsPost(body models.AuthenticationInfo) (*models.UEAuthenticationCtx, *models.ProblemDetails)
}
