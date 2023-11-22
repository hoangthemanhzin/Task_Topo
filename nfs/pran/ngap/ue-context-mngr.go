package ngap

import (
	"etrib5gc/common"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models/n2models"

	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handleInitialContextSetupResponse(ran *ran.Ran, initialContextSetupResponse *ngapType.InitialContextSetupResponse) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var oklist *ngapType.PDUSessionResourceSetupListCxtRes
	var failedlist *ngapType.PDUSessionResourceFailedToSetupListCxtRes
	var critical *ngapType.CriticalityDiagnostics

	log.Info("Receive an Initial Context Setup Response from gnB")

	for _, ie := range initialContextSetupResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Warn("UeContextNgapID is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtRes:
			oklist = ie.Value.PDUSessionResourceSetupListCxtRes
			log.Trace("Decode IE PDUSessionResourceSetupResponseList")
			if oklist == nil {
				log.Warn("PDUSessionResourceSetupResponseList is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes:
			failedlist = ie.Value.PDUSessionResourceFailedToSetupListCxtRes
			log.Trace("Decode IE PDUSessionResourceFailedToSetupList")
			if failedlist == nil {
				log.Warn("PDUSessionResourceFailedToSetupList is nil")
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE Criticality Diagnostics")
			if critical == nil {
				log.Warn("Criticality Diagnostics is nil")
			}
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Errorf("UeContext not found")
		return
	}

	msg := &n2models.InitCtxSetupRsp{
		Diag: sbiCritDiag(critical),
	}
	if oklist != nil && len(oklist.List) > 0 {
		msg.SuccessList = make([]n2models.UlPduSessionResourceInfo, len(oklist.List))
		for i, s := range oklist.List {
			msg.SuccessList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceSetupResponseTransfer,
			}
		}

	}
	if failedlist != nil && len(failedlist.List) > 0 {
		msg.FailedList = make([]n2models.UlPduSessionResourceInfo, len(failedlist.List))
		for i, s := range failedlist.List {
			msg.SuccessList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceSetupUnsuccessfulTransfer,
			}
		}

	}

	log.Tracef("UeContextNgapID[%d] AmfUeNgapID[%d]", uectx.RanNgapId(), uectx.CuNgapId())

	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_UE_SET_RSP,
		Content: msg,
	})
	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
}

func (h *Ngap) handleInitialContextSetupFailure(ran *ran.Ran, initialContextSetupFailure *ngapType.InitialContextSetupFailure) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var failedlist *ngapType.PDUSessionResourceFailedToSetupListCxtFail
	var critical *ngapType.CriticalityDiagnostics
	var cause *ngapType.Cause

	log.Info("Receive an InitialContextSetupFailure from gnB")

	for _, ie := range initialContextSetupFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Warn("UeContextNgapID is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail:
			failedlist = ie.Value.PDUSessionResourceFailedToSetupListCxtFail
			log.Trace("Decode IE PDUSessionResourceFailedToSetupList")
			if failedlist == nil {
				log.Warn("PDUSessionResourceFailedToSetupList is nil")
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Warn("Cause is nil")
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE Criticality Diagnostics")
			if critical == nil {
				log.Warn("CriticalityDiagnostics is nil")
			}
		}
	}
	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Errorf("UeContext not found")
		return
	}

	msg := &n2models.InitCtxSetupFailure{}
	if cause != nil {
		msg.Cause = causeConvert(cause)
	}

	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_UE_SET_FAIL,
		Content: msg,
	})

	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
}

// 8.3.2
func (h *Ngap) handleUEContextReleaseRequest(ran *ran.Ran, uEContextReleaseRequest *ngapType.UEContextReleaseRequest) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var oklist *ngapType.PDUSessionResourceListCxtRelReq
	var cause *ngapType.Cause

	log.Info("Receive an UE Context Release Request from gnB")

	for _, ie := range uEContextReleaseRequest.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelReq:
			oklist = ie.Value.PDUSessionResourceListCxtRelReq
			log.Trace("Decode IE Pdu Session Resource List")
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Warn("Cause is nil")
			}
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		cause = &ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		SendErrorIndication(ran, coreNgapId, ranNgapId, cause, nil)
		return
	}

	log.Tracef("UeContextNgapID[%d] AmfUeNgapID[%d]", uectx.RanNgapId(), uectx.CuNgapId())

	msg := &n2models.UeCtxRelReq{}
	if cause != nil {
		msg.Cause = causeConvert(cause)
	}
	if oklist != nil && len(oklist.List) > 0 {
		msg.SuccessList = make([]int64, len(oklist.List))
		for i, s := range oklist.List {
			msg.SuccessList[i] = int64(s.PDUSessionID.Value)
		}

	}
	uectx.UeCtxRelReq(msg)
}

