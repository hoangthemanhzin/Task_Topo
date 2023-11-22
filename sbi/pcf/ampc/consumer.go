package ampc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "ampc/npcf-am-policy-control/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.PolicyAssociation,
*/
func CreateIndividualAMPolicyAssociation(client sbi.ConsumerClient, body models.PolicyAssociationRequest) (rsp *models.PolicyAssociation, err error) {
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/policies", SERVICE_PATH)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.PolicyAssociation{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param polAssoId Identifier of a policy association
@return
*/
func DeleteIndividualAMPolicyAssociation(client sbi.ConsumerClient, polAssoId string) (err error) {
	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/policies/{polAssoId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"polAssoId"+"}", url.PathEscape(polAssoId), -1)
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

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param polAssoId Identifier of a policy association
@return *models.PolicyAssociation,
*/
func ReadIndividualAMPolicyAssociation(client sbi.ConsumerClient, polAssoId string) (rsp *models.PolicyAssociation, err error) {
	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/policies/{polAssoId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"polAssoId"+"}", url.PathEscape(polAssoId), -1)
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.PolicyAssociation{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param polAssoId Identifier of a policy association
@return *models.PolicyUpdate,
*/
func ReportObservedEventTriggersForIndividualAMPolicyAssociation(client sbi.ConsumerClient, polAssoId string, body models.PolicyAssociationUpdateRequest) (rsp *models.PolicyUpdate, err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/policies/{polAssoId}/update", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"polAssoId"+"}", url.PathEscape(polAssoId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.StatusCode == 200 {
		rsp = &models.PolicyUpdate{}
		response.Body = rsp
		err = client.DecodeResponse(response)
	} else {
		err = fmt.Errorf("%d: %s", response.StatusCode, response.Status)
	}
	return
}
