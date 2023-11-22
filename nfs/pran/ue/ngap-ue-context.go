package ue

import (
	"encoding/hex"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/sbi/utils/ngapConvert"
	"fmt"

	"github.com/free5gc/aper"
	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

// 8.3.1
func (uectx *UeContext) SendInitialContextSetupRequest(msg *n2models.InitCtxSetupReq) (err error) {
	defer uectx.logSendingReport("InitialContextSetupRequest", err)

	var ngapmsg ngapType.NGAPPDU
	ngapmsg.Present = ngapType.NGAPPDUPresentInitiatingMessage
	ngapmsg.InitiatingMessage = new(ngapType.InitiatingMessage)

	initmsg := ngapmsg.InitiatingMessage
	initmsg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	initmsg.Criticality.Value = ngapType.CriticalityPresentReject

	initmsg.Value.Present = ngapType.InitiatingMessagePresentInitialContextSetupRequest
	initmsg.Value.InitialContextSetupRequest = new(ngapType.InitialContextSetupRequest)

	req := initmsg.Value.InitialContextSetupRequest
	ies := &req.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.InitialContextSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = uectx.CuNgapId()
	ies.List = append(ies.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.InitialContextSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = uectx.RanNgapId()
	ies.List = append(ies.List, ie)

	// Old AMF (optional)
	if len(msg.OldAmf) > 0 {
		ie = ngapType.InitialContextSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDOldAMF
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentOldAMF
		ie.Value.OldAMF = new(ngapType.AMFName)
		ie.Value.OldAMF.Value = msg.OldAmf
		ies.List = append(ies.List, ie)
	}
	// UE Aggregate Maximum Bit Rate (conditional: if pdu session resource setup)
	// The subscribed UE-AMBR is a subscription parameter which is
	// retrieved from UDM and provided to the (R)AN by the AMF
	if len(msg.PduList) > 0 {
		uectx.Tracef("pdulist is not empty = %d", len(msg.PduList))
		ie = ngapType.InitialContextSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUEAggregateMaximumBitRate
		ie.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)

		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateUL.Value = msg.UeAmbr.Ul
		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateDL.Value = msg.UeAmbr.Dl

		ies.List = append(ies.List, ie)

		// PDU Session Resource Setup Request List
		ie = ngapType.InitialContextSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtReq
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentPDUSessionResourceSetupListCxtReq
		ie.Value.PDUSessionResourceSetupListCxtReq = new(ngapType.PDUSessionResourceSetupListCxtReq)
		pdulist := ie.Value.PDUSessionResourceSetupListCxtReq
		var item *ngapType.PDUSessionResourceSetupItemCxtReq
		for _, s := range msg.PduList {
			item = new(ngapType.PDUSessionResourceSetupItemCxtReq)
			item.PDUSessionID.Value = int64(s.Id)
			item.SNSSAI = ngapConvert.SNssaiToNgap(s.Snssai)
			item.PDUSessionResourceSetupRequestTransfer = s.Transfer
			if len(s.NasPdu) > 0 {
				item.NASPDU = new(ngapType.NASPDU)
				item.NASPDU.Value = s.NasPdu
			}
			pdulist.List = append(pdulist.List, *item)
		}
		ies.List = append(ies.List, ie)

	}
	// GUAMI
	ie = ngapType.InitialContextSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDGUAMI
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentGUAMI
	ie.Value.GUAMI = new(ngapType.GUAMI)

	uectx.Tracef("Plmnid=%s, Amfid=%s", msg.Guami.PlmnId.String(), msg.Guami.AmfId)
	guami := ie.Value.GUAMI
	guami.PLMNIdentity = ngapConvert.PlmnIdToNgap(msg.Guami.PlmnId)
	guami.AMFRegionID.Value, guami.AMFSetID.Value, guami.AMFPointer.Value = ngapConvert.AmfIdToNgap(msg.Guami.AmfId)

	ies.List = append(ies.List, ie)

	/*
		//tungtq: need to check convertion, otherwise there will be an encoding
		//error
			// Allowed NSSAI
			if len(msg.AllowedNssai) > 0 {
				log.Info("has allowed nssai")
				ie = ngapType.InitialContextSetupRequestIEs{}
				ie.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
				ie.Criticality.Value = ngapType.CriticalityPresentReject
				ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentAllowedNSSAI
				allowednssai := ngapConvert.AllowedNssaiToNgap(msg.AllowedNssai)
				ie.Value.AllowedNSSAI = &allowednssai

				ies.List = append(ies.List, ie)
			}
	*/
	// UE Security Capabilities
	ie = ngapType.InitialContextSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUESecurityCapabilities
	ie.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	ueseccap := ie.Value.UESecurityCapabilities

	if msg.UeSecCap.Nr != nil {
		uectx.Tracef("Has nr seccap")
		ueseccap.NRencryptionAlgorithms.Value = ngapConvert.ByteToBitString(msg.UeSecCap.Nr.Enc[:], 16)
		ueseccap.NRintegrityProtectionAlgorithms.Value = ngapConvert.ByteToBitString(msg.UeSecCap.Nr.Int[:], 16)
	} else {
		uectx.Tracef("Has dummy nr seccap")
		dummy := []byte{0x00, 0x00}
		ueseccap.NRencryptionAlgorithms.Value = ngapConvert.ByteToBitString(dummy, 16)
		ueseccap.NRintegrityProtectionAlgorithms.Value = ngapConvert.ByteToBitString(dummy, 16)
	}

	if msg.UeSecCap.Eutra != nil {
		uectx.Trace("has eutra seccap")
		ueseccap.EUTRAencryptionAlgorithms.Value = ngapConvert.ByteToBitString(msg.UeSecCap.Nr.Enc[:], 16)
		ueseccap.EUTRAintegrityProtectionAlgorithms.Value = ngapConvert.ByteToBitString(msg.UeSecCap.Nr.Int[:], 16)
	} else {
		uectx.Trace("has dummy eutra seccap")
		dummy := []byte{0x00, 0x00}
		ueseccap.EUTRAencryptionAlgorithms.Value = ngapConvert.ByteToBitString(dummy, 16)
		ueseccap.EUTRAintegrityProtectionAlgorithms.Value = ngapConvert.ByteToBitString(dummy, 16)
	}

	ies.List = append(ies.List, ie)

	// Security Key
	ie = ngapType.InitialContextSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDSecurityKey
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentSecurityKey
	ie.Value.SecurityKey = new(ngapType.SecurityKey)

	seckey := ie.Value.SecurityKey
	seckey.Value = ngapConvert.ByteToBitString(msg.SecKey, 256)
	ies.List = append(ies.List, ie)
	uectx.Tracef("key %x", msg.SecKey)

	// NAS-PDU (optional)
	if len(msg.NasPdu) > 0 {
		uectx.Tracef("Set nas pdu, len=%d", len(msg.NasPdu))
		ie = ngapType.InitialContextSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDNASPDU
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentNASPDU
		ie.Value.NASPDU = new(ngapType.NASPDU)

		ie.Value.NASPDU.Value = msg.NasPdu

		ies.List = append(ies.List, ie)
	}

	// UE Radio Capability (optional)
	if len(msg.UeRadCap) > 0 {
		ie = ngapType.InitialContextSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUERadioCapability
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUERadioCapability
		ie.Value.UERadioCapability = new(ngapType.UERadioCapability)
		if ie.Value.UERadioCapability.Value, err = hex.DecodeString(msg.UeRadCap); err != nil {
			return
		}
		ies.List = append(ies.List, ie)
	}

	/*
		// Core Network Assistance Information (optional)
		if coreNetworkAssistanceInfo != nil {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentCoreNetworkAssistanceInformation
			ie.Value.CoreNetworkAssistanceInformation = coreNetworkAssistanceInfo
			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}

		udminfo := amfUe.UdmClient().Info()
		// Trace Activation (optional)
		if udminfo.TraceData != nil {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDTraceActivation
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentTraceActivation
			ie.Value.TraceActivation = new(ngapType.TraceActivation)
			// TS 32.422 4.2.2.9
			// TODO: AMF allocate Trace Recording Session Reference
			traceActivation := ngapConvert.TraceDataToNgap(*udminfo.TraceData, ue.Trsr())
			ie.Value.TraceActivation = &traceActivation
			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}

		// Mobility Restriction List (optional)
		if anType == models.ACCESSTYPE__3_GPP_ACCESS {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDMobilityRestrictionList
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentMobilityRestrictionList
			ie.Value.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)

			mobilityRestrictionList := util.BuildIEMobilityRestrictionList(amfUe)
			ie.Value.MobilityRestrictionList = &mobilityRestrictionList

			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}


				// Masked IMEISV (optional)
		// TS 38.413 9.3.1.54; TS 23.003 6.2; TS 23.501 5.9.3
		// last 4 digits of the SNR masked by setting the corresponding bits to 1.
		// The first to fourth bits correspond to the first digit of the IMEISV,
		// the fifth to eighth bits correspond to the second digit of the IMEISV, and so on
		if amfUe.Pei != "" && strings.HasPrefix(amfUe.Pei, "imeisv") {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDMaskedIMEISV
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentMaskedIMEISV
			ie.Value.MaskedIMEISV = new(ngapType.MaskedIMEISV)

			imeisv := strings.TrimPrefix(amfUe.Pei, "imeisv-")
			imeisvBytes, err := hex.DecodeString(imeisv)
			if err != nil {
				//logger.NgapLog.Errorf("[Build Error] DecodeString imeisv error: %+v", err)
			}

			var maskedImeisv []byte
			maskedImeisv = append(maskedImeisv, imeisvBytes[:5]...)
			maskedImeisv = append(maskedImeisv, []byte{0xff, 0xff}...)
			maskedImeisv = append(maskedImeisv, imeisvBytes[7])
			ie.Value.MaskedIMEISV.Value = aper.BitString{
				BitLength: 64,
				Bytes:     maskedImeisv,
			}
			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}


		// Emergency Fallback indicator (optional)
		if emergencyFallbackIndicator != nil {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDEmergencyFallbackIndicator
			ie.Criticality.Value = ngapType.CriticalityPresentReject
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentEmergencyFallbackIndicator
			ie.Value.EmergencyFallbackIndicator = emergencyFallbackIndicator
			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}

		// RRC Inactive Transition Report Request (optional)
		if rrcInactiveTransitionReportRequest != nil {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentRRCInactiveTransitionReportRequest
			ie.Value.RRCInactiveTransitionReportRequest = rrcInactiveTransitionReportRequest
			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}

		// UE Radio Capability for Paging (optional)
		if amfUe.UeRadioCapabilityForPaging != nil {
			ie = ngapType.InitialContextSetupRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUERadioCapabilityForPaging
			ie.Value.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
			uERadioCapabilityForPaging := ie.Value.UERadioCapabilityForPaging
			var err error
			if amfUe.UeRadioCapabilityForPaging.NR != "" {
				uERadioCapabilityForPaging.UERadioCapabilityForPagingOfNR.Value, err =
					hex.DecodeString(amfUe.UeRadioCapabilityForPaging.NR)
				if err != nil {
					//logger.NgapLog.Errorf("[Build Error] DecodeString amfUe.UeRadioCapabilityForPaging.NR error: %+v", err)
				}
			}
			if amfUe.UeRadioCapabilityForPaging.EUTRA != "" {
				uERadioCapabilityForPaging.UERadioCapabilityForPagingOfEUTRA.Value, err =
					hex.DecodeString(amfUe.UeRadioCapabilityForPaging.EUTRA)
				if err != nil {
					//logger.NgapLog.Errorf("[Build Error] DecodeString amfUe.UeRadioCapabilityForPaging.NR error: %+v", err)
				}
			}
			initialContextSetupRequestIEs.List = append(initialContextSetupRequestIEs.List, ie)
		}
	*/
	// Index to RAT/Frequency Selection Priority (optional)
	if msg.Rfsp != nil {
		ie = ngapType.InitialContextSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDIndexToRFSP
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.InitialContextSetupRequestIEsPresentIndexToRFSP
		ie.Value.IndexToRFSP = new(ngapType.IndexToRFSP)

		ie.Value.IndexToRFSP.Value = *msg.Rfsp

		ies.List = append(ies.List, ie)
	}
	var pdu []byte
	if pdu, err = libngap.Encoder(ngapmsg); err == nil {
		err = uectx.Send(pdu)
	} else {
		uectx.Info(err)
	}
	return
}

// 8.3.3
func (uectx *UeContext) SendUEContextReleaseCommand(msg *n2models.UeCtxRelCmd) (err error) {
	defer uectx.logSendingReport("UEContextReleaseCommand", err)
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeUEContextRelease
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentUEContextReleaseCommand
	initiatingMessage.Value.UEContextReleaseCommand = new(ngapType.UEContextReleaseCommand)

	cmd := initiatingMessage.Value.UEContextReleaseCommand
	ies := &cmd.ProtocolIEs

	// UE NGAP IDs
	ie := ngapType.UEContextReleaseCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDUENGAPIDs
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UEContextReleaseCommandIEsPresentUENGAPIDs
	ie.Value.UENGAPIDs = new(ngapType.UENGAPIDs)

	ueNGAPIDs := ie.Value.UENGAPIDs

	if uectx.RanNgapId() == NGAP_NULL_ID {
		ueNGAPIDs.Present = ngapType.UENGAPIDsPresentAMFUENGAPID
		ueNGAPIDs.AMFUENGAPID = new(ngapType.AMFUENGAPID)

		ueNGAPIDs.AMFUENGAPID.Value = uectx.CuNgapId()
	} else {
		ueNGAPIDs.Present = ngapType.UENGAPIDsPresentUENGAPIDPair
		ueNGAPIDs.UENGAPIDPair = new(ngapType.UENGAPIDPair)

		ueNGAPIDs.UENGAPIDPair.AMFUENGAPID.Value = uectx.CuNgapId()
		ueNGAPIDs.UENGAPIDPair.RANUENGAPID.Value = uectx.RanNgapId()
	}

	ies.List = append(ies.List, ie)

	// Cause
	ie = ngapType.UEContextReleaseCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDCause
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.UEContextReleaseCommandIEsPresentCause
	ngapcause := ngapType.Cause{
		Present: int(msg.Cause.Present),
	}
	cause := aper.Enumerated(msg.Cause.Value)
	switch ngapcause.Present {
	case ngapType.CausePresentNothing:
		err = fmt.Errorf("Cause Present is not set")
		return
	case ngapType.CausePresentRadioNetwork:
		ngapcause.RadioNetwork = new(ngapType.CauseRadioNetwork)
		ngapcause.RadioNetwork.Value = cause
	case ngapType.CausePresentTransport:
		ngapcause.Transport = new(ngapType.CauseTransport)
		ngapcause.Transport.Value = cause
	case ngapType.CausePresentNas:
		ngapcause.Nas = new(ngapType.CauseNas)
		ngapcause.Nas.Value = cause
	case ngapType.CausePresentProtocol:
		ngapcause.Protocol = new(ngapType.CauseProtocol)
		ngapcause.Protocol.Value = cause
	case ngapType.CausePresentMisc:
		ngapcause.Misc = new(ngapType.CauseMisc)
		ngapcause.Misc.Value = cause
	default:
		err = fmt.Errorf("Cause Present is Unknown")
		return
	}
	ie.Value.Cause = &ngapcause

	ies.List = append(ies.List, ie)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}
	return
}

