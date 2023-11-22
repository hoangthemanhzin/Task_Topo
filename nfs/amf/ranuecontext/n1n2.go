package ranuecontext

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/nfs/amf/sessioncontext"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"net/http"

	"github.com/free5gc/nas/nasMessage"
)

func (ranue *RanUe) handleN1N2Transfer(job *common.AsyncJob) {
	//only handle in MM_REGISTERED, in other cases the message is queued for
	//later transfer
	defer job.Done(nil)

	jobinfo, _ := job.Info().(*events.N1N2TransferJob)
	if ranue.CurrentState() != MM_REGISTERED {
		jobinfo.Rsp = &models.N1N2MessageTransferRspData{
			Cause: models.N1N2MESSAGETRANSFERCAUSE_TEMPORARY_REJECT_REGISTRATION_ONGOING,
		}
		return
	}
	// send the N1N2Message
	ranue.Trace("Handle N1N2messageTransfer")
	req := jobinfo.Req
	reqinfo := req.JsonData
	smctx, _ := jobinfo.Extra.(*sessioncontext.SessionContext)

	var n1type uint8
	n1msg := req.BinaryDataN1Message
	n2info := req.BinaryDataN2Information

	if reqinfo.N1MessageContainer != nil {
		ranue.Info("N1N2Message has N1SM")
		n1type = getN1Type(reqinfo.N1MessageContainer.N1MessageClass)
	}

	if reqinfo.N2InfoContainer != nil {
		ranue.Info("N1N2Message has N2Info")
		switch reqinfo.N2InfoContainer.N2InformationClass {
		case models.N2INFORMATIONCLASS_SM:
			ranue.Tracef("N2Info class is SM (PDU Session ID=%d)", reqinfo.PduSessionId)
		default:
			ranue.Warnf("N2 Information type %s  not supported", reqinfo.N2InfoContainer.N2InformationClass)
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusNotImplemented,
					Detail: "Not implemented",
				},
			}
			return
		}
	}
	var naspdu []byte
	var err error
	if len(n1msg) > 0 {
		ranue.Debug("Build DLNAS for the n1 message")
		//build downlink nas message
		if naspdu, err = nas.BuildDLNASTransport(ranue, n1type, n1msg, int32(smctx.Id()), 0, nil, 0); err != nil {
			ranue.Errorf("Build DLNAS failed: %s", err.Error())
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Cause:  err.Error(),
				},
			}
			return
		}
	}

	if len(n2info) == 0 {
		if len(naspdu) > 0 {
			ranue.Debug("No N2Info, send N1SM Downlink ")
			//n1n2.rsp.Cause = models.N1N2MESSAGETRANSFERCAUSE_N1_N2_TRANSFER_INITIATED
			//send nas downlink
			err = ranue.sendNas(naspdu)
			ranue.logSendingReport("DlNas", err)
		} else {
			//corner case (bad n1n2 message)
			ranue.Warnf("Empty N1N2Message")
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusBadRequest,
					Detail: "Empty N1N2 message",
				},
			}
		}
	} else {

		ranue.Trace("N2Info not empty, send both N1SM and N2Info")
		//n2sm is not empty
		//otherwise, send downlink nas embeded in a n2 message
		//NOTE: the n1 sm message must be packed into a DLNAS message. It then can
		//be sent within either a DlPduSessionResourceInfo or an NGap message
		//(PduSessionResourceSetup/InitialContextSetupRequest etc)

		sminfo := reqinfo.N2InfoContainer.SmInfo
		switch sminfo.N2InfoContent.NgapIeType {
		case models.NGAPIETYPE_PDU_RES_SETUP_REQ:
			ranue.Infof("N2Info: PDUSessionResourceSetupRequest")
			pdulist := []n2models.DlPduSessionResourceInfo{
				n2models.DlPduSessionResourceInfo{
					Id:       int64(smctx.Id()),
					NasPdu:   nil,
					Transfer: n2info,
					Snssai:   smctx.Snssai(),
				},
			}
			if ranue.contextsent {
				//Initial Context Setup Request has been sent
				err = ranue.sendPduSessionResourceSetupRequest(pdulist, naspdu)
			} else {
				err = ranue.sendInitialContextSetupRequest(pdulist, naspdu)
			}
		case models.NGAPIETYPE_PDU_RES_MOD_REQ:
			ranue.Infof("N2Info: PDUSessionResourceModifyRequest")
			pdulist := []n2models.DlPduSessionResourceInfo{
				n2models.DlPduSessionResourceInfo{
					Id:       int64(smctx.Id()),
					NasPdu:   n1msg,
					Transfer: n2info,
					Snssai:   smctx.Snssai(),
				},
			}

			err = ranue.sendPduSessionResourceModifyRequest(pdulist, naspdu)
		case models.NGAPIETYPE_PDU_RES_REL_CMD:
			ranue.Infof("N2Info: PDUSessionResourceReleaseCommand")
			pdulist := []n2models.DlPduSessionResourceInfo{
				n2models.DlPduSessionResourceInfo{
					Id:       int64(smctx.Id()),
					NasPdu:   n1msg,
					Transfer: n2info,
					Snssai:   smctx.Snssai(),
				},
			}

			err = ranue.sendPduSessionResourceReleaseCommand(pdulist, naspdu)
		default:
			ranue.Errorf("NGAP IE Type[%s] is not supported for SmInfo", sminfo.N2InfoContent.NgapIeType)
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusForbidden,
					Detail: "UNSPECIFIED",
				},
			}
		}
	}
	if jobinfo.Ersp == nil {
		if err != nil {
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Detail: err.Error(),
				},
			}

		} else {
			jobinfo.Rsp = &models.N1N2MessageTransferRspData{
				Cause: models.N1N2MESSAGETRANSFERCAUSE_N1_N2_TRANSFER_INITIATED,
			}
		}
	}
}

func (ranue *RanUe) handleNotificationCommand(job *common.AsyncJob) {
	defer job.Done(nil)
	jobinfo, _ := job.Info().(*events.N1N2TransferJob)
	if err := ranue.sendNotification(models.ACCESSTYPE_NON_3_GPP_ACCESS); err != nil {
		jobinfo.Ersp = &models.N1N2MessageTransferError{
			ProblemDetails: models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Cause:  err.Error(),
			},
		}
	} else {
		jobinfo.Rsp = &models.N1N2MessageTransferRspData{
			Cause: models.N1N2MESSAGETRANSFERCAUSE_ATTEMPTING_TO_REACH_UE,
		}

		smctx, _ := jobinfo.Extra.(*sessioncontext.SessionContext)
		//NOTE: should use the generated identity of the queued n1n12 for
		//status notification/query
		ranue.n1n2man.add(smctx, jobinfo.Req)
	}
}

func getN1Type(n1msgclass models.N1MessageClass) (n1type uint8) {
	switch n1msgclass {
	case models.N1MESSAGECLASS_SM:
		n1type = nasMessage.PayloadContainerTypeN1SMInfo
	case models.N1MESSAGECLASS_SMS:
		n1type = nasMessage.PayloadContainerTypeSMS
	case models.N1MESSAGECLASS_LPP:
		n1type = nasMessage.PayloadContainerTypeLPP
	case models.N1MESSAGECLASS_UPDP:
		n1type = nasMessage.PayloadContainerTypeUEPolicy
	default:
		n1type = 0
	}
	return
}
