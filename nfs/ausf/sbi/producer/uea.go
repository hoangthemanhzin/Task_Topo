package producer

import (
	"etrib5gc/nfs/ausf/context"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/udm/ueau"
	"fmt"
	"net/http"
)

func (p *Producer) UEA_HandleDelete5gAkaAuthenticationResult(authCtxId string) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEA_HandleDelete5gAkaAuthenticationResult has not been implemented")
	return
}
func (p *Producer) UEA_HandleDeleteEapAuthenticationResult(authCtxId string) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEA_HandleDeleteEapAuthenticationResult has not been implemented")
	return
}
func (p *Producer) UEA_HandleEapAuthMethod(authCtxId string, body *models.EapSession) (rsp *models.EapSession, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEA_HandleEapAuthMethod has not been implemented")
	return
}
func (p *Producer) UEA_HandleRgAuthenticationsPost(body models.RgAuthenticationInfo) (rsp *models.RgAuthCtx, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEA_HandleRgAuthenticationsPost has not been implemented")
	return
}

func (p *Producer) UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId string, body *models.ConfirmationData) (rsp *models.ConfirmationDataResponse, prob *models.ProblemDetails) {
	log.Infof("Receive UeAuthenticationsAuthCtxId5gAkaConfirmationPut for SUPI=%s", authCtxId)
	ctx := p.ctx
	var err error
	if ue := ctx.GetUeContext(authCtxId); ue != nil {
		if ue.CheckResStar(body.ResStar) {
			ue.WithFields(_logfields).Info("UE is authenticated")

			rsp = &models.ConfirmationDataResponse{
				AuthResult: models.AUTHRESULT_SUCCESS,
				Supi:       ue.Supi(),
				Kseaf:      ue.Kseaf(),
			}
			//success = true
		} else {
			err = fmt.Errorf("%s sends a mismatched resstar", authCtxId)
		}
	} else {
		err = fmt.Errorf("Ue context not found for %s", authCtxId)
	}
	if err != nil {
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}
	}
	//TODO: notify UDM of the authentication result
	return
}

func (p *Producer) UEA_HandleUeAuthenticationsDeregisterPost(body models.DeregistrationInfo) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("UEA_HandleUeAuthenticationsDeregisterPost has not been implemented")
	return
}

func (p *Producer) UEA_HandleUeAuthenticationsPost(body models.AuthenticationInfo) (rsp *models.UEAuthenticationCtx, prob *models.ProblemDetails) {
	log.Infof("Receive an UeAuthenticationsPost [supiOrSuci=%s]", body.SupiOrSuci)

	var req models.AuthenticationInfoRequest

	ueid := body.SupiOrSuci
	snName := body.ServingNetworkName
	req.ServingNetworkName = snName
	var ue *context.UeContext
	if !p.ctx.IsNetworkAuthorized(snName) {
		log.Errorf("Network %s not authorized", snName)
		prob = &models.ProblemDetails{
			Detail: "Serving network is not authorized",
			Status: http.StatusForbidden,
		}
		return
	}
	//log.Infof("%s is authorized", snName)

	//var eapId uint8
	if body.ResynchronizationInfo != nil {
		log.Warnf("Resync is requested for %s]", ueid)
		if ue = p.ctx.GetUeContext(ueid); ue == nil {
			log.Errorf("UE is not found for %s", ueid)
			prob = &models.ProblemDetails{
				Detail: "UE is not found",
				Status: http.StatusForbidden,
			}
			return
		} else {
			if len(body.ResynchronizationInfo.Rand) == 0 {
				body.ResynchronizationInfo.Rand = ue.Rand()
			}
			req.ResynchronizationInfo = body.ResynchronizationInfo
			//	eapId = ue.EapId()
		}
	}
	var err error
	if ue == nil {
		if ue, err = p.ctx.NewUeContext(ueid, snName); err != nil {
			prob = &models.ProblemDetails{
				Cause:  "Internal error",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			}
			return
		}
	}
	uelog := ue.WithFields(_logfields)
	uelog.Info("Request authentication vector from UDM")
	if info, err := ueau.GenerateAuthData(ue.Udm(), ueid, req); err != nil {
		uelog.Errorf("Fails to get authentication vector: %s", err.Error())
		prob = &models.ProblemDetails{
			Cause:  "Upstream server error",
			Detail: err.Error(),
			Status: http.StatusInternalServerError,
		}
		return
	} else {
		uelog.Infof("Receive authentication vector")
		if err := p.ctx.AddUeContext(ue, info); err != nil {
			uelog.Errorf("Failed to create UeContext: %s", err.Error())
			prob = &models.ProblemDetails{
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			}
			return
		} else {
			//log.Info("Build authentication information for %s to send to the AMF", ueid)
			rsp = &models.UEAuthenticationCtx{
				ServingNetworkName: snName,
				AuthType:           info.AuthType,
				Links:              make(map[string]models.LinksValueSchema),
			}
			//there are only 2 auth types (constrained by the previous code
			//block)
			link := p.ctx.Url() + "/nausf-auth/v1/ue-authentications/" + ueid
			if info.AuthType == models.AUTHTYPE_EAP_AKA_PRIME {
				rsp.Links["eap-session"] = models.LinksValueSchema{
					Href: link + "/eap-session",
				}
			} else {
				rsp.Var5gAuthData = ue.Var5gAuthData()
				rsp.Links["5g-aka"] = models.LinksValueSchema{
					Href: link + "/5g-aka-confirmation",
				}
			}
		}
	}
	return
}
