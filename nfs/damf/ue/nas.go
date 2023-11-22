package ue

import (
	"encoding/hex"
	"etrib5gc/nas"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	prannas "etrib5gc/sbi/pran/nas"
	"etrib5gc/sbi/utils/nasConvert"
	"fmt"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

func (uectx *UeContext) sendAuthenticationRequest() (err error) {
	uectx.Infof("Send Nas AuthenticationRequest")
	var pdu []byte
	if pdu, err = nas.BuildAuthenticationRequest(nil, uectx.fillAuthReq); err == nil {
		err = prannas.NasDl(uectx.rancli, uectx.ranueid, n2models.NasDlMsg{
			NasPdu: pdu,
		})
	}
	return
}

func (uectx *UeContext) fillAuthReq(req *nasMessage.AuthenticationRequest) (err error) {
	req.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(uectx.ngksi)
	abba := uectx.abba
	req.ABBA.SetLen(uint8(len(abba)))
	req.ABBA.SetABBAContents(abba)

	switch uectx.authtype {
	case models.AUTHTYPE__5_G_AKA:
		var tmp [16]byte
		copy(tmp[:], uectx.rand)
		req.AuthenticationParameterRAND =
			nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
		req.AuthenticationParameterRAND.SetRANDValue(tmp)

		copy(tmp[:], uectx.autn)
		req.AuthenticationParameterAUTN =
			nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
		req.AuthenticationParameterAUTN.SetLen(uint8(len(uectx.autn)))
		req.AuthenticationParameterAUTN.SetAUTN(tmp)
		uectx.Tracef("Set Auth: rand=%s,autn=%s", hex.EncodeToString(uectx.rand), hex.EncodeToString(uectx.autn))
	case models.AUTHTYPE_EAP_AKA_PRIME:
		err = fmt.Errorf("Unsupport authentication method")

	}
	return
}

func (uectx *UeContext) fillRegistrationReject(msg *nasMessage.RegistrationReject) (err error) {
	msg.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationRejectT3502ValueType)
	msg.T3502Value.SetLen(1)
	msg.T3502Value.SetGPRSTimer2Value(uectx.ctx.GetT3502())
	return
}

func (uectx *UeContext) report(err error) {
	msg := n2models.InitUeContextStatus{}
	if err == nil {
		//authenticated, AMF found
		msg.Success = true
		msg.AmfId = uectx.amfid
		msg.UeCtx = &n2models.AuthUeCtx{
			Supi:     uectx.supi,
			Kamf:     uectx.kamf,
			Rand:     uectx.rand,
			AuthType: uectx.authtype,
			PlmnId:   uectx.plmnid,
		}
	} else {
		cause := nasMessage.Cause5GMMProtocolErrorUnspecified
		//either not authenticated or AMF not found
		msg.Success = false
		msg.Error = err.Error()
		if msg.NasPdu, err = nas.BuildRegistrationReject(nil, cause, uectx.eap, uectx.fillRegistrationReject); err != nil {
			uectx.Errorf("Build RegistrationReject failed: %s", err.Error())
		}
	}
	//send the status
	uectx.Infof("Send InitUeContextStatus to PRAN")
	if err = prannas.InitUeContextStatus(uectx.rancli, uectx.ranueid, msg); err != nil {
		uectx.Errorf("Send InitiUeContextStatus failed: %s", err.Error())
	}
	//remove the UeContext anyway
	uectx.ctx.RemoveUe(uectx)
}
