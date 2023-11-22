package ngap

import (
	"etrib5gc/common"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models/n2models"

	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handleInitialUEMessage(ran *ran.Ran, initialUEMessage *ngapType.InitialUEMessage, pdu []byte) {

	var ranNgapId *ngapType.RANUENGAPID
	var naspdu *ngapType.NASPDU
	var locinfo *ngapType.UserLocationInformation
	var rrcCause *ngapType.RRCEstablishmentCause
	var fiveGSTMSI *ngapType.FiveGSTMSI
	var amfSetId *ngapType.AMFSetID
	var ueContextReq *ngapType.UEContextRequest
	var allowedNSSAI *ngapType.AllowedNSSAI

	var critical ngapType.CriticalityDiagnosticsIEList

	log.Info("Receive an Initial UE Message from gnB")

	for _, ie := range initialUEMessage.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDRANUENGAPID, ngapType.TypeOfErrorPresentMissing)
				critical.List = append(critical.List, item)
			}
		case ngapType.ProtocolIEIDNASPDU: // reject
			naspdu = ie.Value.NASPDU
			log.Trace("Decode IE NasPdu")
			if naspdu == nil {
				log.Error("NasPdu is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDNASPDU,
					ngapType.TypeOfErrorPresentMissing)
				critical.List = append(critical.List, item)
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // reject
			locinfo = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
			if locinfo == nil {
				log.Error("UserLocationInformation is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDUserLocationInformation, ngapType.TypeOfErrorPresentMissing)
				critical.List = append(critical.List, item)
			}
		case ngapType.ProtocolIEIDRRCEstablishmentCause: // ignore
			rrcCause = ie.Value.RRCEstablishmentCause
			log.Trace("Decode IE RRCEstablishmentCause")
		case ngapType.ProtocolIEIDFiveGSTMSI: // optional, reject
			fiveGSTMSI = ie.Value.FiveGSTMSI
			log.Trace("Decode IE 5G-S-TMSI")
		case ngapType.ProtocolIEIDAMFSetID: // optional, ignore
			amfSetId = ie.Value.AMFSetID
			log.Trace("Decode IE AmfSetID")
		case ngapType.ProtocolIEIDUEContextRequest: // optional, ignore
			ueContextReq = ie.Value.UEContextRequest
			log.Trace("Decode IE UEContextRequest")
		case ngapType.ProtocolIEIDAllowedNSSAI: // optional, reject
			allowedNSSAI = ie.Value.AllowedNSSAI
			log.Trace("Decode IE Allowed NSSAI")
		}
	}

	if len(critical.List) > 0 {
		log.Trace("Has missing reject IE(s)")

		procedureCode := ngapType.ProcedureCodeInitialUEMessage
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentIgnore
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage, &procedureCriticality,
			&critical)
		SendErrorIndication(ran, nil, ranNgapId, nil, &criticalityDiagnostics)
	}

	uectx := ran.FindUe(ranNgapId, nil)

	//a brand new UeContext
	if uectx == nil {
		if h.ctx.IsClosed() {
			//Context has been closed, don't create any new UeContext
			log.Warnf("Context has been closed, can't create new UeContexts")
			return
		}
		uectx = ue.NewUeContext(ran, h.ctx, ranNgapId.Value, fiveGSTMSI, amfSetId, allowedNSSAI)
		//		ue.FindAmf( /*h*/ )
		//This is where you locate the right AMF to handle this UE
	}
	//up to now, the UeContext shoudl have an attached AmfConsumer that points to an
	//AMF (either a default one or a designated one)

	// TS 23.502 4.2.2.2.3 step 6a Nnrf_NFDiscovery_Request (NF type, AMF Set)
	// if aMFSetID != nil {
	// TODO: This is a rerouted message
	// TS 38.413: AMF shall, if supported, use the IE as described in TS 23.502
	// }

	// ng-ran propagate allowedNssai in the rerouted initial ue message (TS 38.413 8.6.5)
	// TS 23.502 4.2.2.2.3 step 4a Nnssf_NSSelection_Get
	// if allowedNSSAI != nil {
	// TODO: AMF should use it as defined in TS 23.502
	// }
	msg := &n2models.InitUeContextRequest{
		Access:         ran.Access(),
		RanNets:        ran.RanNets(),
		NasPdu:         naspdu.Value,
		ContextRequest: ueContextReq != nil,
		RrcCause:       uint8(rrcCause.Value),
	}
	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_INIT_UE,
		Content: msg,
	})
}

func (h *Ngap) handleUplinkNasTransport(ran *ran.Ran, uplinkNasTransport *ngapType.UplinkNASTransport) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var naspdu *ngapType.NASPDU
	var locinfo *ngapType.UserLocationInformation

	log.Info("Receive an Uplink NAS Transport message from gnB")

	for i := 0; i < len(uplinkNasTransport.ProtocolIEs.List); i++ {
		ie := uplinkNasTransport.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			log.Trace("Decode IE AmfUeNgapID")
			if coreNgapId = ie.Value.AMFUENGAPID; coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			log.Trace("Decode IE UeContextNgapID")
			if ranNgapId = ie.Value.RANUENGAPID; ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDNASPDU:
			log.Trace("Decode IE NasPdu")
			if naspdu = ie.Value.NASPDU; naspdu == nil {
				log.Error("naspdu is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation:
			log.Trace("Decode IE UserLocationInformation")
			if locinfo = ie.Value.UserLocationInformation; locinfo == nil {
				log.Error("UserLocationInformation is nil")
				return
			}
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		return
	}
	msg := &n2models.UlNasTransport{
		NasPdu: naspdu.Value,
	}
	if locinfo != nil {
		msg.Loc = locConvert(locinfo)
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_UL_NAS,
		Content: msg,
	})
}

func (h *Ngap) handleNasNonDeliveryIndication(ran *ran.Ran, nASNonDeliveryIndication *ngapType.NASNonDeliveryIndication) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var naspdu *ngapType.NASPDU
	var cause *ngapType.Cause

	log.Infof("Receive a NasNonDeliveryIndication from gnB")
	for _, ie := range nASNonDeliveryIndication.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			if coreNgapId = ie.Value.AMFUENGAPID; coreNgapId == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			if ranNgapId = ie.Value.RANUENGAPID; ranNgapId == nil {
				log.Error("UeContextNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDNASPDU:
			if naspdu = ie.Value.NASPDU; naspdu == nil {
				log.Error("NasPdu is nil")
				return
			}
		case ngapType.ProtocolIEIDCause:
			if cause = ie.Value.Cause; cause == nil {
				log.Error("Cause is nil")
				return
			}
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		return
	}
	msg := &n2models.NasNonDeliveryIndication{
		NasPdu: naspdu.Value,
		Cause:  causeConvert(cause),
	}
	uectx.HandleNgap(&common.EventData{
		EvType:  ue.NGAP_NAS_NON_DELIVERY,
		Content: msg,
	})
}
