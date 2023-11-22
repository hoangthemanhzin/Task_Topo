package ranuecontext

import (
	"etrib5gc/sbi/models/n2models"
	rann2 "etrib5gc/sbi/pran/nas"
)

// UE
func (ranue *RanUe) sendInitialContextSetupRequest(pdulist []n2models.DlPduSessionResourceInfo, naspdu []byte) (err error) {
	defer ranue.logSendingReport("InitialContextSetupRequest", err)
	ue := ranue.ue
	amfctx := ranue.amf
	msg := &n2models.InitCtxSetupReq{
		PduList:  pdulist,
		NasPdu:   naspdu,
		SecKey:   ranue.RanKey(),
		UeRadCap: ue.RadCap(),
		Guami:    n2models.Guami{
			//PlmnId: ue.PlmnId(),
			//AmfId:  amfctx.AmfId(),
		},
		UeAmbr: n2models.UeAmbr{
			Ul: 1000,
			Dl: 1000,
		},
		AllowedNssai: ue.AllowedNssai(ranue.Access()),
		UeSecCap: n2models.UeSecCap{
			Nr: new(n2models.SecCap),
		},
	}

	msg.Guami.PlmnId = ue.PlmnId()
	msg.Guami.AmfId = amfctx.AmfId()
	ranue.Tracef("Plmnid=%s, Amfid=%s", msg.Guami.PlmnId.String(), msg.Guami.AmfId)
	seccap := ue.SecCap()

	msg.UeSecCap.Nr.Enc[0] |= seccap.GetEA1_128_5G() << 7
	msg.UeSecCap.Nr.Enc[0] |= seccap.GetEA2_128_5G() << 6
	msg.UeSecCap.Nr.Enc[0] |= seccap.GetEA3_128_5G() << 5

	msg.UeSecCap.Nr.Int[0] |= seccap.GetIA1_128_5G() << 7
	msg.UeSecCap.Nr.Int[0] |= seccap.GetIA2_128_5G() << 6
	msg.UeSecCap.Nr.Int[0] |= seccap.GetIA3_128_5G() << 5

	cli := ranue.rancli
	var rsp *n2models.InitCtxSetupRsp
	var ersp *n2models.InitCtxSetupFailure
	//TODO: handle response
	if rsp, ersp, err = rann2.InitCtxSetupReq(cli, ranue.ranUeId, msg); err != nil {
		return err
	} else if rsp != nil {
		err = ranue.sendEvent(CtxSetupRspEvent, rsp)
	} else {
		err = ranue.sendEvent(CtxSetupFailEvent, ersp)
	}
	return
}

//PDU

// TODO
func (ranue *RanUe) sendUeContextModiticationRequest() (err error) {
	return nil
}

// TODO
func (ranue *RanUe) sendUeContextReleaseCommand() (err error) {
	return nil
}

func (ranue *RanUe) sendPduSessionResourceModifyRequest(pdulist []n2models.DlPduSessionResourceInfo, naspdu []byte) (err error) {
	defer ranue.logSendingReport("PduSessionResourceModifyRequest", err)
	msg := &n2models.PduSessResModReq{
		SessionList: pdulist,
	}

	cli := ranue.rancli
	var rsp *n2models.PduSessResModRsp
	if rsp, err = rann2.PduSessResModReq(cli, ranue.ranUeId, msg); err != nil {
		return
	} else {
		err = ranue.sendEvent(PduModRspEvent, rsp)
	}
	return
}
func (ranue *RanUe) sendPduSessionResourceReleaseCommand(pdulist []n2models.DlPduSessionResourceInfo, naspdu []byte) (err error) {
	defer ranue.logSendingReport("PduSessionResourceReleaseCommand", err)
	msg := &n2models.PduSessResRelCmd{
		SessionList: pdulist,
	}

	cli := ranue.rancli
	var rsp *n2models.PduSessResRelRsp
	//TODO: handle response
	if rsp, err = rann2.PduSessResRelCmd(cli, ranue.ranUeId, msg); err != nil {
		return
	} else {
		err = ranue.sendEvent(PduRelRspEvent, rsp)
	}
	return
}
func (ranue *RanUe) sendPduSessionResourceSetupRequest(pdulist []n2models.DlPduSessionResourceInfo, naspdu []byte) (err error) {
	defer ranue.logSendingReport("PduSessionResourceSetupRequest", err)
	msg := &n2models.PduSessResSetReq{
		SessionList: pdulist,
		NasPdu:      naspdu,
		UeAmbr: &n2models.UeAmbr{
			Ul: 1000,
			Dl: 1000,
		},
	}

	cli := ranue.rancli
	var rsp *n2models.PduSessResSetRsp
	if rsp, err = rann2.PduSessResSetReq(cli, ranue.ranUeId, msg); err != nil {
		return
	} else {
		err = ranue.sendEvent(PduSetupRspEvent, rsp)
	}
	return
}
