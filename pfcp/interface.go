package pfcp

import (
	"etrib5gc/pfcp/pfcpmsg"
	"etrib5gc/sbi"
)

type PFCPSession interface {
	UpfCli() sbi.ConsumerClient
	RemoteSeid() uint64
	LocalSeid() uint64
	FillDeletionRequest(*pfcpmsg.PFCPSessionDeletionRequest)
	FillEstablishmentRequest(*pfcpmsg.PFCPSessionEstablishmentRequest)
	FillModificationRequest(*pfcpmsg.PFCPSessionModificationRequest)
}

type PfcpSender interface {
	SendPfcpSessionDeletionRequest(PFCPSession) (rsp *pfcpmsg.PFCPSessionDeletionResponse, err error)
	SendPfcpSessionEstablishmentRequest(PFCPSession) (*pfcpmsg.PFCPSessionEstablishmentResponse, error)
	SendPfcpSessionModificationRequest(PFCPSession) (*pfcpmsg.PFCPSessionModificationResponse, error)
}
