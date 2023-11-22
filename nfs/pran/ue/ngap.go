package ue

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"fmt"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

const (
	NGAP_NULL_ID int64 = -1
)

func (uectx *UeContext) SendDownlinkNasTransport(msg *n2models.NasDlMsg) (err error) {
	defer uectx.logSendingReport("DownlinkNasTransport", err)
	//naspdu []byte,
	//	mobilityRestrictionList *ngapType.MobilityRestrictionList
	var pdu ngapType.NGAPPDU

	if len(msg.NasPdu) == 0 {
		err = fmt.Errorf("Empty naspdu")
		return
	}

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkNASTransport
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentDownlinkNASTransport
	initiatingMessage.Value.DownlinkNASTransport = new(ngapType.DownlinkNASTransport)

	content := initiatingMessage.Value.DownlinkNASTransport
	ies := &content.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.DownlinkNASTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkNASTransportIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = uectx.CuNgapId()

	ies.List = append(ies.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.DownlinkNASTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkNASTransportIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = uectx.RanNgapId()

	ies.List = append(ies.List, ie)

	// NAS PDU
	ie = ngapType.DownlinkNASTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDNASPDU
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkNASTransportIEsPresentNASPDU
	ie.Value.NASPDU = new(ngapType.NASPDU)

	ie.Value.NASPDU.Value = msg.NasPdu

	ies.List = append(ies.List, ie)

	// Old AMF (optional)
	if len(msg.OldAmf) > 0 {
		ie = ngapType.DownlinkNASTransportIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDOldAMF
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.DownlinkNASTransportIEsPresentOldAMF
		ie.Value.OldAMF = new(ngapType.AMFName)

		ie.Value.OldAMF.Value = msg.OldAmf

		ies.List = append(ies.List, ie)
	}
	// RAN Paging Priority (optional)
	// Mobility Restriction List (optional)
	if msg.MobiRestrictList != nil {
		/*
			ie = ngapType.DownlinkNASTransportIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDMobilityRestrictionList
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.DownlinkNASTransportIEsPresentMobilityRestrictionList
			ie.Value.MobilityRestrictionList = //buil the list here
			ies.List = append(ies.List, ie)
		*/
	}
	// Index to RAT/Frequency Selection Priority (optional)
	// UE Aggregate Maximum Bit Rate (optional)
	// Allowed NSSAI (optional)

	var packet []byte
	if packet, err = libngap.Encoder(pdu); err == nil {
		err = uectx.Send(packet)
	}
	return
}

// TS 23.502 4.2.2.2.3
// anType: indicate amfUe send this msg for which accessType
// amfUeNgapID: initial AMF get it from target AMF
// ngapMessage: initial UE Message to reroute
// allowedNSSAI: provided by AMF, and AMF get it from NSSF (4.2.2.2.3 step 4b)
func (uectx *UeContext) SendRerouteNasRequest(anType models.AccessType, amfUeNgapID *int64, ngapMessage []byte,
	allowedNSSAI *ngapType.AllowedNSSAI) (err error) {
	defer uectx.logSendingReport("RerouteNasRequest", err)
	/*
		log.Info("Send Reroute Nas Request")

		if len(ngapMessage) == 0 {
			log.Error("Ngap Message is nil")
			return
		}

		pkt, err := s.buildRerouteNasRequest(ue, anType, amfUeNgapID, ngapMessage, allowedNSSAI)
		if err != nil {
			log.Errorf("Build RerouteNasRequest failed : %s", err.Error())
			return
		}
		s.NasSendToRan(ue, anType, pkt)
	*/
	return
}

func buildRerouteNasRequest(uectx *UeContext, anType models.AccessType, amfUeNgapID *int64,
	ngapMessage []byte, allowedNSSAI *ngapType.AllowedNSSAI) ([]byte, error) {
	/*
		var pdu ngapType.NGAPPDU

		pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
		pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

		initiatingMessage := pdu.InitiatingMessage
		initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeRerouteNASRequest
		initiatingMessage.Criticality.Value = ngapType.CriticalityPresentReject

		initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentRerouteNASRequest
		initiatingMessage.Value.RerouteNASRequest = new(ngapType.RerouteNASRequest)

		rerouteNasRequest := initiatingMessage.Value.RerouteNASRequest
		rerouteNasRequestIEs := &rerouteNasRequest.ProtocolIEs

		// RAN UE NGAP ID
		ie := ngapType.RerouteNASRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.RerouteNASRequestIEsPresentRANUENGAPID
		ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

		rANUENGAPID := ie.Value.RANUENGAPID
		rANUENGAPID.Value = UeContext[anType].RanNgapId()

		rerouteNasRequestIEs.List = append(rerouteNasRequestIEs.List, ie)

		// AMF UE NGAP ID (optional)
		if amfUeNgapID != nil {
			ie = ngapType.RerouteNASRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
			ie.Criticality.Value = ngapType.CriticalityPresentIgnore
			ie.Value.Present = ngapType.RerouteNASRequestIEsPresentAMFUENGAPID
			ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

			aMFUENGAPID := ie.Value.AMFUENGAPID
			aMFUENGAPID.Value = *amfUeNgapID

			rerouteNasRequestIEs.List = append(rerouteNasRequestIEs.List, ie)
		}

		// NGAP Message (Contains the initial ue message)
		ie = ngapType.RerouteNASRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDNGAPMessage
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.RerouteNASRequestIEsPresentNGAPMessage

		msg := aper.OctetString(ngapMessage)
		ie.Value.NGAPMessage = &msg

		rerouteNasRequestIEs.List = append(rerouteNasRequestIEs.List, ie)

		// AMF Set ID
		ie = ngapType.RerouteNASRequestIEs{}
		ie.Id.Value = ngapType.ProtocolIEIDAMFSetID
		ie.Criticality.Value = ngapType.CriticalityPresentReject
		ie.Value.Present = ngapType.RerouteNASRequestIEsPresentAMFSetID

		// <MCC><MNC><AMF Region ID><AMF Set ID><AMF Pointer><5G-TMSI>
		// <MCC><MNC> is 3 bytes, <AMF Region ID><AMF Set ID><AMF Pointer> is 3 bytes
		// 1 byte is 2 characters
		var amfID string
		if len(ue.Guti) == 19 { // MNC is 2 char
			amfID = ue.Guti[5:11]
		} else {
			amfID = ue.Guti[6:12]
		}
		_, amfSetID, _ := ngapConvert.AmfIdToNgap(amfID)

		ie.Value.AMFSetID = new(ngapType.AMFSetID)
		ie.Value.AMFSetID.Value = amfSetID

		rerouteNasRequestIEs.List = append(rerouteNasRequestIEs.List, ie)

		// Allowed NSSAI
		if allowedNSSAI != nil {
			ie = ngapType.RerouteNASRequestIEs{}
			ie.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
			ie.Criticality.Value = ngapType.CriticalityPresentReject
			ie.Value.Present = ngapType.RerouteNASRequestIEsPresentAllowedNSSAI

			ie.Value.AllowedNSSAI = allowedNSSAI

			rerouteNasRequestIEs.List = append(rerouteNasRequestIEs.List, ie)
		}

		return libngap.Encoder(pdu)
	*/
	return nil, nil
}
