package upmf2fe

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upmf2fe"

var _routes = sbi.SbiRoutes{
	{
		Label:   "TopoQuery",
		Method:  http.MethodGet,
		Path:    "topo",
		Handler: OnGetTopo,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upmf2fe",
		Routes:  _routes,
		Handler: p,
	}
}

func GetTopo(client sbi.ConsumerClient) (topo *n43.TopoUpf, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/path", SERVICE_PATH)
	//req.Body = &body

	req.HeaderParams["Accept"] = "application/json, application/json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		topo = &n43.TopoUpf{}
		response.Body = topo
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
