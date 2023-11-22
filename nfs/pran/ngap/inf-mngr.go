package ngap

import (
	"etrib5gc/nfs/pran/context"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/sbi/utils/ngapConvert"
	"fmt"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handleNGSetupRequest(ran *ran.Ran, req *ngapType.NGSetupRequest) {
	var ranid *ngapType.GlobalRANNodeID
	var ranname *ngapType.RANNodeName
	var supportedtalist *ngapType.SupportedTAList
	var drx *ngapType.PagingDRX

	//var cause ngapType.Cause

	log.Infof("Receive a  NG Setup request")

	for i := 0; i < len(req.ProtocolIEs.List); i++ {
		ie := req.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDGlobalRANNodeID:
			ranid = ie.Value.GlobalRANNodeID
			log.Trace("Decode IE GlobalRANNodeID")
			if ranid == nil {
				log.Error("GlobalRANNodeID is nil")
				return
			}
		case ngapType.ProtocolIEIDSupportedTAList:
			supportedtalist = ie.Value.SupportedTAList
			log.Trace("Decode IE SupportedTAList")
			if supportedtalist == nil {
				log.Error("SupportedTAList is nil")
				return
			}
		case ngapType.ProtocolIEIDRANNodeName:
			ranname = ie.Value.RANNodeName
			log.Trace("Decode IE RANNodeName")
			if ranname == nil {
				log.Error("RANNodeName is nil")
				return
			}
		case ngapType.ProtocolIEIDDefaultPagingDRX:
			drx = ie.Value.DefaultPagingDRX
			log.Trace("Decode IE DefaultPagingDRX")
			if drx == nil {
				log.Error("DefaultPagingDRX is nil")
				return
			}
		}
	}
	ran.Setup(ranid, ranname, drx, supportedtalist)
	h.ranpool.Add(ran)
	SendNGSetupResponse(h.ctx, ran)
	/*
		cause.Misc = &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnknownPLMN,
		}

		SendNGSetupFailure(ran, cause)
	*/
}
func (h *Ngap) handleNGResetAcknowledge(ran *ran.Ran, ack *ngapType.NGResetAcknowledge) {
	var uEAssociatedLogicalNGConnectionList *ngapType.UEAssociatedLogicalNGConnectionList
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	log.Info("Receive an NG Reset Acknowledge")

	for _, ie := range ack.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDUEAssociatedLogicalNGConnectionList:
			uEAssociatedLogicalNGConnectionList = ie.Value.UEAssociatedLogicalNGConnectionList
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
		}
	}

	if uEAssociatedLogicalNGConnectionList != nil {
		log.Tracef("%d UE association(s) has been reset", len(uEAssociatedLogicalNGConnectionList.List))
		for i, item := range uEAssociatedLogicalNGConnectionList.List {
			if item.AMFUENGAPID != nil && item.RANUENGAPID != nil {
				log.Tracef("%d: AmfUeNgapID[%d] UeContextNgapID[%d]", i+1, item.AMFUENGAPID.Value, item.RANUENGAPID.Value)
			} else if item.AMFUENGAPID != nil {
				log.Tracef("%d: AmfUeNgapID[%d] UeContextNgapID[-1]", i+1, item.AMFUENGAPID.Value)
			} else if item.RANUENGAPID != nil {
				log.Tracef("%d: AmfUeNgapID[-1] UeContextNgapID[%d]", i+1, item.RANUENGAPID.Value)
			}
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *Ngap) handleNGReset(ran *ran.Ran, reset *ngapType.NGReset) {
	var cause *ngapType.Cause
	var rtype *ngapType.ResetType

	log.Info("Receive an NG Reset")

	for _, ie := range reset.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		case ngapType.ProtocolIEIDResetType:
			rtype = ie.Value.ResetType
			log.Trace("Decode IE ResetType")
			if rtype == nil {
				log.Error("ResetType is nil")
				return
			}
		}
	}

	switch rtype.Present {
	case ngapType.ResetTypePresentNGInterface:
		log.Trace("ResetType Present: NG Interface")
		ran.RemoveUes()
		SendNGResetAcknowledge(ran, nil, nil)
	case ngapType.ResetTypePresentPartOfNGInterface:
		log.Trace("ResetType Present: Part of NG Interface")

		uelist := rtype.PartOfNGInterface
		if uelist == nil {
			log.Error("PartOfNGInterface is nil")
			return
		}

		for _, ueItem := range uelist.List {
			ran.RemoveUe(ueItem.RANUENGAPID, ueItem.AMFUENGAPID)
		}
		SendNGResetAcknowledge(ran, uelist, nil)
	default:
		log.Warnf("Invalid ResetType[%d]", rtype.Present)
	}
}
func (h *Ngap) handleErrorIndication(ran *ran.Ran, errorIndication *ngapType.ErrorIndication) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	for _, ie := range errorIndication.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ran_ngapid == nil {
				log.Error("UeContextNgapID is nil")
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	log.Infof("Receive an Error Indication: RAN_UE_NGAP_ID:%v AMF_UE_NGAP_ID:%v", ran_ngapid, amf_ngapid)

	if cause == nil && criticalityDiagnostics == nil {
		log.Error("[ErrorIndication] both Cause IE and CriticalityDiagnostics IE are nil, should have at least one")
		return
	}

	if cause != nil {
		printAndGetCause(ran, cause)
	}
	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}

	// TODO: handle error based on cause/criticalityDiagnostics
}