func (h *Ngap) handleUEContextReleaseComplete(ran *ran.Ran, uEContextReleaseComplete *ngapType.UEContextReleaseComplete) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var locinfo *ngapType.UserLocationInformation
	var paginginfo *ngapType.InfoOnRecommendedCellsAndRANNodesForPaging
	var slist *ngapType.PDUSessionResourceListCxtRelCpl
	var critical *ngapType.CriticalityDiagnostics

	log.Info("Receiven an UE Context Release Complete from gnB")

	for _, ie := range uEContextReleaseComplete.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation:
			locinfo = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging:
			paginginfo = ie.Value.InfoOnRecommendedCellsAndRANNodesForPaging
			log.Trace("Decode IE InfoOnRecommendedCellsAndRANNodesForPaging")
			if paginginfo != nil {
				log.Warn("IE infoOnRecommendedCellsAndRANNodesForPaging is not support")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelCpl:
			slist = ie.Value.PDUSessionResourceListCxtRelCpl
			log.Trace("Decode IE PDUSessionResourceList")
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		SendErrorIndication(ran, coreNgapId, ranNgapId, &cause, nil)
		return
	}

	msg := &n2models.UeCtxRelCmpl{}
	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}
	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
	if paginginfo != nil {
		msg.RecRanNodes = recRanNodeListConvert(paginginfo.RecommendRANNodesForPaging.RecommendedRANNodeList.List)
		msg.RecCells = recCellListConvert(paginginfo.RecommendedCellsForPaging.RecommendedCellList.List)
	}
	if slist != nil && len(slist.List) > 0 {
		for i, s := range slist.List {
			msg.Sessions[i] = int64(s.PDUSessionID.Value)
		}
	}

	uectx.UeCtxRelCmpl(msg)
}

func (h *Ngap) handleUEContextModificationResponse(ran *ran.Ran, uEContextModificationResponse *ngapType.UEContextModificationResponse) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var rrcstate *ngapType.RRCState
	var locinfo *ngapType.UserLocationInformation
	var critical *ngapType.CriticalityDiagnostics

	log.Info("Receive an UE Context Modification Response from gnB")

	for _, ie := range uEContextModificationResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Warn("UeContextNgapID is nil")
			}
		case ngapType.ProtocolIEIDRRCState: // optional, ignore
			rrcstate = ie.Value.RRCState
			log.Trace("Decode IE RRCState")
		case ngapType.ProtocolIEIDUserLocationInformation: // optional, ignore
			locinfo = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		return
	}

	log.Tracef("AmfUeNgapID[%d] UeContextNgapID[%d]", uectx.CuNgapId(), uectx.RanNgapId())

	msg := &n2models.UeCtxModRsp{}

	if rrcstate != nil {
		msg.RrcState = uint16(rrcstate.Value)
		switch rrcstate.Value {
		case ngapType.RRCStatePresentInactive:
			log.Trace("UE RRC State: Inactive")
		case ngapType.RRCStatePresentConnected:
			log.Trace("UE RRC State: Connected")
		}
	}

	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}

	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_UE_MOD_RSP,
		Content: msg,
	})
}

func (h *Ngap) handleUEContextModificationFailure(ran *ran.Ran, uEContextModificationFailure *ngapType.UEContextModificationFailure) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var cause *ngapType.Cause
	var critical *ngapType.CriticalityDiagnostics

	log.Info("Receive an UE Context Modification Failure from gnB")
	for _, ie := range uEContextModificationFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Warn("UeContextNgapID is nil")
			}
		case ngapType.ProtocolIEIDCause: // ignore
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Warn("Cause is nil")
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}
	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Errorf("UeContext not found")
		return
	}
	msg := &n2models.UeCtxModFail{
		Cause: causeConvert(cause),
	}
	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}

	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_UE_MOD_FAIL,
		Content: msg,
	})
}

// 8.3.5
func (h *Ngap) handleRRCInactiveTransitionReport(ran *ran.Ran, rRCInactiveTransitionReport *ngapType.RRCInactiveTransitionReport) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var rrcstate *ngapType.RRCState
	var locinfo *ngapType.UserLocationInformation

	log.Info("Receive an RRC Inactive Transition Report fomr gnB")
	for i := 0; i < len(rRCInactiveTransitionReport.ProtocolIEs.List); i++ {
		ie := rRCInactiveTransitionReport.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRRCState: // ignore
			rrcstate = ie.Value.RRCState
			log.Trace("Decode IE RRCState")
			if rrcstate == nil {
				log.Error("RRCState is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // ignore
			locinfo = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
			if locinfo == nil {
				log.Error("UserLocationInformation is nil")
				return
			}
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Warnf("No UE Context[UeContextNgapID: %d]", ranNgapId.Value)
		return
	}
	msg := &n2models.RrcInactTranRep{}
	log.Tracef("RANUENGAPID[%d] AMFUENGAPID[%d]", uectx.RanNgapId(), uectx.CuNgapId())

	if rrcstate != nil {
		msg.RrcState = uint16(rrcstate.Value)
		switch rrcstate.Value {
		case ngapType.RRCStatePresentInactive:
			log.Trace("UE RRC State: Inactive")
		case ngapType.RRCStatePresentConnected:
			log.Trace("UE RRC State: Connected")
		}
	}
	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}

	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_UE_RRC_REP,
		Content: msg,
	})

}
