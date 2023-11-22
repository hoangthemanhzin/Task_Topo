package ngap

import (
	"etrib5gc/nfs/pran/ran"
	"etrib5gc/nfs/pran/ue"
	"fmt"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

// AOI List is from SMF
// The SMF may subscribe to the UE mobility event notification from the AMF
// (e.g. location reporting, UE moving into or out of Area Of Interest) TS 23.502 4.3.2.2.1 Step.17
// The Location Reporting Control message shall identify the UE for which reports are requested and may include
// Reporting Type, Location Reporting Level, Area Of Interest and Request Reference ID
// TS 23.502 4.10 LocationReportingProcedure
// The AMF may request the NG-RAN location reporting with event reporting type (e.g. UE location or UE presence
// in Area of Interest), reporting mode and its related parameters (e.g. number of reporting) TS 23.501 5.4.7
// Location Reference ID To Be Cancelled IE shall be present if the Event Type IE is set to "Stop UE presence
// in the area of interest". otherwise set it to 0
func SendLocationReportingControl(
	ue *ue.UeContext,
	AOIList *ngapType.AreaOfInterestList,
	LocationReportingReferenceIDToBeCancelled int64,
	eventType ngapType.EventType) (err error) {
	defer logSendingReport("LocationReportingControl", err)

	if eventType.Value == ngapType.EventTypePresentStopUePresenceInAreaOfInterest {
		if LocationReportingReferenceIDToBeCancelled < 1 || LocationReportingReferenceIDToBeCancelled > 64 {
			err = fmt.Errorf("LocationReportingReferenceIDToBeCancelled out of range (should be 1 ~ 64)")
			log.Error(err.Error())
			return
		}
	}

	var pkt []byte
	if pkt, err = buildLocationReportingControl(ue, AOIList, LocationReportingReferenceIDToBeCancelled, eventType); err != nil {
		err = fmt.Errorf("Build LocationReportingControl failed : %s", err.Error())
		log.Error(err.Error())
	}
	err = ue.Send(pkt)
	return
}

func buildLocationReportingControl(
	ue *ue.UeContext,
	AOIList *ngapType.AreaOfInterestList,
	LocationReportingReferenceIDToBeCancelled int64,
	eventType ngapType.EventType) ([]byte, error) {

	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeLocationReportingControl
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentLocationReportingControl
	initiatingMessage.Value.LocationReportingControl = new(ngapType.LocationReportingControl)

	locationReportingControl := initiatingMessage.Value.LocationReportingControl
	locationReportingControlIEs := &locationReportingControl.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.LocationReportingControlIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.LocationReportingControlIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	locationReportingControlIEs.List = append(locationReportingControlIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.LocationReportingControlIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.LocationReportingControlIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()

	locationReportingControlIEs.List = append(locationReportingControlIEs.List, ie)

	// Location Reporting Request Type
	ie = ngapType.LocationReportingControlIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	ie.Criticality.Value = ngapType.CriticalityPresentIgnore
	ie.Value.Present = ngapType.LocationReportingControlIEsPresentLocationReportingRequestType
	ie.Value.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)

	locationReportingRequestType := ie.Value.LocationReportingRequestType

	// Event Type
	locationReportingRequestType.EventType = eventType

	// Report Area in Location Reporting Request Type
	locationReportingRequestType.ReportArea.Value = ngapType.ReportAreaPresentCell // only this enum

	// AOI List in Location Reporting Request Type
	if AOIList != nil {
		locationReportingRequestType.AreaOfInterestList = new(ngapType.AreaOfInterestList)
		areaOfInterestList := locationReportingRequestType.AreaOfInterestList
		areaOfInterestList.List = AOIList.List
	}

	// location reference ID to be Cancelled [Conditional]
	if locationReportingRequestType.EventType.Value ==
		ngapType.EventTypePresentStopUePresenceInAreaOfInterest {
		locationReportingRequestType.LocationReportingReferenceIDToBeCancelled =
			new(ngapType.LocationReportingReferenceID)
		locationReportingRequestType.LocationReportingReferenceIDToBeCancelled.Value =
			LocationReportingReferenceIDToBeCancelled
	}

	locationReportingControlIEs.List = append(locationReportingControlIEs.List, ie)

	return libngap.Encoder(pdu)
}
func (h *Ngap) handleLocationReportingFailureIndication(ran *ran.Ran, locationReportingFailureIndication *ngapType.LocationReportingFailureIndication) {

	log.Info("Handle Location Reporting Failure Indication")

	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var ue *ue.UeContext

	var cause *ngapType.Cause

	for i := 0; i < len(locationReportingFailureIndication.ProtocolIEs.List); i++ {
		ie := locationReportingFailureIndication.ProtocolIEs.List[i]
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
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		}
	}

	ue = ran.FindUe(ranNgapId, coreNgapId)
	if ue == nil {
		return
	}
	//TODO: tqtung - it seems free5gc have not implement the handling of this event
}

