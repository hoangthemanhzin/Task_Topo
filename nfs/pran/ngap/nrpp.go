package ngap

import (
	"encoding/hex"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"fmt"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

// NRPPa PDU is a pdu from LMF to RAN defined in TS 23.502 4.13.5.5 step 3
// NRPPa PDU is by pass
func SendDownlinkUEAssociatedNRPPaTransport(ue *ue.UeContext, nRPPaPDU ngapType.NRPPaPDU) (err error) {
	log.Info("Send Downlink UE Associated NRPPa Transport")

	if len(nRPPaPDU.Value) == 0 {
		err = fmt.Errorf("length of NRPPA-PDU is 0")
		log.Error(err.Error())
		return
	}

	var pkt []byte
	if pkt, err = buildDownlinkUEAssociatedNRPPaTransport(ue, nRPPaPDU); err != nil {
		log.Errorf("Build DownlinkUEAssociatedNRPPaTransport failed : %s", err.Error())
		return
	}
	err = ue.Send(pkt)
	return
}

// NRPPa PDU is a pdu from LMF to RAN defined in TS 23.502 4.13.5.5 step 3
// NRPPa PDU is by pass
func buildDownlinkUEAssociatedNRPPaTransport(ue *ue.UeContext, nRPPaPDU ngapType.NRPPaPDU) ([]byte, error) {
	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkUEAssociatedNRPPaTransport
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport
	initiatingMessage.Value.DownlinkUEAssociatedNRPPaTransport = new(ngapType.DownlinkUEAssociatedNRPPaTransport)

	downlinkUEAssociatedNRPPaTransport := initiatingMessage.Value.DownlinkUEAssociatedNRPPaTransport
	downlinkUEAssociatedNRPPaTransportIEs := &downlinkUEAssociatedNRPPaTransport.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	downlinkUEAssociatedNRPPaTransportIEs.List = append(downlinkUEAssociatedNRPPaTransportIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()

	downlinkUEAssociatedNRPPaTransportIEs.List = append(downlinkUEAssociatedNRPPaTransportIEs.List, ie)

	// Routing ID
	ie = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRoutingID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRoutingID
	ie.Value.RoutingID = new(ngapType.RoutingID)

	var err error
	routingID := ie.Value.RoutingID
	routingID.Value, err = hex.DecodeString(ue.RoutingId())
	if err != nil {
		//logger.NgapLog.Errorf("[Build Error] DecodeString ue.RoutingId() error: %+v", err)
	}

	downlinkUEAssociatedNRPPaTransportIEs.List = append(downlinkUEAssociatedNRPPaTransportIEs.List, ie)

	// NRPPa-PDU
	ie = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDNRPPaPDU
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU
	ie.Value.NRPPaPDU = new(ngapType.NRPPaPDU)

	ie.Value.NRPPaPDU = &nRPPaPDU

	downlinkUEAssociatedNRPPaTransportIEs.List = append(downlinkUEAssociatedNRPPaTransportIEs.List, ie)

	return libngap.Encoder(pdu)
}
func (h *Ngap) handleUplinkUEAssociatedNRPPATransport(ran *ran.Ran, uplinkUEAssociatedNRPPaTransport *ngapType.UplinkUEAssociatedNRPPaTransport) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var routingID *ngapType.RoutingID
	var nRPPaPDU *ngapType.NRPPaPDU

	for _, ie := range uplinkUEAssociatedNRPPaTransport.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE coreNgapId")
			if coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE ranNgapId")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRoutingID: // reject
			routingID = ie.Value.RoutingID
			log.Trace("Decode IE routingID")
			if routingID == nil {
				log.Error("routingID is nil")
				return
			}
		case ngapType.ProtocolIEIDNRPPaPDU: // reject
			nRPPaPDU = ie.Value.NRPPaPDU
			log.Trace("Decode IE nRPPaPDU")
			if nRPPaPDU == nil {
				log.Error("nRPPaPDU is nil")
				return
			}
		}
	}

	ue := ran.FindUe(ranNgapId, coreNgapId)
	if ue == nil {
		return
	}

	log.Info("Handle Uplink UE Associated NRPPA Transpor")

	ue.SetRoutingId(hex.EncodeToString(routingID.Value))

	// TODO: Forward NRPPaPDU to LMF
}

