package ngap

import (
	"encoding/hex"
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/ngapConvert"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

func SendDeactivateTrace(ue *ue.UeContext, anType models.AccessType) (err error) {
	log.Info("Send Deactivate Trace")

	var pkt []byte
	if pkt, err = buildDeactivateTrace(ue, anType); err != nil {
		log.Errorf("Build DeactivateTrace failed : %s", err.Error())
		return
	}
	err = ue.Send(pkt)
	return
}

func buildDeactivateTrace(ue *ue.UeContext, anType models.AccessType) ([]byte, error) {
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeDeactivateTrace
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentDeactivateTrace
	initiatingMessage.Value.DeactivateTrace = new(ngapType.DeactivateTrace)

	deactivateTrace := initiatingMessage.Value.DeactivateTrace
	deactivateTraceIEs := &deactivateTrace.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.DeactivateTraceIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DeactivateTraceIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	deactivateTraceIEs.List = append(deactivateTraceIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.DeactivateTraceIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.DeactivateTraceIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()
	/*
		TODO: need information from AMF
			deactivateTraceIEs.List = append(deactivateTraceIEs.List, ie)
			udminfo := amfUe.UdmClient().Info()
			if udminfo.TraceData != nil {
				// NG-RAN TraceID
				ie = ngapType.DeactivateTraceIEs{}
				ie.Id.Value = ngapType.ProtocolIEIDNGRANTraceID
				ie.Criticality.Value = ngapType.CriticalityPresentIgnore
				ie.Value.Present = ngapType.DeactivateTraceIEsPresentNGRANTraceID
				ie.Value.NGRANTraceID = new(ngapType.NGRANTraceID)

				// TODO:composed of the following TS:32.422
				traceData := *udminfo.TraceData
				subStringSlice := strings.Split(traceData.TraceRef, "-")

				if len(subStringSlice) != 2 {
					//logger.NgapLog.Warningln("TraceRef format is not correct")
				}

				plmnID := models.PlmnId{}
				plmnID.Mcc = subStringSlice[0][:3]
				plmnID.Mnc = subStringSlice[0][3:]
				traceID, err := hex.DecodeString(subStringSlice[1])
				if err != nil {
					//logger.NgapLog.Errorf("[Build Error] DecodeString traceID error: %+v", err)
				}

				tmp := ngapConvert.PlmnIdToNgap(plmnID)
				traceReference := append(tmp.Value, traceID...)
				trsr := ue.Trsr()
				trsrNgap, err := hex.DecodeString(trsr)
				if err != nil {
					//logger.NgapLog.Errorf(
					//	"[Build Error] DecodeString trsr error: %+v", err)
				}
				ie.Value.NGRANTraceID.Value = append(traceReference, trsrNgap...)
				deactivateTraceIEs.List = append(deactivateTraceIEs.List, ie)
			}
	*/
	return libngap.Encoder(pdu)
}

func (h *Ngap) handleCellTrafficTrace(ran *ran.Ran, cellTrafficTrace *ngapType.CellTrafficTrace) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var nGRANTraceID *ngapType.NGRANTraceID
	var nGRANCGI *ngapType.NGRANCGI
	var traceCollectionEntityIPAddress *ngapType.TransportLayerAddress

	var ue *ue.UeContext

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	log.Info("Handle Cell Traffic Trace")

	for _, ie := range cellTrafficTrace.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			coreNgapId = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ranNgapId = ie.Value.RANUENGAPID
			log.Trace("Decode IE UeContextNgapID")

		case ngapType.ProtocolIEIDNGRANTraceID: // ignore
			nGRANTraceID = ie.Value.NGRANTraceID
			log.Trace("Decode IE NGRANTraceID")
		case ngapType.ProtocolIEIDNGRANCGI: // ignore
			nGRANCGI = ie.Value.NGRANCGI
			log.Trace("Decode IE NGRANCGI")
		case ngapType.ProtocolIEIDTraceCollectionEntityIPAddress: // ignore
			traceCollectionEntityIPAddress = ie.Value.TraceCollectionEntityIPAddress
			log.Trace("Decode IE TraceCollectionEntityIPAddress")
		}
	}
	if coreNgapId == nil {
		log.Error("AmfUeNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDAMFUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if ranNgapId == nil {
		log.Error("UeContextNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDRANUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procedureCode := ngapType.ProcedureCodeCellTrafficTrace
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentIgnore
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage, &procedureCriticality,
			&iesCriticalityDiagnostics)
		SendErrorIndication(ran, coreNgapId, ranNgapId, nil, &criticalityDiagnostics)
		return
	}

	if ue = ran.FindUe(ranNgapId, coreNgapId); ue == nil {
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		SendErrorIndication(ran, coreNgapId, ranNgapId, &cause, nil)
		return
	}

	log.Debugf("UE: AmfUeNgapID[%d], UeContextNgapID[%d]", ue.CuNgapId(), ue.RanNgapId())

	ue.SetTrsr(hex.EncodeToString(nGRANTraceID.Value[6:]))

	log.Tracef("TRSR[%s]", ue.Trsr())

	switch nGRANCGI.Present {
	case ngapType.NGRANCGIPresentNRCGI:
		plmnID := ngapConvert.PlmnIdToModels(nGRANCGI.NRCGI.PLMNIdentity)
		cellID := ngapConvert.BitStringToHex(&nGRANCGI.NRCGI.NRCellIdentity.Value)
		log.Debugf("NRCGI[plmn: %s, cellID: %s]", plmnID, cellID)
	case ngapType.NGRANCGIPresentEUTRACGI:
		plmnID := ngapConvert.PlmnIdToModels(nGRANCGI.EUTRACGI.PLMNIdentity)
		cellID := ngapConvert.BitStringToHex(&nGRANCGI.EUTRACGI.EUTRACellIdentity.Value)
		log.Debugf("EUTRACGI[plmn: %s, cellID: %s]", plmnID, cellID)
	}

	tceIpv4, tceIpv6 := ngapConvert.IPAddressToString(*traceCollectionEntityIPAddress)
	if tceIpv4 != "" {
		log.Debugf("TCE IP Address[v4: %s]", tceIpv4)
	}
	if tceIpv6 != "" {
		log.Debugf("TCE IP Address[v6: %s]", tceIpv6)
	}

	// TODO: TS 32.422 4.2.2.10
	// When AMF receives this new NG signalling message containing the Trace Recording Session Reference (TRSR)
	// and Trace Reference (TR), the AMF shall look up the SUPI/IMEI(SV) of the given call from its database and
	// shall send the SUPI/IMEI(SV) numbers together with the Trace Recording Session Reference and Trace Reference
	// to the Trace Collection Entity.
}
func buildTraceStart() ([]byte, error) {

	//var pdu ngapType.NGAPPDU
	//return libngap.Encoder(pdu)

	return nil, nil
}
