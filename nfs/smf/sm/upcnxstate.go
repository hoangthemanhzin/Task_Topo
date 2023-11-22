package sm

import (
	"etrib5gc/nfs/smf/upman/up"
	"etrib5gc/sbi/models"
	"net/http"
)

//session must be in SM_ACTIVE

func (smctx *SmContext) handleUpCnxState(job *UpdateSmContextJob) {
	switch job.Req.JsonData.UpCnxState {
	case models.UPCNXSTATE_ACTIVATING:
		smctx.handleUpCnxStateActivating(job)
	case models.UPCNXSTATE_DEACTIVATED:
		smctx.handleUpCnxStateDeactivated(job)
	default:
		smctx.Warnf("Unknown UPCNXSTATE: %s", job.Req.JsonData.UpCnxState)
	}

}

// the session is active, the SMF tell RAN to setup the resource for the session
func (smctx *SmContext) handleUpCnxStateActivating(job *UpdateSmContextJob) {
	var err error
	//smContext.SMContextState = smf_context.ModificationPending
	if job.n2info, err = smctx.buildPduSessionResourceSetupRequestTransfer(); err != nil {
		smctx.Errorf("Build PDUSession Resource Setup Request Transfer Error(%s)", err.Error())
		job.err = &models.ExtProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Type:   "N2InfoError",
		}
	} else {
		job.n2infoid = "PDUSessionResourceSetupRequestTransfer"
		job.UpCnxState = models.UPCNXSTATE_ACTIVATING
		job.n2infotype = models.N2SMINFOTYPE_PDU_RES_SETUP_REQ
	}
	// smContext.UpCnxState = models.UpCnxState_ACTIVATING
}

// In case session is in SM_INACTIVE or
// SM_INACTIVATING, skip sending pfcp modification request
// otherwise it must be in SM_ACTIVE state, then send the pfcp modification
// request
// NOTE: it looks like the UE tell SMF to ask RAN to release the resource for
// the PDU
func (smctx *SmContext) handleUpCnxStateDeactivated(job *UpdateSmContextJob) {
	job.UpCnxState = models.UPCNXSTATE_DEACTIVATED
	// If the PDU session has been released, skip sending PFCP Session
	// Modification Request
	state := smctx.CurrentState()
	if state == SM_INACTIVATING || state == SM_INACTIVE {
		smctx.Tracef("Skip sending PFCP Session Modification Request for %s", smctx.ref)
		//NOTE: it seems we need to response now
		return
	}

	//smctx.UpCnxState = status.req.JsonData.UpCnxState
	//smctx.UeLocation = status.req.JsonData.UeLocation
	// TODO: Deactivate N2 downlink tunnel Set FAR and An, N3 Release Info

	dlpdr := smctx.tunnel.AnNode().DlPdr()
	dlpdr.FAR.State = up.RULE_UPDATE
	dlpdr.FAR.ApplyAction.Forw = false
	dlpdr.FAR.ApplyAction.Buff = true
	dlpdr.FAR.ApplyAction.Nocp = true
	job.rulechange = true
	//smContext.PendingUPF[ANUPF.GetNodeIP()] = true
}
