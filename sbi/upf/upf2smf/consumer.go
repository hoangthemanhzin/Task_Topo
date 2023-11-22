package upf2smf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n41"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upf2smf"

var _routes = sbi.SbiRoutes{
	{
		Label:   "SessionEstablishment",
		Method:  http.MethodPost,
		Path:    "sess-create",
		Handler: OnSessionEstablishment,
	},
	{
		Label:   "SessionModification",
		Method:  http.MethodPost,
		Path:    "sess-modify/:seid",
		Handler: OnSessionModification,
	},
	{
		Label:   "SessionDeletion",
		Method:  http.MethodPost,
		Path:    "sess-delete/:seid",
		Handler: OnSessionDeletion,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upf2smf",
		Routes:  _routes,
		Handler: p,
	}
}

func SessionEstablishment(client sbi.ConsumerClient, body n41.SessionEstablishmentRequest) (rsp *n41.SessionEstablishmentResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sess-create", SERVICE_PATH)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 201 {
		rsp = &n41.SessionEstablishmentResponse{}
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

func SessionModification(client sbi.ConsumerClient, seid uint64, body n41.SessionModificationRequest) (rsp *n41.SessionModificationResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sess-modify/%d", SERVICE_PATH, seid)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 201 {
		rsp = &n41.SessionModificationResponse{}
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

func SessionDeletion(client sbi.ConsumerClient, seid uint64, body n41.SessionDeletionRequest) (rsp *n41.SessionDeletionResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/sess-delete/%d", SERVICE_PATH, seid)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 201 {
		rsp = &n41.SessionDeletionResponse{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		prob := &models.ProblemDetails{}
		response.Body = prob
		if err = client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.Status, prob.Detail)
		} else {
			err = fmt.Errorf("%s: Decode problem failed: %+v", response.Status, err)
		}
	}
	return
}