// NRPPa PDU is by pass
// NRPPa PDU is from LMF define in 4.13.5.6
func SendDownlinkNonUEAssociatedNRPPATransport(ue *ue.UeContext, nRPPaPDU ngapType.NRPPaPDU) (err error) {
	log.Info("Send Downlink Non UE Associated NRPPA Transport")

	if len(nRPPaPDU.Value) == 0 {
		err = fmt.Errorf("length of NRPPA-PDU is 0")
		log.Error(err.Error())
		return
	}
	var pkt []byte
	if pkt, err = buildDownlinkNonUEAssociatedNRPPATransport(ue, nRPPaPDU); err != nil {
		log.Errorf("Build DownlinkNonUEAssociatedNRPPATransport failed : %s", err.Error())
		return
	}
	err = ue.Send(pkt)
	return
}

func buildDownlinkNonUEAssociatedNRPPATransport(
	ue *ue.UeContext, nRPPaPDU ngapType.NRPPaPDU) ([]byte, error) {
	// NRPPa PDU is by pass
	// NRPPa PDU is from LMF define in 4.13.5.6

	var pdu ngapType.NGAPPDU
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport
	initiatingMessage.Value.DownlinkNonUEAssociatedNRPPaTransport =
		new(ngapType.DownlinkNonUEAssociatedNRPPaTransport)

	downlinkNonUEAssociatedNRPPaTransport := initiatingMessage.Value.DownlinkNonUEAssociatedNRPPaTransport
	downlinkNonUEAssociatedNRPPaTransportIEs := &downlinkNonUEAssociatedNRPPaTransport.ProtocolIEs

	// Routing ID
	// Routing id in the ran context
	ie := ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRoutingID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID
	ie.Value.RoutingID = new(ngapType.RoutingID)

	var err error
	routingID := ie.Value.RoutingID
	routingID.Value, err = hex.DecodeString(ue.RoutingId())
	if err != nil {
		log.Errorf("[Build Error] DecodeString ue.RoutingId() error: %+v", err)
	}

	downlinkNonUEAssociatedNRPPaTransportIEs.List = append(downlinkNonUEAssociatedNRPPaTransportIEs.List, ie)

	// NRPPa-PDU
	ie = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDNRPPaPDU
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU
	ie.Value.NRPPaPDU = new(ngapType.NRPPaPDU)

	ie.Value.NRPPaPDU = &nRPPaPDU

	downlinkNonUEAssociatedNRPPaTransportIEs.List = append(downlinkNonUEAssociatedNRPPaTransportIEs.List, ie)
	return libngap.Encoder(pdu)
}

func (h *Ngap) handleUplinkNonUEAssociatedNRPPATransport(ran *ran.Ran, uplinkNonUEAssociatedNRPPATransport *ngapType.UplinkNonUEAssociatedNRPPaTransport) {
	var routingID *ngapType.RoutingID
	var nRPPaPDU *ngapType.NRPPaPDU

	log.Info("Handle Uplink Non UE Associated NRPPA Transport")

	for i := 0; i < len(uplinkNonUEAssociatedNRPPATransport.ProtocolIEs.List); i++ {
		ie := uplinkNonUEAssociatedNRPPATransport.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRoutingID:
			routingID = ie.Value.RoutingID
			log.Trace("Decode IE RoutingID")

		case ngapType.ProtocolIEIDNRPPaPDU:
			nRPPaPDU = ie.Value.NRPPaPDU
			log.Trace("Decode IE NRPPaPDU")
		}
	}

	if routingID == nil {
		log.Error("RoutingID is nil")
		return
	}
	// Forward routingID to LMF
	// Described in (23.502 4.13.5.6)

	if nRPPaPDU == nil {
		log.Error("NRPPaPDU is nil")
		return
	}
	// TODO: Forward NRPPaPDU to LMF
}
