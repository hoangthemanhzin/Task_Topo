package producer

import (
	"etrib5gc/common"
	"etrib5gc/nfs/damf/ue"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"net/http"
)

func (p *Producer) HandleInitUeContext(callback models.Callback, msg *n2models.InitUeContextRequest) (rsp *n2models.InitUeContextResponse, prob *models.ProblemDetails) {
	log.Info("Receive InitialUeMessage")
	var uectx *ue.UeContext
	var err error
	log.Tracef("Message with callback: %s", string(callback))
	if !p.ctx.HasUe(msg.RanUeId) {
		if uectx, err = ue.CreateUeContext(p.ctx, callback, msg); err != nil {
			prob = &models.ProblemDetails{
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			}
			log.Errorf("Create UeContext [ranUeId=%d] failed: %s", msg.RanUeId, err.Error())
			return
		}
	} else {
		log.Errorf("UeContext[ranUeId=%d] notexisted", msg.RanUeId)
		prob = &models.ProblemDetails{
			Detail: "UeContext existed",
			Status: http.StatusConflict,
		}
	}

	//add ue to the list
	p.ctx.AddUe(uectx)

	//send event to the UeContext for handling
	if err := uectx.HandleSbi(&common.EventData{
		EvType:  ue.NAS_INIT_UE,
		Content: msg,
	}); err != nil {
		log.Errorf("Message not handled: %s", err.Error())
		prob = &models.ProblemDetails{
			Detail: err.Error(),
			Status: http.StatusConflict,
		}
	} else {
		rsp = &n2models.InitUeContextResponse{
			AmfUeId: uectx.AmfUeId(),
		}
	}

	return
}
func (p *Producer) HandleNasNonDeliveryIndication(ueid int64, msg *n2models.NasNonDeliveryIndication) (prob *models.ProblemDetails) {
	log.Info("Receive NasNonDelivery")
	if uectx := p.ctx.FindUe(ueid); uectx == nil {
		log.Errorf("UeContext not found [ueid=%d]", ueid)
		prob = &models.ProblemDetails{
			Detail: "Uecontext not found",
			Status: http.StatusNotFound,
		}
	} else if err := uectx.HandleSbi(&common.EventData{
		EvType:  ue.NAS_NON_DELIVERY,
		Content: msg,
	}); err != nil {
		log.Errorf("Message not handled: %s", err.Error())
		prob = &models.ProblemDetails{
			Detail: err.Error(),
			Status: http.StatusConflict,
		}

	}

	return
}

func (p *Producer) HandleUlNasTransport(ueid int64, msg *n2models.UlNasTransport) (prob *models.ProblemDetails) {
	log.Info("Receive UplinkNasTransport")
	if uectx := p.ctx.FindUe(ueid); uectx == nil {
		log.Errorf("UeContext not found [ueid=%d]", ueid)
		prob = &models.ProblemDetails{
			Detail: "Uecontext not found",
			Status: http.StatusNotFound,
		}
	} else if err := uectx.HandleSbi(&common.EventData{
		EvType:  ue.NAS_UL_TRANSPORT,
		Content: msg,
	}); err != nil {
		log.Errorf("Message not handled: %s", err.Error())
		prob = &models.ProblemDetails{
			Detail: err.Error(),
			Status: http.StatusConflict,
		}

	}
	return
}
