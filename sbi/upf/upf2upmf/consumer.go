package upf2upmf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upf2upmf"

var _routes = sbi.SbiRoutes{
	{
		Label:   "Heartbeat",
		Method:  http.MethodGet,
		Path:    "heartbeat",
		Handler: OnHeartbeat,
	},
	{
		Label: 	 "Activate",
		Method:   http.MethodGet,
		Path:    "activate",
		Handler:  OnActivate,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upf2upmf",
		Routes:  _routes,
		Handler: p,
	}
}
func Heartbeat(client sbi.ConsumerClient, body n42.HeartbeatRequest) (rsp *n42.HeartbeatResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/heartbeat", SERVICE_PATH)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json, application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &n42.HeartbeatResponse{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		prob := &models.ProblemDetails{}
		response.Body = prob
		if err = client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.Status, prob.Detail)
		}
	}
	return
}

func UpfActivate(client sbi.ConsumerClient, body n42.UpfActivateQuery) (rsp *n42.UpfActivate, err error){

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/activate", SERVICE_PATH)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json, application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	//handle the response
	if response.StatusCode == 200 {
		rsp = &n42.UpfActivate{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		prob := &models.ProblemDetails{}
		response.Body = prob
		if err = client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.Status, prob.Detail)
		}
	}
	return
}
