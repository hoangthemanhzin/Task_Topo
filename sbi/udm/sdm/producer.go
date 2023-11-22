package sdm

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

// sbi producer handler for GetAmData
func OnGetAmData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	/*
		prod := handler.(Producer)

		supi := ctx.Param("supi")
		if len(supi) == 0 {
			//supi is required
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: "supi is required",
			}))
			return
		}
		supportedFeatures := ctx.Param("supported-features")
		plmnIdStr := ctx.Param("plmn-id")
		var plmnId *models.PlmnId
		var plmnIdErr error
		if plmnId, plmnIdErr = utils.String2PlmnId(plmnIdStr); plmnIdErr != nil {
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: plmnIdErr.Error(),
			}))
			return
		}

		ifNoneMatch := ctx.Param("If-None-Match")
		ifModifiedSince := ctx.Param("If-Modified-Since")

		var apierr *sbi.ApiError
		var successCode int32
		var result models.AccessAndMobilitySubscriptionData

		result, err = prod.SDM_HandleGetAmData(supi, supportedFeatures, plmnId, ifNoneMatch, ifModifiedSince)

		if apierr != nil {
			resp.SetApiError(apierr)
		} else {
			resp.SetBody(int(successCode), &result)
		}
	*/
	return
}

// sbi producer handler for GetSupiOrGpsi
func OnGetSupiOrGpsi(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	/*
		prod := handler.(Producer)

		ueId := ctx.Param("ueId")
		if len(ueId) == 0 {
			//ueId is required
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: "ueId is required",
			}))
			return
		}
		supportedFeatures := ctx.Param("supported-features")
		appPortIdStr := ctx.Param("app-port-id")
		var appPortId *models.AppPortId
		var appPortIdErr error
		if appPortId, appPortIdErr = utils.String2AppPortId(appPortIdStr); appPortIdErr != nil {
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: appPortIdErr.Error(),
			}))
			return
		}

		ifNoneMatch := ctx.Param("If-None-Match")
		ifModifiedSince := ctx.Param("If-Modified-Since")

		var apierr *sbi.ApiError
		var successCode int32
		var result models.IdTranslationResult

		successCode, result, apierr = prod.SDM_HandleGetSupiOrGpsi(ueId, supportedFeatures, appPortId, ifNoneMatch, ifModifiedSince)

		if apierr != nil {
			resp.SetApiError(apierr)
		} else {
			resp.SetBody(int(successCode), &result)
		}
	*/
	return
}

// sbi producer handler for GetSmfSelData
func OnGetSmfSelData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	/*
		prod := handler.(Producer)

		supi := ctx.Param("supi")
		if len(supi) == 0 {
			//supi is required
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: "supi is required",
			}))
			return
		}
		supportedFeatures := ctx.Param("supported-features")
		plmnIdStr := ctx.Param("plmn-id")
		var plmnId *models.PlmnId
		var plmnIdErr error
		if plmnId, plmnIdErr = utils.String2PlmnId(plmnIdStr); plmnIdErr != nil {
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: plmnIdErr.Error(),
			}))
			return
		}

		ifNoneMatch := ctx.Param("If-None-Match")
		ifModifiedSince := ctx.Param("If-Modified-Since")

		var apierr *sbi.ApiError
		var successCode int32
		var result models.SmfSelectionSubscriptionData

		successCode, result, apierr = prod.SDM_HandleGetSmfSelData(supi, supportedFeatures, plmnId, ifNoneMatch, ifModifiedSince)

		if apierr != nil {
			resp.SetApiError(apierr)
		} else {
			resp.SetBody(int(successCode), &result)
		}
	*/
	return
}

// sbi producer handler for GetSmData
func OnGetSmData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	/*
		prod := handler.(Producer)

		supi := ctx.Param("supi")
		if len(supi) == 0 {
			//supi is required
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: "supi is required",
			}))
			return
		}
		supportedFeatures := ctx.Param("supported-features")
		singleNssaiStr := ctx.Param("single-nssai")
		var singleNssai *models.Snssai
		var singleNssaiErr error
		if singleNssai, singleNssaiErr = utils.String2Snssai(singleNssaiStr); singleNssaiErr != nil {
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: singleNssaiErr.Error(),
			}))
			return
		}

		dnn := ctx.Param("dnn")
		plmnIdStr := ctx.Param("plmn-id")
		var plmnId *models.PlmnId
		var plmnIdErr error
		if plmnId, plmnIdErr = utils.String2PlmnId(plmnIdStr); plmnIdErr != nil {
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: plmnIdErr.Error(),
			}))
			return
		}

		ifNoneMatch := ctx.Param("If-None-Match")
		ifModifiedSince := ctx.Param("If-Modified-Since")

		var apierr *sbi.ApiError
		var successCode int32
		var result []models.SessionManagementSubscriptionData

		successCode, result, apierr = prod.SDM_HandleGetSmData(supi, supportedFeatures, singleNssai, dnn, plmnId, ifNoneMatch, ifModifiedSince)

		if apierr != nil {
			resp.SetApiError(apierr)
		} else {
			resp.SetBody(int(successCode), &result)
		}
	*/
	return
}

// sbi producer handler for GetUeCtxInAmfData
func OnGetUeCtxInAmfData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	/*
		prod := handler.(Producer)

		supi := ctx.Param("supi")
		if len(supi) == 0 {
			//supi is required
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: "supi is required",
			}))
			return
		}
		supportedFeatures := ctx.Param("supported-features")

		var apierr *sbi.ApiError
		var successCode int32
		var result models.UeContextInAmfData

		successCode, result, apierr = prod.SDM_HandleGetUeCtxInAmfData(supi, supportedFeatures)

		if apierr != nil {
			resp.SetApiError(apierr)
		} else {
			resp.SetBody(int(successCode), &result)
		}
	*/
	return
}

// sbi producer handler for GetUeCtxInSmfData
func OnGetUeCtxInSmfData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	/*
		prod := handler.(Producer)

		supi := ctx.Param("supi")
		if len(supi) == 0 {
			//supi is required
			resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
				Title:  "Bad request",
				Status: http.StatusBadRequest,
				Detail: "supi is required",
			}))
			return
		}
		supportedFeatures := ctx.Param("supported-features")

		var apierr *sbi.ApiError
		var successCode int32
		var result models.UeContextInSmfData

		successCode, result, apierr = prod.SDM_HandleGetUeCtxInSmfData(supi, supportedFeatures)

		if apierr != nil {
			resp.SetApiError(apierr)
		} else {
			resp.SetBody(int(successCode), &result)
		}
	*/
	return
}

type Producer interface {
	SDM_HandleGetAmData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (*models.AccessAndMobilitySubscriptionData, *models.ProblemDetails)
	SDM_HandleGetSupiOrGpsi(ueId string, supportedFeatures string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (*models.IdTranslationResult, *models.ProblemDetails)
	SDM_HandleGetSmfSelData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (*models.SmfSelectionSubscriptionData, *models.ProblemDetails)
	SDM_HandleGetSmData(supi string, supportedFeatures string, singleNssai *models.Snssai, dnn string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) ([]models.SessionManagementSubscriptionData, *models.ProblemDetails)
	SDM_HandleGetUeCtxInAmfData(supi string, supportedFeatures string) (*models.UeContextInAmfData, *models.ProblemDetails)
	SDM_HandleGetUeCtxInSmfData(supi string, supportedFeatures string) (*models.UeContextInSmfData, *models.ProblemDetails)
}
