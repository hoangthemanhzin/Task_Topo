package ngap

import (
	"etrib5gc/common"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models/n2models"

	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handlePDUSessionResourceSetupResponse(ran *ran.Ran, rsp *ngapType.PDUSessionResourceSetupResponse) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var critical *ngapType.CriticalityDiagnostics
	var oklist *ngapType.PDUSessionResourceSetupListSURes
	var failedlist *ngapType.PDUSessionResourceFailedToSetupListSURes

	for _, ie := range rsp.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListSURes: // ignore
			oklist = ie.Value.PDUSessionResourceSetupListSURes
			log.Trace("Decode IE PDUSessionResourceSetupListSURes")
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListSURes: // ignore
			failedlist = ie.Value.PDUSessionResourceFailedToSetupListSURes
			log.Trace("Decode IE PDUSessionResourceFailedToSetupListSURes")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	var uectx *ue.UeContext
	uectx = ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Errorf("UeContext not found")
		return
	}

	log.Info("Receive a PDU Session Resource Setup Response from gnB ")

	msg := &n2models.PduSessResSetRsp{}
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
			msg.FailedList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceSetupUnsuccessfulTransfer,
			}
		}

	}

	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_PDU_SET_RSP,
		Content: msg,
	})

	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
}

func (h *Ngap) handlePDUSessionResourceReleaseResponse(ran *ran.Ran, pDUSessionResourceReleaseResponse *ngapType.PDUSessionResourceReleaseResponse) {
	//var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var critical *ngapType.CriticalityDiagnostics
	var oklist *ngapType.PDUSessionResourceReleasedListRelRes
	var locinfo *ngapType.UserLocationInformation

	log.Info("Receive a PDU Session Resource Release Response from gnB")

	for _, ie := range pDUSessionResourceReleaseResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			//	coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID:
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUENgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceReleasedListRelRes:
			oklist = ie.Value.PDUSessionResourceReleasedListRelRes
			log.Trace("Decode IE PDUSessionResourceReleasedList")
		case ngapType.ProtocolIEIDUserLocationInformation:
			locinfo = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	uectx := ran.FindUe(ranNgapId, nil)
	if uectx == nil {
		log.Errorf("UeContext not found")
		return
	}

	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}

	msg := &n2models.PduSessResRelRsp{}
	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}

	if oklist != nil && len(oklist.List) > 0 {
		msg.List = make([]n2models.UlPduSessionResourceInfo, len(oklist.List))
		for i, s := range oklist.List {
			msg.List[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceReleaseResponseTransfer,
			}
		}
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_PDU_REL_RSP,
		Content: msg,
	})

}

func (h *Ngap) handlePDUSessionResourceModifyResponse(ran *ran.Ran, pDUSessionResourceModifyResponse *ngapType.PDUSessionResourceModifyResponse) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var oklist *ngapType.PDUSessionResourceModifyListModRes
	var failedlist *ngapType.PDUSessionResourceFailedToModifyListModRes
	var locinfo *ngapType.UserLocationInformation
	var critical *ngapType.CriticalityDiagnostics

	log.Info("Receive a PDU Session Resource Modify Response from gnB")

	for _, ie := range pDUSessionResourceModifyResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceModifyListModRes: // ignore
			oklist = ie.Value.PDUSessionResourceModifyListModRes
			log.Trace("Decode IE PDUSessionResourceModifyListModRes")
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModRes: // ignore
			failedlist = ie.Value.PDUSessionResourceFailedToModifyListModRes
			log.Trace("Decode IE PDUSessionResourceFailedToModifyListModRes")
		case ngapType.ProtocolIEIDUserLocationInformation: // optional, ignore
			locinfo = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	var uectx *ue.UeContext

	if uectx = ran.FindUe(ranNgapId, coreNgapId); uectx == nil {
		log.Errorf("UeContext not found")
		return
	}

	msg := &n2models.PduSessResModRsp{}

	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}

	log.Tracef("AmfUeNgapID[%d] UeContextNgapID[%d]", uectx.CuNgapId(), uectx.RanNgapId())
	if oklist != nil && len(oklist.List) > 0 {
		msg.SuccessList = make([]n2models.UlPduSessionResourceInfo, len(oklist.List))
		for i, s := range oklist.List {
			msg.SuccessList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceModifyResponseTransfer,
			}
		}
	}
	if failedlist != nil && len(failedlist.List) > 0 {
		msg.FailedList = make([]n2models.UlPduSessionResourceInfo, len(failedlist.List))
		for i, s := range failedlist.List {
			msg.FailedList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceModifyUnsuccessfulTransfer,
			}
		}
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_PDU_MOD_RSP,
		Content: msg,
	})

	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
}

