package group

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils"
	"fmt"
	"net/http"
)

const (
	SERVICE_PATH = "{apiRoot}/nudr-group-id-map/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param nfType Type of NF
@param subscriberId Identifier of the subscriber
@return map[string]string,
*/
func GetNfGroupIDs(client sbi.ConsumerClient, nfType []models.NFType, subscriberId string) (rsp map[string]string, err error) {

	nfTypeStr := utils.Param2String(nfType)
	if len(nfTypeStr) == 0 {
		err = fmt.Errorf("nfType is required")
		return
	}
	if len(subscriberId) == 0 {
		err = fmt.Errorf("subscriberId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/nf-group-ids", SERVICE_PATH)
	req.QueryParams.Add("nf-type", nfTypeStr)

	req.QueryParams.Add("subscriberId", subscriberId)

	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = make(map[string]string)
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}
