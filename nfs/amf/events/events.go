package events

import "etrib5gc/sbi/models"

// EventType for UeContext
const (
	PAGING_REQ           uint8 = iota
	N1N2_TRANSFER              //SBI - SMF
	REGISTRATION_REQUEST       //RanUe
	REGISTRATION_CMPL          //RanUe
	UPDATE_SECMODE             //RanUe
)

// EventType for RanUe
const (
	INIT_UE_CONTEXT       uint8 = iota //SBI
	NAS_NON_DELIVERY                   //SBI
	NAS_UL_TRANSPORT                   //SBI
	UECTX_REL_REQ                      //SBI
	PDU_NOTIFY                         //SBI
	PDU_MOD_IND                        //SBI
	RRC_INACT_TRAN_REP                 //SBI
	SEND_N1N2_TRANSFER                 //from UeContext
	SEND_NOTIFICATION                  //from UeContext
	SEND_PAGING                        //from UeContext
	REGISTRATION_GRANTED               //from UeContext to start registration
	REGISTRATION_REJECTED              //from UeContext to reject the registration
)

type N1N2TransferJob struct {
	Req   *models.N1N2MessageTransferRequest
	Rsp   *models.N1N2MessageTransferRspData
	Ersp  *models.N1N2MessageTransferError
	Extra interface{} //to hold found SmContext
}