// TS139.413-V15.3.0 8.2.4
func (h *Ngap) handlePDUSessionResourceNotify(ran *ran.Ran, PDUSessionResourceNotify *ngapType.PDUSessionResourceNotify) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var notlist *ngapType.PDUSessionResourceNotifyList
	var rellist *ngapType.PDUSessionResourceReleasedListNot
	var locinfo *ngapType.UserLocationInformation

	log.Info("Receive a PDU Session Resource Notify from gnB")

	for _, ie := range PDUSessionResourceNotify.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			coreNgapId = ie.Value.AMFUENGAPID // reject
			//log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID:
			ranNgapId = ie.Value.RANUENGAPID // reject
			//log.Trace("Decode IE UeContextNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceNotifyList: // reject
			notlist = ie.Value.PDUSessionResourceNotifyList
			//log.Trace("Decode IE pDUSessionResourceNotifyList")
		case ngapType.ProtocolIEIDPDUSessionResourceReleasedListNot: // ignore
			rellist = ie.Value.PDUSessionResourceReleasedListNot
			//log.Trace("Decode IE PDUSessionResourceReleasedListNot")
		case ngapType.ProtocolIEIDUserLocationInformation: // optional, ignore
			locinfo = ie.Value.UserLocationInformation
			//log.Trace("Decode IE userLocationInformation")
		}
	}
	var uectx *ue.UeContext

	uectx = ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Errorf("UeContext not found")
		return
	}

	msg := &n2models.PduSessResNot{}
	log.Tracef("AmfUeNgapID[%d] UeContextNgapID[%d]", uectx.CuNgapId(), uectx.RanNgapId())

	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}

	if notlist != nil && len(notlist.List) > 0 {
		msg.NotifyList = make([]n2models.UlPduSessionResourceInfo, len(notlist.List))
		for i, s := range notlist.List {
			msg.NotifyList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceNotifyTransfer,
			}
		}
	}
	if rellist != nil && len(rellist.List) > 0 {
		msg.ReleasedList = make([]n2models.UlPduSessionResourceInfo, len(rellist.List))
		for i, s := range rellist.List {
			msg.ReleasedList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceNotifyReleasedTransfer,
			}
		}
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_PDU_NOT,
		Content: msg,
	})

}

// TS139.413-V15.3.0 8.2.5
func (h *Ngap) handlePDUSessionResourceModifyIndication(ran *ran.Ran, pDUSessionResourceModifyIndication *ngapType.PDUSessionResourceModifyIndication) {

	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var modlist *ngapType.PDUSessionResourceModifyListModInd
	var critical ngapType.CriticalityDiagnosticsIEList

	log.Info("Receive a PDU Session Resource Modify Indication from gnB")

	for _, ie := range pDUSessionResourceModifyIndication.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDAMFUENGAPID, ngapType.TypeOfErrorPresentMissing)
				critical.List = append(critical.List, item)
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDRANUENGAPID, ngapType.TypeOfErrorPresentMissing)
				critical.List = append(critical.List, item)
			}
		case ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd: // reject
			modlist = ie.Value.PDUSessionResourceModifyListModInd
			log.Trace("Decode IE PDUSessionResourceModifyListModInd")
			if modlist == nil {
				//log.Error("PDUSessionResourceModifyListModInd is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd, ngapType.TypeOfErrorPresentMissing)
				critical.List = append(critical.List, item)
			}
		}
	}

	var uectx *ue.UeContext
	if len(critical.List) > 0 {
		log.Error("Has missing reject IE(s)")

		procedureCode := ngapType.ProcedureCodePDUSessionResourceModifyIndication
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage, &procedureCriticality,
			&critical)
		SendErrorIndication(ran, coreNgapId, ranNgapId, nil, &criticalityDiagnostics)
		return
	}

	uectx = ran.FindUe(ranNgapId, coreNgapId)
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

	msg := &n2models.PduSessResModInd{}

	log.Tracef("UE Context AmfUeNgapID[%d] UeContextNgapID[%d]", uectx.CuNgapId(), uectx.RanNgapId())
	if modlist != nil && len(modlist.List) > 0 {
		msg.ModifyList = make([]n2models.UlPduSessionResourceInfo, len(modlist.List))
		for i, s := range modlist.List {
			msg.ModifyList[i] = n2models.UlPduSessionResourceInfo{
				Id:       int64(s.PDUSessionID.Value),
				Transfer: s.PDUSessionResourceModifyIndicationTransfer,
			}
		}

	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_PDU_MOD_IND,
		Content: msg,
	})

}
