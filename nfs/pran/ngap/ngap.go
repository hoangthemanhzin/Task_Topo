package ngap

import (
	"etrib5gc/nfs/pran/context"
	"etrib5gc/nfs/pran/ran"
	"net"

	"etrib5gc/sctp"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

type Ngap struct {
	ctx     *context.CuContext
	ranpool *ran.RanPool
}

func NewNgap(ctx *context.CuContext) *Ngap {
	ret := &Ngap{
		ctx:     ctx,
		ranpool: ran.NewRanPool(ctx),
	}
	return ret
}

func (h *Ngap) HandleMessage(conn net.Conn, pdu []byte) {
	ran := h.ranpool.ByConn(conn)
	if ran == nil {
		log.Infof("Create a new NG connection for: %s", conn.RemoteAddr().String())
		//TungTQ note: should we not add a new ran to the pool. Only add
		//it after handling of a Setup request message
		ran = h.ranpool.NewRan(conn)
	}

	if len(pdu) == 0 {
		log.Infof("RAN close the connection.")
		h.ranpool.Remove(ran)
		return
	}

	ngapmsg, err := libngap.Decoder(pdu)
	if err != nil {
		log.Errorf("NGAP decode error : %+v", err)
		return
	}

	warning := false
	switch ngapmsg.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		iMsg := ngapmsg.InitiatingMessage
		if iMsg == nil {
			log.Errorln("Initiating Message is nil")
			return
		}
		switch iMsg.ProcedureCode.Value {
		case ngapType.ProcedureCodeNGSetup:
			if iMsg.Value.NGSetupRequest != nil {
				h.handleNGSetupRequest(ran, iMsg.Value.NGSetupRequest)
				return
			}
		case ngapType.ProcedureCodeInitialUEMessage:
			if iMsg.Value.InitialUEMessage != nil {
				h.handleInitialUEMessage(ran, iMsg.Value.InitialUEMessage, pdu)
			}
		case ngapType.ProcedureCodeUplinkNASTransport:
			if iMsg.Value.UplinkNASTransport != nil {
				h.handleUplinkNasTransport(ran, iMsg.Value.UplinkNASTransport)
				return
			}
		case ngapType.ProcedureCodeNGReset:
			if iMsg.Value.NGReset != nil {
				h.handleNGReset(ran, iMsg.Value.NGReset)
				return
			}
		case ngapType.ProcedureCodeHandoverCancel:
			if iMsg.Value.HandoverCancel != nil {
				h.handleHandoverCancel(ran, iMsg.Value.HandoverCancel)
				return
			}
		case ngapType.ProcedureCodeUEContextReleaseRequest:
			if iMsg.Value.UEContextReleaseRequest != nil {
				h.handleUEContextReleaseRequest(ran, iMsg.Value.UEContextReleaseRequest)
				return
			}
		case ngapType.ProcedureCodeNASNonDeliveryIndication:
			if iMsg.Value.NASNonDeliveryIndication != nil {
				h.handleNasNonDeliveryIndication(ran, iMsg.Value.NASNonDeliveryIndication)
				return
			}
		case ngapType.ProcedureCodeLocationReportingFailureIndication:
			if iMsg.Value.LocationReportingFailureIndication != nil {
				h.handleLocationReportingFailureIndication(ran, iMsg.Value.LocationReportingFailureIndication)
				return
			}

		case ngapType.ProcedureCodeErrorIndication:
			if iMsg.Value.ErrorIndication != nil {
				h.handleErrorIndication(ran, iMsg.Value.ErrorIndication)
				return
			}
		case ngapType.ProcedureCodeUERadioCapabilityInfoIndication:
			if iMsg.Value.UERadioCapabilityInfoIndication != nil {
				h.handleUERadioCapabilityInfoIndication(ran, iMsg.Value.UERadioCapabilityInfoIndication)
				return
			}
		case ngapType.ProcedureCodeHandoverNotification:
			if iMsg.Value.HandoverNotify != nil {
				h.handleHandoverNotify(ran, iMsg.Value.HandoverNotify)
				return
			}
		case ngapType.ProcedureCodeHandoverPreparation:
			if iMsg.Value.HandoverRequired != nil {
				h.handleHandoverRequired(ran, iMsg.Value.HandoverRequired)
				return
			}
		case ngapType.ProcedureCodeRANConfigurationUpdate:
			if iMsg.Value.RANConfigurationUpdate != nil {
				h.handleRanConfigurationUpdate(ran, iMsg.Value.RANConfigurationUpdate)
				return
			}
		case ngapType.ProcedureCodeRRCInactiveTransitionReport:
			if iMsg.Value.RRCInactiveTransitionReport != nil {
				h.handleRRCInactiveTransitionReport(ran, iMsg.Value.RRCInactiveTransitionReport)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceNotify:
			if iMsg.Value.PDUSessionResourceNotify != nil {
				h.handlePDUSessionResourceNotify(ran, iMsg.Value.PDUSessionResourceNotify)
				return
			}
		case ngapType.ProcedureCodePathSwitchRequest:
			if iMsg.Value.PathSwitchRequest != nil {
				h.handlePathSwitchRequest(ran, iMsg.Value.PathSwitchRequest)
				return
			}
		case ngapType.ProcedureCodeLocationReport:
			if iMsg.Value.LocationReport != nil {
				h.handleLocationReport(ran, iMsg.Value.LocationReport)
				return
			}
		case ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport:
			if iMsg.Value.UplinkUEAssociatedNRPPaTransport != nil {
				h.handleUplinkUEAssociatedNRPPATransport(ran, iMsg.Value.UplinkUEAssociatedNRPPaTransport)
				return
			}
		case ngapType.ProcedureCodeUplinkRANConfigurationTransfer:
			if iMsg.Value.UplinkRANConfigurationTransfer != nil {
				h.handleUplinkRanConfigurationTransfer(ran, iMsg.Value.UplinkRANConfigurationTransfer)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
			if iMsg.Value.PDUSessionResourceModifyIndication != nil {
				h.handlePDUSessionResourceModifyIndication(ran, iMsg.Value.PDUSessionResourceModifyIndication)
				return
			}
		case ngapType.ProcedureCodeCellTrafficTrace:
			if iMsg.Value.CellTrafficTrace != nil {
				h.handleCellTrafficTrace(ran, iMsg.Value.CellTrafficTrace)
				return
			}
		case ngapType.ProcedureCodeUplinkRANStatusTransfer:
			if iMsg.Value.UplinkRANStatusTransfer != nil {
				h.handleUplinkRanStatusTransfer(ran, iMsg.Value.UplinkRANStatusTransfer)
				return
			}
		case ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
			if iMsg.Value.UplinkNonUEAssociatedNRPPaTransport != nil {
				h.handleUplinkNonUEAssociatedNRPPATransport(ran, iMsg.Value.UplinkNonUEAssociatedNRPPaTransport)
				return
			}
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", ngapmsg.Present, iMsg.ProcedureCode.Value)
			warning = true
		}
	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		sMsg := ngapmsg.SuccessfulOutcome
		if sMsg == nil {
			log.Errorln("successful Outcome is nil")
			return
		}
		switch sMsg.ProcedureCode.Value {
		case ngapType.ProcedureCodeNGReset:
			if sMsg.Value.NGResetAcknowledge != nil {
				h.handleNGResetAcknowledge(ran, sMsg.Value.NGResetAcknowledge)
				return
			}
		case ngapType.ProcedureCodeUEContextRelease:
			if sMsg.Value.UEContextReleaseComplete != nil {
				h.handleUEContextReleaseComplete(ran, sMsg.Value.UEContextReleaseComplete)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceRelease:
			if sMsg.Value.PDUSessionResourceReleaseResponse != nil {
				h.handlePDUSessionResourceReleaseResponse(ran, sMsg.Value.PDUSessionResourceReleaseResponse)
				return
			}
		case ngapType.ProcedureCodeUERadioCapabilityCheck:
			if sMsg.Value.UERadioCapabilityCheckResponse != nil {
				h.handleUERadioCapabilityCheckResponse(ran, sMsg.Value.UERadioCapabilityCheckResponse)
				return
			}
		case ngapType.ProcedureCodeAMFConfigurationUpdate:
			if sMsg.Value.AMFConfigurationUpdateAcknowledge != nil {
				h.handleAMFconfigurationUpdateAcknowledge(ran, sMsg.Value.AMFConfigurationUpdateAcknowledge)
				return
			}
		case ngapType.ProcedureCodeInitialContextSetup:
			if sMsg.Value.InitialContextSetupResponse != nil {
				h.handleInitialContextSetupResponse(ran, sMsg.Value.InitialContextSetupResponse)
				return
			}
		case ngapType.ProcedureCodeUEContextModification:
			if sMsg.Value.UEContextModificationResponse != nil {
				h.handleUEContextModificationResponse(ran, sMsg.Value.UEContextModificationResponse)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceSetup:
			if sMsg.Value.PDUSessionResourceSetupResponse != nil {
				h.handlePDUSessionResourceSetupResponse(ran, sMsg.Value.PDUSessionResourceSetupResponse)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceModify:
			if sMsg.Value.PDUSessionResourceModifyResponse != nil {
				h.handlePDUSessionResourceModifyResponse(ran, sMsg.Value.PDUSessionResourceModifyResponse)
				return
			}
		case ngapType.ProcedureCodeHandoverResourceAllocation:
			if content := sMsg.Value.HandoverRequestAcknowledge; content != nil {
				h.handleHandoverRequestAcknowledge(ran, content)
				return
			}
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", ngapmsg.Present, sMsg.ProcedureCode.Value)
			warning = true
		}
	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		uMsg := ngapmsg.UnsuccessfulOutcome
		if uMsg == nil {
			log.Errorln("unsuccessful Outcome is nil")
			return
		}
		switch uMsg.ProcedureCode.Value {
		case ngapType.ProcedureCodeAMFConfigurationUpdate:
			if uMsg.Value.AMFConfigurationUpdateFailure != nil {
				h.handleAMFconfigurationUpdateFailure(ran, uMsg.Value.AMFConfigurationUpdateFailure)
				return
			}
		case ngapType.ProcedureCodeInitialContextSetup:
			if uMsg.Value.InitialContextSetupFailure != nil {
				h.handleInitialContextSetupFailure(ran, uMsg.Value.InitialContextSetupFailure)
				return
			}
		case ngapType.ProcedureCodeUEContextModification:
			if uMsg.Value.UEContextModificationFailure != nil {
				h.handleUEContextModificationFailure(ran, uMsg.Value.UEContextModificationFailure)
				return
			}
		case ngapType.ProcedureCodeHandoverResourceAllocation:
			if uMsg.Value.HandoverFailure != nil {
				h.handleHandoverFailure(ran, uMsg.Value.HandoverFailure)
				return
			}
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", ngapmsg.Present, uMsg.ProcedureCode.Value)
			warning = true
		}
	}
	if !warning {
		//TODO: message has no content to process, fire an error message
	}
}

func (h *Ngap) HandleSCTPNotification(conn net.Conn, notification sctp.Notification) {
	log.Infof("Handle SCTP Notification[addr: %+v]", conn.RemoteAddr())
	ran := h.ranpool.ByConn(conn)
	if ran == nil {
		log.Warnf("RAN context has been removed[addr: %+v]", conn.RemoteAddr())
		return
	}

	switch notification.Type() {
	case sctp.SCTP_ASSOC_CHANGE:
		log.Infof("SCTP_ASSOC_CHANGE notification")
		event := notification.(*sctp.SCTPAssocChangeEvent)
		switch event.State() {
		case sctp.SCTP_COMM_LOST:
			log.Infof("SCTP state is SCTP_COMM_LOST, close the connection")
			h.ranpool.Remove(ran)
		case sctp.SCTP_SHUTDOWN_COMP:
			log.Infof("SCTP state is SCTP_SHUTDOWN_COMP, close the connection")
			h.ranpool.Remove(ran)
		default:
			log.Warnf("SCTP state[%+v] is not handled", event.State())
		}
	case sctp.SCTP_SHUTDOWN_EVENT:
		log.Infof("SCTP_SHUTDOWN_EVENT notification, close the connection")
		h.ranpool.Remove(ran)
	default:
		log.Warnf("Non handled notification type: 0x%x", notification.Type())
	}
}