func (h *Ngap) handleRanConfigurationUpdate(ran *ran.Ran, update *ngapType.RANConfigurationUpdate) {
	var ranname *ngapType.RANNodeName
	var supportedtalist *ngapType.SupportedTAList
	var drx *ngapType.PagingDRX

	//var cause ngapType.Cause

	log.Info("Receive a Ran Configuration Update")
	for i := 0; i < len(update.ProtocolIEs.List); i++ {
		ie := update.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRANNodeName:
			ranname = ie.Value.RANNodeName
			if ranname == nil {
				log.Error("RAN Node Name is nil")
				return
			}
			log.Tracef("Decode IE RANNodeName = [%s]", ranname.Value)
		case ngapType.ProtocolIEIDSupportedTAList:
			supportedtalist = ie.Value.SupportedTAList
			log.Trace("Decode IE SupportedTAList")
			if supportedtalist == nil {
				log.Error("Supported TA List is nil")
				return
			}
		case ngapType.ProtocolIEIDDefaultPagingDRX:
			drx = ie.Value.DefaultPagingDRX
			if drx == nil {
				log.Error("PagingDRX is nil")
				return
			}
			log.Tracef("Decode IE PagingDRX = [%d]", drx.Value)
		}
	}

	ran.Setup(nil, ranname, drx, supportedtalist)
	h.ranpool.Add(ran)
	log.Info("Receive a RanConfigurationUpdateAcknowledge")
	SendRanConfigurationUpdateAcknowledge(ran, nil)
	/*
		log.Info("Handle RanConfigurationUpdateAcknowledgeFailure")
		SendRanConfigurationUpdateFailure(ran, cause, nil)
	*/
}

