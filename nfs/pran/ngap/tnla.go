package ngap

import (
	"etrib5gc/nfs/pran/ue"

	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

func SendUETNLABindingReleaseRequest(ue *ue.UeContext) (err error) {
	log.Info("Send UE TNLA Binging Release Request")
	var pkt []byte
	if pkt, err = buildUETNLABindingReleaseRequest(ue); err != nil {
		log.Errorf("Build UETNLABindingReleaseRequest failed : %s", err.Error())
		return
	}
	err = ue.Send(pkt)
	return
}

func buildUETNLABindingReleaseRequest(ue *ue.UeContext) ([]byte, error) {
	var pdu ngapType.NGAPPDU

	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)

	initiatingMessage := pdu.InitiatingMessage
	initiatingMessage.ProcedureCode.Value = ngapType.ProcedureCodeUETNLABindingRelease
	initiatingMessage.Criticality.Value = ngapType.CriticalityPresentIgnore

	initiatingMessage.Value.Present = ngapType.InitiatingMessagePresentUETNLABindingReleaseRequest
	initiatingMessage.Value.UETNLABindingReleaseRequest = new(ngapType.UETNLABindingReleaseRequest)

	uETNLABindingReleaseRequest := initiatingMessage.Value.UETNLABindingReleaseRequest
	uETNLABindingReleaseRequestIEs := &uETNLABindingReleaseRequest.ProtocolIEs

	// AMF UE NGAP ID
	ie := ngapType.UETNLABindingReleaseRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UETNLABindingReleaseRequestIEsPresentAMFUENGAPID
	ie.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)

	aMFUENGAPID := ie.Value.AMFUENGAPID
	aMFUENGAPID.Value = ue.CuNgapId()

	uETNLABindingReleaseRequestIEs.List = append(uETNLABindingReleaseRequestIEs.List, ie)

	// RAN UE NGAP ID
	ie = ngapType.UETNLABindingReleaseRequestIEs{}
	ie.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	ie.Criticality.Value = ngapType.CriticalityPresentReject
	ie.Value.Present = ngapType.UETNLABindingReleaseRequestIEsPresentRANUENGAPID
	ie.Value.RANUENGAPID = new(ngapType.RANUENGAPID)

	rANUENGAPID := ie.Value.RANUENGAPID
	rANUENGAPID.Value = ue.RanNgapId()

	uETNLABindingReleaseRequestIEs.List = append(uETNLABindingReleaseRequestIEs.List, ie)

	return libngap.Encoder(pdu)
}
