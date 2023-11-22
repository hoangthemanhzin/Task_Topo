package pdu

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "pdu/nsmf-pdusession/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param smContextRef SM context reference
@return *models.SmContextReleasedData,
*/
func ReleaseSmContext(client sbi.ConsumerClient, smContextRef string, body *models.ReleaseSmContextRequest) (rsp *models.SmContextReleasedData, err error) {

	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sm-contexts/{smContextRef}/release", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"smContextRef"+"}", url.PathEscape(smContextRef), -1)
	req.Body = body
	req.HeaderParams["Content-Type"] = "multipart/related"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.SmContextReleasedData{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		var prob models.ProblemDetails
		response.Body = &prob
		if err := client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%d-%s: Problem = %s", response.StatusCode, response.Status, prob.Detail)
		}

	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param smContextRef SM context reference
@return *models.SmContextUpdatedData,
*/
func UpdateSmContext(client sbi.ConsumerClient, smContextRef string, body models.UpdateSmContextRequest) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {

	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sm-contexts/{smContextRef}/modify", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"smContextRef"+"}", url.PathEscape(smContextRef), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch response.StatusCode {
	case 200:
		rsp = &models.UpdateSmContextResponse{}
		response.Body = &rsp
		err = client.DecodeResponse(response)
	case 400:
		ersp = &models.UpdateSmContextErrorResponse{}
		response.Body = ersp
		err = client.DecodeResponse(response)
	default:
		prob := &models.ProblemDetails{}
		response.Body = prob
		if err = client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%d-%s: Problem = %s", response.StatusCode, response.Status, prob.Detail)
		}
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.SmContextCreatedData,
*/
func PostSmContexts(client sbi.ConsumerClient, body models.PostSmContextsRequest, callback models.Callback) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost
	req.Callback = callback

	req.Path = fmt.Sprintf("%s/sm-contexts", SERVICE_PATH)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json"
	req.HeaderParams["Callback"] = string(callback)
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		err = fmt.Errorf("Send Sbi: %s", err.Error())
		return
	}

	switch response.StatusCode {
	case 201:
		rsp = &models.PostSmContextsResponse{}
		response.Body = &rsp
		err = client.DecodeResponse(response)
	case 400:
		ersp = &models.PostSmContextsErrorResponse{}
		response.Body = ersp
		err = client.DecodeResponse(response)
	default:
		var prob models.ProblemDetails
		response.Body = &prob
		if err := client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%d-%s: Problem = %s", response.StatusCode, response.Status, prob.Detail)
		}
	}

	return
}
