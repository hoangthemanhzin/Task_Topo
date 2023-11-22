package ueau

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "ueau/nudm-ueau/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi SUPI of the user
@return *models.AuthEvent,
*/
func ConfirmAuth(client sbi.ConsumerClient, supi string, body models.AuthEvent) (rsp *models.AuthEvent, err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/{supi}/auth-events", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
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
		rsp = &models.AuthEvent{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi SUPI of the user
@param authEventId authEvent Id
@return
*/
func DeleteAuth(client sbi.ConsumerClient, supi string, authEventId string, body models.AuthEvent) (err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if len(authEventId) == 0 {
		err = fmt.Errorf("authEventId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/{supi}/auth-events/{authEventId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
	req.Path = strings.Replace(req.Path, "{"+"authEventId"+"}", url.PathEscape(authEventId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/problem+json"
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
@param supiOrSuci SUPI or SUCI of the user
@return *models.AuthenticationInfoResult,
*/
func GenerateAuthData(client sbi.ConsumerClient, supiOrSuci string, body models.AuthenticationInfoRequest) (rsp *models.AuthenticationInfoResult, err error) {

	if len(supiOrSuci) == 0 {
		err = fmt.Errorf("supiOrSuci is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/{supiOrSuci}/security-information/generate-auth-data", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supiOrSuci"+"}", url.PathEscape(supiOrSuci), -1)
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
		rsp = &models.AuthenticationInfoResult{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}
