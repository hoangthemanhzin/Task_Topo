package producer

import (
	"etrib5gc/common"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"net/http"
)

const (
	SBI_JOB_TIMEOUT = 200 //millisecond
)

func (p *Producer) HandleInitUeContextStatus(ueid int64, msg *n2models.InitUeContextStatus) (prob *models.ProblemDetails) {
	log.Infof("Receive a InitUeContextStatus from AMF for UE [cuNgapId=%d]", ueid)
	var err error
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		job := ue.NewSbiJob(&ue.InitUeContextStatusJob{
			Msg: msg,
		}, -1)

		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_INIT_UE_STATUS,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}
		if err != nil {
			uectx.WithFields(_logfields).Errorf("Handle InitUeContextStatus failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
			}
		}

	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "Context not found",
		}
	}

	return
}

func (p *Producer) HandleNasDl(ueid int64, msg *n2models.NasDlMsg) (prob *models.ProblemDetails) {
	log.Infof("Receive a Downlink Nas message from AMF for UE [id=%d]", ueid)
	var err error
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		job := ue.NewSbiJob(&ue.NasDlJob{
			Msg: msg,
		}, -1)
		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_NAS_DL,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}

		if err != nil {
			uectx.WithFields(_logfields).Errorf("Handle NasDl failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusBadGateway,
				Cause:  err.Error(),
			}
		}
	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "Context not found",
		}
	}
	return

}

func (p *Producer) HandleInitCtxSetupReq(ueid int64, msg *n2models.InitCtxSetupReq) (rsp *n2models.InitCtxSetupRsp, ersp *n2models.InitCtxSetupFailure) {
	log.Infof("Receive a InitCtxSetupReq from AMF for UE [id=%d]", ueid)
	var err error
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		jobinfo := &ue.InitCtxSetupReqJob{
			Msg: msg,
		}
		job := ue.NewSbiJob(jobinfo, SBI_JOB_TIMEOUT)
		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_UE_SET_REQ,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}
		if err != nil {
			uectx.WithFields(_logfields).Errorf("Handle InitCtxSetupReq failed: %s", err.Error())
			ersp = &n2models.InitCtxSetupFailure{
				Cause: n2models.Cause{}, //TODO: set cause values
			}
		} else {
			rsp = jobinfo.Rsp
			ersp = jobinfo.Ersp
		}
	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		ersp = &n2models.InitCtxSetupFailure{
			Cause: n2models.Cause{}, //TODO: set cause values
		}
	}
	return
}

func (p *Producer) HandleUeCtxModReq(ueid int64, msg *n2models.UeCtxModReq) (rsp *n2models.UeCtxModRsp, ersp *n2models.UeCtxModFail) {
	log.Infof("Receive a UeCtxModRe from AMF for UE [id=%d]", ueid)
	var err error
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		jobinfo := &ue.UeCtxModReqJob{
			Msg: msg,
		}
		job := ue.NewSbiJob(jobinfo, SBI_JOB_TIMEOUT)
		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_UE_MOD_REQ,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}
		if err != nil {
			uectx.WithFields(_logfields).Errorf("Handle UeCtxModReq failed: %s]", err.Error())
			ersp = &n2models.UeCtxModFail{
				Cause: n2models.Cause{}, //TODO: set cause values
			}
		} else {
			rsp = jobinfo.Rsp
			ersp = jobinfo.Ersp
		}
	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		ersp = &n2models.UeCtxModFail{
			Cause: n2models.Cause{}, //TODO: set cause values
		}
	}
	return

}

func (p *Producer) HandleUeCtxRelCmd(ueid int64, body *n2models.UeCtxRelCmd) (rsp *n2models.UeCtxRelCmpl, prob *models.ProblemDetails) {
	panic("HandleUeCtxRelCmd not implemented")
	return
}

func (p *Producer) HandlePduSessResSetReq(ueid int64, msg *n2models.PduSessResSetReq) (rsp *n2models.PduSessResSetRsp, prob *models.ProblemDetails) {
	log.Infof("Receive a PduSessResSetReq from AMF for UE [id=%d]", ueid)
	var err error
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		jobinfo := &ue.PduSessResSetReqJob{
			Msg: msg,
		}
		job := ue.NewSbiJob(jobinfo, SBI_JOB_TIMEOUT)
		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_PDU_SET_REQ,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}
		if err != nil {
			uectx.WithFields(_logfields).Errorf("Handle PduSessResSetReq failed: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
			}
		} else {
			rsp = jobinfo.Rsp
		}

	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Detail: "Context not found",
		}
	}

	return
}

func (p *Producer) HandlePduSessResModReq(ueid int64, msg *n2models.PduSessResModReq) (rsp *n2models.PduSessResModRsp, prob *models.ProblemDetails) {
	var err error
	log.Infof("Receive a PduSessResModReq from AMF for UE [id=%d]", ueid)
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		jobinfo := &ue.PduSessResModReqJob{
			Msg: msg,
		}
		job := ue.NewSbiJob(jobinfo, SBI_JOB_TIMEOUT)
		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_PDU_MOD_REQ,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}
		if err != nil {
			uectx.WithFields(_logfields).Errorf("Fail to send PduSessionResourceModifyRequest: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
			}
		} else {
			rsp = jobinfo.Rsp
		}

	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "Context not found",
		}
	}

	return
}

func (p *Producer) HandlePduSessResRelCmd(ueid int64, msg *n2models.PduSessResRelCmd) (rsp *n2models.PduSessResRelRsp, prob *models.ProblemDetails) {
	log.Infof("Receive a PduSessResRelCmd from AMF for UE [id=%d]", ueid)
	var err error
	if uectx := p.ctx.FindByCuNgapId(ueid); uectx != nil {
		jobinfo := &ue.PduSessResRelCmdJob{
			Msg: msg,
		}
		job := ue.NewSbiJob(jobinfo, SBI_JOB_TIMEOUT)
		if err = uectx.HandleSbi(&common.EventData{
			EvType:  ue.SBI_PDU_REL_CMD,
			Content: job,
		}); err == nil {
			err = job.Wait()
		}
		if err != nil {
			uectx.WithFields(_logfields).Errorf("Fail to send PduSessionResourceReleaseCommand: %s", err.Error())
			prob = &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
			}
		} else {
			rsp = jobinfo.Rsp
		}

	} else {
		log.Errorf("UeContext not found [id= %d]", ueid)
		prob = &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "Context not found",
		}
	}

	return
}
