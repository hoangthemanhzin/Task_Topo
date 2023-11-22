package sessioncontext

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/smf/pdu"
)

type Causes struct {
	Cause        *models.Cause
	NgapCause    *models.NgApCause
	Var5gMmCause *int32
}

func (sc *SessionContext) sendUpdate(msg *models.SmContextUpdateData, n1msg []byte, n2info []byte) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	rsp, ersp, err = pdu.UpdateSmContext(sc.smfcli, sc.ref, models.UpdateSmContextRequest{
		JsonData:                  *msg,
		BinaryDataN1SmMessage:     n1msg,
		BinaryDataN2SmInformation: n2info,
	})
	if err != nil {
		sc.Errorf("Send SmContextUpdate to SMF failed: %s", err.Error())
	} else if ersp != nil {
		sc.Errorf("SmContextUpdate not processed by SMF: %s", ersp.JsonData.Error.Detail)
	}
	return
}

func (sc *SessionContext) SendN1SmMsg(n1msg []byte) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sc.Info("Foward N1Sm to SMF")
	msg := &models.SmContextUpdateData{
		//	Procedure: models.UPDATE_SM_CONTEXT_N1N2,
		N1SmMsg: models.RefToBinaryData{
			ContentId: "N1SmMsg",
		},
		Pei: sc.ue.Pei(),
		//Gpsi:sc.ue.Gpsi(),
	}
	return sc.sendUpdate(msg, n1msg, nil)
}

func (sc *SessionContext) DuplicationRelease() (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sc.Info("Ask SMF to release duplicated session")
	//sc.MarkDuplicated(msg)
	msg := &models.SmContextUpdateData{
		//	Procedure:          models.UPDATE_SM_CONTEXT_REL_DUP,
		Release:            true,
		Cause:              models.CAUSE_REL_DUE_TO_DUPLICATE_SESSION_ID,
		SmContextStatusUri: sc.statusUri,
	}
	return sc.sendUpdate(msg, nil, nil)
}

func (sc *SessionContext) RequestSmf2CreateSession(smpdu []byte, callback models.Callback) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse, err error) {
	req := models.PostSmContextsRequest{
		BinaryDataN1SmMessage: smpdu,
	}

	req.JsonData.N1SmMsg = models.RefToBinaryData{
		ContentId: "N1SmMsg",
	}
	sc.fillSessionContextCreateData(&req.JsonData)
	rsp, ersp, err = pdu.PostSmContexts(sc.smfcli, req, callback)
	if err != nil {
		sc.Errorf("Send PostSmContexts failed: %s", err.Error())
	} else if ersp != nil {
		sc.Errorf("PostSmContexts not handled at SMF: %s", ersp.JsonData.Error.Detail)
	}
	return
}

func (sc *SessionContext) SendSmContextN2Info(n2infotype models.N2SmInfoType, n2info []byte) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sc.Info("Forward N2Info to SMF")
	msg := &models.SmContextUpdateData{
		//	Procedure:    models.UPDATE_SM_CONTEXT_N1N2,
		N2SmInfoType: n2infotype,
		N2SmInfo: models.RefToBinaryData{
			ContentId: "N2SmInfo",
		},
		UeLocation: sc.loc,
	}
	return sc.sendUpdate(msg, nil, n2info)
}

func (sc *SessionContext) ReleaseSmContext(causes Causes, n2smtype models.N2SmInfoType, n2info []byte) (rsp *models.SmContextReleasedData, err error) {
	msg := models.SmContextReleaseData{
		N2SmInfoType: n2smtype,
		N2SmInfo: models.RefToBinaryData{
			ContentId: "n2SmInfo",
		},
	}
	if causes.Cause != nil {
		msg.Cause = *causes.Cause
	}
	if causes.NgapCause != nil {
		msg.NgApCause = *causes.NgapCause
	}
	if causes.Var5gMmCause != nil {
		msg.Var5gMmCauseValue = *causes.Var5gMmCause
	}
	//TODO: time zone

	rsp, err = pdu.ReleaseSmContext(sc.smfcli, sc.ref, &models.ReleaseSmContextRequest{
		JsonData:                  msg,
		BinaryDataN2SmInformation: n2info,
	})

	if err != nil {
		sc.Errorf("Send ReleaseSmContext failed: %s", err.Error())
	}
	return
}

func (sc *SessionContext) ChangeAccess3gpp() (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sc.Info("Ask SMF to change access")
	msg := &models.SmContextUpdateData{
		//	Procedure:          models.UPDATE_SM_CONTEXT_AN_CHANGE,
		AnTypeCanBeChanged: true,
	}

	return sc.sendUpdate(msg, nil, nil)
}

func (sc *SessionContext) ActivateCnxState(access models.AccessType) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sc.Info("Ask SMF to activate CnxState")
	msg := &models.SmContextUpdateData{
		//	Procedure:  models.UPDATE_SM_CONTEXT_UPCNXSTATE,
		UpCnxState: models.UPCNXSTATE_ACTIVATING,
		//TODO: add access, location, etc
	}

	return sc.sendUpdate(msg, nil, nil)
}

func (sc *SessionContext) DeactivateCnxState() (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
	sc.Info("Ask SMF to deactivate CnxState")
	msg := &models.SmContextUpdateData{
		//	Procedure:  models.UPDATE_SM_CONTEXT_UPCNXSTATE,
		UpCnxState: models.UPCNXSTATE_DEACTIVATED,
		//TODO: add cause
	}
	return sc.sendUpdate(msg, nil, nil)
}
