package ue

import (
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/sbi/utils/ngapConvert"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

///PDU

// TS138.413-V15.3.0 8.2.1
// nasPDU: from nas layer
// pduSessionResourceSetupRequestList: provided by AMF, and transfer data is from SMF
func (uectx *UeContext) SendPduSessionResourceSetupRequest(msg *n2models.PduSessResSetReq) (err error) {
	// TODO: Ran Paging Priority (optional)
	defer uectx.logSendingReport("PduSessionResourceSetupRequest", err)

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceSetup
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceSetupRequest
	initiatingMessage.Value.PDUSessionResourceSetupRequest = new(ngapType.PDUSessionResourceSetupRequest)

	req := initiatingMessage.Value.PDUSessionResourceSetupRequest
	ies := &req.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.PDUSessionResourceSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = uectx.CuNgapId()

	ies.List = append(ies.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.PDUSessionResourceSetupRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = uectx.RanNgapId()

	ies.List = append(ies.List, ie)

	// Ran Paging Priority (optional)

	// NAS-PDU (optional)
	if len(msg.NasPdu) > 0 {
		ie = ngapType.PDUSessionResourceSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDNASPDU
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentNASPDU
		ie.Value.NASPDU = new(ngapType.NASPDU)

		ie.Value.NASPDU.Value = msg.NasPdu

		ies.List = append(ies.List, ie)
	}

	if len(msg.SessionList) > 0 {
		// PDU Session Resource Setup Request list
		pdulist := make([]ngapType.PDUSessionResourceSetupItemSUReq, len(msg.SessionList))
		for i, s := range msg.SessionList {
			pdulist[i] = ngapType.PDUSessionResourceSetupItemSUReq{
				PDUSessionID: ngapType.PDUSessionID{
					Value: s.Id,
				},
				SNSSAI:                                 ngapConvert.SNssaiToNgap(s.Snssai),
				PDUSessionResourceSetupRequestTransfer: s.Transfer,
			}
			if len(s.NasPdu) > 0 {
				pdulist[i].PDUSessionNASPDU = &ngapType.NASPDU{
					Value: s.NasPdu,
				}
			}
		}
		ie = ngapType.PDUSessionResourceSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentPDUSessionResourceSetupListSUReq
		ie.Value.PDUSessionResourceSetupListSUReq = &ngapType.PDUSessionResourceSetupListSUReq{
			List: pdulist,
		}
		ies.List = append(ies.List, ie)
	}
	// UE AggreateMaximum Bit Rate
	if msg.UeAmbr != nil {
		ie = ngapType.PDUSessionResourceSetupRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentUEAggregateMaximumBitRate
		ie.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateUL.Value = msg.UeAmbr.Ul
		ie.Value.UEAggregateMaximumBitRate.UEAggregateMaximumBitRateDL.Value = msg.UeAmbr.Dl
		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}
	return
}

//TS138.413-V15.3.0 8.2.2

func (uectx *UeContext) SendPduSessionResourceReleaseCommand(msg *n2models.PduSessResRelCmd) (err error) {
	defer uectx.logSendingReport("PduSessionResourceReleaseCommand", err)

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceRelease
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject
	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceReleaseCommand
	initiatingMessage.Value.PDUSessionResourceReleaseCommand = new(ngapType.PDUSessionResourceReleaseCommand)

	cmd := initiatingMessage.Value.PDUSessionResourceReleaseCommand
	ies := &cmd.ProtocolIEs

	// AMFUENGAPID
	ie := ngapType.PDUSessionResourceReleaseCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = uectx.CuNgapId()

	ies.List = append(ies.List, ie)

	// RANUENGAPID
	ie = ngapType.PDUSessionResourceReleaseCommandIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = uectx.RanNgapId()

	ies.List = append(ies.List, ie)

	// NAS-PDU (optional)
	if len(msg.NasPdu) > 0 {
		ie = ngapType.PDUSessionResourceReleaseCommandIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDNASPDU
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentNASPDU
		ie.Value.NASPDU = new(ngapType.NASPDU)

		ie.Value.NASPDU.Value = msg.NasPdu

		ies.List = append(ies.List, ie)
	}

	// PDUSessionResourceToReleaseListRelCmd
	if len(msg.SessionList) > 0 {
		pdulist := make([]ngapType.PDUSessionResourceToReleaseItemRelCmd, len(msg.SessionList))
		for i, s := range msg.SessionList {
			pdulist[i] = ngapType.PDUSessionResourceToReleaseItemRelCmd{
				PDUSessionID: ngapType.PDUSessionID{
					Value: s.Id,
				},
				PDUSessionResourceReleaseCommandTransfer: s.Transfer,
			}
		}
		ie = ngapType.PDUSessionResourceReleaseCommandIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceToReleaseListRelCmd
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentPDUSessionResourceToReleaseListRelCmd
		ie.Value.PDUSessionResourceToReleaseListRelCmd = &ngapType.PDUSessionResourceToReleaseListRelCmd{
			List: pdulist,
		}

		ies.List = append(ies.List, ie)
	}

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}

	return
}

// TS138.413-V15.3.0 8.2.3
// pduSessionResourceModifyRequestList: from SMF
func (uectx *UeContext) SendPduSessionResourceModifyRequest(msg *n2models.PduSessResModReq) (err error) {
	defer uectx.logSendingReport("PduSessionResourceModifyRequest", err)
	// TODO: Ran Paging Priority (optional)
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModify
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceModifyRequest
	initiatingMessage.Value.PDUSessionResourceModifyRequest = new(ngapType.PDUSessionResourceModifyRequest)

	req := initiatingMessage.Value.PDUSessionResourceModifyRequest
	ies := &req.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.PDUSessionResourceModifyRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = uectx.CuNgapId()

	ies.List = append(ies.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.PDUSessionResourceModifyRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = uectx.RanNgapId()

	ies.List = append(ies.List, ie)

	// Ran Paging Priority (optional)

	// PDU Session Resource Modify Request List
	if len(msg.SessionList) > 0 {
		pdulist := make([]ngapType.PDUSessionResourceModifyItemModReq, len(msg.SessionList))
		for i, s := range msg.SessionList {
			pdulist[i] = ngapType.PDUSessionResourceModifyItemModReq{
				PDUSessionID: ngapType.PDUSessionID{
					Value: s.Id,
				},
				PDUSessionResourceModifyRequestTransfer: s.Transfer,
			}
			if len(s.NasPdu) > 0 {
				pdulist[i].NASPDU = &ngapType.NASPDU{
					Value: s.NasPdu,
				}
			}
		}
		ie = ngapType.PDUSessionResourceModifyRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModReq
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentPDUSessionResourceModifyListModReq
		ie.Value.PDUSessionResourceModifyListModReq = &ngapType.PDUSessionResourceModifyListModReq{
			List: pdulist,
		}
		ies.List = append(ies.List, ie)
	}
	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}
	return
}

// pduSessionResourceFailedToModifyList: provided by AMF, and transfer data is return from SMF
func (uectx *UeContext) SendPduSessionResourceModifyConfirm(msg *n2models.PduSessResModCfm) (err error) {
	defer uectx.logSendingReport("PduSessionResourceModifyConfirm", err)

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)

	successfulOutcome := pdu.SuccessfulOutcome
	successfulOutcome.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModifyIndication
	successfulOutcome.Criticality.Value = ngapType.CriticalityPresentReject

	successfulOutcome.Value.Present = ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyConfirm
	successfulOutcome.Value.PDUSessionResourceModifyConfirm = new(ngapType.PDUSessionResourceModifyConfirm)

	cfm := successfulOutcome.Value.PDUSessionResourceModifyConfirm
	ies := &cfm.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.PDUSessionResourceModifyConfirmIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = uectx.CuNgapId()

	ies.List = append(ies.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.PDUSessionResourceModifyConfirmIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = uectx.RanNgapId()

	ies.List = append(ies.List, ie)

	// PDU Session Resource Modify Confirm List
	if len(msg.ConfirmedList) > 0 {
		pdulist := make([]ngapType.PDUSessionResourceModifyItemModCfm, len(msg.ConfirmedList))
		for i, s := range msg.ConfirmedList {
			pdulist[i] = ngapType.PDUSessionResourceModifyItemModCfm{
				PDUSessionID: ngapType.PDUSessionID{
					Value: s.Id,
				},
				PDUSessionResourceModifyConfirmTransfer: s.Transfer,
			}
		}
		ie = ngapType.PDUSessionResourceModifyConfirmIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModCfm
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceModifyListModCfm
		ie.Value.PDUSessionResourceModifyListModCfm = &ngapType.PDUSessionResourceModifyListModCfm{
			List: pdulist,
		}
		ies.List = append(ies.List, ie)
	}

	//	criticalityDiagnostics *ngapType.CriticalityDiagnostics
	// PDU Session Resource Failed to Modify List
	if len(msg.FailedList) > 0 {
		pdulist := make([]ngapType.PDUSessionResourceFailedToModifyItemModCfm, len(msg.FailedList))
		for i, s := range msg.FailedList {
			pdulist[i] = ngapType.PDUSessionResourceFailedToModifyItemModCfm{
				PDUSessionID: ngapType.PDUSessionID{
					Value: s.Id,
				},
				PDUSessionResourceModifyIndicationUnsuccessfulTransfer: s.Transfer,
			}
		}
		ie = ngapType.PDUSessionResourceModifyConfirmIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm
		ie.Criticality.Value = ngapType.CriticalityPresentIgnore
		ie.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceFailedToModifyListModCfm
		ie.Value.PDUSessionResourceFailedToModifyListModCfm = &ngapType.PDUSessionResourceFailedToModifyListModCfm{
			List: pdulist,
		}
		ies.List = append(ies.List, ie)
	}

	/*
		// Criticality Diagnostics (optional)
		if criticalityDiagnostics != nil {
			ie = ngapType.PDUSessionResourceModifyConfirmIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentCriticalityDiagnostics
			ie.Value.CriticalityDiagnostics = criticalityDiagnostics
			pDUSessionResourceModifyConfirmIEs.List = append(pDUSessionResourceModifyConfirmIEs.List, ie)
		}
	*/
	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}
	return
}
