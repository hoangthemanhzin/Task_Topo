package authentication

import (
	"encoding/hex"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

// send nas authentication request
func (proc *AuthProc) authRequest() {
	proc.t3560.Start()
	proc.ranue.SendAuthenticationRequest(proc.fillAuthenticationRequest)
}

func (proc *AuthProc) fillAuthenticationRequest(req *nasMessage.AuthenticationRequest) (err error) {
	uectx := proc.ranue.UeContext()
	req.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(*uectx.NgKsi())
	abba := uectx.Abba()
	req.ABBA.SetLen(uint8(len(abba)))
	req.ABBA.SetABBAContents(abba)

	switch proc.authtype {
	case models.AUTHTYPE__5_G_AKA:
		var tmp [16]byte
		copy(tmp[:], proc.rand)
		req.AuthenticationParameterRAND =
			nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
		req.AuthenticationParameterRAND.SetRANDValue(tmp)

		copy(tmp[:], proc.autn)
		req.AuthenticationParameterAUTN =
			nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
		req.AuthenticationParameterAUTN.SetLen(uint8(len(proc.autn)))
		req.AuthenticationParameterAUTN.SetAUTN(tmp)
		proc.Tracef("Set Auth: rand=%s,autn=%s", hex.EncodeToString(proc.rand), hex.EncodeToString(proc.autn))
	case models.AUTHTYPE_EAP_AKA_PRIME:
		/*
			eapMsg := ue.AuthenticationCtx.Var5gAuthData.(string)
			rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
			if err != nil {
				return nil, err
			}
			req.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRequestEAPMessageType)
			req.APMessage.SetLen(uint16(len(rawEapMsg)))

			req.EAPMessage.SetEAPMessage(rawEapMsg)
		*/
	}
	return
}
