package producer

import (
	"etrib5gc/sbi/models"
)

func (p *Producer) COMM_HandleCreateUEContext(ueContextId string, body models.CreateUEContextRequest) (rsp *models.CreateUEContextResponse, ersp *models.CreateUEContextErrorResponse) {
	//TODO: to be implemented
	panic("COMM_HandleCreateUEContext has not been implemented")
	return
}
func (p *Producer) COMM_HandleReleaseUEContext(ueContextId string, body models.UEContextRelease) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("COMM_HandleReleaseUEContext has not been implemented")
	return
}
func (p *Producer) COMM_HandleUEContextTransfer(ueContextId string, body models.UEContextTransferRequest) (rsp *models.UEContextTransferResponse, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("COMM_HandleUEContextTransfer has not been implemented")
	return
}
func (p *Producer) COMM_HandleN1N2MessageTransfer(ueContextId string, body models.N1N2MessageTransferRequest) (rsp *models.N1N2MessageTransferRspData, ersp *models.N1N2MessageTransferError) {
	//TODO: to be implemented
	panic("COMM_HandleN1N2MessageTransfer has not been implemented")
	return
}
