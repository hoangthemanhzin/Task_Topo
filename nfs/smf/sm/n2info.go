package sm

import (
	"encoding/binary"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
)

func (smctx *SmContext) handleN2Info(job *UpdateSmContextJob) {
	smctx.Infof("Handle N2SmInfo")
	if len(job.Req.BinaryDataN2SmInformation) == 0 {
		smctx.Warnf("N2SmInfo empty")
		return
	}
	switch job.Req.JsonData.N2SmInfoType {
	case models.N2SMINFOTYPE_PDU_RES_SETUP_RSP:
		smctx.handlePduResSetRsp(job)
	case models.N2SMINFOTYPE_PDU_RES_REL_RSP: //inactive/inactivating (dup rel) or inactivating (normal)
		//TODO: isolate the dupplication case
		smctx.handlePduResRelRsp(job)
	case models.N2SMINFOTYPE_PDU_RES_SETUP_FAIL: //all cases
		smctx.handlePduResSetFail(job)
	case models.N2SMINFOTYPE_PATH_SWITCH_REQ: //active
		smctx.handlePathSwReq(job)
	case models.N2SMINFOTYPE_PATH_SWITCH_SETUP_FAIL: //active
		smctx.handlePathSwSetFail(job)
	case models.N2SMINFOTYPE_HANDOVER_REQUIRED: //active
		smctx.handleHandoverRequired(job)
	default:
		smctx.Infof("Unknown N2SmInfoType=%s", job.Req.JsonData.N2SmInfoType)
	}

}

// all states
func (smctx *SmContext) handlePduResSetFail(job *UpdateSmContextJob) {
	smctx.Infof("Handle N2 PduSessionResourceSetupFail")
}

// SM_INACTIVE or SM_INACTIVATING
func (smctx *SmContext) handlePduResRelRsp(job *UpdateSmContextJob) {
	smctx.Infof("Handle PduSessionResourceReleaseResponse")
	/*
		logger.PduSessLog.Infoln("[SMF] N2 PDUSession Release Complete ")
					if smContext.PDUSessionRelease_DUE_TO_DUP_PDU_ID {
						state := smContext.SMContextState
						if !(state == smf_context.InActivePending || state == smf_context.InActive) {
							logger.PduSessLog.Warnf("SMContext[%s-%02d] should be InActivePending, but actual %s",
								smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
							return &httpwrapper.Response{
								Status: http.StatusForbidden,
								Body: models.UpdateSmContextErrorResponse{
									JsonData: &models.SmContextUpdateError{
										Error: &Nsmf_PDUSession.N2SmError,
									},
								},
							}
						}
						smContext.SMContextState = smf_context.InActive
						logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
						logger.PduSessLog.Infoln("[SMF] Send Update SmContext Response")
						response.JsonData.UpCnxState = models.UpCnxState_DEACTIVATED

						smContext.PDUSessionRelease_DUE_TO_DUP_PDU_ID = false
						smf_context.RemoveSMContext(smContext.Ref)
						problemDetails, err := consumer.SendSMContextStatusNotification(smContext.SmStatusNotifyUri)
						if problemDetails != nil || err != nil {
							if problemDetails != nil {
								logger.PduSessLog.Warnf("Send SMContext Status Notification Problem[%+v]", problemDetails)
							}

							if err != nil {
								logger.PduSessLog.Warnf("Send SMContext Status Notification Error[%v]", err)
							}
						} else {
							logger.PduSessLog.Traceln("Send SMContext Status Notification successfully")
						}
					} else { // normal case
						if smContext.SMContextState != smf_context.InActivePending {
							logger.PduSessLog.Warnf("SMContext[%s-%02d] should be InActivePending, but actual %s",
								smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
							return &httpwrapper.Response{
								Status: http.StatusForbidden,
								Body: models.UpdateSmContextErrorResponse{
									JsonData: &models.SmContextUpdateError{
										Error: &Nsmf_PDUSession.N2SmError,
									},
								},
							}
						}
						logger.PduSessLog.Infoln("[SMF] Send Update SmContext Response")
					}

	*/

}

