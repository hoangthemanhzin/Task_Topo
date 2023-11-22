package uea

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "uea/nausf-auth/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param authCtxId
@return
*/
func Delete5gAkaAuthenticationResult(client sbi.ConsumerClient, authCtxId string) (err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/5g-aka-confirmation", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param authCtxId
@return
*/
func DeleteEapAuthenticationResult(client sbi.ConsumerClient, authCtxId string) (err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/eap-session", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param authCtxId
@return *models.EapSession,
*/
func EapAuthMethod(client sbi.ConsumerClient, authCtxId string, body *models.EapSession) (rsp *models.EapSession, err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/eap-session", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/3gppHal+json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.EapSession{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.RgAuthCtx,
*/
func RgAuthenticationsPost(client sbi.ConsumerClient, body models.RgAuthenticationInfo) (rsp *models.RgAuthCtx, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/rg-authentications", SERVICE_PATH)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.RgAuthCtx{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param authCtxId
@return *models.ConfirmationDataResponse,
*/
func UeAuthenticationsAuthCtxId5gAkaConfirmationPut(client sbi.ConsumerClient, authCtxId string, body *models.ConfirmationData) (rsp *models.ConfirmationDataResponse, err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/5g-aka-confirmation", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.ConfirmationDataResponse{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return
*/
func UeAuthenticationsDeregisterPost(client sbi.ConsumerClient, body models.DeregistrationInfo) (err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-authentications/deregister", SERVICE_PATH)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.UEAuthenticationCtx,
*/
func UeAuthenticationsPost(client sbi.ConsumerClient, body models.AuthenticationInfo) (rsp *models.UEAuthenticationCtx, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-authentications", SERVICE_PATH)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/3gppHal+json, application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.UEAuthenticationCtx{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}
