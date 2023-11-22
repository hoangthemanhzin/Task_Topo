package comm

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "comm/namf-comm/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueContextId UE Context Identifier
@return *models.UeContextCreatedData,
*/
func CreateUEContext(client sbi.ConsumerClient, ueContextId string, body models.CreateUEContextRequest) (rsp *models.CreateUEContextResponse, ersp *models.CreateUEContextErrorResponse, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/ue-contexts/{ueContextId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"ueContextId"+"}", url.PathEscape(ueContextId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "multipart/related"
	req.HeaderParams["Accept"] = "application/json, multipart/related, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	switch response.StatusCode {
	case 400, 403, 500:
		ersp = &models.CreateUEContextErrorResponse{}
		response.Body = ersp
	case 201:
		rsp = &models.CreateUEContextResponse{}
		response.Body = rsp
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}

	if response.Body != nil {
		err = client.DecodeResponse(response)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueContextId UE Context Identifier
@return *models.UeContextTransferRspData,
*/
func UEContextTransfer(client sbi.ConsumerClient, ueContextId string, body models.UEContextTransferRequest) (rsp *models.UEContextTransferResponse, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-contexts/{ueContextId}/transfer", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"ueContextId"+"}", url.PathEscape(ueContextId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "multipart/related"
	req.HeaderParams["Accept"] = "application/json, multipart/related, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	if response.StatusCode == 200 {
		rsp = &models.UEContextTransferResponse{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueContextId UE Context Identifier
@return *models.N1N2MessageTransferRspData,
*/
func N1N2MessageTransfer(client sbi.ConsumerClient, ueContextId string, body models.N1N2MessageTransferRequest) (rsp *models.N1N2MessageTransferRspData, ersp *models.N1N2MessageTransferError, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-contexts/{ueContextId}/n1-n2-messages", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"ueContextId"+"}", url.PathEscape(ueContextId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "multipart/related"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = &models.N1N2MessageTransferRspData{}
		response.Body = rsp
	case 409, 504:
		ersp = &models.N1N2MessageTransferError{}
		response.Body = ersp
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}

	if response.Body != nil {
		err = client.DecodeResponse(response)
	}
	return
}
