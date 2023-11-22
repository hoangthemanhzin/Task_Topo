package ueau

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

// sbi producer handler for ConfirmAuth
func OnConfirmAuth(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	supi := ctx.Param("ueid")
	if len(supi) == 0 {
		//supi is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "supi is required",
		})
		return
	}

	var input models.AuthEvent

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEAU_HandleConfirmAuth(supi, input); prob != nil {
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

// sbi producer handler for DeleteAuth
func OnDeleteAuth(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	supi := ctx.Param("ueid")
	if len(supi) == 0 {
		//supi is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "supi is required",
		})
		return
	}
	authEventId := ctx.Param("authEventId")
	if len(authEventId) == 0 {
		//authEventId is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authEventId is required",
		})
		return
	}

	var input models.AuthEvent

	if err := ctx.DecodeRequest(&input); err == nil {
		if prob := prod.UEAU_HandleDeleteAuth(supi, authEventId, input); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, nil)
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

// sbi producer handler for GenerateAuthData
func OnGenerateAuthData(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	supiOrSuci := ctx.Param("ueid")
	if len(supiOrSuci) == 0 {
		//supiOrSuci is required
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "supiOrSuci is required",
		})

		return
	}

	var input models.AuthenticationInfoRequest

	if err := ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.UEAU_HandleGenerateAuthData(supiOrSuci, input); prob != nil {
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

type Producer interface {
	UEAU_HandleConfirmAuth(supi string, body models.AuthEvent) (*models.AuthEvent, *models.ProblemDetails)
	UEAU_HandleDeleteAuth(supi string, authEventId string, body models.AuthEvent) *models.ProblemDetails
	UEAU_HandleGenerateAuthData(supiOrSuci string, body models.AuthenticationInfoRequest) (*models.AuthenticationInfoResult, *models.ProblemDetails)
}
