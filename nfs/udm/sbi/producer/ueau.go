package producer

import (
	"etrib5gc/nfs/udm/context"
	"etrib5gc/sbi/models"
	"net/http"
)

func (p *Producer) UEAU_HandleConfirmAuth(supi string, body models.AuthEvent) (rsp *models.AuthEvent, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEAU_HandleConfirmAuth has not been implemented")
	return
}
func (p *Producer) UEAU_HandleDeleteAuth(supi string, authEventId string, body models.AuthEvent) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEAU_HandleDeleteAuth has not been implemented")
	return
}
func (p *Producer) UEAU_HandleGenerateAv(supi string, hssAuthType models.HssAuthTypeInUri, body models.HssAuthenticationInfoRequest) (rsp *models.HssAuthenticationInfoResult, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEAU_HandleGenerateAv has not been implemented")
	return
}
func (p *Producer) UEAU_HandleGetRgAuthData(supiOrSuci string, authenticatedInd bool, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (rsp *models.RgAuthCtx, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEAU_HandleGetRgAuthData has not been implemented")
	return
}

func (p *Producer) UEAU_HandleGenerateAuthData(supiOrSuci string, body models.AuthenticationInfoRequest) (rsp *models.AuthenticationInfoResult, prob *models.ProblemDetails) {
	log.Infof("Receive a GenerateAuthData request for SUCI=%s", supiOrSuci)
	var ue *context.UeContext
	var err error
	//recover suci then get the ue's subscription data
	//need to talk to UDR
	if ue, err = p.ctx.GetUeContext(supiOrSuci); err != nil {
		log.Errorf("Get UeContext[%s] faileds: %s", supiOrSuci, err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusForbidden,
			Detail: err.Error(),
			Cause:  "Authentication rejected",
		}
		return
	}

	//update ue context if needed (sequence number?)
	//need to talk to UDR
	if body.ResynchronizationInfo != nil {
		info := body.ResynchronizationInfo
		if err = ue.Resync(info.Auts, info.Rand); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusForbidden, //TODO: set an appropriate status
				Detail: err.Error(),
				Cause:  "Authentication rejected",
			}
			return
		}
	}
	rsp = &models.AuthenticationInfoResult{}
	//build vector and ask UDR to update the sequence number for the UE
	if rsp.AuthenticationVector, err = ue.BuildAuthenticationVector(body.ServingNetworkName); err != nil {
		prob = &models.ProblemDetails{
			Status: http.StatusForbidden, //TODO: set an appropriate status
			Detail: err.Error(),
			Cause:  "Authentication rejected",
		}

		return
	} else {
		rsp.AuthType = ue.AuthType()
		rsp.Supi = ue.Supi()
	}
	return
}
