package sdm

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

const (
	SERVICE_PATH = "{apiRoot}/nudm-sdm/v2"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi Identifier of the UE
@param supportedFeatures Supported Features
@param plmnId serving PLMN ID
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return *models.AccessAndMobilitySubscriptionData,
*/
func GetAmData(client sbi.ConsumerClient, supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp *models.AccessAndMobilitySubscriptionData, err error) {
	/*
		if len(supi) == 0 {
			err = fmt.Errorf("supi is required")
			return
		}
		//create a request
		req := sbi.DefaultRequest()
		req.Method = http.MethodGet

		req.Path = fmt.Sprintf("%s/{supi}/am-data", SERVICE_PATH)
		req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
		if len(supportedFeatures) > 0 {
			req.QueryParams.Add("supported-features", supportedFeatures)
		}
		plmnIdStr := utils.Param2String(plmnId)
		if len(plmnIdStr) > 0 {
			req.QueryParams.Add("plmn-id", plmnIdStr)
		}
		if len(ifNoneMatch) > 0 {
			req.HeaderParams["If-None-Match"] = ifNoneMatch
		}
		if len(ifModifiedSince) > 0 {
			req.HeaderParams["If-Modified-Since"] = ifModifiedSince
		}
		req.HeaderParams["Accept"] = "application/json, application/problem+json"
		//send the request
		var resp *sbi.Response
		if resp, err = client.Send(req); err != nil {
			return
		}

		//handle the response
		if resp.StatusCode >= 300 {
			if resp.StatusCode == 400 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 404 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 500 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 503 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.Body != nil {
				if err = client.DecodeResponse(resp); err == nil {
					err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
				}
				return
			} else {
				err = fmt.Errorf("%d is unknown to GetAmData", resp.StatusCode)
				return
			}
		}

		resp.Body = &result
		if err = client.DecodeResponse(resp); err == nil {
			err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
		}
	*/
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId Identifier of the UE
@param supportedFeatures Supported Features
@param appPortId Application port identifier
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return *models.IdTranslationResult,
*/
func GetSupiOrGpsi(client sbi.ConsumerClient, ueId string, supportedFeatures string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (rsp *models.IdTranslationResult, err error) {
	/*
		if len(ueId) == 0 {
			err = fmt.Errorf("ueId is required")
			return
		}
		//create a request
		req := sbi.DefaultRequest()
		req.Method = http.MethodGet

		req.Path = fmt.Sprintf("%s/{ueId}/id-translation-result", SERVICE_PATH)
		req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
		if len(supportedFeatures) > 0 {
			req.QueryParams.Add("supported-features", supportedFeatures)
		}
		appPortIdStr := utils.Param2String(appPortId)
		if len(appPortIdStr) > 0 {
			req.QueryParams.Add("app-port-id", appPortIdStr)
		}
		if len(ifNoneMatch) > 0 {
			req.HeaderParams["If-None-Match"] = ifNoneMatch
		}
		if len(ifModifiedSince) > 0 {
			req.HeaderParams["If-Modified-Since"] = ifModifiedSince
		}
		req.HeaderParams["Accept"] = "application/json, application/problem+json"
		//send the request
		var resp *sbi.Response
		if resp, err = client.Send(req); err != nil {
			return
		}

		//handle the response
		if resp.StatusCode >= 300 {
			if resp.StatusCode == 400 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 404 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 500 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 503 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.Body != nil {
				if err = client.DecodeResponse(resp); err == nil {
					err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
				}
				return
			} else {
				err = fmt.Errorf("%d is unknown to GetSupiOrGpsi", resp.StatusCode)
				return
			}
		}

		resp.Body = &result
		if err = client.DecodeResponse(resp); err == nil {
			err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
		}
	*/
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi Identifier of the UE
@param supportedFeatures Supported Features
@param plmnId serving PLMN ID
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return *models.SmfSelectionSubscriptionData,
*/
func GetSmfSelData(client sbi.ConsumerClient, supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp *models.SmfSelectionSubscriptionData, err error) {
	/*
		if len(supi) == 0 {
			err = fmt.Errorf("supi is required")
			return
		}
		//create a request
		req := sbi.DefaultRequest()
		req.Method = http.MethodGet

		req.Path = fmt.Sprintf("%s/{supi}/smf-select-data", SERVICE_PATH)
		req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
		if len(supportedFeatures) > 0 {
			req.QueryParams.Add("supported-features", supportedFeatures)
		}
		plmnIdStr := utils.Param2String(plmnId)
		if len(plmnIdStr) > 0 {
			req.QueryParams.Add("plmn-id", plmnIdStr)
		}
		if len(ifNoneMatch) > 0 {
			req.HeaderParams["If-None-Match"] = ifNoneMatch
		}
		if len(ifModifiedSince) > 0 {
			req.HeaderParams["If-Modified-Since"] = ifModifiedSince
		}
		req.HeaderParams["Accept"] = "application/json, application/problem+json"
		//send the request
		var resp *sbi.Response
		if resp, err = client.Send(req); err != nil {
			return
		}

		//handle the response
		if resp.StatusCode >= 300 {
			if resp.StatusCode == 400 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 404 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 500 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 503 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.Body != nil {
				if err = client.DecodeResponse(resp); err == nil {
					err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
				}
				return
			} else {
				err = fmt.Errorf("%d is unknown to GetSmfSelData", resp.StatusCode)
				return
			}
		}

		resp.Body = &result
		if err = client.DecodeResponse(resp); err == nil {
			err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
		}
	*/
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi Identifier of the UE
@param supportedFeatures Supported Features
@param singleNssai
@param dnn
@param plmnId
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return []models.SessionManagementSubscriptionData,
*/
func GetSmData(client sbi.ConsumerClient, supi string, supportedFeatures string, singleNssai *models.Snssai, dnn string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp []models.SessionManagementSubscriptionData, err error) {
	/*
		if len(supi) == 0 {
			err = fmt.Errorf("supi is required")
			return
		}
		//create a request
		req := sbi.DefaultRequest()
		req.Method = http.MethodGet

		req.Path = fmt.Sprintf("%s/{supi}/sm-data", SERVICE_PATH)
		req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
		if len(supportedFeatures) > 0 {
			req.QueryParams.Add("supported-features", supportedFeatures)
		}
		singleNssaiStr := utils.Param2String(singleNssai)
		if len(singleNssaiStr) > 0 {
			req.QueryParams.Add("single-nssai", singleNssaiStr)
		}
		if len(dnn) > 0 {
			req.QueryParams.Add("dnn", dnn)
		}
		plmnIdStr := utils.Param2String(plmnId)
		if len(plmnIdStr) > 0 {
			req.QueryParams.Add("plmn-id", plmnIdStr)
		}
		if len(ifNoneMatch) > 0 {
			req.HeaderParams["If-None-Match"] = ifNoneMatch
		}
		if len(ifModifiedSince) > 0 {
			req.HeaderParams["If-Modified-Since"] = ifModifiedSince
		}
		req.HeaderParams["Accept"] = "application/json, application/problem+json"
		//send the request
		var resp *sbi.Response
		if resp, err = client.Send(req); err != nil {
			return
		}

		//handle the response
		if resp.StatusCode >= 300 {
			if resp.StatusCode == 400 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 404 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 500 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 503 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.Body != nil {
				if err = client.DecodeResponse(resp); err == nil {
					err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
				}
				return
			} else {
				err = fmt.Errorf("%d is unknown to GetSmData", resp.StatusCode)
				return
			}
		}

		resp.Body = &result
		if err = client.DecodeResponse(resp); err == nil {
			err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
		}
	*/
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi Identifier of the UE
@param supportedFeatures Supported Features
@return *models.UeContextInAmfData,
*/
func GetUeCtxInAmfData(client sbi.ConsumerClient, supi string, supportedFeatures string) (rsp *models.UeContextInAmfData, err error) {
	/*
		if len(supi) == 0 {
			err = fmt.Errorf("supi is required")
			return
		}
		//create a request
		req := sbi.DefaultRequest()
		req.Method = http.MethodGet

		req.Path = fmt.Sprintf("%s/{supi}/ue-context-in-amf-data", SERVICE_PATH)
		req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
		if len(supportedFeatures) > 0 {
			req.QueryParams.Add("supported-features", supportedFeatures)
		}
		req.HeaderParams["Accept"] = "application/json, application/problem+json"
		//send the request
		var resp *sbi.Response
		if resp, err = client.Send(req); err != nil {
			return
		}

		//handle the response
		if resp.StatusCode >= 300 {
			if resp.StatusCode == 400 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 404 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 500 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 503 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.Body != nil {
				if err = client.DecodeResponse(resp); err == nil {
					err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
				}
				return
			} else {
				err = fmt.Errorf("%d is unknown to GetUeCtxInAmfData", resp.StatusCode)
				return
			}
		}

		resp.Body = &result
		if err = client.DecodeResponse(resp); err == nil {
			err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
		}
	*/
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi Identifier of the UE
@param supportedFeatures Supported Features
@return *models.UeContextInSmfData,
*/
func GetUeCtxInSmfData(client sbi.ConsumerClient, supi string, supportedFeatures string) (rsp *models.UeContextInSmfData, err error) {
	/*
		if len(supi) == 0 {
			err = fmt.Errorf("supi is required")
			return
		}
		//create a request
		req := sbi.DefaultRequest()
		req.Method = http.MethodGet

		req.Path = fmt.Sprintf("%s/{supi}/ue-context-in-smf-data", SERVICE_PATH)
		req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
		if len(supportedFeatures) > 0 {
			req.QueryParams.Add("supported-features", supportedFeatures)
		}
		req.HeaderParams["Accept"] = "application/json, application/problem+json"
		//send the request
		var resp *sbi.Response
		if resp, err = client.Send(req); err != nil {
			return
		}

		//handle the response
		if resp.StatusCode >= 300 {
			if resp.StatusCode == 400 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 404 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 500 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.StatusCode == 503 {
				resp.Body = &models.ProblemDetails{}
			}
			if resp.Body != nil {
				if err = client.DecodeResponse(resp); err == nil {
					err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
				}
				return
			} else {
				err = fmt.Errorf("%d is unknown to GetUeCtxInSmfData", resp.StatusCode)
				return
			}
		}

		resp.Body = &result
		if err = client.DecodeResponse(resp); err == nil {
			err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
		}
	*/
	return
}
