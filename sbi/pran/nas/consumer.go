package nas

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models/n2models"
	"fmt"
	"net/http"
)

const (
	SERVICE_PATH = "amf"
)

func InitUeContextStatus(client sbi.ConsumerClient, ueid int64, body n2models.InitUeContextStatus) (err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/status/%d", SERVICE_PATH, ueid)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
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

func NasDl(client sbi.ConsumerClient, ueid int64, body n2models.NasDlMsg) (err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/dl/%d", SERVICE_PATH, ueid)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
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

func InitCtxSetupReq(client sbi.ConsumerClient, ueid int64, body *n2models.InitCtxSetupReq) (rsp *n2models.InitCtxSetupRsp, ersp *n2models.InitCtxSetupFailure, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/uectx/%d/set", SERVICE_PATH, ueid)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch response.StatusCode {
	case 200:
		rsp = &n2models.InitCtxSetupRsp{}
		response.Body = rsp
	case 400:
		ersp = &n2models.InitCtxSetupFailure{}
		response.Body = ersp
	default:
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	if response.Body != nil {
		err = client.DecodeResponse(response)
	}
	return
}

func UeCtxModReq(client sbi.ConsumerClient, ueid int64, body *n2models.UeCtxModReq) (rsp *n2models.UeCtxModRsp, ersp *n2models.UeCtxModFail, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/uectx/%d/mod", SERVICE_PATH, ueid)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch response.StatusCode {
	case 200:
		rsp = &n2models.UeCtxModRsp{}
		response.Body = rsp
	case 400:
		ersp = &n2models.UeCtxModFail{}
		response.Body = ersp
	default:
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	if response.Body != nil {
		err = client.DecodeResponse(response)
	}
	return
}

func UeCtxRelCmd(client sbi.ConsumerClient, ueid int64, body *n2models.UeCtxRelCmd) (rsp *n2models.UeCtxRelCmpl, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/uectx/%d/rel", SERVICE_PATH, ueid)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch response.StatusCode {
	case 200:
		rsp = &n2models.UeCtxRelCmpl{}
		response.Body = rsp
	default:
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	if response.Body != nil {
		err = client.DecodeResponse(response)
	}

	return
}

func PduSessResSetReq(client sbi.ConsumerClient, ueid int64, body *n2models.PduSessResSetReq) (rsp *n2models.PduSessResSetRsp, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/pdu/%d/set", SERVICE_PATH, ueid)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	} else {
		rsp = &n2models.PduSessResSetRsp{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	}

	return
}

func PduSessResModReq(client sbi.ConsumerClient, ueid int64, body *n2models.PduSessResModReq) (rsp *n2models.PduSessResModRsp, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/pdu/%d/mod", SERVICE_PATH, ueid)
	req.Body = body
	req.HeaderParams["content-type"] = "application/json-patch+json"
	req.HeaderParams["accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	} else {
		rsp = &n2models.PduSessResModRsp{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	}
	return
}

func PduSessResRelCmd(client sbi.ConsumerClient, ueid int64, body *n2models.PduSessResRelCmd) (rsp *n2models.PduSessResRelRsp, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/pdu/%d/rel", SERVICE_PATH, ueid)
	req.Body = body
	req.HeaderParams["content-type"] = "application/json-patch+json"
	req.HeaderParams["accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	} else {
		rsp = &n2models.PduSessResRelRsp{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	}

	return
}
