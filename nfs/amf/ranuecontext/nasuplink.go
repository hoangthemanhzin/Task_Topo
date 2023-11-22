package ranuecontext

import (
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/sessioncontext"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"

	"github.com/free5gc/nas/nasMessage"
)

// NOTE: ignore any uplink nas messages which are supported now
func (ranue *RanUe) handleUlNasTransport(msg *nasMessage.ULNASTransport) {
	ranue.Info("Handle UplinkNasTransport")
	switch msg.GetPayloadContainerType() {
	// TS 24.501 5.4.5.2.3 case a)
	case nasMessage.PayloadContainerTypeN1SMInfo:
		//only handle this case
		ranue.forwardSmMsg(msg)
	case nasMessage.PayloadContainerTypeSMS:
		ranue.Error("PayloadContainerTypeSMS not supported")
	case nasMessage.PayloadContainerTypeLPP:
		ranue.Error("PayloadContainerTypeLPP not supported")
	case nasMessage.PayloadContainerTypeSOR:
		ranue.Error("PayloadContainerTypeSOR not supported")
	case nasMessage.PayloadContainerTypeUEPolicy:
		ranue.Warn("AMF Transfering UEPolicy To PCF not implemented")
		//TODO: to be implemented
		//callback.SendN1MessageNotify(ue, models.N1MessageClass_UPDP,
		//	msg.PayloadContainer.GetPayloadContainerContents(), nil)
	case nasMessage.PayloadContainerTypeUEParameterUpdate:
		ranue.Warn("AMF Transfering UEParameterUpdate To UDM not implemented")
		//TODO: to be implemented
		/*
			if upuMac, err := nasConvert.UpuAckToModels(ulNasTransport.PayloadContainer.GetPayloadContainerContents()); err == nil {
				err = consumer.PutUpuAck(ue, upuMac)
			}
			if err != nil {
				return err
			}
		*/
	case nasMessage.PayloadContainerTypeMultiplePayload:
		//TODO: send DL message to tell the cause (Pdu session Id is nil for an
		//UL N1SM message
		ranue.Error("PayloadContainerTypeMultiplePayload not supported")
	}
}

// Look for the SessionContext or create a new one then forward the Nas message
func (ranue *RanUe) forwardSmMsg(msg *nasMessage.ULNASTransport) {
	n1smpdu := msg.PayloadContainer.GetPayloadContainerContents()

	var sid int32 //pdu session identity

	if id := msg.PduSessionID2Value; id != nil {
		sid = int32(id.GetPduSessionID2Value())
	} else {
		ranue.Errorf("PDU Session ID is nil")
		//ignore
		return
	}

	if msg.OldPDUSessionID != nil {
		ranue.Errorf("SSC mode 3 is not supported")
		//ignore
		return
	}

	//TS24.501: 5.4.5.2.2 The request type is not provided along 5GSM messages other than the PDU SESSION ESTABLISHMENT REQUEST
	rtype := msg.RequestType
	if rtype != nil {
		switch rtype.GetRequestTypeValue() {
		case nasMessage.ULNASTransportRequestTypeInitialEmergencyRequest:
			fallthrough
		case nasMessage.ULNASTransportRequestTypeExistingEmergencyPduSession:
			ranue.Errorf("Emergency PDU Session not supported")
			ranue.sendDlN1SmError(sid, n1smpdu)
			return
		}
	}
	sc := findSession(ranue.ue, sid)

	//has the session context
	if sc != nil {
		//sc := scref.(*sessioncontext.SessionContext)
		ranue.Infof("SessionContext found for id=%d", sc.Id())
		// no request type
		if rtype == nil {
			ranue.Warnf("RequestType empty")
			ranue.doForwardSmMsg(sc, n1smpdu)
			return
		}

		switch rtype.GetRequestTypeValue() {
		//a duplicated pdu session (for initial pdu session request)
		//need to release it
		case nasMessage.ULNASTransportRequestTypeInitialRequest:
			ranue.Info("RequestType = INITIAL")
			//NOTE: should handle the multiple-duplicating case
			if rsp, ersp, err := sc.DuplicationRelease(); ersp != nil || err != nil {
				if ersp != nil || ersp.BinaryDataN1SmMessage != nil {
					//send N1Sm downlink
					ranue.logSendingReport("Downlink N1Sm", ranue.sendDlN1Sm(sid, ersp.BinaryDataN1SmMessage, 0))
				} else {
					//N1Sm was not handled at SMF, send an error
					ranue.sendDlN1SmError(sid, n1smpdu)
				}
			} else { //rsp != nil
				n2info := rsp.BinaryDataN2SmInformation
				if n2info != nil {
					switch rsp.JsonData.N2SmInfoType {
					case models.N2SMINFOTYPE_PDU_RES_REL_CMD:
						ranue.Debug("AMF Transfer NGAP PDU Session Resource Release Command from SMF")
						pdulist := []n2models.DlPduSessionResourceInfo{
							n2models.DlPduSessionResourceInfo{
								Id:       int64(sid),
								Transfer: n2info,
							},
						}
						ranue.sendPduSessionResourceReleaseCommand(pdulist, nil)
					default:
						//TODO: other cases
						ranue.Warnf("Unknown N2SmInfo: %s", rsp.JsonData.N2SmInfoType)
					}
				}
			}

		//existing pdu session
		case nasMessage.ULNASTransportRequestTypeExistingPduSession:
			ranue.Info("RequestType = EXISTING")
			if ranue.ue.IsSnssaiAllowed(sc.Snssai(), ranue.access) {
				ranue.doForwardSmMsg(sc, n1smpdu)
			} else {
				ranue.Errorf("Session slice is not allowed")
				ranue.sendDlN1SmError(sid, n1smpdu)
			}

		// other types, just forward
		default:
			ranue.Info("RequestType = OTHERS")
			ranue.doForwardSmMsg(sc, n1smpdu)
			return
		}
	} else { // SessionContext does not exist
		ranue.Infof("SessionContext not found for id=%d", sid)
		if rtype == nil {
			ranue.Errorf("RequestType is nil")
			ranue.sendDlN1SmError(sid, n1smpdu)
			return
		}
		switch rtype.GetRequestTypeValue() {
		//initial request, find SMF and create an SessionContext
		case nasMessage.ULNASTransportRequestTypeInitialRequest:
			if sc, err := sessioncontext.CreateSessionContext(ranue, msg, false); err == nil {
				//SessionContextRequest
				//n1smpdu will be forwarded when sending this request
				if rsp, ersp, err := sc.RequestSmf2CreateSession(n1smpdu, ranue.amf.Callback()); err != nil || ersp != nil {
					if ersp != nil {
						ranue.logSendingReport("Downlink N1Sm", ranue.sendDlN1Sm(sid, ersp.BinaryDataN1SmMessage, 0))
						//NOTE: handle N2Info, is there any?
					} else {
						ranue.sendDlN1SmError(sid, n1smpdu)
					}
				} else {
					//update SessionContext
					scref := rsp.JsonData.SmContextRef
					ranue.Infof("SessionContext created at SMF: %s", scref)
					sc.SetRef(scref)
					ranue.ue.StoreSessionContext(sc)
					/*
						//TODO: handle N2Info
						n2info := rsp.BinaryDataN2SmInformation
						if n2info != nil {
							switch rsp.JsonData.N2SmInfoType {
							case models.N2SMINFOTYPE_PDU_RES_REL_CMD:
							}
						}
					*/
				}
			} else { //fail to create a SessionContext
				ranue.Errorf("Create SessionContext [sid=%d] failed: %s", sid, err.Error())
				ranue.sendDlN1SmError(sid, n1smpdu)
			}

		//existing sm context has no presence on AMF, try to create it
		case nasMessage.ULNASTransportRequestTypeModificationRequest:
			fallthrough

		case nasMessage.ULNASTransportRequestTypeExistingPduSession:
			//sc will be created with information from UDM
			if sc, err := sessioncontext.CreateSessionContext(ranue, msg, true); err == nil {
				// TS 24.501 5.4.5.2.3 case a) 1) iv)
				ranue.ue.StoreSessionContext(sc)
				ranue.doForwardSmMsg(sc, n1smpdu)
			} else {
				ranue.Errorf("Create SessionContext [sid=%d] failed: %s", sid, err.Error())
				ranue.sendDlN1SmError(sid, n1smpdu)
			}
		default:
			ranue.Warnf("Unknown RequestType %d", rtype.GetRequestTypeValue())
		}
	}
	return
}

