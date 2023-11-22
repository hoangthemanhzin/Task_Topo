package producer

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/nfs/amf/ranuecontext"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"fmt"
	"net/http"

	libnas "github.com/free5gc/nas"
)

func (p *Producer) HandleInitUeContext(callback models.Callback, msg *n2models.InitUeContextRequest) (rsp *n2models.InitUeContextResponse, prob *models.ProblemDetails) {
	log.Infof("Receive InitUeContextRequest from %s", callback)
	var nasMsg libnas.Message
	var err error
	var ranstate *ranuecontext.RanUe
	var uectx *uecontext.UeContext

	//in case of a ServiceRequest, the RanUe is still existed at the AMF
	ranstate = p.ctx.FindRanUeByIdAtRan(msg.RanUeId)

	//NOTE: this invariance must be enforced: a RanUe is always attached to an
	//UeContext. Once it is detached, it must be gone (can't look up for it).
	//UeContext must not detach a RanUe that is stuck in a registration
	//procedure. New registration is not accepted if a previous one is on-going

	if ranstate == nil {
		//decode plain text content
		if nasMsg, err = nas.Decode(nil, msg.NasPdu); err != nil {
			log.Errorf("Decode plaintext Nas failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusBadRequest,
				Detail: err.Error(),
				Cause:  "Nas Decoding",
			}
			return
		}

		log.Infof("RanUe not found [ranUeId=%d]", msg.RanUeId)

		//extract UE identity from the message then find its context (or create
		//a new one).

		if uectx, err = p.ctx.SearchUeContext(nasMsg.GmmMessage); err != nil {
			log.Error("Search/Create UeContext failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
				Cause:  "Search/Create UeContext",
			}

			return
		}
		//Afterward, the UeContext is added to the pool

		//create a new RanUe, it points to the UeContext. However, the
		//UeContext has not attached the RanUe yet
		if ranstate, err = ranuecontext.NewRanUe(p.ctx, uectx, msg.Access, msg.RanNets, callback, msg.RanUeId); err != nil {
			log.Error("Create RanUeContext failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
				Cause:  "Create RanUe",
			}
			return
		}
	} else {
		//need to decode with a SecurityModeContext
		if nasMsg, err = nas.Decode(ranstate, msg.NasPdu); err != nil {
			log.Errorf("Deccode Nas message failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusBadRequest,
				Detail: err.Error(),
				Cause:  "Nas Decoding",
			}
			return
		} else {
			//make sure to have a valid N1Msg
			switch nasMsg.GmmMessage.GetMessageType() {
			case libnas.MsgTypeRegistrationRequest:
			case libnas.MsgTypeServiceRequest:
			default:
				log.Error("Unknown Nas message in InitialUeMsg")
				prob = &models.ProblemDetails{
					Status: http.StatusBadRequest,
					Detail: "Unknown N1Msg in InitialUeMsg",
					Cause:  "Nas Message",
				}
				return
			}
		}
	}

	if err = ranstate.HandleEvent(&common.EventData{
		EvType: events.INIT_UE_CONTEXT,
		Content: &events.InitUeContextData{
			InitUeMsg: msg,
			GmmMsg:    nasMsg.GmmMessage,
		},
	}); err != nil {
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "Invalid RanUe state",
		}
		return
	}
	//everything is fine, set the AmfUeId for the response
	rsp = &n2models.InitUeContextResponse{
		AmfUeId: ranstate.AmfUeId(),
	}
	return
}

func (p *Producer) HandleNasNonDeliveryIndication(ueid int64, msg *n2models.NasNonDeliveryIndication) (prob *models.ProblemDetails) {
	log.Infof("Receive NasNonDeliveryIndication")
	var err error
	if ue := p.ctx.FindRanUe(ueid); ue == nil {
		err = fmt.Errorf("RanUe not found [amfUeid=%d]", ueid)
		log.Error(err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: err.Error()}
		return
	} else {
		if err := ue.HandleEvent(&common.EventData{
			EvType:  events.NAS_NON_DELIVERY,
			Content: msg,
		}); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error()}
		}

	}

	return
}

func (p *Producer) HandleUlNasTransport(ueid int64, msg *n2models.UlNasTransport) (prob *models.ProblemDetails) {
	log.Infof("Receive UplinkNasTransport")
	var err error
	if ue := p.ctx.FindRanUe(ueid); ue == nil {
		err = fmt.Errorf("RanUe not found [amfUeid=%d]", ueid)
		log.Error(err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: err.Error()}
		log.Error(err.Error())
		return
	} else {
		if err := ue.HandleEvent(&common.EventData{
			EvType:  events.NAS_UL_TRANSPORT,
			Content: msg,
		}); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error()}
		}

	}

	return
}

func (p *Producer) HandleUeCtxRelReq(ueid int64, msg *n2models.UeCtxRelReq) (prob *models.ProblemDetails) {
	log.Infof("Receive  UeContextReleaseRequest")
	var err error
	if ue := p.ctx.FindRanUe(ueid); ue == nil {
		err = fmt.Errorf("RanUe not found [amfUeid=%d]", ueid)
		log.Error(err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: err.Error()}
		log.Error(err.Error())
		return
	} else {
		if err := ue.HandleEvent(&common.EventData{
			EvType:  events.UECTX_REL_REQ,
			Content: msg,
		}); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error()}
		}

	}

	return
}

func (p *Producer) HandleRrcInactTranRep(ueid int64, msg *n2models.RrcInactTranRep) (prob *models.ProblemDetails) {
	log.Infof("Receive RrcInactiveTranResponse")
	var err error
	if ue := p.ctx.FindRanUe(ueid); ue == nil {
		err = fmt.Errorf("RanUe not found [amfUeid=%d]", ueid)
		log.Error(err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: err.Error()}
		log.Error(err.Error())
		return
	} else {
		if err := ue.HandleEvent(&common.EventData{
			EvType:  events.RRC_INACT_TRAN_REP,
			Content: msg,
		}); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error()}
		}

	}
	return
}

func (p *Producer) HandlePduSessResNot(ueid int64, msg *n2models.PduSessResNot) (prob *models.ProblemDetails) {
	log.Infof("Receive PduSessionResourceNotify")
	var err error
	if ue := p.ctx.FindRanUe(ueid); ue == nil {
		err = fmt.Errorf("RanUe not found [amfUeid=%d]", ueid)
		log.Error(err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: err.Error()}
		log.Error(err.Error())
		return
	} else {
		if err := ue.HandleEvent(&common.EventData{
			EvType:  events.PDU_NOTIFY,
			Content: msg,
		}); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error()}
		}

	}
	return
}

func (p *Producer) HandlePduSessResModInd(ueid int64, msg *n2models.PduSessResModInd) (prob *models.ProblemDetails) {
	log.Infof("Receive PduSessionResourceModifyIndication")
	var err error
	if ue := p.ctx.FindRanUe(ueid); ue == nil {
		err = fmt.Errorf("RanUe not found [amfUeid=%d]", ueid)
		log.Error(err.Error())
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: err.Error()}
		log.Error(err.Error())
		return
	} else {
		if err := ue.HandleEvent(&common.EventData{
			EvType:  events.PDU_MOD_IND,
			Content: msg,
		}); err != nil {
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error()}
		}

	}
	return
}
