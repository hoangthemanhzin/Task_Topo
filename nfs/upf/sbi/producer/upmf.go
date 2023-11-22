package producer

import (
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/pfcp/pfcptypes"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"time"
)

func (prod *Producer) HandleHeartbeat(req *n42.HeartbeatRequest) (rsp *n42.HeartbeatResponse, prob *models.ProblemDetails) {
	prod.Infof("Receive HeartbeatRequest from UPMF")
	prod.Infof("Nonce=%d, Time=%s", req.Nonce, req.Msg.RecoveryTimeStamp.RecoveryTimeStamp.String())
	rsp = &n42.HeartbeatResponse{
		Nonce: req.Nonce,
		Msg: pfcpmsg.HeartbeatResponse{
			RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
				RecoveryTimeStamp: time.Now(),
			},
		},
	}
	return
}

func (prod *Producer) HandleActivate(req *n42.UpfActivateQuery) (rsp *n42.UpfActivate, prob *models.ProblemDetails) {
	prod.Infof("Receive ActivateRequest from UPMF")
	// Should I add logic code here ???
	rsp.Msg["upf"] = "test"
	return
}