// 8.3.4
func (uectx *UeContext) SendUEContextModificationRequest(msg *n2models.UeCtxModReq) (err error) {
	defer uectx.logSendingReport("UEContextModificationRequest", err)

	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentUEContextModificationRequest
	initiatingMessage.Value.UEContextModificationRequest = new(ngapType.UEContextModificationRequest)

	req := initiatingMessage.Value.UEContextModificationRequest
	ies := &req.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.UEContextModificationRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	if msg.OldAmfNgapId != nil {
		ie.Value.AMFUENGAPID.Value = *msg.OldAmfNgapId
	} else {
		ie.Value.AMFUENGAPID.Value = uectx.CuNgapId()
	}
	ies.List = append(ies.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.UEContextModificationRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	ie.Value.RANUENGAPID.Value = uectx.RanNgapId()

	ies.List = append(ies.List, ie)

	// Ran Paging Priority (optional)

	// Security Key (optional)

	// Index to RAT/Frequency Selection Priority (optional)
	if msg.Rfsp != nil {
		//model data structure has been changed, need more investigation
		ie = ngapType.UEContextModificationRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDIndexToRFSP
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentIndexToRFSP
		ie.Value.IndexToRFSP = new(ngapType.IndexToRFSP)
		ie.Value.IndexToRFSP.Value = *msg.Rfsp

		ies.List = append(ies.List, ie)
	}
	// UE Aggregate Maximum Bit Rate (optional)
	if msg.UeAmbr != nil {
		ie = ngapType.UEContextModificationRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentUEAggregateMaximumBitRate
		ie.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)

		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateUL.Value = msg.UeAmbr.Ul
		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateDL.Value = msg.UeAmbr.Dl

		ies.List = append(ies.List, ie)
	}
	// UE Security Capabilities (optional)

	// Core Network Assistance Information (optional)
	if msg.CoreAssist != nil {
		/*
			ie = ngapType.UEContextModificationRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentCoreNetworkAssistanceInformation
			ie.Value.CoreNetworkAssistanceInformation = coreNetworkAssistanceInfo
			ies.List = append(ies.List, ie)
		*/
	}
	// Emergency Fallback Indicator (optional)
	if msg.Emerg != nil {
		/*
			ie = ngapType.UEContextModificationRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDEmergencyFallbackIndicator
			ie.Criticality.Value = ngapType.CriticalityPresentReject
			ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentEmergencyFallbackIndicator
			ie.Value.EmergencyFallbackIndicator = emergencyFallbackIndicator
			ies.List = append(ies.List, ie)
		*/
	}
	// New AMF UE NGAP ID (optional)
	if msg.OldAmfNgapId != nil {
		ie = ngapType.UEContextModificationRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDNewAMFUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentNewAMFUENGAPID
		ie.Value.NewAMFUENGAPID = new(ngapType.AMFUENGAPID)

		ie.Value.NewAMFUENGAPID.Value = uectx.CuNgapId()

		ies.List = append(ies.List, ie)
	}
	// RRC Inactive Transition Report Request (optional)
	if msg.RrcInactTranRepReq != nil {
		ie = ngapType.UEContextModificationRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.UEContextModificationRequestIEsPresentRRCInactiveTransitionReportRequest
		ie.Value.RRCInactiveTransitionReportRequest = &ngapType.RRCInactiveTransitionReportRequest{
			Value: aper.Enumerated(*msg.RrcInactTranRepReq),
		}
		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}

	return
}
