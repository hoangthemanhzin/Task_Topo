package ngap

import (
	"encoding/hex"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

func (h *Ngap) handleUERadioCapabilityInfoIndication(ran *ran.Ran, uERadioCapabilityInfoIndication *ngapType.UERadioCapabilityInfoIndication) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var radioCap *ngapType.UERadioCapability
	var radioCap4Paging *ngapType.UERadioCapabilityForPaging

	for i := 0; i < len(uERadioCapabilityInfoIndication.ProtocolIEs.List); i++ {
		ie := uERadioCapabilityInfoIndication.ProtocolIEs.List[i]
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
		case ngapType.ProtocolIEIDUERadioCapability:
			radioCap = ie.Value.UERadioCapability
			log.Trace("Decode IE UERadioCapability")
			if radioCap == nil {
				log.Error("UERadioCapability is nil")
				return
			}
		case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
			radioCap4Paging = ie.Value.UERadioCapabilityForPaging
			log.Trace("Decode IE UERadioCapabilityForPaging")
			if radioCap4Paging == nil {
				log.Error("UERadioCapabilityForPaging is nil")
				return
			}
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		log.Errorf("No UE Context[UeContextNgapID: %d]", ranNgapId.Value)
		return
	}
	log.Tracef("UeContextNgapID[%d] AmfUeNgapID[%d]", uectx.RanNgapId(), uectx.CuNgapId())
	log.Info("Handle UE Radio Capability Info Indication")

	var dat models.RadioCapabilityInfoIndication
	if radioCap != nil {
		dat.RadioCap = hex.EncodeToString(radioCap.Value)
	}
	if radioCap4Paging != nil {
		if radioCap4Paging.UERadioCapabilityForPagingOfNR != nil {
			dat.RadioCap4PagingNr = hex.EncodeToString(radioCap4Paging.UERadioCapabilityForPagingOfNR.Value)
		}
		if radioCap4Paging.UERadioCapabilityForPagingOfEUTRA != nil {
			dat.RadioCap4PagingEutra = hex.EncodeToString(radioCap4Paging.UERadioCapabilityForPagingOfEUTRA.Value)
		}
	}
	dummy(&dat)

	// TS 38.413 8.14.1.2/TS 23.502 4.2.8a step5/TS 23.501, clause 5.4.4.1.
	// send its most up to date UE Radio Capability information to the RAN in the N2 REQUEST message.

}

func SendUERadioCapabilityCheckRequest(ue *ue.UeContext) (err error) {

	log.Info("Send UE Radio Capability Check Request")
	var pkt []byte
	if pkt, err = buildUERadioCapabilityCheckRequest(ue); err != nil {
		log.Errorf("Build UERadioCapabilityCheckRequest failed : %s", err.Error())
		return
	}
	err = ue.Send(pkt)
	return
}

func buildUERadioCapabilityCheckRequest(ue *ue.UeContext) ([]byte, error) {

	msg := new(ngapType.InitiatingMessage)

	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityCheck
	msg.Criticality.Value = ngapType.CriticalityPresentReject

	msg.Value.Present = ngapType.InitiatingMessagePresentUERadioCapabilityCheckRequest
	msg.Value.UERadioCapabilityCheckRequest = new(ngapType.UERadioCapabilityCheckRequest)

	uERadioCapabilityCheckRequest := msg.Value.UERadioCapabilityCheckRequest
	uERadioCapabilityCheckRequestIEs := &uERadioCapabilityCheckRequest.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.UERadioCapabilityCheckRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UERadioCapabilityCheckRequestIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	uERadioCapabilityCheckRequestIEs.List = append(uERadioCapabilityCheckRequestIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.UERadioCapabilityCheckRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UERadioCapabilityCheckRequestIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()

	uERadioCapabilityCheckRequestIEs.List = append(uERadioCapabilityCheckRequestIEs.List, ie)

	pdu := ngapType.NGAPPDU{
		Present:           ngapType.NGAPPDUPresentInitiatingMessage,
		InitiatingMessage: msg,
	}

	// TODO:UE Radio Capability(optional)
	return libngap.Encoder(pdu)
}

func (h *Ngap) handleUERadioCapabilityCheckResponse(ran *ran.Ran, rsp *ngapType.UERadioCapabilityCheckResponse) {
	var (
		coreNgapId *ngapType.AMFUENGAPID
		ranNgapId  *ngapType.RANUENGAPID
		critical   *ngapType.CriticalityDiagnostics
		ims        *ngapType.IMSVoiceSupportIndicator
	)
	for _, ie := range rsp.ProtocolIEs.List {
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
		case ngapType.ProtocolIEIDIMSVoiceSupportIndicator:
			ims = ie.Value.IMSVoiceSupportIndicator
			log.Trace("Decode IE IMSVoiceSupportIndicator")
			if ims == nil {
				log.Error("iMSVoiceSupportIndicator is nil")
				return
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			critical = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	uectx := ran.FindUe(ranNgapId, coreNgapId)
	if uectx == nil {
		return
	}

	log.Info("Handle UE Radio Capability Check Response")

	// TODO: handle iMSVoiceSupportIndicator
	if critical != nil {
		printCriticalityDiagnostics(ran, critical)
	}
}
