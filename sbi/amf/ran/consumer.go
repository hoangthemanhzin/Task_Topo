package ran

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"fmt"
	"net/http"
)

const (
	SERVICE_PATH = "ran"
)

func InitUeContext(client sbi.ConsumerClient, body *n2models.InitUeContextRequest, callback models.Callback) (rsp *n2models.InitUeContextResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Callback = callback
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/init", SERVICE_PATH)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	req.HeaderParams["Callback"] = string(callback)
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	if response.StatusCode != 200 {
		prob := &models.ProblemDetails{}
		response.Body = prob
		if err = client.DecodeResponse(response); err == nil {
			err = prob.MakeError()
		}

	} else {
		rsp = &n2models.InitUeContextResponse{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	}
	return
}

func UlNasTransport(client sbi.ConsumerClient, ueId int64, body *n2models.UlNasTransport) (err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/%d/ul", SERVICE_PATH, ueId)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	if response.StatusCode != 200 {
		prob := &models.ProblemDetails{}
		if err = client.DecodeResponse(response); err == nil {
			err = prob.MakeError()
		}
	}
	return
}

func NasNonDeliveryIndication(client sbi.ConsumerClient, ueId int64, body *n2models.NasNonDeliveryIndication) (err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/%d/naserr", SERVICE_PATH, ueId)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	if response.StatusCode != 200 {
		prob := &models.ProblemDetails{}
		if err = client.DecodeResponse(response); err == nil {
			err = prob.MakeError()
		}
	}

	return
}

func PduSessionResourceNotification(client sbi.ConsumerClient, ueId int64, body *n2models.PduSessResNot) (err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/%d/rep/pdu", SERVICE_PATH, ueId)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	if response.StatusCode != 200 {
		prob := &models.ProblemDetails{}
		if err = client.DecodeResponse(response); err == nil {
			err = prob.MakeError()
		}
	}

	return
}

func PduSessionResourceModifyIndication(client sbi.ConsumerClient, ueId int64, body *n2models.PduSessResModInd) (err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/%d/rep/modind", SERVICE_PATH, ueId)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	if response.StatusCode != 200 {
		prob := &models.ProblemDetails{}
		if err = client.DecodeResponse(response); err == nil {
			err = prob.MakeError()
		}
	}

	return
}

func RrcInactiveTransactionReport(client sbi.ConsumerClient, ueId int64, body *n2models.RrcInactTranRep) (err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/%d/rep/rrc", SERVICE_PATH, ueId)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	if response.StatusCode != 200 {
		prob := &models.ProblemDetails{}
		if err = client.DecodeResponse(response); err == nil {
			err = prob.MakeError()
		}
	}

	return
}
