package producer

import (
	"etrib5gc/common"
	"etrib5gc/nfs/smf/sm"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
)

func (p *Producer) PDU_HandleReleasePduSession(pduSessionRef string, body *models.ReleaseData) (rsp *models.ReleasedData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleReleasePduSession has not been implemented")
	return
}
func (p *Producer) PDU_HandleRetrievePduSession(pduSessionRef string, body models.RetrieveData) (rsp *models.RetrievedData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleRetrievePduSession has not been implemented")
	return
}
func (p *Producer) PDU_HandleTransferMoData(pduSessionRef string, body models.TransferMoDataRequest) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleTransferMoData has not been implemented")
	return
}
func (p *Producer) PDU_HandleUpdatePduSession(pduSessionRef string, body models.HsmfUpdateData) (rsp *models.HsmfUpdatedData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleUpdatePduSession has not been implemented")
	return
}
func (p *Producer) PDU_HandleRetrieveSmContext(smContextRef string, body *models.SmContextRetrieveData) (rsp *models.SmContextRetrievedData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleRetrieveSmContext has not been implemented")
	return
}
func (p *Producer) PDU_HandleSendMoData(smContextRef string, body models.SendMoDataRequest) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleSendMoData has not been implemented")
	return
}

func (p *Producer) PDU_HandlePostPduSessions(body models.PduSessionCreateData) (rsp *models.PduSessionCreatedData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandlePostPduSessions has not been implemented")
	return
}

func (p *Producer) PDU_HandleReleaseSmContext(smContextRef string, body *models.ReleaseSmContextRequest) (rsp *models.SmContextReleasedData, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("PDU_HandleReleaseSmContext has not been implemented")
	return
}

func (p *Producer) PDU_HandleUpdateSmContext(ref string, body models.UpdateSmContextRequest) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse) {
	p.Infof("Receive UpdateSmContext from %s", ref)
	if smctx := p.ctx.FindSmContext(ref); smctx == nil {
		err := fmt.Errorf("SmContext[%s] not found", ref)
		p.Errorf(err.Error())
		ersp = &models.UpdateSmContextErrorResponse{
			JsonData: models.SmContextUpdateError{
				Error: &models.ExtProblemDetails{
					Status: http.StatusNotFound,
					Detail: err.Error(),
				},
				UpCnxState: models.UPCNXSTATE_DEACTIVATED,
			},
		}
	} else {
		var err error
		info := &sm.UpdateSmContextJob{
			Req: &body,
		}
		job := common.NewAsyncJob(info, 0)
		//Release the session
		if err = smctx.HandleEvent(&common.EventData{
			EvType:  sm.UPDATE_SMCONTEXT,
			Content: job,
		}); err == nil {
			if err = job.Wait(); err == nil {
				rsp = info.Rsp
				ersp = info.Ersp
			} else {
				p.Errorf("Handle UpdateSmContext return error: %s", err.Error())
			}
		}
		if err != nil {
			ersp = &models.UpdateSmContextErrorResponse{
				JsonData: models.SmContextUpdateError{
					Error: &models.ExtProblemDetails{
						Status: http.StatusInternalServerError,
						Detail: err.Error(),
					},
				},
			}

		}
	}
	return
}

func (p *Producer) PDU_HandlePostSmContexts(body models.PostSmContextsRequest, callback models.Callback) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse) {
	supi := body.JsonData.Supi
	sid := uint32(body.JsonData.PduSessionId)
	ref := common.SmContextRef(supi, sid)

	p.Infof("Receive PostSmContextRequests for %s", ref)

	var smctx *sm.SmContext
	var err error
	if smctx = p.ctx.FindSmContext(ref); smctx != nil {
		p.Infof("SmContext for %s exists, release it", ref)
		info := &sm.ReleaseSmContextJob{}
		job := common.NewAsyncJob(info, 0)
		//Release the session
		if err = smctx.HandleEvent(&common.EventData{
			EvType:  sm.RELEASE_SMCONTEXT,
			Content: job,
		}); err == nil {
			if err = job.Wait(); err != nil {
				p.Errorf("Release existing Smcontext return error: %s", err.Error())
			}
		}
		if err != nil {
			//fail to release existing session
			ersp = &models.PostSmContextsErrorResponse{
				JsonData: models.SmContextCreateError{
					Error: &models.ExtProblemDetails{
						Status: http.StatusInternalServerError,
						Detail: err.Error(),
					},
				},
			}
			return
		}

	} else {
		//err := fmt.Errorf("SmContext[%s] not found")
		//create a new SmContext
		if p.ctx.IsClosed() {
			err = fmt.Errorf("SMF is terminated")
		} else {
			if smctx, err = sm.CreateSmContext(p.ctx, p.upmanager, callback, &body.JsonData); err != nil {
				p.Errorf("Create SmContext failed: %s", err.Error())
			}
		}

		if err != nil {
			ersp = &models.PostSmContextsErrorResponse{
				JsonData: models.SmContextCreateError{
					Error: &models.ExtProblemDetails{
						Status: http.StatusInternalServerError,
						Detail: err.Error(),
					},
				},
			}
			return
		}
	}

	info := &sm.PostSmContextsJob{
		Req: &body,
	}
	job := common.NewAsyncJob(info, 0)
	if err = smctx.HandleEvent(&common.EventData{
		EvType:  sm.POST_SMCONTEXTS,
		Content: job,
	}); err == nil {
		if err = job.Wait(); err == nil {
			rsp = info.Rsp
			ersp = info.Ersp
		} else {
			p.Errorf("Handle PostSmContexts return error: %s", err.Error())
		}
	}
	if err != nil {
		ersp = &models.PostSmContextsErrorResponse{
			JsonData: models.SmContextCreateError{
				Error: &models.ExtProblemDetails{
					Status: http.StatusInternalServerError,
					Detail: err.Error(),
				},
			},
		}

	}
	return
}