// SM_ACTIVE
func (smctx *SmContext) handlePduResSetRsp(job *UpdateSmContextJob) {
	smctx.Infof("Handle PduSessionResourceSetupResponse")
	msg := ngapType.PDUSessionResourceSetupResponseTransfer{}
	var err error
	if err = aper.UnmarshalWithParams(job.Req.BinaryDataN2SmInformation, &msg, "valueExt"); err != nil {
		//TODO: uncomment
		//job.status = http.StatusBadRequest
		//status.errcause = "N2 Decoding error"
		smctx.Errorf("Decode PduSessionResourceSetupResponse failed: %s", err.Error())
		return
	}

	qosflow := msg.DLQosFlowPerTNLInformation

	if qosflow.UPTransportLayerInformation.Present !=
		ngapType.UPTransportLayerInformationPresentGTPTunnel {
		//TODO: uncomment
		job.err = &models.ExtProblemDetails{
			Status: http.StatusBadRequest,
			Detail: fmt.Sprintf("No UPTransportLayerInformationPresentGTPTunnel"),
			Cause:  "Invalid N2Sm message",
		}
		return
	}

	GTPTunnel := qosflow.UPTransportLayerInformation.GTPTunnel
	//update raninfo for the tunnel
	smctx.tunnel.UpdateRanInfo(GTPTunnel.TransportLayerAddress.Value.Bytes, binary.BigEndian.Uint32(GTPTunnel.GTPTEID.Value))
	//send pfcp session modification to UPFs
	smctx.tunnel.Update()
}

// SM_ACTIVE
func (smctx *SmContext) handlePathSwReq(job *UpdateSmContextJob) {
	smctx.Infof("Handle PathSwitchRequest")
	/*
		logger.PduSessLog.Traceln("Handle Path Switch Request")
		if smContext.SMContextState != smf_context.Active {
			// Wait till the state becomes Active again TODO: implement sleep
			// wait in concurrent architecture
			logger.PduSessLog.Warnf("SMContext[%s-%02d] should be Active, but actual %s",
				smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
		}
		smContext.SMContextState = smf_context.ModificationPending
		logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())

		if err := smf_context.HandlePathSwitchRequestTransfer(body.BinaryDataN2SmInformation, smContext); err != nil {
			logger.PduSessLog.Errorf("Handle PathSwitchRequestTransfer: %+v", err)
		}

		if n2Buf, err := smf_context.BuildPathSwitchRequestAcknowledgeTransfer(smContext); err != nil {
			logger.PduSessLog.Errorf("Build Path Switch Transfer Error(%+v)", err)
		} else {
			response.JsonData.N2SmInfoType = models.N2SmInfoType_PATH_SWITCH_REQ_ACK
			response.BinaryDataN2SmInformation = n2Buf
			response.JsonData.N2SmInfo = &models.RefToBinaryData{
				ContentId: "PATH_SWITCH_REQ_ACK",
			}
		}

		smContext.PendingUPF = make(smf_context.PendingUPF)
		for _, dataPath := range tunnel.DataPathPool {
			if dataPath.Activated {
				ANUPF := dataPath.FirstDPNode
				DLPDR := ANUPF.DownLinkTunnel.PDR

				pdrList = append(pdrList, DLPDR)
				farList = append(farList, DLPDR.FAR)

				if _, exist := smContext.PendingUPF[ANUPF.GetNodeIP()]; !exist {
					smContext.PendingUPF[ANUPF.GetNodeIP()] = true
				}
			}
		}

		sendPFCPModification = true
		smContext.SMContextState = smf_context.PFCPModification
		logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
	*/
}

// SM_ACTIVE
func (smctx *SmContext) handlePathSwSetFail(job *UpdateSmContextJob) {
	smctx.Infof("Handle PathSwitchSetFail")
	/*
		if smContext.SMContextState != smf_context.Active {
						// Wait till the state becomes Active again TODO: implement sleep
						// wait in concurrent architecture
						logger.PduSessLog.Warnf("SMContext[%s-%02d] should be Active, but actual %s",
							smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
					}
					smContext.SMContextState = smf_context.ModificationPending
					logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
					if err := smf_context.HandlePathSwitchRequestSetupFailedTransfer(
						body.BinaryDataN2SmInformation, smContext); err != nil {
						logger.PduSessLog.Error()
					}
	*/
}
func (smctx *SmContext) handleHandoverRequired(job *UpdateSmContextJob) {
	smctx.Infof("Handle HandoverRequired")
	/*
		if smContext.SMContextState != smf_context.Active {
						// Wait till the state becomes Active again TODO: implement sleep
						// wait in concurrent architecture
						logger.PduSessLog.Warnf("SMContext[%s-%02d] should be Active, but actual %s",
							smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
					}
					smContext.SMContextState = smf_context.ModificationPending
					logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
					response.JsonData.N2SmInfo = &models.RefToBinaryData{ContentId: "Handover"}
				}

	*/
}
