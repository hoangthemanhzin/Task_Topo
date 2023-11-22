package upmf2upf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upmf2upf"

var _routes = sbi.SbiRoutes{
	{
		Label:   "Register",
		Method:  http.MethodPost,
		Path:    "register",
		Handler: OnRegister,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upmf2upf",
		Routes:  _routes,
		Handler: p,
	}
}
func Register(client sbi.ConsumerClient, body n42.RegistrationRequest) (rsp *n42.RegistrationResponse, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/register", SERVICE_PATH)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		fmt.Println("Send not ok")
		return
	}

	//handle the response
	if response.StatusCode == 201 {
		rsp = &n42.RegistrationResponse{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		prob := &models.ProblemDetails{}
		response.Body = prob
		if err = client.DecodeResponse(response); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.Status, prob.Detail)
		} else {
			err = fmt.Errorf("%s", response.Status)
		}
	}
	return
}
