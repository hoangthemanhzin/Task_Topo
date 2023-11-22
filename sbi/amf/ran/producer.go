package ran

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"net/http"
	"strconv"
)

func OnInitUeContext(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(DProducer)
	callbackstr := ctx.Header("Callback")
	if len(callbackstr) == 0 {
		response.SetProblem(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "Callback is not set in the request header",
		})
		return
	}
	callback := models.Callback(callbackstr)
	var input n2models.InitUeContextRequest
	var err error
	if err = ctx.DecodeRequest(&input); err == nil {
		if rsp, prob := prod.HandleInitUeContext(callback, &input); prob != nil {
			response.SetProblem(prob)
		} else {
			response.SetBody(200, rsp)
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}
func OnUlNasTransport(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(DProducer)

	var input n2models.UlNasTransport
	var ueid int64
	var err error
	ueidstr := ctx.Param("ueId")
	if ueid, err = strconv.ParseInt(ueidstr, 10, 64); err == nil {
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandleUlNasTransport(ueid, &input); prob != nil {
				response.SetProblem(prob)
			} else {
				response.SetBody(200, nil)
			}
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

func OnNasNonDeliveryIndication(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(DProducer)

	var input n2models.NasNonDeliveryIndication
	var ueid int64
	var err error
	ueidstr := ctx.Param("ueId")
	if ueid, err = strconv.ParseInt(ueidstr, 10, 64); err == nil {
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandleNasNonDeliveryIndication(ueid, &input); prob != nil {
				response.SetProblem(prob)
			} else {
				response.SetBody(200, nil)
			}
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

func OnUeCtxRelReq(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input n2models.UeCtxRelReq
	var ueid int64
	var err error
	ueidstr := ctx.Param("ueId")
	if ueid, err = strconv.ParseInt(ueidstr, 10, 64); err == nil {
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandleUeCtxRelReq(ueid, &input); prob != nil {
				response.SetProblem(prob)
			} else {
				response.SetBody(200, nil)
			}
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

func OnRrcInactTranRep(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input n2models.RrcInactTranRep
	var ueid int64
	var err error
	ueidstr := ctx.Param("ueId")
	if ueid, err = strconv.ParseInt(ueidstr, 10, 64); err == nil {
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandleRrcInactTranRep(ueid, &input); prob != nil {
				response.SetProblem(prob)
			} else {
				response.SetBody(200, nil)
			}
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

func OnPduSessResNot(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input n2models.PduSessResNot
	var ueid int64
	var err error
	ueidstr := ctx.Param("ueId")
	if ueid, err = strconv.ParseInt(ueidstr, 10, 64); err == nil {
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandlePduSessResNot(ueid, &input); prob != nil {
				response.SetProblem(prob)
			} else {
				response.SetBody(200, nil)
			}
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

func OnPduSessResModInd(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	var input n2models.PduSessResModInd
	var ueid int64
	var err error
	ueidstr := ctx.Param("ueId")
	if ueid, err = strconv.ParseInt(ueidstr, 10, 64); err == nil {
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandlePduSessResModInd(ueid, &input); prob != nil {
				response.SetProblem(prob)
			} else {
				response.SetBody(200, nil)
			}
		}
	}

	if err != nil {
		response.SetProblem(&models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}

	return
}

type DProducer interface {
	HandleInitUeContext(models.Callback, *n2models.InitUeContextRequest) (*n2models.InitUeContextResponse, *models.ProblemDetails)
	HandleUlNasTransport(int64, *n2models.UlNasTransport) *models.ProblemDetails
	HandleNasNonDeliveryIndication(int64, *n2models.NasNonDeliveryIndication) *models.ProblemDetails
}
type Producer interface {
	DProducer

	HandleUeCtxRelReq(int64, *n2models.UeCtxRelReq) *models.ProblemDetails
	HandleRrcInactTranRep(int64, *n2models.RrcInactTranRep) *models.ProblemDetails

	//notification
	HandlePduSessResNot(int64, *n2models.PduSessResNot) *models.ProblemDetails
	HandlePduSessResModInd(int64, *n2models.PduSessResModInd) *models.ProblemDetails
}