// Forward a N1Sm message toward its SMF
func (ranue *RanUe) doForwardSmMsg(sc *sessioncontext.SessionContext, n1smpdu []byte) {
	sid := int32(sc.Id())

	if rsp, ersp, err := sc.SendN1SmMsg(n1smpdu); err != nil {
		//Fail to forward N1Sm to SMF
		ranue.sendDlN1SmError(sid, n1smpdu)
	} else if ersp != nil {
		if n1msg := ersp.BinaryDataN1SmMessage; len(n1msg) > 0 {
			ranue.logSendingReport("Downlink N1Sm", ranue.sendDlN1Sm(sid, n1msg, 0))
			//TODO: handle N2Info?
		} else {
			ranue.Errorf("Receive empty N1SM")
			ranue.sendDlN1SmError(sid, n1smpdu)
		}
	} else {
		//update access type and location for the sc
		sc.Update(ranue.Access())
		var dlnaspdu []byte
		n2sminfo := rsp.BinaryDataN2SmInformation

		if rsp.BinaryDataN1SmMessage != nil {
			ranue.Info("Receive N1Sm from SMF")
			if dlnaspdu, err = nas.BuildDLNASTransport(ranue, nasMessage.PayloadContainerTypeN1SMInfo, rsp.BinaryDataN1SmMessage, sid, 0, nil, 0); err != nil {
				ranue.Error("Build DLNasTransport from N1Sm failed: ", err.Error())
				return
			}
		}

		if n2sminfo != nil {
			pdulist := []n2models.DlPduSessionResourceInfo{
				n2models.DlPduSessionResourceInfo{
					Id:       int64(sid),
					Snssai:   sc.Snssai(),
					NasPdu:   dlnaspdu,
					Transfer: n2sminfo,
				},
			}

			n2infotype := rsp.JsonData.N2SmInfoType
			ranue.Infof("Receive N2SmInfo [%s] from SMF", n2infotype)
			switch n2infotype {
			case models.N2SMINFOTYPE_PDU_RES_MOD_REQ:
				ranue.sendPduSessionResourceModifyRequest(pdulist, dlnaspdu)
			case models.N2SMINFOTYPE_PDU_RES_REL_CMD:
				ranue.sendPduSessionResourceReleaseCommand(pdulist, dlnaspdu)
			default:
				ranue.Errorf("N2SmInfo [%s] not supported", n2infotype)
				ranue.sendDlN1SmError(sid, n1smpdu)
			}
		} else {
			ranue.logSendingReport("DlNasTransport", ranue.sendNas(dlnaspdu))
		}
	}

}

func findSession(ue *uecontext.UeContext, sid int32) (sc *sessioncontext.SessionContext) {
	if scref := ue.FindSessionContext(sid); scref != nil {
		sc, _ = scref.(*sessioncontext.SessionContext)
	}
	return
}
