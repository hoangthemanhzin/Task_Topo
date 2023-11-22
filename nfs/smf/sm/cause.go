package sm

import (
	"etrib5gc/sbi/models"
)

func (smctx *SmContext) handleCause(job *UpdateSmContextJob) {
	switch job.Req.JsonData.Cause {
	case models.CAUSE_REL_DUE_TO_DUPLICATE_SESSION_ID:
		smctx.handleCauseRelDue2DupSessId(job)
	default:
	}
}

// NOTE: session must be in SM_INACTIVE, SM_ACTIVE or SM_INACTIVATING
// only send Pfcp modification request if the session is in SM_ACTIVE state
func (smctx *SmContext) handleCauseRelDue2DupSessId(job *UpdateSmContextJob) {
	var err error
	if job.n2info, err = smctx.buildPduSessionResourceReleaseCommandTransfer(); err != nil {
		smctx.Error(err.Error())
		//TODO: create error here
	} else {
		job.n2infoid = "PDUResourceReleaseCommand"
		job.n2infotype = models.N2SMINFOTYPE_PDU_RES_REL_CMD
	}

	state := smctx.CurrentState()
	switch state {
	case SM_ACTIVE:
		//need to release PFCP session
		//	smContext.PDUSessionRelease_DUE_TO_DUP_PDU_ID = true
		smctx.tunnel.Release()
	case SM_ACTIVATING:
		//need to send PduSessionEstablishmentReject
		//NOTE: free5gc: if session is not in SM_ACTIVE, then send a Nas
		//PduSessionEstablishmentReject in SmContextUpdateError
		//BuildGSMPDUSessionEstablishmentReject(nasMessage.Cause5GSMRequestRejectedUnspecified)
		//		job.status = http.StatusForbidden
	default:
		//
	}

}