func (h *Ngap) handleLocationReport(ran *ran.Ran, locationReport *ngapType.LocationReport) {
	var coreNgapId *ngapType.AMFUENGAPID
	var ranNgapId *ngapType.RANUENGAPID
	var userLocationInformation *ngapType.UserLocationInformation
	var uEPresenceInAreaOfInterestList *ngapType.UEPresenceInAreaOfInterestList
	var locationReportingRequestType *ngapType.LocationReportingRequestType

	log.Info("Handle Location Report")

	for _, ie := range locationReport.ProtocolIEs.List {
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
		case ngapType.ProtocolIEIDUserLocationInformation: // ignore
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE userLocationInformation")
			if userLocationInformation == nil {
				log.Warn("userLocationInformation is nil")
			}
		case ngapType.ProtocolIEIDUEPresenceInAreaOfInterestList: // optional, ignore
			uEPresenceInAreaOfInterestList = ie.Value.UEPresenceInAreaOfInterestList
			log.Trace("Decode IE uEPresenceInAreaOfInterestList")
			if uEPresenceInAreaOfInterestList == nil {
				log.Warn("uEPresenceInAreaOfInterestList is nil [optional]")
			}
		case ngapType.ProtocolIEIDLocationReportingRequestType: // ignore
			locationReportingRequestType = ie.Value.LocationReportingRequestType
			log.Trace("Decode IE LocationReportingRequestType")
			if locationReportingRequestType == nil {
				log.Warn("LocationReportingRequestType is nil")
			}
		}
	}

	ue := ran.FindUe(ranNgapId, coreNgapId)
	if ue == nil {
		return
	}

	//ue.UpdateLocInfo(userLocationInformation)

	log.Tracef("Report Area[%d]", locationReportingRequestType.ReportArea.Value)

	switch locationReportingRequestType.EventType.Value {
	case ngapType.EventTypePresentDirect:
		log.Trace("To report directly")

	case ngapType.EventTypePresentChangeOfServeCell:
		log.Trace("To report upon change of serving cell")

	case ngapType.EventTypePresentUePresenceInAreaOfInterest:
		log.Trace("To report UE presence in the area of interest")
		for _, uEPresenceInAreaOfInterestItem := range uEPresenceInAreaOfInterestList.List {
			uEPresence := uEPresenceInAreaOfInterestItem.UEPresence.Value
			referenceID := uEPresenceInAreaOfInterestItem.LocationReportingReferenceID.Value

			for _, AOIitem := range locationReportingRequestType.AreaOfInterestList.List {
				if referenceID == AOIitem.LocationReportingReferenceID.Value {
					log.Tracef("uEPresence[%d], presence AOI ReferenceID[%d]", uEPresence, referenceID)
				}
			}
		}

	case ngapType.EventTypePresentStopChangeOfServeCell:
		log.Trace("To stop reporting at change of serving cell")
		SendLocationReportingControl(ue, nil, 0, locationReportingRequestType.EventType)
		// TODO: Clear location report

	case ngapType.EventTypePresentStopUePresenceInAreaOfInterest:
		log.Trace("To stop reporting UE presence in the area of interest")
		log.Tracef("ReferenceID To Be Cancelled[%d]",
			locationReportingRequestType.LocationReportingReferenceIDToBeCancelled.Value)
		// TODO: Clear location report

	case ngapType.EventTypePresentCancelLocationReportingForTheUe:
		log.Trace("To cancel location reporting for the UE")
		// TODO: Clear location report
	}
}
