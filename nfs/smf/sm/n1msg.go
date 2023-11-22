package sm

import (
	"etrib5gc/nas"
	"etrib5gc/sbi/models"
	"net/http"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

func (smctx *SmContext) handleN1Msg(job *UpdateSmContextJob) {
	if len(job.Req.BinaryDataN1SmMessage) == 0 {
		smctx.Warnf("N1Sm message empty")
		return
	}

	n1msg := libnas.NewMessage()
	if err := n1msg.GsmMessageDecode(&job.Req.BinaryDataN1SmMessage); err != nil {
		job.err = &models.ExtProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
			Cause:  "N1Msg Decoding",
		}
		return
	}

	switch n1msg.GsmHeader.GetMessageType() {
	case libnas.MsgTypePDUSessionReleaseRequest:
		smctx.handlePduSessionReleaseRequest(n1msg.PDUSessionReleaseRequest, job)
	case libnas.MsgTypePDUSessionReleaseComplete:
		smctx.handlePduSessionReleaseComplete(n1msg.PDUSessionReleaseComplete, job)
	default:
	}
}

func (smctx *SmContext) handlePduSessionEstablishmentRequest(req *nasMessage.PDUSessionEstablishmentRequest) (err error, n1cause uint8) {

	// Retrieve PDUSessionID
	smctx.sid = uint32(req.PDUSessionID.GetPDUSessionID())

	smctx.Infof("Handle PDUSessionEstablishmentRequest")

	// Retrieve PTI (Procedure transaction identity)
	smctx.pti = req.GetPTI()

	// Retrieve MaxIntegrityProtectedDataRate of UE for
	// UP Security
	switch req.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink() {
	case 0x00:
		smctx.maxul = models.MAXINTEGRITYPROTECTEDDATARATE__64_KBPS
	case 0xff:
		smctx.maxul = models.MAXINTEGRITYPROTECTEDDATARATE_MAX_UE_RATE
	}

	switch req.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink() {
	case 0x00:
		smctx.maxdl = models.MAXINTEGRITYPROTECTEDDATARATE__64_KBPS
	case 0xff:
		smctx.maxdl = models.MAXINTEGRITYPROTECTEDDATARATE_MAX_UE_RATE
	}

	// Handle PDUSessionType
	if req.PDUSessionType != nil {
		rtype := req.PDUSessionType.GetPDUSessionTypeValue()
		if err = smctx.checkPduSessionType(rtype); err != nil {
			smctx.Errorf("Check PduSessionType return error: %s", err.Error())
			//TODO: set a right cause
			n1cause = nasMessage.Cause5GSMProtocolErrorUnspecified
			return
		}
	} else {
		// Set to default supported PDU Session Type
		smctx.sessiontype = smctx.ctx.DefaultPduSessionType()
		smctx.Infof("Set default session type to %v", smctx.sessiontype)
	}

	if req.ExtendedProtocolConfigurationOptions != nil {
		smctx.Trace("Decode extended PCO")
		content := req.ExtendedProtocolConfigurationOptions.GetExtendedProtocolConfigurationOptionsContents()
		smctx.pco = pcoFromNas(content)
	}
	return
}

// NOTE: session is in SM_ACTIVE or SM_INACTIVATING states
func (smctx *SmContext) handlePduSessionReleaseRequest(req *nasMessage.PDUSessionReleaseRequest, job *UpdateSmContextJob) {
	state := smctx.CurrentState()
	if state != SM_ACTIVE && state != SM_INACTIVATING {
		//TODO: any other action?
		return
	}
	smctx.Info("Handle PduSessionReleaseRequest")

	smctx.pti = req.GetPTI()

	//TODO release UEIP
	// remove SM Policy Association
	//smctx.sendPolAssTermination()

	cause := nasMessage.Cause5GSMRegularDeactivation
	if req.Cause5GSM != nil {
		cause = req.Cause5GSM.GetCauseValue()
	}
	var err error
	if job.n1msgbyte, err = nas.BuildPduSessionReleaseCommand(smctx, cause); err != nil {
		smctx.Errorf("Build PDUSessionReleaseCommand failed: %+v", err)
		job.err = &models.ExtProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Type:   "N1MsgError",
		}

	} else {
		job.n1msgid = "PDUSessionReleaseCommand"
	}

	if job.n2info, err = smctx.buildPduSessionResourceReleaseCommandTransfer(); err != nil {
		smctx.Errorf("Build PDUSessionResourceReleaseCommandTransfer failed: %+v", err)
		job.err = &models.ExtProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Type:   "N2InfoError",
		}
	} else {
		job.n2infotype = models.N2SMINFOTYPE_PDU_RES_REL_CMD
		job.n2infoid = "PDUResourceReleaseCommand"
	}

	// In the case of dupulicated PDUSessionReleaseRequest, skip
	// deleting PFCP sessions
	if smctx.CurrentState() == SM_INACTIVATING {
		smctx.Tracef("Skip deleting the PFCP session")
		return
	}

	smctx.tunnel.Release()
}

// NOTE: session should be in SM_INACTIVATING
func (smctx *SmContext) handlePduSessionReleaseComplete(req *nasMessage.PDUSessionReleaseComplete, job *UpdateSmContextJob) {
	state := smctx.CurrentState()
	if state != SM_INACTIVATING {
		//TODO: any other action?
		return
	}

	//req := status.n1msg.PDUSessionReleaseComplete
	smctx.Infof("Handle PduSessionReleaseComplete")
	// Send Release Notify to AMF
	smctx.notifyAmf()

	//NOTE: terminate PolAsso at PCF
	job.UpCnxState = models.UPCNXSTATE_DEACTIVATED

	//send ReleasedEvent to move to SM_INACTIVE state
	smctx.sendEvent(SessDeactCmplEvent, nil)
}
