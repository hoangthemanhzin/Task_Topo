package sm

import "etrib5gc/sbi/models"

func (smctx *SmContext) handleHoState(job *UpdateSmContextJob) {
	if smctx.CurrentState() != SM_ACTIVE {
		//TODO: set an error
		return
	}

	switch job.Req.JsonData.HoState {
	case models.HOSTATE_PREPARING: //active
		smctx.handleHoStatePreparing(job)
	case models.HOSTATE_PREPARED: //active
		smctx.handleHoStatePrepared(job)
	case models.HOSTATE_COMPLETED: //active
		smctx.handleHoStateCompleted(job)
	default:
		//TODO: should we return an error?
		//or just ignore?
	}

}

func (smctx *SmContext) handleHoStatePreparing(status *UpdateSmContextJob) {
	/*
		logger.PduSessLog.Traceln("In HoState_PREPARING")
		if smContext.SMContextState != smf_context.Active {
			// Wait till the state becomes Active again TODO: implement sleep
			// wait in concurrent architecture
			logger.PduSessLog.Warnf("SMContext[%s-%02d] should be Active, but actual %s",
				smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
		}
		smContext.SMContextState = smf_context.ModificationPending
		logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
		smContext.HoState = models.HoState_PREPARING
		if err := smf_context.HandleHandoverRequiredTransfer(body.BinaryDataN2SmInformation, smContext); err != nil {
			logger.PduSessLog.Errorf("Handle HandoverRequiredTransfer failed: %+v", err)
		}
		response.JsonData.N2SmInfoType = models.N2SmInfoType_PDU_RES_SETUP_REQ

		if n2Buf, err := smf_context.BuildPDUSessionResourceSetupRequestTransfer(smContext); err != nil {
			logger.PduSessLog.Errorf("Build PDUSession Resource Setup Request Transfer Error(%s)", err.Error())
		} else {
			response.BinaryDataN2SmInformation = n2Buf
			response.JsonData.N2SmInfoType = models.N2SmInfoType_PDU_RES_SETUP_REQ
			response.JsonData.N2SmInfo = &models.RefToBinaryData{
				ContentId: "PDU_RES_SETUP_REQ",
			}
		}
		response.JsonData.HoState = models.HoState_PREPARING
	*/
}

func (smctx *SmContext) handleHoStateCompleted(status *UpdateSmContextJob) {
	/*
		logger.PduSessLog.Traceln("In HoState_COMPLETED")
		if smContext.SMContextState != smf_context.Active {
			// Wait till the state becomes Active again TODO: implement sleep
			// wait in concurrent architecture
			logger.PduSessLog.Warnf("SMContext[%s-%02d] should be Active, but actual %s",
				smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
		}
		smContext.SMContextState = smf_context.ModificationPending
		logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
		smContext.HoState = models.HoState_COMPLETED
		response.JsonData.HoState = models.HoState_COMPLETED
	*/
}

func (smctx *SmContext) handleHoStatePrepared(status *UpdateSmContextJob) {
	/*
		logger.PduSessLog.Traceln("In HoState_PREPARED")
		if smContext.SMContextState != smf_context.Active {
			// Wait till the state becomes Active again TODO: implement sleep
			// wait in concurrent architecture
			logger.PduSessLog.Warnf("SMContext[%s-%02d] should be Active, but actual %s",
				smContext.Supi, smContext.PDUSessionID, smContext.SMContextState.String())
		}
		smContext.SMContextState = smf_context.ModificationPending
		logger.CtxLog.Traceln("SMContextState Change State: ", smContext.SMContextState.String())
		smContext.HoState = models.HoState_PREPARED
		response.JsonData.HoState = models.HoState_PREPARED
		if err := smf_context.HandleHandoverRequestAcknowledgeTransfer(
			body.BinaryDataN2SmInformation, smContext); err != nil {
			logger.PduSessLog.Errorf("Handle HandoverRequestAcknowledgeTransfer failed: %+v", err)
		}

		if n2Buf, err := smf_context.BuildHandoverCommandTransfer(smContext); err != nil {
			logger.PduSessLog.Errorf("Build PDUSession Resource Setup Request Transfer Error(%s)", err.Error())
		} else {
			response.BinaryDataN2SmInformation = n2Buf
			response.JsonData.N2SmInfoType = models.N2SmInfoType_HANDOVER_CMD
			response.JsonData.N2SmInfo = &models.RefToBinaryData{
				ContentId: "HANDOVER_CMD",
			}
		}
		response.JsonData.HoState = models.HoState_PREPARING
	*/
}
