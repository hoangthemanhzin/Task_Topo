package producer

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n41"
	"fmt"
	"net/http"
)

func (prod *Producer) HandleSessionEstablishment(req *n41.SessionEstablishmentRequest) (rsp *n41.SessionEstablishmentResponse, prob *models.ProblemDetails) {
	prod.Infof("Receive SessionEstablishmentRequest from SMF")
	prob = &models.ProblemDetails{
		Status: http.StatusInternalServerError,
		Detail: fmt.Sprintf("Not implemented"),
	}
	//TODO: call the handler to handle this message from free5gc UPF
	return
}

func (prod *Producer) HandleSessionModification(seid uint64, req *n41.SessionModificationRequest) (rsp *n41.SessionModificationResponse, prob *models.ProblemDetails) {
	prod.Infof("Receive SessionModificationRequest from SMF")
	prob = &models.ProblemDetails{
		Status: http.StatusInternalServerError,
		Detail: fmt.Sprintf("Not implemented"),
	}
	//TODO: call the handler to handle this message from free5gc UPF
	return
}

func (prod *Producer) HandleSessionDeletion(seid uint64, req *n41.SessionDeletionRequest) (rsp *n41.SessionDeletionResponse, prob *models.ProblemDetails) {
	prod.Infof("Receive SessionModificationRequest from SMF")
	prob = &models.ProblemDetails{
		Status: http.StatusInternalServerError,
		Detail: fmt.Sprintf("Not implemented"),
	}
	//TODO: call the handler to handle this message from free5gc UPF
	return
}
