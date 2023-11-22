package upmf2smf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upmf2smf"

var _routes = sbi.SbiRoutes{
	{
		Label:   "PathQuery",
		Method:  http.MethodGet,
		Path:    "path",
		Handler: OnGetUpfPath,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upmf2smf",
		Routes:  _routes,
		Handler: p,
	}
}
func GetUpfPath(client sbi.ConsumerClient, body n43.UpfPathQuery) (path *n43.UpfPath, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/path", SERVICE_PATH)
	req.Body = &body

	req.HeaderParams["Accept"] = "application/json, application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		path = &n43.UpfPath{}
		response.Body = path
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
