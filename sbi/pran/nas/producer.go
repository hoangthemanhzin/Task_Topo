package nas

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"net/http"
	"strconv"
)

func OnInitUeContextStatus(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {
		var input n2models.InitUeContextStatus
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandleInitUeContextStatus(ueId, &input); prob == nil {
				response.SetBody(200, nil)
			} else {
				response.SetProblem(prob)
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

func OnNasDl(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)

	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {
		var input n2models.NasDlMsg
		if err = ctx.DecodeRequest(&input); err == nil {
			if prob := prod.HandleNasDl(ueId, &input); prob == nil {
				response.SetBody(200, nil)
			} else {
				response.SetProblem(prob)
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

func OnInitCtxSet(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {

		var input n2models.InitCtxSetupReq
		if err = ctx.DecodeRequest(&input); err == nil {
			if rsp, ersp := prod.HandleInitCtxSetupReq(ueId, &input); ersp == nil {
				response.SetBody(200, rsp)
			} else {
				response.SetBody(400, ersp)
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

func OnUeCtxMod(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {

		var input n2models.UeCtxModReq
		if err = ctx.DecodeRequest(&input); err == nil {
			if rsp, ersp := prod.HandleUeCtxModReq(ueId, &input); ersp == nil {
				response.SetBody(200, rsp)
			} else {
				response.SetBody(400, ersp)
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

func OnUeCtxRel(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {

		var input n2models.UeCtxRelCmd
		if err = ctx.DecodeRequest(&input); err == nil {
			if rsp, prob := prod.HandleUeCtxRelCmd(ueId, &input); prob == nil {
				response.SetBody(200, rsp)
			} else {
				response.SetProblem(prob)
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

func OnPduSessResSet(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {
		var input n2models.PduSessResSetReq
		if err = ctx.DecodeRequest(&input); err == nil {
			if rsp, prob := prod.HandlePduSessResSetReq(ueId, &input); prob == nil {
				response.SetBody(200, rsp)
			} else {
				response.SetProblem(prob)
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

func OnPduSessResMod(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {
		var input n2models.PduSessResModReq
		if err = ctx.DecodeRequest(&input); err == nil {
			if rsp, prob := prod.HandlePduSessResModReq(ueId, &input); prob == nil {
				rsp = &n2models.PduSessResModRsp{}
				response.SetBody(200, rsp)
			} else {
				response.SetProblem(prob)
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

func OnPduSessResRel(ctx sbi.RequestContext, handler interface{}) (response sbi.Response) {
	prod := handler.(Producer)
	ueIdStr := ctx.Param("ueId")
	var err error
	var ueId int64
	if ueId, err = strconv.ParseInt(ueIdStr, 10, 64); err == nil {

		var input n2models.PduSessResRelCmd

		if err = ctx.DecodeRequest(&input); err == nil {
			if rsp, prob := prod.HandlePduSessResRelCmd(ueId, &input); prob == nil {
				response.SetBody(200, rsp)
			} else {
				response.SetProblem(prob)
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

type Producer interface {
	//a non-delivery notification should be send to the consumer if an error
	//occurs
	HandleInitUeContextStatus(int64, *n2models.InitUeContextStatus) *models.ProblemDetails
	HandleNasDl(int64, *n2models.NasDlMsg) *models.ProblemDetails

	HandleInitCtxSetupReq(int64, *n2models.InitCtxSetupReq) (*n2models.InitCtxSetupRsp, *n2models.InitCtxSetupFailure)
	HandleUeCtxModReq(int64, *n2models.UeCtxModReq) (*n2models.UeCtxModRsp, *n2models.UeCtxModFail)
	HandleUeCtxRelCmd(int64, *n2models.UeCtxRelCmd) (*n2models.UeCtxRelCmpl, *models.ProblemDetails)

	HandlePduSessResSetReq(int64, *n2models.PduSessResSetReq) (*n2models.PduSessResSetRsp, *models.ProblemDetails)
	HandlePduSessResModReq(int64, *n2models.PduSessResModReq) (*n2models.PduSessResModRsp, *models.ProblemDetails)
	HandlePduSessResRelCmd(int64, *n2models.PduSessResRelCmd) (*n2models.PduSessResRelRsp, *models.ProblemDetails)
}