func (h *Ngap) handleAMFconfigurationUpdateAcknowledge(ran *ran.Ran, ack *ngapType.AMFConfigurationUpdateAcknowledge) {
	var successlist *ngapType.AMFTNLAssociationSetupList
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics
	var failedlist *ngapType.TNLAssociationList

	log.Info("Receive an AMF Configuration Update Acknowledge")

	for i := 0; i < len(ack.ProtocolIEs.List); i++ {
		ie := ack.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFTNLAssociationSetupList:
			successlist = ie.Value.AMFTNLAssociationSetupList
			log.Trace("Decode IE AMFTNLAssociationSetupList")
			if successlist == nil {
				log.Error("AMFTNLAssociationSetupList is nil")
				return
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE Criticality Diagnostics")

		case ngapType.ProtocolIEIDAMFTNLAssociationFailedToSetupList:
			failedlist = ie.Value.AMFTNLAssociationFailedToSetupList
			log.Trace("Decode IE AMFTNLAssociationFailedToSetupList")
			if failedlist == nil {
				log.Error("AMFTNLAssociationFailedToSetupList is nil")
				return
			}
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *Ngap) handleAMFconfigurationUpdateFailure(ran *ran.Ran, failure *ngapType.AMFConfigurationUpdateFailure) {
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	log.Info("Receive an AMF Confioguration Update Failure")

	for _, ie := range failure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	//	TODO: Time To Wait

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func SendNGSetupResponse(ctx *context.CuContext, ran *ran.Ran) (err error) {
	defer logSendingReport("NGSetupResponse", err)
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

	sc := pdu.SuccessfulOutcome
	sc.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	sc.Criticality.Value = ngapType.CriticalityPresentReject
	sc.Value.Present = ngapType.SuccessfulOutcomePresentNGSetupResponse
	sc.Value.NGSetupResponse = new(ngapType.NGSetupResponse)

	rsp := sc.Value.NGSetupResponse
	ies := &rsp.ProtocolIEs

	// AMFName
	ie := ngapType.NGSetupResponseIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFName
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.NGSetupResponseIEsPresentAMFName
	ie.Value.AMFName = &ngapType.AMFName{
		Value: ctx.Name(),
	}

	log.Tracef("AmfName = %s", ctx.Name())
	ies.List = append(ies.List, ie)

	// ServedGUAMIList
	ie = ngapType.NGSetupResponseIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDServedGUAMIList
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.NGSetupResponseIEsPresentServedGUAMIList
	ie.Value.ServedGUAMIList = new(ngapType.ServedGUAMIList)
	guamlist := ie.Value.ServedGUAMIList

	log.Tracef("AmfId = %s", ctx.AmfId())
	log.Tracef("Mnc %s - Mcc: %s", ctx.PlmnId().Mnc, ctx.PlmnId().Mcc)

	item := ngapType.ServedGUAMIItem{}
	item.GUAMI.PLMNIdentity =
		ngapConvert.PlmnIdToNgap(*ctx.PlmnId()) //tungtq
	regionId, setId, prtId := ngapConvert.AmfIdToNgap(ctx.AmfId())
	item.GUAMI.AMFRegionID.Value = regionId
	item.GUAMI.AMFSetID.Value = setId
	item.GUAMI.AMFPointer.Value = prtId
	guamlist.List = append(guamlist.List, item)

	ies.List = append(ies.List, ie)
	// relativeAMFCapacity
	ie = ngapType.NGSetupResponseIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRelativeAMFCapacity
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.NGSetupResponseIEsPresentRelativeAMFCapacity
	ie.Value.RelativeAMFCapacity = &ngapType.RelativeAMFCapacity{
		Value: ctx.RelativeCapacity(),
	}

	ies.List = append(ies.List, ie)

	// PlmnList
	ie = ngapType.NGSetupResponseIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDPLMNSupportList
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.NGSetupResponseIEsPresentPLMNSupportList
	ie.Value.PLMNSupportList = new(ngapType.PLMNSupportList)

	plmnlist := ie.Value.PLMNSupportList
	for plmnid, slices := range ctx.PlmnList() {
		item := ngapType.PLMNSupportItem{}
		item.PLMNIdentity = ngapConvert.PlmnIdToNgap(plmnid)
		for _, snssai := range slices {
			ssitem := ngapType.SliceSupportItem{}
			ssitem.SNSSAI = ngapConvert.SNssaiToNgap(snssai)
			item.SliceSupportList.List = append(item.SliceSupportList.List, ssitem)
		}
		plmnlist.List = append(plmnlist.List, item)
	}

	ies.List = append(ies.List, ie)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

func SendNGSetupFailure(ran *ran.Ran, cause ngapType.Cause) (err error) {
	defer logSendingReport("NGSetupFailure", err)

	if cause.Present == ngapType.CausePresentNothing {
		err = fmt.Errorf("Cause present is nil")
		log.Errorf(err.Error())
		return
	}
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)

	uo := pdu.UnsuccessfulOutcome
	uo.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	uo.Criticality.Value = ngapType.CriticalityPresentReject
	uo.Value.Present = ngapType.UnsuccessfulOutcomePresentNGSetupFailure
	uo.Value.NGSetupFailure = new(ngapType.NGSetupFailure)

	ies := &uo.Value.NGSetupFailure.ProtocolIEs

	// Cause
	ie := ngapType.NGSetupFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDCause
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.NGSetupFailureIEsPresentCause
	ie.Value.Cause = &cause

	ies.List = append(ies.List, ie)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

// criticality ->from received node when received node can't comprehend the IE or missing IE
func SendRanConfigurationUpdateAcknowledge(ran *ran.Ran, criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {
	defer logSendingReport("RanConfigurationUpdateAcknowledge", err)

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

	sc := pdu.SuccessfulOutcome
	sc.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	sc.Criticality.Value = ngapType.CriticalityPresentReject
	sc.Value.Present = ngapType.SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge
	sc.Value.RANConfigurationUpdateAcknowledge = new(ngapType.RANConfigurationUpdateAcknowledge)

	ies := sc.Value.RANConfigurationUpdateAcknowledge.ProtocolIEs

	// Criticality Doagnostics(Optional)
	if criticalityDiagnostics != nil {
		ie := ngapType.RANConfigurationUpdateAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.RANConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = criticalityDiagnostics
		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

// criticality ->from received node when received node can't comprehend the IE or missing IE
// If the AMF cannot accept the update,
// it shall respond with a RAN CONFIGURATION UPDATE FAILURE message and appropriate cause value.
func SendRanConfigurationUpdateFailure(ran *ran.Ran, cause ngapType.Cause,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {

	defer logSendingReport("RanConfigurationUpdateFailure", err)

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)

	uo := pdu.UnsuccessfulOutcome
	uo.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	uo.Criticality.Value = ngapType.CriticalityPresentReject
	uo.Value.Present = ngapType.UnsuccessfulOutcomePresentRANConfigurationUpdateFailure
	uo.Value.RANConfigurationUpdateFailure = new(ngapType.RANConfigurationUpdateFailure)

	ies := uo.Value.RANConfigurationUpdateFailure.ProtocolIEs

	// Cause
	ie := ngapType.RANConfigurationUpdateFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDCause
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.RANConfigurationUpdateFailureIEsPresentCause
	ie.Value.Cause = &cause

	ies.List = append(ies.List, ie)

	// Time To Wait(Optional)
	ie = ngapType.RANConfigurationUpdateFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDTimeToWait
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.RANConfigurationUpdateFailureIEsPresentTimeToWait
	ie.Value.TimeToWait = new(ngapType.TimeToWait)

	timeToWait := ie.Value.TimeToWait
	timeToWait.Value = ngapType.TimeToWaitPresentV1s

	ies.List = append(ies.List, ie)

	// Criticality Doagnostics(Optional)
	if criticalityDiagnostics != nil {
		ie = ngapType.RANConfigurationUpdateFailureIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.RANConfigurationUpdateFailureIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = criticalityDiagnostics
		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

//NOTE: inhereted from free5gc, we may not need this in cloud-native
//deployments

// Weight Factor associated with each of the TNL association within the AMF
func SendAMFConfigurationUpdate(ctx *context.CuContext, ran *ran.Ran, tNLassociationUsage ngapType.TNLAssociationUsage,
	tNLAddressWeightFactor ngapType.TNLAddressWeightFactor) (err error) {
	defer logSendingReport("AMFConfigurationUpdate", err)

	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentReject
	msg.Value.Present = ngapType.InitiatingMessagePresentAMFConfigurationUpdate
	msg.Value.AMFConfigurationUpdate = new(ngapType.AMFConfigurationUpdate)

	ies := msg.Value.AMFConfigurationUpdate.ProtocolIEs

	//	AMF Name(optional)
	ie := ngapType.AMFConfigurationUpdateIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFName
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFName
	ie.Value.AMFName = new(ngapType.AMFName)

	aMFName := ie.Value.AMFName
	aMFName.Value = ctx.Name()

	ies.List = append(ies.List, ie)

	//	Served GUAMI List
	/*
		ie = ngapType.AMFConfigurationUpdateIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDServedGUAMIList
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentServedGUAMIList
		ie.Value.ServedGUAMIList = new(ngapType.ServedGUAMIList)

		guamilist := ie.Value.ServedGUAMIList
		for _, guami := range ctx.ServedGuamiList() {
			guamiitem := ngapType.ServedGUAMIItem{}
			guamiitem.GUAMI.PLMNIdentity =
				ngapConvert.PlmnIdToNgap(models.PlmnId{
					Mcc: guami.PlmnId.Mcc,
					Mnc: guami.PlmnId.Mnc,
				})
			regionId, setId, prtId := ngapConvert.AmfIdToNgap(guami.AmfId)
			guamiitem.GUAMI.AMFRegionID.Value = regionId
			guamiitem.GUAMI.AMFSetID.Value = setId
			guamiitem.GUAMI.AMFPointer.Value = prtId
			guamilist.List = append(guamilist.List, guamiitem)
		}

		ies.List = append(ies.List, ie)
	*/
	//	relative AMF Capability
	ie = ngapType.AMFConfigurationUpdateIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRelativeAMFCapacity
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.NGSetupResponseIEsPresentRelativeAMFCapacity
	ie.Value.RelativeAMFCapacity = new(ngapType.RelativeAMFCapacity)
	relativeAMFCapacity := ie.Value.RelativeAMFCapacity
	relativeAMFCapacity.Value = ctx.RelativeCapacity()

	ies.List = append(ies.List, ie)

	//	PLMN Support List
	ie = ngapType.AMFConfigurationUpdateIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDPLMNSupportList
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentPLMNSupportList
	ie.Value.PLMNSupportList = new(ngapType.PLMNSupportList)

	plmnlist := ie.Value.PLMNSupportList
	for plmnid, slices := range ctx.PlmnList() {
		plmnitem := ngapType.PLMNSupportItem{}
		plmnitem.PLMNIdentity = ngapConvert.PlmnIdToNgap(plmnid)
		for _, snssai := range slices {
			ssitem := ngapType.SliceSupportItem{}
			ssitem.SNSSAI = ngapConvert.SNssaiToNgap(snssai)
			plmnitem.SliceSupportList.List = append(plmnitem.SliceSupportList.List, ssitem)
		}
		plmnlist.List = append(plmnlist.List, plmnitem)
	}

	ies.List = append(ies.List, ie)

	//	AMF TNL Association to Add List
	ie = ngapType.AMFConfigurationUpdateIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationToAddList
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToAddList
	ie.Value.AMFTNLAssociationToAddList = new(ngapType.AMFTNLAssociationToAddList)

	aMFTNLAssociationToAddList := ie.Value.AMFTNLAssociationToAddList

	//	AMFTNLAssociationToAddItem in AMFTNLAssociationToAddList
	aMFTNLAssociationToAddItem := ngapType.AMFTNLAssociationToAddItem{}
	aMFTNLAssociationToAddItem.AMFTNLAssociationAddress.Present =
		ngapType.CPTransportLayerInformationPresentEndpointIPAddress
	aMFTNLAssociationToAddItem.AMFTNLAssociationAddress.EndpointIPAddress =
		new(ngapType.TransportLayerAddress)
	*aMFTNLAssociationToAddItem.AMFTNLAssociationAddress.EndpointIPAddress =
		//ngapConvert.IPAddressToNgap(ctx.Ipv4Addr(), ctx.Ipv6Addr())
		ngapConvert.IPAddressToNgap("", "")

	//	AMF TNL Association Usage[optional]
	if aMFTNLAssociationToAddItem.TNLAssociationUsage != nil {
		aMFTNLAssociationToAddItem.TNLAssociationUsage = new(ngapType.TNLAssociationUsage)
		aMFTNLAssociationToAddItem.TNLAssociationUsage = &tNLassociationUsage
	}

	//	AMF TNL Address Weight Factor
	aMFTNLAssociationToAddItem.TNLAddressWeightFactor = tNLAddressWeightFactor

	aMFTNLAssociationToAddList.List = append(aMFTNLAssociationToAddList.List, aMFTNLAssociationToAddItem)
	ies.List = append(ies.List, ie)

	//	AMF TNL Association to Remove List
	ie = ngapType.AMFConfigurationUpdateIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationToRemoveList
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToRemoveList
	ie.Value.AMFTNLAssociationToRemoveList = new(ngapType.AMFTNLAssociationToRemoveList)

	aMFTNLAssociationToRemoveList := ie.Value.AMFTNLAssociationToRemoveList

	//	AMFTNLAssociationToRemoveItem
	aMFTNLAssociationToRemoveItem := ngapType.AMFTNLAssociationToRemoveItem{}
	aMFTNLAssociationToRemoveItem.AMFTNLAssociationAddress.Present =
		ngapType.CPTransportLayerInformationPresentEndpointIPAddress
	aMFTNLAssociationToRemoveItem.AMFTNLAssociationAddress.EndpointIPAddress =
		new(ngapType.TransportLayerAddress)
	*aMFTNLAssociationToRemoveItem.AMFTNLAssociationAddress.EndpointIPAddress =
		//ngapConvert.IPAddressToNgap(ctx.Ipv4Addr(), ctx.Ipv6Addr())
		ngapConvert.IPAddressToNgap("", "")

	aMFTNLAssociationToRemoveList.List = append(aMFTNLAssociationToRemoveList.List, aMFTNLAssociationToRemoveItem)
	ies.List = append(ies.List, ie)

	//	AMFTNLAssociationToUpdateList
	ie = ngapType.AMFConfigurationUpdateIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationToUpdateList
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToUpdateList
	ie.Value.AMFTNLAssociationToUpdateList = new(ngapType.AMFTNLAssociationToUpdateList)

	aMFTNLAssociationToUpdateList := ie.Value.AMFTNLAssociationToUpdateList

	//	AMFTNLAssociationAddress in AMFTNLAssociationtoUpdateItem
	aMFTNLAssociationToUpdateItem := ngapType.AMFTNLAssociationToUpdateItem{}
	aMFTNLAssociationToUpdateItem.AMFTNLAssociationAddress.Present =
		ngapType.CPTransportLayerInformationPresentEndpointIPAddress
	aMFTNLAssociationToUpdateItem.AMFTNLAssociationAddress.EndpointIPAddress =
		new(ngapType.TransportLayerAddress)
	*aMFTNLAssociationToUpdateItem.AMFTNLAssociationAddress.EndpointIPAddress =
		//ngapConvert.IPAddressToNgap(ctx.Ipv4Addr(), ctx.Ipv6Addr())
		ngapConvert.IPAddressToNgap("", "")

	//	TNLAssociationUsage in AMFTNLAssociationtoUpdateItem [optional]
	if aMFTNLAssociationToUpdateItem.TNLAssociationUsage != nil {
		aMFTNLAssociationToUpdateItem.TNLAssociationUsage = new(ngapType.TNLAssociationUsage)
		aMFTNLAssociationToUpdateItem.TNLAssociationUsage = &tNLassociationUsage
	}
	//	TNLAddressWeightFactor in AMFTNLAssociationtoUpdateItem [optional]
	if aMFTNLAssociationToUpdateItem.TNLAddressWeightFactor != nil {
		aMFTNLAssociationToUpdateItem.TNLAddressWeightFactor = new(ngapType.TNLAddressWeightFactor)
		aMFTNLAssociationToUpdateItem.TNLAddressWeightFactor = &tNLAddressWeightFactor
	}
	aMFTNLAssociationToUpdateList.List = append(aMFTNLAssociationToUpdateList.List, aMFTNLAssociationToUpdateItem)
	ies.List = append(ies.List, ie)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

// partOfNGInterface: if reset type is "reset all", set it to nil TS 38.413 9.2.6.11
func SendNGReset(ran *ran.Ran, cause ngapType.Cause,
	partOfNGInterface *ngapType.UEAssociatedLogicalNGConnectionList) (err error) {
	defer logSendingReport("NGReset", err)
	var pdu ngapType.NGAPPDU

	//log.Trace("Build NG Reset message")

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGReset
	msg.Criticality.Value = ngapType.CriticalityPresentReject

	msg.Value.Present = ngapType.InitiatingMessagePresentNGReset
	msg.Value.NGReset = new(ngapType.NGReset)

	ies := msg.Value.NGReset.ProtocolIEs

	// Cause
	ie := ngapType.NGResetIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDCause
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.NGResetIEsPresentCause
	ie.Value.Cause = &cause

	ies.List = append(ies.List, ie)

	// Reset Type
	ie = ngapType.NGResetIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDResetType
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.NGResetIEsPresentResetType
	ie.Value.ResetType = new(ngapType.ResetType)

	resetType := ie.Value.ResetType

	if partOfNGInterface == nil {
		resetType.Present = ngapType.ResetTypePresentNGInterface
		resetType.NGInterface = new(ngapType.ResetAll)
		resetType.NGInterface.Value = ngapType.ResetAllPresentResetAll
	} else {
		resetType.Present = ngapType.ResetTypePresentPartOfNGInterface
		resetType.PartOfNGInterface = new(ngapType.UEAssociatedLogicalNGConnectionList)
		resetType.PartOfNGInterface = partOfNGInterface
	}

	ies.List = append(ies.List, ie)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

func SendNGResetAcknowledge(ran *ran.Ran, partOfNGInterface *ngapType.UEAssociatedLogicalNGConnectionList,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {
	defer logSendingReport("NGResetAcknowledge", err)
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

	sc := pdu.SuccessfulOutcome
	sc.ProcedureCode.Value = ngapType.ProcedureCodeNGReset
	sc.Criticality.Value = ngapType.CriticalityPresentReject

	sc.Value.Present = ngapType.SuccessfulOutcomePresentNGResetAcknowledge
	sc.Value.NGResetAcknowledge = new(ngapType.NGResetAcknowledge)

	ies := sc.Value.NGResetAcknowledge.ProtocolIEs

	// UE-associated Logical NG-connection List (optional)
	if partOfNGInterface != nil && len(partOfNGInterface.List) > 0 {
		ie := ngapType.NGResetAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUEAssociatedLogicalNGConnectionList
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.NGResetAcknowledgeIEsPresentUEAssociatedLogicalNGConnectionList
		ie.Value.UEAssociatedLogicalNGConnectionList = new(ngapType.UEAssociatedLogicalNGConnectionList)

		connlist := ie.Value.UEAssociatedLogicalNGConnectionList

		for i, item := range partOfNGInterface.List {
			if item.AMFUENGAPID == nil && item.RANUENGAPID == nil {
				log.Warn("[Build NG Reset Ack] No AmfUeNgapID & UeContextNgapID")
				continue
			}

			connitem := ngapType.UEAssociatedLogicalNGConnectionItem{}

			if item.AMFUENGAPID != nil {
				connitem.AMFUENGAPID = new(ngapType.AMFUENGAPID)
				connitem.AMFUENGAPID = item.AMFUENGAPID
				log.Tracef(
					"[Build NG Reset Ack] (pair %d) AmfUeNgapID[%d]", i, connitem.AMFUENGAPID)
			}
			if item.RANUENGAPID != nil {
				connitem.RANUENGAPID = new(ngapType.RANUENGAPID)
				connitem.RANUENGAPID = item.RANUENGAPID
				log.Tracef(
					"[Build NG Reset Ack] (pair %d) UeContextNgapID[%d]", i, connitem.RANUENGAPID)
			}

			connlist.List = append(connlist.List, connitem)
		}

		ies.List = append(ies.List, ie)
	}

	// Criticality Diagnostics (optional)
	if criticalityDiagnostics != nil {
		ie := ngapType.NGResetAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.NGResetAcknowledgeIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = criticalityDiagnostics

		ies.List = append(ies.List, ie)
	}

	if partOfNGInterface != nil && len(partOfNGInterface.List) == 0 {
		err = fmt.Errorf("length of partOfNGInterface is 0")
		log.Error(err.Error())
		return
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}

func SendErrorIndication(ran *ran.Ran, amfUeNgapId *ngapType.AMFUENGAPID, ranUeNgapId *ngapType.RANUENGAPID,
	cause *ngapType.Cause, criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {
	defer logSendingReport("ErrorIndication", err)
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeErrorIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore

	msg.Value.Present = ngapType.InitiatingMessagePresentErrorIndication
	msg.Value.ErrorIndication = new(ngapType.ErrorIndication)

	ies := msg.Value.ErrorIndication.ProtocolIEs

	if cause == nil && criticalityDiagnostics == nil {
		log.Error("[Build Error Indication] shall contain at least either the Cause or the Criticality Diagnostics")
	}

	if amfUeNgapId != nil {
		ie := ngapType.ErrorIndicationIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.ErrorIndicationIEsPresentAMFUENGAPID
		ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

		aMFUENGAPID := ie.Value.AMFUENGAPID
		aMFUENGAPID.Value = amfUeNgapId.Value

		ies.List = append(ies.List, ie)
	}

	if ranUeNgapId != nil {
		ie := ngapType.ErrorIndicationIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.ErrorIndicationIEsPresentRANUENGAPID
		ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

		rANUENGAPID := ie.Value.RANUENGAPID
		rANUENGAPID.Value = ranUeNgapId.Value

		ies.List = append(ies.List, ie)
	}

	if cause != nil {
		ie := ngapType.ErrorIndicationIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCause
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.ErrorIndicationIEsPresentCause
		ie.Value.Cause = new(ngapType.Cause)

		ie.Value.Cause = cause

		ies.List = append(ies.List, ie)
	}

	if criticalityDiagnostics != nil {
		ie := ngapType.ErrorIndicationIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.ErrorIndicationIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = criticalityDiagnostics

		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}

	return
}

// An AMF shall be able to instruct other peer CP NFs, subscribed to receive such a notification,
// that it will be unavailable on this AMF and its corresponding target AMF(s).
// If CP NF does not subscribe to receive AMF unavailable notification, the CP NF may attempt
// forwarding the transaction towards the old AMF and detect that the AMF is unavailable. When
// it detects unavailable, it marks the AMF and its associated GUAMI(s) as unavailable.
// Defined in 23.501 5.21.2.2.2
func SendAMFStatusIndication(ran *ran.Ran, unavailableGUAMIList ngapType.UnavailableGUAMIList) (err error) {
	defer logSendingReport("AMFStatusIndication", err)
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFStatusIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore

	msg.Value.Present = ngapType.InitiatingMessagePresentAMFStatusIndication
	msg.Value.AMFStatusIndication = new(ngapType.AMFStatusIndication)

	ies := msg.Value.AMFStatusIndication.ProtocolIEs

	//	Unavailable GUAMI List
	ie := ngapType.AMFStatusIndicationIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDUnavailableGUAMIList
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.AMFStatusIndicationIEsPresentUnavailableGUAMIList
	ie.Value.UnavailableGUAMIList = new(ngapType.UnavailableGUAMIList)

	ie.Value.UnavailableGUAMIList = &unavailableGUAMIList

	ies.List = append(ies.List, ie)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
		return
	}
	return
}

// TS 23.501 5.19.5.2
// amfOverloadResponse: the required behaviour of NG-RAN, provided by AMF
// amfTrafficLoadReductionIndication(int 1~99): indicates the percentage of the type, set to 0 if does not need this ie
// of traffic relative to the instantaneous incoming rate at the NG-RAN node, provided by AMF
// overloadStartNSSAIList: overload slices, provide by AMF
func SendOverloadStart(
	ran *ran.Ran,
	amfOverloadResponse *ngapType.OverloadResponse,
	amfTrafficLoadReductionIndication int64,
	overloadStartNSSAIList *ngapType.OverloadStartNSSAIList) (err error) {

	defer logSendingReport("OverloadStart", err)

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeOverloadStart
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore

	msg.Value.Present = ngapType.InitiatingMessagePresentOverloadStart
	msg.Value.OverloadStart = new(ngapType.OverloadStart)

	ies := msg.Value.OverloadStart.ProtocolIEs

	// AMF Overload Response (optional)
	if amfOverloadResponse != nil {
		ie := ngapType.OverloadStartIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAMFOverloadResponse
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.OverloadStartIEsPresentAMFOverloadResponse
		ie.Value.AMFOverloadResponse = amfOverloadResponse
		ies.List = append(ies.List, ie)
	}

	// AMF Traffic Load Reduction Indication (optional)
	if amfTrafficLoadReductionIndication != 0 {
		ie := ngapType.OverloadStartIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAMFTrafficLoadReductionIndication
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.OverloadStartIEsPresentAMFTrafficLoadReductionIndication
		ie.Value.AMFTrafficLoadReductionIndication = &ngapType.TrafficLoadReductionIndication{
			Value: amfTrafficLoadReductionIndication,
		}
		ies.List = append(ies.List, ie)
	}

	// Overload Start NSSAI List (optional)
	if overloadStartNSSAIList != nil {
		ie := ngapType.OverloadStartIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDOverloadStartNSSAIList
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.OverloadStartIEsPresentOverloadStartNSSAIList
		ie.Value.OverloadStartNSSAIList = overloadStartNSSAIList
		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {

		err = ran.Send(packet)
	}

	return
}

func SendOverloadStop(ran *ran.Ran) (err error) {
	defer logSendingReport("OverloadStop", err)
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeOverloadStop
	msg.Criticality.Value = ngapType.CriticalityPresentReject

	msg.Value.Present = ngapType.InitiatingMessagePresentOverloadStop
	msg.Value.OverloadStop = new(ngapType.OverloadStop)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = ran.Send(packet)
	}
	return
}
