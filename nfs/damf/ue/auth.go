package ue

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"etrib5gc/logctx"
	"etrib5gc/sbi/ausf/uea"
	"etrib5gc/sbi/models"
	"etrib5gc/util/sec"
	"fmt"

	"github.com/free5gc/nas/nasMessage"
)

func (uectx *UeContext) authenticate() (err error) {
	uectx.Trace("Start authenticating")
	info := &models.AuthenticationInfo{
		SupiOrSuci:         uectx.suci, //NOTE: may use supi as well
		ServingNetworkName: uectx.ServingNetwork(),
	}

	uectx.Infof("Send an UeAuthenticationsPost")
	var authctx *models.UEAuthenticationCtx
	if authctx, err = uea.UeAuthenticationsPost(uectx.ausfcli, *info); err != nil {
		err = fmt.Errorf("Authentication with AUSF failed: %s", err.Error())
		return
	} else {
		uectx.Info("Receive Authentication Context")
		if err = uectx.loadCtx(authctx); err != nil {
			return
		}
		if err = uectx.sendAuthenticationRequest(); err != nil { //send challenging to UE
			err = fmt.Errorf("Send AuthenticationRequest failed: %s", err.Error())
		}
	}
	return
}

func (uectx *UeContext) loadCtx(authctx *models.UEAuthenticationCtx) (err error) {
	uectx.authtype = authctx.AuthType
	dat := &authctx.Var5gAuthData
	if uectx.rand, err = hex.DecodeString(dat.Rand); err != nil {
		return
	}
	if uectx.autn, err = hex.DecodeString(dat.Autn); err != nil {
		return
	}
	if uectx.hxresstar, err = hex.DecodeString(dat.HxresStar); err != nil {
		return
	}
	return
}

func (uectx *UeContext) handleAuthenticationResponse(msg *nasMessage.AuthenticationResponse) (err error) {
	uectx.Infof("Receive AuthenticationResponse")
	switch uectx.authtype {
	case models.AUTHTYPE__5_G_AKA:
		resstar := msg.AuthenticationResponseParameter.GetRES()
		hx := sha256.Sum256(append(uectx.rand, resstar[:]...))
		hresstar := hx[16:]
		uectx.Tracef("rand=%x,resstar=%x,hash=%x", uectx.rand, resstar[:], hresstar)
		//check resstar
		if !bytes.Equal(hresstar, uectx.hxresstar) {
			err = fmt.Errorf("Mismatch hxresstar %x vs %x", hresstar, uectx.hxresstar)
			return
		}
		uectx.Infof("Send UeAuthenticationsAuthCtxId5gAkaConfirmationPut")
		body := &models.ConfirmationData{
			ResStar: hex.EncodeToString(resstar[:]),
		}
		var resp *models.ConfirmationDataResponse
		if resp, err = uea.UeAuthenticationsAuthCtxId5gAkaConfirmationPut(uectx.ausfcli, uectx.suci, body); err != nil {
			return
		}
		switch resp.AuthResult {
		case models.AUTHRESULT_SUCCESS:
			uectx.supi = resp.Supi
			//update uectx logger
			uectx.LogWriter = uectx.WithFields(logctx.Fields{
				"ue-supi": uectx.supi,
			})
			if uectx.createKamf(resp.Kseaf) == nil {
				uectx.Tracef("KAMF= %x", uectx.kamf)
			}

		case models.AUTHRESULT_FAILURE:
			err = fmt.Errorf("Authentication fails!")
		}
	case models.AUTHTYPE_EAP_AKA_PRIME:
		err = fmt.Errorf("Eap-Aka-Primt authentication is not supported")
		/*
			var resp *models.EapSession
			eapmsg := base64.StdEncoding.EncodeToString(msg.EAPMessage.GetEAPMessage())
			body := &models.EapSession{
				EapPayload: eapmsg,
			}
			if resp, err = uea.EapAuthMethod(ue.Ausf(), status.id, body); err != nil {
				proc.sendEvent(RejectEvent, err)
				return
			}
			switch resp.AuthResult {
			case models.AUTHRESULT_SUCCESS:
				if err = proc.createKamf(resp.Supi, resp.KSeaf); err != nil {
					proc.sendEvent(RejectEvent, err)
					return
				}
				proc.status.eap = resp.EapPayload
				proc.status.supi = resp.Supi

			case models.AUTHRESULT_FAILURE:
				proc.sendEvent(RejectEvent, err)

			case models.AUTHRESULT_ONGOING:
				proc.status.eap = resp.EapPayload
				proc.authRequest()
			}
		*/
	}
	return
}

func (uectx *UeContext) createKamf(kseaf string) (err error) {
	if err = checkSupi(uectx.supi); err != nil {
		return
	}
	abba := uectx.abba
	var kseafbyte []byte
	if kseafbyte, err = hex.DecodeString(kseaf); err != nil {
		return
	}
	uectx.kamf, err = sec.KAMF(kseafbyte, []byte(uectx.supi[5:]), abba)
	uectx.Tracef("supi=%x, abba=%x", []byte(uectx.supi[4:]), abba)
	uectx.Tracef("kseaf=%x, kamf=%x", kseafbyte, uectx.kamf)
	return
}

// NOTE: Supi should be represented in its strict format (not string), then we
// don't need this check
func checkSupi(supi string) error {
	//TODO: may need to elaborate more
	if len(supi) < 4 {
		return fmt.Errorf("Invalid SUPI")
	}
	return nil
}
