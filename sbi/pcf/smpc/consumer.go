package smpc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "smpc/npcf-smpolicycontrol/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param smPolicyId Identifier of a policy association
@return
*/
func DeleteSMPolicy(client sbi.ConsumerClient, smPolicyId string, body models.SmPolicyDeleteData) (err error) {
	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sm-policies/{smPolicyId}/delete", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"smPolicyId"+"}", url.PathEscape(smPolicyId), -1)
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
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param smPolicyId Identifier of a policy association
@return *models.SmPolicyControl,
*/
func GetSMPolicy(client sbi.ConsumerClient, smPolicyId string) (rsp *models.SmPolicyControl, err error) {
	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/sm-policies/{smPolicyId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"smPolicyId"+"}", url.PathEscape(smPolicyId), -1)
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.SmPolicyControl{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param smPolicyId Identifier of a policy association
@return *models.SmPolicyDecision,
*/
func UpdateSMPolicy(client sbi.ConsumerClient, smPolicyId string, body models.SmPolicyUpdateContextData) (rsp *models.SmPolicyDecision, err error) {
	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sm-policies/{smPolicyId}/update", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"smPolicyId"+"}", url.PathEscape(smPolicyId), -1)
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
		rsp = &models.SmPolicyDecision{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.SmPolicyDecision,
*/
func CreateSMPolicy(client sbi.ConsumerClient, body models.SmPolicyContextData) (rsp *models.SmPolicyDecision, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sm-policies", SERVICE_PATH)
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
		rsp = &models.SmPolicyDecision{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}

	return
}
