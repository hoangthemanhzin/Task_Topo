package producer

import (
	"etrib5gc/common"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi/models"
	"net/http"
)

func (p *Producer) COMM_HandleN1N2MessageTransfer(ueid string, body models.N1N2MessageTransferRequest) (rsp *models.N1N2MessageTransferRspData, ersp *models.N1N2MessageTransferError) {
	log.Infof("Receive N1N2MessageTransfer [supi=%s]", ueid)

	//1. extract message, (get target access type)
	var uectx *uecontext.UeContext
	if uectx = p.ctx.FindUeBySupi(ueid); uectx == nil {
		log.Errorf("No UeContext for supi=", ueid)
		ersp = &models.N1N2MessageTransferError{
			ProblemDetails: models.ProblemDetails{
				Status: http.StatusNotFound,
				Detail: "Ue Context not found",
			},
		}
		return
	}
	info := &events.N1N2TransferJob{
		Req: &body,
	}

	job := common.NewAsyncJob(info, 500)
	var err error
	if err = uectx.HandleEvent(&common.EventData{
		EvType:  events.N1N2_TRANSFER,
		Content: job,
	}); err == nil {
		if err = job.Wait(); err == nil {
			rsp = info.Rsp
			ersp = info.Ersp
		}
	}
	if err != nil {
		log.Errorf(err.Error())
		ersp = &models.N1N2MessageTransferError{
			ProblemDetails: models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Detail: err.Error(),
			},
		}

	}

	return
}

func (p *Producer) COMM_HandleCreateUEContext(ueContextId string, body models.CreateUEContextRequest) (rsp *models.CreateUEContextResponse, ersp *models.CreateUEContextErrorResponse) {
	//TODO: to be implemented
	panic("COMM_HandleCreateUEContext has not been implemented")
	return
}

func (p *Producer) COMM_HandleUEContextTransfer(ueContextId string, body models.UEContextTransferRequest) (rsp *models.UEContextTransferResponse, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("COMM_HandleUEContextTransfer has not been implemented")
	return
}
