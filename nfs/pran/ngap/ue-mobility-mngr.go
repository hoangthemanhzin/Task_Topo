package ngap

import (
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"
	"fmt"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handleHandoverRequired(ran *ran.Ran, HandoverRequired *ngapType.HandoverRequired) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var handoverType *ngapType.HandoverType
	var cause *ngapType.Cause
	var targetID *ngapType.TargetID
	var sessionList *ngapType.PDUSessionResourceListHORqd
	var container *ngapType.SourceToTargetTransparentContainer
	var critical ngapType.CriticalityDiagnosticsIEList

	for i := 0; i < len(HandoverRequired.ProtocolIEs.List); i++ {
		ie := HandoverRequired.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			coreNgapId = ie.Value.AMFUENGAPID // reject
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
		case ngapType.ProtocolIEIDHandoverType: // reject
			handoverType = ie.Value.HandoverType
			log.Trace("Decode IE HandoverType")
		case ngapType.ProtocolIEIDCause: // ignore
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
		case ngapType.ProtocolIEIDTargetID: // reject
			targetID = ie.Value.TargetID
			log.Trace("Decode IE TargetID")
		case ngapType.ProtocolIEIDPDUSessionResourceListHORqd: // reject
			sessionList = ie.Value.PDUSessionResourceListHORqd
			log.Trace("Decode IE PDUSessionResourceListHORqd")
		case ngapType.ProtocolIEIDSourceToTargetTransparentContainer: // reject
			container = ie.Value.SourceToTargetTransparentContainer
			log.Trace("Decode IE SourceToTargetTransparentContainer")
		}
	}

	if coreNgapId == nil {
		log.Error("AmfUeNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDAMFUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		critical.List = append(critical.List, item)
	}
	if ranNgapId == nil {
		log.Error("UeContextNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDRANUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		critical.List = append(critical.List, item)
	}

	if handoverType == nil {
		log.Error("handoverType is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDHandoverType,
			ngapType.TypeOfErrorPresentMissing)
		critical.List = append(critical.List, item)
	}
	if targetID == nil {
		log.Error("targetID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDTargetID,
			ngapType.TypeOfErrorPresentMissing)
		critical.List = append(critical.List, item)
	}
	if sessionList == nil {
		log.Error("pDUSessionResourceListHORqd is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
			ngapType.ProtocolIEIDPDUSessionResourceListHORqd, ngapType.TypeOfErrorPresentMissing)
		critical.List = append(critical.List, item)
	}
	if container == nil {
		log.Error("sourceToTargetTransparentContainer is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
			ngapType.ProtocolIEIDSourceToTargetTransparentContainer, ngapType.TypeOfErrorPresentMissing)
		critical.List = append(critical.List, item)
	}

	if len(critical.List) > 0 {
		procedureCode := ngapType.ProcedureCodeHandoverPreparation
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage,
			&procedureCriticality, &critical)
		SendErrorIndication(ran, coreNgapId, ranNgapId, nil, &criticalityDiagnostics)
		return
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
	if targetID.Present != ngapType.TargetIDPresentTargetRANNodeID {
		log.Errorf("targetID type[%d] is not supported", targetID.Present)
		return
	}

	if targetRan := h.ranpool.ByRanId(&targetID.TargetRANNodeID.GlobalRANNodeID); targetRan == nil {
		//Unknown Ran, find other AMF
		//		sourceUe.Log.Warnf("Handover required : cannot find target Ran Node Id[%+v] in this AMF", targetRanNodeId)
		//		sourceUe.Log.Error("Handover between different AMF has not been implemented yet")
		return
		// TODO: Send to T-AMF
		// Described in (23.502 4.9.1.3.2) step 3.Namf_Communication_CreateUEContext Request

	} else {
		dat := models.HandoverRequire{
			HandoverType: handoverType,
			TargetId:     targetID,
			Sessions:     sessionList,
			Container:    container,
		}
		dummy(&dat)
		//TODO: depending on the outcomes we should take different actions

	}

}

// pduSessionResourceHandoverList: provided by amf and transfer is return from smf
// pduSessionResourceToReleaseList: provided by amf and transfer is return from smf
// criticalityDiagnostics = criticalityDiagonstics IE in receiver node's error indication
// when received node can't comprehend the IE or missing IE
func SendHandoverCommand(
	sourceUe *ue.UeContext,
	handoverList ngapType.PDUSessionResourceHandoverList,
	releaseList ngapType.PDUSessionResourceToReleaseListHOCmd,
	container ngapType.TargetToSourceTransparentContainer,
	critical *ngapType.CriticalityDiagnostics) (err error) {

	if sourceUe == nil {
		err = fmt.Errorf("SourceUe is nil")
		log.Error(err.Error())
		return
	}

	log.Info("Send Handover Command")
	/*
		if len(handoverList.List) > context.MaxNumOfPDUSessions {
			log.Error("Pdu List out of range")
			return
		}

		if len(releaseList.List) > context.MaxNumOfPDUSessions {
			log.Error("Pdu List out of range")
			return
		}
	*/
	var pkt []byte
	if pkt, err = buildHandoverCommand(sourceUe, handoverList, releaseList, container, critical); err != nil {
		log.Errorf("Build HandoverCommand failed : %s", err.Error())
		return
	}
	err = sourceUe.Send(pkt)
	return
}

// pduSessionResourceHandoverList: provided by amf and transfer is return from smf
// pduSessionResourceToReleaseList: provided by amf and transfer is return from smf
// criticalityDiagnostics = criticalityDiagonstics IE in receiver node's error indication
// when received node can't comprehend the IE or missing IE
func buildHandoverCommand(
	sourceUe *ue.UeContext,
	handoverList ngapType.PDUSessionResourceHandoverList,
	releaseList ngapType.PDUSessionResourceToReleaseListHOCmd,
	container ngapType.TargetToSourceTransparentContainer,
	critical *ngapType.CriticalityDiagnostics) ([]byte, error) {

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

	successfulOutcome := pdu.SuccessfulOutcome
	successfulOutcome.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	successfulOutcome.Criticality.Value = ngapType.CriticalityPresentReject
	successfulOutcome.Value.Present = ngapType.SuccessfulOutcomePresentHandoverCommand
	successfulOutcome.Value.HandoverCommand = new(ngapType.HandoverCommand)

	handoverCommand := successfulOutcome.Value.HandoverCommand
	handoverCommandIEs := &handoverCommand.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.HandoverCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.HandoverCommandIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = sourceUe.CuNgapId()

	handoverCommandIEs.List = append(handoverCommandIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.HandoverCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = sourceUe.RanNgapId()

	handoverCommandIEs.List = append(handoverCommandIEs.List, ie)

	// Handover Type
	ie = ngapType.HandoverCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDHandoverType
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.HandoverCommandIEsPresentHandoverType
	ie.Value.HandoverType = new(ngapType.HandoverType)

	handoverType := ie.Value.HandoverType
	handoverType.Value = sourceUe.HandoverInfo().HandoverType.Value

	handoverCommandIEs.List = append(handoverCommandIEs.List, ie)

	// NAS Security Parameters from NG-RAN [C-iftoEPS]
	if handoverType.Value == ngapType.HandoverTypePresentFivegsToEps {
		ie = ngapType.HandoverCommandIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDNASSecurityParametersFromNGRAN
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverCommandIEsPresentNASSecurityParametersFromNGRAN
		ie.Value.NASSecurityParametersFromNGRAN = new(ngapType.NASSecurityParametersFromNGRAN)

		handoverCommandIEs.List = append(handoverCommandIEs.List, ie)
	}

	// PDU Session Resource Handover List
	ie = ngapType.HandoverCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceHandoverList
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.HandoverCommandIEsPresentPDUSessionResourceHandoverList
	ie.Value.PDUSessionResourceHandoverList = &handoverList
	handoverCommandIEs.List = append(handoverCommandIEs.List, ie)

	// PDU Session Resource to Release List
	if len(releaseList.List) > 0 {
		ie = ngapType.HandoverCommandIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.HandoverCommandIEsPresentPDUSessionResourceToReleaseListHOCmd
		ie.Value.PDUSessionResourceToReleaseListHOCmd = &releaseList
		handoverCommandIEs.List = append(handoverCommandIEs.List, ie)
	}

	// Target to Source Transparent Container
	ie = ngapType.HandoverCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDTargetToSourceTransparentContainer
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.HandoverCommandIEsPresentTargetToSourceTransparentContainer
	ie.Value.TargetToSourceTransparentContainer = &container

	handoverCommandIEs.List = append(handoverCommandIEs.List, ie)

	// Criticality Diagnostics [optional]
	if critical != nil {
		ie := ngapType.HandoverCommandIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = critical

		handoverCommandIEs.List = append(handoverCommandIEs.List, ie)
	}

	return libngap.Encoder(pdu)
}

// cause = initiate the Handover Cancel procedure with the appropriate value for the Cause IE.
// criticalityDiagnostics = criticalityDiagonstics IE in receiver node's error indication
// when received node can't comprehend the IE or missing IE
func SendHandoverPreparationFailure(sourceUe *ue.UeContext, cause ngapType.Cause,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {
	if sourceUe == nil {
		err = fmt.Errorf("SourceUe is nil")
		log.Info(err.Error())
		return
	}

	log.Info("Send Handover Preparation Failure")

	/*
		//TODO: need to invetigate this procedure
		amfUe := sourceUe.AmfUe()
		if amfUe == nil {
			log.Error("amfUe is nil")
			return
		}
		amfUe.SetOnGoing(sourceUe.Ran().AnType(), &context.OnGoing{
			Procedure: context.OnGoingProcedureNothing,
		})
	*/
	var pkt []byte
	if pkt, err = buildHandoverPreparationFailure(sourceUe, cause, criticalityDiagnostics); err != nil {
		log.Errorf("Build HandoverPreparationFailure failed : %s", err.Error())
		return
	}
	err = sourceUe.Send(pkt)
	return
}
func buildHandoverPreparationFailure(sourceUe *ue.UeContext, cause ngapType.Cause,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) ([]byte, error) {
	// cause = initiate the Handover Cancel procedure with the appropriate value for the Cause IE.

	// criticalityDiagnostics = criticalityDiagonstics IE in receiver node's error indication
	// when received node can't comprehend the IE or missing IE

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)

	unsuccessfulOutcome := pdu.UnsuccessfulOutcome
	unsuccessfulOutcome.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	unsuccessfulOutcome.Criticality.Value = ngapType.CriticalityPresentReject
	unsuccessfulOutcome.Value.Present = ngapType.UnsuccessfulOutcomePresentHandoverPreparationFailure
	unsuccessfulOutcome.Value.HandoverPreparationFailure = new(ngapType.HandoverPreparationFailure)

	handoverPreparationFailure := unsuccessfulOutcome.Value.HandoverPreparationFailure
	handoverPreparationFailureIEs := &handoverPreparationFailure.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.HandoverPreparationFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = sourceUe.CuNgapId()

	handoverPreparationFailureIEs.List = append(handoverPreparationFailureIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.HandoverPreparationFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = sourceUe.RanNgapId()

	handoverPreparationFailureIEs.List = append(handoverPreparationFailureIEs.List, ie)

	// Cause
	ie = ngapType.HandoverPreparationFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDCause
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics
	ie.Value.Cause = new(ngapType.Cause)

	ie.Value.Cause = &cause

	handoverPreparationFailureIEs.List = append(handoverPreparationFailureIEs.List, ie)

	// Criticality Diagnostics [optional]
	if criticalityDiagnostics != nil {
		ie := ngapType.HandoverPreparationFailureIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = criticalityDiagnostics

		handoverPreparationFailureIEs.List = append(handoverPreparationFailureIEs.List, ie)
	}

	return libngap.Encoder(pdu)
}

/*The PGW-C+SMF (V-SMF in the case of home-routed roaming scenario only) sends
a Nsmf_PDUSession_CreateSMContext Response(N2 SM Information (PDU Session ID, cause code)) to the AMF.*/
// Cause is from SMF
// sessionList provided by AMF, and the transfer data is from SMF
// container is received from S-RAN
// nsci: new security context indicator, if amfUe has updated security context, set nsci to true, otherwise set to false
// N2 handover in same AMF
func SendHandoverRequest(sourceUe *ue.UeContext, targetRan *ran.Ran, cause ngapType.Cause,
	sessionList ngapType.PDUSessionResourceSetupListHOReq,
	container ngapType.SourceToTargetTransparentContainer, nsci bool) (err error) {
	/*
		if sourceUe == nil {
			log.Error("sourceUe is nil")
			return
		}

		log.Info("Send Handover Request")

		amfUe := sourceUe.AmfUe
		if amfUe == nil {
			log.Error("amfUe is nil")
			return
		}
		if targetRan == nil {
			log.Error("targetRan is nil")
			return
		}

		if sourceUe.GetHandoverInfo().TargetUe != nil {
			log.Error("Handover Required Duplicated")
			return
		}

		if len(sessionList.List) > context.MaxNumOfPDUSessions {
			log.Error("Pdu List out of range")
			return
		}

		if len(container.Value) == 0 {
			log.Error("Source To Target TransparentContainer is nil")
			return
		}

		var targetUe UeContext
		if targetUeTmp, err := targetRan.NewUeContext(ue.UeContextNgapIdUnspecified); err != nil {
			log.Errorf("Create target UE error: %+v", err)
		} else {
			targetUe = targetUeTmp
		}

		log.Tracef("Source : AMF_UE_NGAP_ID[%d], RAN_UE_NGAP_ID[%d]", sourceUe.AmfUeNgapId, sourceUe.UeContextNgapId)
		log.Tracef("Target : AMF_UE_NGAP_ID[%d], RAN_UE_NGAP_ID[Unknown]", targetUe.AmfUeNgapId)
		context.AttachSourceUeTargetUe(sourceUe, targetUe)

		pkt, err := s.buildHandoverRequest(targetUe, cause, pduSessionResourceSetupListHOReq,
			sourceToTargetTransparentContainer, nsci)
		if err != nil {
			log.Errorf("Build HandoverRequest failed : %s", err.Error())
			return
		}
		SendToUeContext(targetUe, pkt)
	*/
	return
}

/*The PGW-C+SMF (V-SMF in the case of home-routed roaming scenario only) sends
a Nsmf_PDUSession_CreateSMContext Response(N2 SM Information (PDU Session ID, cause code)) to the AMF.*/
// Cause is from SMF
// pduSessionResourceSetupList provided by AMF, and the transfer data is from SMF
// sourceToTargetTransparentContainer is received from S-RAN
// nsci: new security context indicator, if amfUe has updated security context,
// set nsci to true, otherwise set to false
func buildHandoverRequest(ue *ue.UeContext, cause ngapType.Cause,
	pduSessionResourceSetupListHOReq ngapType.PDUSessionResourceSetupListHOReq,
	sourceToTargetTransparentContainer ngapType.SourceToTargetTransparentContainer, nsci bool) ([]byte, error) {
	/*
		amf := s.backend.Context()
		amfUe := ue.AmfUe()
		if amfUe == nil {
			return nil, fmt.Errorf("AmfUe is nil")
		}

		var pdu ngapType.NGAPPDU

		pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
		pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

		initiatingMessage := pdu.InitiatingMessage
		initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
		initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject

		initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentHandoverRequest
		initiatingMessage.Value.HandoverRequest = new(ngapType.HandoverRequest)

		handoverRequest := initiatingMessage.Value.HandoverRequest
		handoverRequestIEs := &handoverRequest.ProtocolIEs

		// AMF UE NGAP ID
		ie := ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentAMFUENGAPID
		ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

		aMFUENGAPID := ie.Value.AMFUENGAPID
		aMFUENGAPID.Value = ue.CuNgapId()

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// Handover Type
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDHandoverType
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentHandoverType
		ie.Value.HandoverType = new(ngapType.HandoverType)

		handoverType := ie.Value.HandoverType
		handoverType.Value = ue.GetHandoverInfo().HandOverType.Value

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// Cause
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCause
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.HandoverRequestIEsPresentCause
		ie.Value.Cause = &cause

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// UE Aggregate Maximum Bit Rate
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentUEAggregateMaximumBitRate
		ie.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)

		udminfo := amfUe.UdmClient().Info()
		ueAmbrUL := ngapConvert.UEAmbrToInt64(udminfo.AccessAndMobilitySubscriptionData.SubscribedUeAmbr.Uplink)
		ueAmbrDL := ngapConvert.UEAmbrToInt64(udminfo.AccessAndMobilitySubscriptionData.SubscribedUeAmbr.Downlink)
		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateUL.Value = ueAmbrUL
		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateDL.Value = ueAmbrDL

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// UE Security Capabilities
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentUESecurityCapabilities
		ie.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)

		ueSecurityCapabilities := ie.Value.UESecurityCapabilities

		secinfo := amfUe.AusfClient().SecInfo()

		nrEncryptionAlgorighm := []byte{0x00, 0x00}
		nrEncryptionAlgorighm[0] |= secinfo.UESecurityCapability.GetEA1_128_5G() << 7
		nrEncryptionAlgorighm[0] |= secinfo.UESecurityCapability.GetEA2_128_5G() << 6
		nrEncryptionAlgorighm[0] |= secinfo.UESecurityCapability.GetEA3_128_5G() << 5
		ueSecurityCapabilities.NRencryptionAlgorithms.Value = ngapConvert.ByteToBitString(nrEncryptionAlgorighm, 16)

		nrIntegrityAlgorithm := []byte{0x00, 0x00}
		nrIntegrityAlgorithm[0] |= secinfo.UESecurityCapability.GetIA1_128_5G() << 7
		nrIntegrityAlgorithm[0] |= secinfo.UESecurityCapability.GetIA2_128_5G() << 6
		nrIntegrityAlgorithm[0] |= secinfo.UESecurityCapability.GetIA3_128_5G() << 5
		ueSecurityCapabilities.NRintegrityProtectionAlgorithms.Value =
			ngapConvert.ByteToBitString(nrIntegrityAlgorithm, 16)

		// only support NR algorithms
		eutraEncryptionAlgorithm := []byte{0x00, 0x00}
		ueSecurityCapabilities.EUTRAencryptionAlgorithms.Value =
			ngapConvert.ByteToBitString(eutraEncryptionAlgorithm, 16)

		eutraIntegrityAlgorithm := []byte{0x00, 0x00}
		ueSecurityCapabilities.EUTRAintegrityProtectionAlgorithms.Value =
			ngapConvert.ByteToBitString(eutraIntegrityAlgorithm, 16)

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// Security Context
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDSecurityContext
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentSecurityContext
		ie.Value.SecurityContext = new(ngapType.SecurityContext)

		securityContext := ie.Value.SecurityContext
		securityContext.NextHopChainingCount.Value = int64(secinfo.NCC)
		securityContext.NextHopNH.Value = ngapConvert.HexToBitString(hex.EncodeToString(secinfo.NH), 256)

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// PDU Session Resource Setup List
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListHOReq
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentPDUSessionResourceSetupListHOReq
		ie.Value.PDUSessionResourceSetupListHOReq = &pduSessionResourceSetupListHOReq
		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// Allowed NSSAI
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentAllowedNSSAI
		ie.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)

		allowedNSSAI := ie.Value.AllowedNSSAI
		plmnList := amf.PlmnSupportList()
		for _, snssaiItem := range plmnList[0].SNssaiList {
			allowedNSSAIItem := ngapType.AllowedNSSAIItem{}

			ngapSnssai := ngapConvert.SNssaiToNgap(snssaiItem)
			allowedNSSAIItem.SNSSAI = ngapSnssai
			allowedNSSAI.List = append(allowedNSSAI.List, allowedNSSAIItem)
		}

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// Source To Target Transparent Container
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDSourceToTargetTransparentContainer
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentSourceToTargetTransparentContainer
		ie.Value.SourceToTargetTransparentContainer = new(ngapType.SourceToTargetTransparentContainer)

		sourceToTargetTransparentContaine := ie.Value.SourceToTargetTransparentContainer
		sourceToTargetTransparentContaine.Value = sourceToTargetTransparentContainer.Value

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)
		// GUAMI
		ie = ngapType.HandoverRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDGUAMI
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.HandoverRequestIEsPresentGUAMI
		ie.Value.GUAMI = new(ngapType.GUAMI)

		guami := ie.Value.GUAMI
		//plmnID := &guami.PLMNIdentity
		amfRegionID := &guami.AMFRegionID
		amfSetID := &guami.AMFSetID
		amfPtrID := &guami.AMFPointer

		servedGuami := amf.ServedGuamiList()[0]

		//	*plmnID = ngapConvert.PlmnIdToNgap(*servedGuami.PlmnId) tungtq
		amfRegionID.Value, amfSetID.Value, amfPtrID.Value = ngapConvert.AmfIdToNgap(servedGuami.AmfId)

		handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// //Core Network Assistance Information(optional)
		// ie = ngapType.HandoverRequestIEs{}
		// ie.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
		// ie.Criticality.Value = ngapType.CriticalityPresentReject
		// ie.Value.Present = ngapType.HandoverRequestIEsPresentCoreNetworkAssistanceInformation
		// ie.Value.CoreNetworkAssistanceInformation = new(ngapType.CoreNetworkAssistanceInformation)
		// handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// New Security ContextInd(optional)
		if nsci {
			ie = ngapType.HandoverRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDNewSecurityContextInd
			ie.Criticality.Value = ngapType.CriticalityPresentReject
			ie.Value.Present = ngapType.HandoverRequestIEsPresentNewSecurityContextInd
			ie.Value.NewSecurityContextInd = new(ngapType.NewSecurityContextInd)
			ie.Value.NewSecurityContextInd.Value = ngapType.NewSecurityContextIndPresentTrue
			handoverRequestIEs.List = append(handoverRequestIEs.List, ie)
		}

		// NASC(optional)
		// ie.Criticality.Value = ngapType.CriticalityPresentReject
		// ie.Value.Present = ngapType.HandoverRequestIEsPresentNASC
		// ie.Id.Value = ngapType.ProtocolIEIDNASC
		// ie.Criticality.Value = ngapType.CriticalityPresentReject
		// ie.Value.Present = ngapType.HandoverRequestIEsPresentNASC
		// ie.Value.NASC = new(ngapType.)
		// handoverRequestIEs.List = append(handoverRequestIEs.List, ie)

		// Trace Activation(optional)
		// Masked IMEISV(optional)
		// Mobility Restriction List(optional)
		// Location Reporting Request Type(optional)
		// RRC Inactive Transition Report Reques(optional)
		return libngap.Encoder(pdu)
	*/
	return nil, nil
}

func (h *Ngap) handleHandoverRequestAcknowledge(ran *ran.Ran /*message *ngapType.NGAPPDU*/, handoverRequestAcknowledge *ngapType.HandoverRequestAcknowledge) {
	/*
		var coreNgapId *ngapType.AMFUENGAPID
		var ranNgapId *ngapType.RANUENGAPID
		var pDUSessionResourceAdmittedList *ngapType.PDUSessionResourceAdmittedList
		var pDUSessionResourceFailedToSetupListHOAck *ngapType.PDUSessionResourceFailedToSetupListHOAck
		var targetToSourceTransparentContainer *ngapType.TargetToSourceTransparentContainer
		var criticalityDiagnostics *ngapType.CriticalityDiagnostics

		var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList
	*/
	/*
		successfulOutcome := message.SuccessfulOutcome
		handoverRequestAcknowledge := successfulOutcome.Value.HandoverRequestAcknowledge // reject
		if handoverRequestAcknowledge == nil {
			log.Error("HandoverRequestAcknowledge is nil")
			return
		}
	*/
	/*
		for _, ie := range handoverRequestAcknowledge.ProtocolIEs.List {
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
				coreNgapId = ie.Value.AMFUENGAPID
				log.Trace("Decode IE AmfUeNgapID")
			case ngapType.ProtocolIEIDRANUENGAPID: // ignore
				ranNgapId = ie.Value.RANUENGAPID
				log.Trace("Decode IE UeContextNgapID")
			case ngapType.ProtocolIEIDPDUSessionResourceAdmittedList: // ignore
				pDUSessionResourceAdmittedList = ie.Value.PDUSessionResourceAdmittedList
				log.Trace("Decode IE PduSessionResourceAdmittedList")
			case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck: // ignore
				pDUSessionResourceFailedToSetupListHOAck = ie.Value.PDUSessionResourceFailedToSetupListHOAck
				log.Trace("Decode IE PduSessionResourceFailedToSetupListHOAck")
			case ngapType.ProtocolIEIDTargetToSourceTransparentContainer: // reject
				targetToSourceTransparentContainer = ie.Value.TargetToSourceTransparentContainer
				log.Trace("Decode IE TargetToSourceTransparentContainer")
			case ngapType.ProtocolIEIDCriticalityDiagnostics: // ignore
				criticalityDiagnostics = ie.Value.CriticalityDiagnostics
				log.Trace("Decode IE CriticalityDiagnostics")
			}
		}
		if targetToSourceTransparentContainer == nil {
			log.Error("TargetToSourceTransparentContainer is nil")
			item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
				ngapType.ProtocolIEIDTargetToSourceTransparentContainer, ngapType.TypeOfErrorPresentMissing)
			iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
		}
		if len(iesCriticalityDiagnostics.List) > 0 {
			log.Error("Has missing reject IE(s)")

			procedureCode := ngapType.ProcedureCodeHandoverResourceAllocation
			triggeringMessage := ngapType.TriggeringMessagePresentSuccessfulOutcome
			procedureCriticality := ngapType.CriticalityPresentReject
			criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage,
				&procedureCriticality, &iesCriticalityDiagnostics)
			SendErrorIndication(ran, coreNgapId, ranNgapId, nil, &criticalityDiagnostics)
		}

		if criticalityDiagnostics != nil {
			printCriticalityDiagnostics(ran, criticalityDiagnostics)
		}

		targetUe := h.backend.Context().UeContextFindByAmfUeNgapID(coreNgapId.Value)
		if targetUe == nil {
			log.Errorf("No UE Context[AMFUENGAPID: %d]", coreNgapId.Value)
			return
		}

		log.Info("Handle Handover Request Acknowledge")

		if ranNgapId != nil {
			targetUe.SetUeContextNgapId(ranNgapId.Value)
			//		targetUe.UpdateLogFields()
		}
		log.Debugf("Target Ue UeContextNgapID[%d] AmfUeNgapID[%d]", targetUe.UeContextNgapId(), targetUe.AmfUeNgapId())

		amfUe := targetUe.AmfUe()
		if amfUe == nil {
			log.Error("amfUe is nil")
			return
		}

		var pduSessionResourceHandoverList ngapType.PDUSessionResourceHandoverList
		var pduSessionResourceToReleaseList ngapType.PDUSessionResourceToReleaseListHOCmd

		// describe in 23.502 4.9.1.3.2 step11
		if pDUSessionResourceAdmittedList != nil {
			for _, item := range pDUSessionResourceAdmittedList.List {
				pduSessionID := item.PDUSessionID.Value
				transfer := item.HandoverRequestAcknowledgeTransfer
				pduSessionId := int32(pduSessionID)
				if smContext, exist := amfUe.SmContextFindByPDUSessionID(pduSessionId); exist {
					response, errResponse, problemDetails, err := smContext.SmfClient().SendUpdateSmContextN2HandoverPrepared(models.N2SMINFOTYPE_HANDOVER_REQ_ACK, transfer)
					if err != nil {
						log.Errorf("Send HandoverRequestAcknowledgeTransfer error: %v", err)
					}
					if problemDetails != nil {
						log.Warnf("ProblemDetails[status: %d, Cause: %s]", problemDetails.Status, problemDetails.Cause)
					}
					if response != nil && response.BinaryDataN2SmInformation != nil {
						handoverItem := ngapType.PDUSessionResourceHandoverItem{}
						handoverItem.PDUSessionID = item.PDUSessionID
						handoverItem.HandoverCommandTransfer = response.BinaryDataN2SmInformation
						pduSessionResourceHandoverList.List = append(pduSessionResourceHandoverList.List, handoverItem)
						targetUe.GetHandoverInfo().SuccessPduSessionId = append(targetUe.GetHandoverInfo().SuccessPduSessionId, pduSessionId)
					}
					if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
						releaseItem := ngapType.PDUSessionResourceToReleaseItemHOCmd{}
						releaseItem.PDUSessionID = item.PDUSessionID
						releaseItem.HandoverPreparationUnsuccessfulTransfer = errResponse.BinaryDataN2SmInformation
						pduSessionResourceToReleaseList.List = append(pduSessionResourceToReleaseList.List, releaseItem)
					}
				}
			}
		}

		if pDUSessionResourceFailedToSetupListHOAck != nil {
			for _, item := range pDUSessionResourceFailedToSetupListHOAck.List {
				pduSessionID := item.PDUSessionID.Value
				transfer := item.HandoverResourceAllocationUnsuccessfulTransfer
				pduSessionId := int32(pduSessionID)
				if smContext, exist := amfUe.SmContextFindByPDUSessionID(pduSessionId); exist {
					_, _, problemDetails, err := smContext.SmfClient().SendUpdateSmContextN2HandoverPrepared(models.N2SMINFOTYPE_HANDOVER_RES_ALLOC_FAIL, transfer)
					if err != nil {
						log.Errorf("Send HandoverResourceAllocationUnsuccessfulTransfer error: %v", err)
					}
					if problemDetails != nil {
						log.Warnf("ProblemDetails[status: %d, Cause: %s]", problemDetails.Status, problemDetails.Cause)
					}
				}
			}
		}

		sourceUe := targetUe.GetHandoverInfo().SourceUe
		if sourceUe == nil {
			// TODO: Send Namf_Communication_CreateUEContext Response to S-AMF
			log.Error("handover between different Ue has not been implement yet")
		} else {
			log.Tracef("Source: UeContextNgapID[%d] AmfUeNgapID[%d]", sourceUe.UeContextNgapId(), sourceUe.AmfUeNgapId())
			log.Tracef("Target: UeContextNgapID[%d] AmfUeNgapID[%d]", targetUe.UeContextNgapId(), targetUe.AmfUeNgapId())
			if len(pduSessionResourceHandoverList.List) == 0 {
				log.Info("Handle Handover Preparation Failure [HoFailure In Target5GC NgranNode Or TargetSystem]")
				cause := &ngapType.Cause{
					Present: ngapType.CausePresentRadioNetwork,
					RadioNetwork: &ngapType.CauseRadioNetwork{
						Value: ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem,
					},
				}
				SendHandoverPreparationFailure(sourceUe, *cause, nil)
				return
			}
			SendHandoverCommand(sourceUe, pduSessionResourceHandoverList, pduSessionResourceToReleaseList,
				*targetToSourceTransparentContainer, nil)
		}
	*/
}

func (h *Ngap) handleHandoverFailure(ran *ran.Ran, handoverFailure *ngapType.HandoverFailure) {
	/*
		var (
			coreNgapId *ngapType.AMFUENGAPID
			cause      *ngapType.Cause
			critical   *ngapType.CriticalDiagnotics
		)
		for _, ie := range handoverFailure.ProtocolIEs.List {
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
				coreNgapId = ie.Value.AMFUENGAPID
				log.Trace("Decode IE AmfUeNgapID")
			case ngapType.ProtocolIEIDCause: // ignore
				Cause = ie.Value.Cause
				log.Trace("Decode IE Cause")
			case ngapType.ProtocolIEIDCriticalityDiagnostics: // ignore
				critical = ie.Value.CriticalityDiagnostics
				log.Trace("Decode IE CriticalityDiagnostics")
			}
		}

		causePresent := ngapType.CausePresentRadioNetwork
		causeValue := ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem
		if cause != nil {
			causePresent, causeValue = printAndGetCause(ran, cause)
		}

		if criticalityDiagnostics != nil {
			printCriticalityDiagnostics(ran, criticalityDiagnostics)
		}
		ue = ran.FindUe(nil, coreNgapId)

		if targetUe == nil {
			log.Errorf("No UE Context[AmfUENGAPID: %d]", coreNgapId.Value)
			cause := ngapType.Cause{
				Present: ngapType.CausePresentRadioNetwork,
				RadioNetwork: &ngapType.CauseRadioNetwork{
					Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
				},
			}
			SendErrorIndication(ran, coreNgapId, nil, &cause, nil)
			return
		}

		log.Info("Handle Handover Failure")

		sourceUe := targetUe.GetHandoverInfo().SourceUe
		if sourceUe == nil {
			// TODO: handle N2 Handover between AMF
			log.Error("N2 Handover between AMF has not been implemented yet")
		} else {
			amfUe := targetUe.AmfUe()
			if amfUe != nil {
				amfUe.SmContextList.Range(func(key, value interface{}) bool {
					pduSessionID := key.(int32)
					smContext := value.(*context.SmContext)
					causeAll := context.CauseAll{
						NgapCause: &models.NgApCause{
							Group: int32(causePresent),
							Value: int32(causeValue),
						},
					}
					_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverCanceled(causeAll)
					if err != nil {
						log.Errorf("Send UpdateSmContextN2HandoverCanceled Error for PduSessionId[%d]", pduSessionID)
					}
					return true
				})
			}
		//	SendHandoverPreparationFailure(sourceUe, *cause, criticalityDiagnostics)
		}

		//			SendUEContextReleaseCommand(targetUe, ue.UeContextReleaseHandover, causePresent, causeValue)
	*/
}

func (h *Ngap) handleHandoverNotify(ran *ran.Ran, HandoverNotify *ngapType.HandoverNotify) {
	/*
		var coreNgapId *ngapType.AMFUENGAPID
		var ranNgapId *ngapType.RANUENGAPID
		var userLocationInformation *ngapType.UserLocationInformation

		for i := 0; i < len(HandoverNotify.ProtocolIEs.List); i++ {
			ie := HandoverNotify.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				coreNgapId = ie.Value.AMFUENGAPID
				log.Trace("Decode IE AmfUeNgapID")
				if coreNgapId == nil {
					log.Error("AMFUENGAPID is nil")
					return
				}
			case ngapType.ProtocolIEIDRANUENGAPID:
				ranNgapId = ie.Value.RANUENGAPID
				log.Trace("Decode IE UeContextNgapID")
				if ranNgapId == nil {
					log.Error("RANUENGAPID is nil")
					return
				}
			case ngapType.ProtocolIEIDUserLocationInformation:
				userLocationInformation = ie.Value.UserLocationInformation
				log.Trace("Decode IE userLocationInformation")
				if userLocationInformation == nil {
					log.Error("userLocationInformation is nil")
					return
				}
			}
		}

		targetUe := ran.UeContextFindByUeContextNgapID(ranNgapId.Value)

		if targetUe == nil {
			log.Errorf("No UeContext Context[AmfUeNgapID: %d]", coreNgapId.Value)
			cause := ngapType.Cause{
				Present: ngapType.CausePresentRadioNetwork,
				RadioNetwork: &ngapType.CauseRadioNetwork{
					Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
				},
			}
			SendErrorIndication(ran, coreNgapId, ranNgapId, &cause, nil)
			return
		}

		//	targetUe.Log.Info("Handle Handover notification")

		if userLocationInformation != nil {
			targetUe.UpdateLocation(userLocationInformation)
		}
		amfUe := targetUe.AmfUe()
		if amfUe == nil {
			log.Error("AmfUe is nil")
			return
		}
		sourceUe := targetUe.GetHandoverInfo().SourceUe
		if sourceUe == nil {
			// TODO: Send to S-AMF
			// Desciibed in (23.502 4.9.1.3.3) [conditional] 6a.Namf_Communication_N2InfoNotify.
			log.Error("N2 Handover between AMF has not been implemented yet")
		} else {
			//		targetUe.Log.Info("Handle Handover notification Finshed ")
			for _, pduSessionid := range targetUe.GetHandoverInfo().SuccessPduSessionId {
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionid)
				if !ok {
					//				sourceUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionid)
				}
				_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverComplete("", nil)
				if err != nil {
					log.Errorf("Send UpdateSmContextN2HandoverComplete Error[%s]", err.Error())
				}
			}
			amfUe.AttachUeContext(targetUe)

			SendUEContextReleaseCommand(sourceUe, ue.UeContextReleaseHandover, ngapType.CausePresentNas,
				ngapType.CauseNasPresentNormalRelease)
		}

		// TODO: The UE initiates Mobility Registration Update procedure as described in clause 4.2.2.2.2.
	*/
}

// TS 23.502 4.9.1
func (h *Ngap) handlePathSwitchRequest(ran *ran.Ran, pathSwitchRequest *ngapType.PathSwitchRequest) {
	/*
		var ranNgapId *ngapType.RANUENGAPID
		var sourceAMFUENGAPID *ngapType.AMFUENGAPID
		var userLocationInformation *ngapType.UserLocationInformation
		var uESecurityCapabilities *ngapType.UESecurityCapabilities
		var pduSessionResourceToBeSwitchedInDLList *ngapType.PDUSessionResourceToBeSwitchedDLList
		var pduSessionResourceFailedToSetupList *ngapType.PDUSessionResourceFailedToSetupListPSReq

		var ue *ue.UeContext

		for _, ie := range pathSwitchRequest.ProtocolIEs.List {
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDRANUENGAPID: // reject
				ranNgapId = ie.Value.RANUENGAPID
				log.Trace("Decode IE UeContextNgapID")
				if ranNgapId == nil {
					log.Error("UeContextNgapID is nil")
					return
				}
			case ngapType.ProtocolIEIDSourceAMFUENGAPID: // reject
				sourceAMFUENGAPID = ie.Value.SourceAMFUENGAPID
				log.Trace("Decode IE SourceAmfUeNgapID")
				if sourceAMFUENGAPID == nil {
					log.Error("SourceAmfUeNgapID is nil")
					return
				}
			case ngapType.ProtocolIEIDUserLocationInformation: // ignore
				userLocationInformation = ie.Value.UserLocationInformation
				log.Trace("Decode IE UserLocationInformation")
			case ngapType.ProtocolIEIDUESecurityCapabilities: // ignore
				uESecurityCapabilities = ie.Value.UESecurityCapabilities
				log.Trace("Decode IE UESecurityCapabilities")
			case ngapType.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList: // reject
				pduSessionResourceToBeSwitchedInDLList = ie.Value.PDUSessionResourceToBeSwitchedDLList
				log.Trace("Decode IE PDUSessionResourceToBeSwitchedDLList")
				if pduSessionResourceToBeSwitchedInDLList == nil {
					log.Error("PDUSessionResourceToBeSwitchedDLList is nil")
					return
				}
			case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq: // ignore
				pduSessionResourceFailedToSetupList = ie.Value.PDUSessionResourceFailedToSetupListPSReq
				log.Trace("Decode IE PDUSessionResourceFailedToSetupListPSReq")
			}
		}

		if sourceAMFUENGAPID == nil {
			log.Error("SourceAmfUeNgapID is nil")
			return
		}
		ue = h.backend.Context().UeContextFindByAmfUeNgapID(sourceAMFUENGAPID.Value)
		if ue == nil {
			log.Errorf("Cannot find UE from sourceAMfUeNgapID[%d]", sourceAMFUENGAPID.Value)
			SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ranNgapId.Value, nil, nil)
			return
		}

		log.Tracef("AmfUeNgapID[%d] UeContextNgapID[%d]", ue.AmfUeNgapId(), ue.UeContextNgapId())
		log.Info("Handle Path Switch Request")

		amfUe := ue.AmfUe()
		ausf := amfUe.AusfClient()

		if amfUe == nil {
			log.Error("AmfUe is nil")
			SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ranNgapId.Value, nil, nil)
			return
		}

		if ausf.SecurityContextIsValid() {
			// Update NH
			ausf.UpdateNH()
		} else {
			log.Errorf("No Security Context : SUPI[%s]", amfUe.Supi)
			SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ranNgapId.Value, nil, nil)
			return
		}

		ausf.SetUeSecCap(uESecurityCapabilities)

		if ranNgapId != nil {
			ue.SetUeContextNgapId(ranNgapId.Value)
		}

		ue.UpdateLocation(userLocationInformation)

		var pduSessionResourceSwitchedList ngapType.PDUSessionResourceSwitchedList
		var pduSessionResourceReleasedListPSAck ngapType.PDUSessionResourceReleasedListPSAck
		var pduSessionResourceReleasedListPSFail ngapType.PDUSessionResourceReleasedListPSFail

		if pduSessionResourceToBeSwitchedInDLList != nil {
			for _, item := range pduSessionResourceToBeSwitchedInDLList.List {
				pduSessionID := int32(item.PDUSessionID.Value)
				transfer := item.PathSwitchRequestTransfer
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
					log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				}
				response, errResponse, _, err := smContext.SmfClient().SendUpdateSmContextXnHandover(models.N2SMINFOTYPE_PATH_SWITCH_REQ, transfer)
				if err != nil {
					log.Errorf("SendUpdateSmContextXnHandover[PathSwitchRequestTransfer] Error:\n%s", err.Error())
				}
				if response != nil && response.BinaryDataN2SmInformation != nil {
					pduSessionResourceSwitchedItem := ngapType.PDUSessionResourceSwitchedItem{}
					pduSessionResourceSwitchedItem.PDUSessionID.Value = int64(pduSessionID)
					pduSessionResourceSwitchedItem.PathSwitchRequestAcknowledgeTransfer = response.BinaryDataN2SmInformation
					pduSessionResourceSwitchedList.List = append(pduSessionResourceSwitchedList.List, pduSessionResourceSwitchedItem)
				}
				if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
					pduSessionResourceReleasedItem := ngapType.PDUSessionResourceReleasedItemPSFail{}
					pduSessionResourceReleasedItem.PDUSessionID.Value = int64(pduSessionID)
					pduSessionResourceReleasedItem.PathSwitchRequestUnsuccessfulTransfer = errResponse.BinaryDataN2SmInformation
					pduSessionResourceReleasedListPSFail.List = append(pduSessionResourceReleasedListPSFail.List,
						pduSessionResourceReleasedItem)
				}
			}
		}

		if pduSessionResourceFailedToSetupList != nil {
			for _, item := range pduSessionResourceFailedToSetupList.List {
				pduSessionID := int32(item.PDUSessionID.Value)
				transfer := item.PathSwitchRequestSetupFailedTransfer
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
					log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				}
				response, errResponse, _, err := smContext.SmfClient().SendUpdateSmContextXnHandoverFailed(models.N2SMINFOTYPE_PATH_SWITCH_SETUP_FAIL, transfer)
				if err != nil {
					log.Errorf("SendUpdateSmContextXnHandoverFailed[PathSwitchRequestSetupFailedTransfer] Error: %+v", err)
				}
				if response != nil && response.BinaryDataN2SmInformation != nil {
					pduSessionResourceReleasedItem := ngapType.PDUSessionResourceReleasedItemPSAck{}
					pduSessionResourceReleasedItem.PDUSessionID.Value = int64(pduSessionID)
					pduSessionResourceReleasedItem.PathSwitchRequestUnsuccessfulTransfer = response.BinaryDataN2SmInformation
					pduSessionResourceReleasedListPSAck.List = append(pduSessionResourceReleasedListPSAck.List,
						pduSessionResourceReleasedItem)
				}
				if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
					pduSessionResourceReleasedItem := ngapType.PDUSessionResourceReleasedItemPSFail{}
					pduSessionResourceReleasedItem.PDUSessionID.Value = int64(pduSessionID)
					pduSessionResourceReleasedItem.PathSwitchRequestUnsuccessfulTransfer = errResponse.BinaryDataN2SmInformation
					pduSessionResourceReleasedListPSFail.List = append(pduSessionResourceReleasedListPSFail.List,
						pduSessionResourceReleasedItem)
				}
			}
		}

		// TS 23.502 4.9.1.2.2 step 7: send ack to Target NG-RAN. If none of the requested PDU Sessions have been switched
		// successfully, the AMF shall send an N2 Path Switch Request Failure message to the Target NG-RAN
		if len(pduSessionResourceSwitchedList.List) > 0 {
			// TODO: set newSecurityContextIndicator to true if there is a new security context
			err := ue.SwitchToRan(ran, ranNgapId.Value)
			if err != nil {
				log.Error(err.Error())
				return
			}
			SendPathSwitchRequestAcknowledge(ue, pduSessionResourceSwitchedList,
				pduSessionResourceReleasedListPSAck, false, nil, nil, nil)
		} else if len(pduSessionResourceReleasedListPSFail.List) > 0 {
			SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ranNgapId.Value,
				&pduSessionResourceReleasedListPSFail, nil)
		} else {
			SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ranNgapId.Value, nil, nil)
		}
	*/
}

// pduSessionResourceSwitchedList: provided by AMF, and the transfer data is from SMF
// pduSessionResourceReleasedList: provided by AMF, and the transfer data is from SMF
// newSecurityContextIndicator: if AMF has activated a new 5G NAS security context, set it to true,
// otherwise set to false
// coreNetworkAssistanceInformation: provided by AMF, based on collection of UE behaviour statistics
// and/or other available
// information about the expected UE behaviour. TS 23.501 5.4.6, 5.4.6.2
// rrcInactiveTransitionReportRequest: configured by amf
// criticalityDiagnostics: from received node when received not comprehended IE or missing IE
func SendPathSwitchRequestAcknowledge(
	ue *ue.UeContext,
	pduSessionResourceSwitchedList ngapType.PDUSessionResourceSwitchedList,
	pduSessionResourceReleasedList ngapType.PDUSessionResourceReleasedListPSAck,
	newSecurityContextIndicator bool,
	coreNetworkAssistanceInformation *ngapType.CoreNetworkAssistanceInformation,
	rrcInactiveTransitionReportRequest *ngapType.RRCInactiveTransitionReportRequest,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {
	/*
	   log.Info("Send Path Switch Request Acknowledge")

	   	if len(pduSessionResourceSwitchedList.List) > context.MaxNumOfPDUSessions {
	   		log.Error("Pdu List out of range")
	   		return
	   	}

	   	if len(pduSessionResourceReleasedList.List) > context.MaxNumOfPDUSessions {
	   		log.Error("Pdu List out of range")
	   		return
	   	}

	   pkt, err := s.buildPathSwitchRequestAcknowledge(ue, pduSessionResourceSwitchedList, pduSessionResourceReleasedList,

	   	newSecurityContextIndicator, coreNetworkAssistanceInformation, rrcInactiveTransitionReportRequest,
	   	criticalityDiagnostics)

	   	if err != nil {
	   		log.Errorf("Build PathSwitchRequestAcknowledge failed : %s", err.Error())
	   		return
	   	}

	   SendToUeContext(ue, pkt)
	*/
	return
}

func buildPathSwitchRequestAcknowledge(
	ue *ue.UeContext,
	pduSessionResourceSwitchedList ngapType.PDUSessionResourceSwitchedList,
	pduSessionResourceReleasedList ngapType.PDUSessionResourceReleasedListPSAck,
	newSecurityContextIndicator bool,
	coreNetworkAssistanceInformation *ngapType.CoreNetworkAssistanceInformation,
	rrcInactiveTransitionReportRequest *ngapType.RRCInactiveTransitionReportRequest,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) ([]byte, error) {
	/*
		amf := s.backend.Context()

		var pdu ngapType.NGAPPDU
		pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
		pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

		successfulOutcome := pdu.SuccessfulOutcome
		successfulOutcome.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
		successfulOutcome.Criticality.Value = ngapType.CriticalityPresentReject

		successfulOutcome.Value.Present = ngapType.SuccessfulOutcomePresentPathSwitchRequestAcknowledge
		successfulOutcome.Value.PathSwitchRequestAcknowledge = new(ngapType.PathSwitchRequestAcknowledge)

		pathSwitchRequestAck := successfulOutcome.Value.PathSwitchRequestAcknowledge
		pathSwitchRequestAckIEs := &pathSwitchRequestAck.ProtocolIEs

		// AMF UE NGAP ID
		ie := ngapType.PathSwitchRequestAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentAMFUENGAPID
		ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

		aMFUENGAPID := ie.Value.AMFUENGAPID
		aMFUENGAPID.Value = ue.CuNgapId()

		pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)

		// RAN UE NGAP ID
		ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentRANUENGAPID
		ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

		rANUENGAPID := ie.Value.RANUENGAPID
		rANUENGAPID.Value = ue.RanNgapId()

		pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)

		// UE Security Capabilities (optional)
		ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentUESecurityCapabilities
		ie.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)

		secinfo := ue.AmfUe().AusfClient().SecInfo()
		ueSecurityCapabilities := ie.Value.UESecurityCapabilities
		nrEncryptionAlgorighm := []byte{0x00, 0x00}
		nrEncryptionAlgorighm[0] |= secinfo.UESecurityCapability.GetEA1_128_5G() << 7
		nrEncryptionAlgorighm[0] |= secinfo.UESecurityCapability.GetEA2_128_5G() << 6
		nrEncryptionAlgorighm[0] |= secinfo.UESecurityCapability.GetEA3_128_5G() << 5
		ueSecurityCapabilities.NRencryptionAlgorithms.Value = ngapConvert.ByteToBitString(nrEncryptionAlgorighm, 16)

		nrIntegrityAlgorithm := []byte{0x00, 0x00}
		nrIntegrityAlgorithm[0] |= secinfo.UESecurityCapability.GetIA1_128_5G() << 7
		nrIntegrityAlgorithm[0] |= secinfo.UESecurityCapability.GetIA2_128_5G() << 6
		nrIntegrityAlgorithm[0] |= secinfo.UESecurityCapability.GetIA3_128_5G() << 5
		ueSecurityCapabilities.NRintegrityProtectionAlgorithms.Value =
			ngapConvert.ByteToBitString(nrIntegrityAlgorithm, 16)

		// only support NR algorithms
		eutraEncryptionAlgorithm := []byte{0x00, 0x00}
		ueSecurityCapabilities.EUTRAencryptionAlgorithms.Value =
			ngapConvert.ByteToBitString(eutraEncryptionAlgorithm, 16)

		eutraIntegrityAlgorithm := []byte{0x00, 0x00}
		ueSecurityCapabilities.EUTRAintegrityProtectionAlgorithms.Value =
			ngapConvert.ByteToBitString(eutraIntegrityAlgorithm, 16)

		pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)

		// Security Context
		ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDSecurityContext
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentSecurityContext
		ie.Value.SecurityContext = new(ngapType.SecurityContext)

		securityContext := ie.Value.SecurityContext
		securityContext.NextHopChainingCount.Value = int64(secinfo.NCC)
		securityContext.NextHopNH.Value = ngapConvert.HexToBitString(hex.EncodeToString(secinfo.NH), 256)

		pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)

		// New Security Context Indicator (optional)
		if newSecurityContextIndicator {
			ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDNewSecurityContextInd
			ie.Criticality.Value = ngapType.CriticalityPresentReject
			ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentNewSecurityContextInd
			ie.Value.NewSecurityContextInd = new(ngapType.NewSecurityContextInd)
			ie.Value.NewSecurityContextInd.Value = ngapType.NewSecurityContextIndPresentTrue
			pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)
		}

		// PDU Session Resource Switched List
		ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSwitchedList
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceSwitchedList
		ie.Value.PDUSessionResourceSwitchedList = &pduSessionResourceSwitchedList
		pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)

		// PDU Session Resource Released List
		if len(pduSessionResourceReleasedList.List) > 0 {
			ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSAck
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceReleasedListPSAck
			ie.Value.PDUSessionResourceReleasedListPSAck = &pduSessionResourceReleasedList
			pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)
		}

		// Allowed NSSAI
		ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentAllowedNSSAI
		ie.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)

		allowedNSSAI := ie.Value.AllowedNSSAI
		// plmnSupportList[0] is serving plmn
		for _, modelSnssai := range amf.PlmnSupportList()[0].SNssaiList {
			allowedNSSAIItem := ngapType.AllowedNSSAIItem{}

			ngapSnssai := ngapConvert.SNssaiToNgap(modelSnssai)
			allowedNSSAIItem.SNSSAI = ngapSnssai
			allowedNSSAI.List = append(allowedNSSAI.List, allowedNSSAIItem)
		}
		pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)

		// Core Network Assistance Information (optional)
		if coreNetworkAssistanceInformation != nil {
			ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentCoreNetworkAssistanceInformation
			ie.Value.CoreNetworkAssistanceInformation = coreNetworkAssistanceInformation
			pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)
		}

		// RRC Inactive Transition Report Request (optional)
		if rrcInactiveTransitionReportRequest != nil {
			ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentRRCInactiveTransitionReportRequest
			ie.Value.RRCInactiveTransitionReportRequest = rrcInactiveTransitionReportRequest
			pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)
		}

		// Criticality Diagnostics (optional)
		if criticalityDiagnostics != nil {
			ie = ngapType.PathSwitchRequestAcknowledgeIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentCriticalityDiagnostics
			ie.Value.CriticalityDiagnostics = criticalityDiagnostics
			pathSwitchRequestAckIEs.List = append(pathSwitchRequestAckIEs.List, ie)
		}

		return libngap.Encoder(pdu)
	*/
	return nil, nil
}

// pduSessionResourceReleasedList: provided by AMF, and the transfer data is from SMF
// criticalityDiagnostics: from received node when received not comprehended IE or missing IE
func SendPathSwitchRequestFailure(
	ran *ran.Ran,
	amfUeNgapId,
	ueNgapId int64,
	pduSessionResourceReleasedList *ngapType.PDUSessionResourceReleasedListPSFail,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {
	/*
		log.Info("Send Path Switch Request Failure")

		if pduSessionResourceReleasedList != nil && len(pduSessionResourceReleasedList.List) > context.MaxNumOfPDUSessions {
			log.Error("Pdu List out of range")
			return
		}

		pkt, err := s.buildPathSwitchRequestFailure(amfUeNgapId, ueNgapId, pduSessionResourceReleasedList,
			criticalityDiagnostics)
		if err != nil {
			log.Errorf("Build PathSwitchRequestFailure failed : %s", err.Error())
			return
		}
		SendToRan(ran, pkt)
	*/
	return
}

func buildPathSwitchRequestFailure(
	amfUeNgapId,
	ueNgapId int64,
	pduSessionResourceReleasedList *ngapType.PDUSessionResourceReleasedListPSFail,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) ([]byte, error) {
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)

	unsuccessfulOutcome := pdu.UnsuccessfulOutcome
	unsuccessfulOutcome.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	unsuccessfulOutcome.Criticality.Value = ngapType.CriticalityPresentReject

	unsuccessfulOutcome.Value.Present = ngapType.UnsuccessfulOutcomePresentPathSwitchRequestFailure
	unsuccessfulOutcome.Value.PathSwitchRequestFailure = new(ngapType.PathSwitchRequestFailure)

	pathSwitchRequestFailure := unsuccessfulOutcome.Value.PathSwitchRequestFailure
	pathSwitchRequestFailureIEs := &pathSwitchRequestFailure.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.PathSwitchRequestFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = amfUeNgapId

	pathSwitchRequestFailureIEs.List = append(pathSwitchRequestFailureIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.PathSwitchRequestFailureIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ueNgapId

	pathSwitchRequestFailureIEs.List = append(pathSwitchRequestFailureIEs.List, ie)

	// PDU Session Resource Released List
	if pduSessionResourceReleasedList != nil {
		ie = ngapType.PathSwitchRequestFailureIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSFail
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentPDUSessionResourceReleasedListPSFail
		ie.Value.PDUSessionResourceReleasedListPSFail = pduSessionResourceReleasedList
		pathSwitchRequestFailureIEs.List = append(pathSwitchRequestFailureIEs.List, ie)
	}

	if criticalityDiagnostics != nil {
		ie = ngapType.PathSwitchRequestFailureIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = criticalityDiagnostics
		pathSwitchRequestFailureIEs.List = append(pathSwitchRequestFailureIEs.List, ie)
	}

	return libngap.Encoder(pdu)
}

func (h *Ngap) handleHandoverCancel(ran *ran.Ran, HandoverCancel *ngapType.HandoverCancel) {
	/*
		var coreNgapId *ngapType.AMFUENGAPID
		var ranNgapId *ngapType.RANUENGAPID
		var cause *ngapType.Cause

		for i := 0; i < len(HandoverCancel.ProtocolIEs.List); i++ {
			ie := HandoverCancel.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				coreNgapId = ie.Value.AMFUENGAPID
				log.Trace("Decode IE AmfUeNgapID")
				if coreNgapId == nil {
					log.Error("AMFUENGAPID is nil")
					return
				}
			case ngapType.ProtocolIEIDRANUENGAPID:
				ranNgapId = ie.Value.RANUENGAPID
				log.Trace("Decode IE UeContextNgapID")
				if ranNgapId == nil {
					log.Error("RANUENGAPID is nil")
					return
				}
			case ngapType.ProtocolIEIDCause:
				cause = ie.Value.Cause
				log.Trace("Decode IE Cause")
				if cause == nil {
					log.Error(cause, "cause is nil")
					return
				}
			}
		}

		sourceUe := ran.UeContextFindByUeContextNgapID(ranNgapId.Value)
		if sourceUe == nil {
			log.Errorf("No UE Context[UeContextNgapID: %d]", ranNgapId.Value)
			cause := ngapType.Cause{
				Present: ngapType.CausePresentRadioNetwork,
				RadioNetwork: &ngapType.CauseRadioNetwork{
					Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
				},
			}
			SendErrorIndication(ran, coreNgapId, ranNgapId, &cause, nil)
			return
		}

		//	sourceUe.Log.Info("Handle Handover Cancel")

		if sourceUe.AmfUeNgapId() != coreNgapId.Value {
			log.Warnf("Conflict AMF_UE_NGAP_ID : %d != %d", sourceUe.AmfUeNgapId(), coreNgapId.Value)
		}
		log.Tracef("Source: RAN_UE_NGAP_ID[%d] AMF_UE_NGAP_ID[%d]", sourceUe.UeContextNgapId(), sourceUe.AmfUeNgapId())

		causePresent := ngapType.CausePresentRadioNetwork
		causeValue := ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem
		if cause != nil {
			causePresent, causeValue = printAndGetCause(ran, cause)
		}
		targetUe := sourceUe.GetHandoverInfo().TargetUe
		if targetUe == nil {
			// Described in (23.502 4.11.1.2.3) step 2
			// Todo : send to T-AMF invoke Namf_UeContextReleaseRequest(targetUe)
			log.Error("N2 Handover between AMF has not been implemented yet")
		} else {
			log.Tracef("Target : RAN_UE_NGAP_ID[%d] AMF_UE_NGAP_ID[%d]", targetUe.UeContextNgapId(), targetUe.AmfUeNgapId())
			amfUe := sourceUe.AmfUe()
			if amfUe != nil {
				amfUe.SmContextList.Range(func(key, value interface{}) bool {
					//pduSessionID := key.(int32)
					smContext := value.(*context.SmContext)
					causeAll := context.CauseAll{
						NgapCause: &models.NgApCause{
							Group: int32(causePresent),
							Value: int32(causeValue),
						},
					}
					_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverCanceled(causeAll)
					if err != nil {
						//					sourceUe.Log.Errorf("Send UpdateSmContextN2HandoverCanceled Error for PduSessionId[%d]", pduSessionID)
					}
					return true
				})
			}
			SendUEContextReleaseCommand(targetUe, ue.UeContextReleaseHandover, causePresent, causeValue)
			SendHandoverCancelAcknowledge(sourceUe, nil)
		}
	*/
}

func SendHandoverCancelAcknowledge(ue *ue.UeContext, criticalityDiagnostics *ngapType.CriticalityDiagnostics) (err error) {

	log.Info("Send Handover Cancel Acknowledge")

	var pkt []byte
	if pkt, err = buildHandoverCancelAcknowledge(ue, criticalityDiagnostics); err != nil {
		log.Errorf("Build HandoverCancelAcknowledge failed : %s", err.Error())
		return
	}
	err = ue.Send(pkt)
	return
}
func buildHandoverCancelAcknowledge(
	ue *ue.UeContext, criticalityDiagnostics *ngapType.CriticalityDiagnostics) ([]byte, error) {
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

	successfulOutcome := pdu.SuccessfulOutcome
	successfulOutcome.ProcedureCode.Value = ngapType.ProcedureCodeHandoverCancel
	successfulOutcome.Criticality.Value = ngapType.CriticalityPresentReject
	successfulOutcome.Value.Present = ngapType.SuccessfulOutcomePresentHandoverCancelAcknowledge
	successfulOutcome.Value.HandoverCancelAcknowledge = new(ngapType.HandoverCancelAcknowledge)

	handoverCancelAcknowledge := successfulOutcome.Value.HandoverCancelAcknowledge
	handoverCancelAcknowledgeIEs := &handoverCancelAcknowledge.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.HandoverCancelAcknowledgeIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	handoverCancelAcknowledgeIEs.List = append(handoverCancelAcknowledgeIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.HandoverCancelAcknowledgeIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()

	handoverCancelAcknowledgeIEs.List = append(handoverCancelAcknowledgeIEs.List, ie)

	// Criticality Diagnostics [optional]
	if criticalityDiagnostics != nil {
		ie := ngapType.HandoverCancelAcknowledgeIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics
		ie.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)

		ie.Value.CriticalityDiagnostics = criticalityDiagnostics

		handoverCancelAcknowledgeIEs.List = append(handoverCancelAcknowledgeIEs.List, ie)
	}

	return libngap.Encoder(pdu)
}

func (h *Ngap) handleUplinkRanStatusTransfer(ran *ran.Ran, uplinkRanStatusTransfer *ngapType.UplinkRANStatusTransfer) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var rANStatusTransferTransparentContainer *ngapType.RANStatusTransferTransparentContainer
	var ue *ue.UeContext

	for _, ie := range uplinkRanStatusTransfer.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANStatusTransferTransparentContainer: // reject
			rANStatusTransferTransparentContainer = ie.Value.RANStatusTransferTransparentContainer
			log.Trace("Decode IE RANStatusTransferTransparentContainer")
			if rANStatusTransferTransparentContainer == nil {
				log.Error("RANStatusTransferTransparentContainer is nil")
			}
		}
	}

	ue = ran.FindUe(ranNgapId, coreNgapId)
	if ue == nil {
		return
	}

	log.Tracef("UE Context AmfUeNgapID[%d] UeContextNgapID[%d]", ue.CuNgapId(), ue.RanNgapId())
	log.Info("Handle Uplink Ran Status Transfer")
}

// RanStatusTransferTransparentContainer from Uplink Ran Configuration Transfer
func SendDownlinkRanStatusTransfer(ue *ue.UeContext, container ngapType.RANStatusTransferTransparentContainer) (err error) {
	log.Info("Send Downlink Ran Status Transfer")
	/*
		if len(container.DRBsSubjectToStatusTransferList.List) > context.MaxNumOfDRBs {
			log.Error("Pdu List out of range")
			return
		}
	*/
	var pkt []byte
	if pkt, err = buildDownlinkRanStatusTransfer(ue, container); err != nil {
		log.Errorf("Build DownlinkRanStatusTransfer failed : %s", err.Error())
		return
	}

	err = ue.Send(pkt)
	return
}

func buildDownlinkRanStatusTransfer(ue *ue.UeContext,
	ranStatusTransferTransparentContainer ngapType.RANStatusTransferTransparentContainer) ([]byte, error) {
	// ranStatusTransferTransparentContainer from Uplink Ran Configuration Transfer
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkRANStatusTransfer
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore
	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentDownlinkRANStatusTransfer
	initiatingMessage.Value.DownlinkRANStatusTransfer = new(ngapType.DownlinkRANStatusTransfer)

	downlinkRanStatusTransfer := initiatingMessage.Value.DownlinkRANStatusTransfer
	downlinkRanStatusTransferIEs := &downlinkRanStatusTransfer.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.DownlinkRANStatusTransferIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkRANStatusTransferIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	downlinkRanStatusTransferIEs.List = append(downlinkRanStatusTransferIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.DownlinkRANStatusTransferIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkRANStatusTransferIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()

	downlinkRanStatusTransferIEs.List = append(downlinkRanStatusTransferIEs.List, ie)

	// RAN Status Transfer Transparent Container
	ie = ngapType.DownlinkRANStatusTransferIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANStatusTransferTransparentContainer
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer

	ie.Value.RANStatusTransferTransparentContainer = &ranStatusTransferTransparentContainer

	downlinkRanStatusTransferIEs.List = append(downlinkRanStatusTransferIEs.List, ie)

	return libngap.Encoder(pdu)
}
