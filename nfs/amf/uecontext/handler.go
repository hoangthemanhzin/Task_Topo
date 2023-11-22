package uecontext

import (
	"etrib5gc/common"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
)

// receive external events for handling
func (uectx *UeContext) HandleEvent(ev *common.EventData) (err error) {
	switch ev.EvType {
	case events.N1N2_TRANSFER:
		err = uectx.sendEvent(N1N2TransferEvent, ev.Content)
	case events.REGISTRATION_REQUEST:
		err = uectx.sendEvent(RegistrationRequestEvent, ev.Content)
	case events.REGISTRATION_CMPL:
		err = uectx.sendEvent(RegistrationDoneEvent, ev.Content)
	case events.UPDATE_SECMODE:
		err = uectx.sendEvent(UpdateSecmodeEvent, ev.Content)
	default:
		err = fmt.Errorf("Unknown event")
	}
	if err != nil {
		err = common.WrapError("HandleEvent failed", err)
		uectx.Error(err.Error())
	}
	return
}

func (uectx *UeContext) handleN1N2Transfer(job *common.AsyncJob) {
	var err error
	jobinfo, _ := job.Info().(*events.N1N2TransferJob)
	req := jobinfo.Req

	uectx.Infof("Handle N1N2messageTransfer")

	sid := int32(req.JsonData.PduSessionId)
	access := req.JsonData.TargetAccess

	var sc SessionContext
	if sc = uectx.FindSessionContext(sid); sc == nil {
		uectx.Errorf("SmContext not found [%d]", sid)

		jobinfo.Ersp = &models.N1N2MessageTransferError{
			ProblemDetails: models.ProblemDetails{
				Status: http.StatusNotFound,
				Detail: "Sm context not found",
			},
		}
		return
	}
	jobinfo.Extra = sc
	//ue.Infof("Found an SessionContext[%s] to  handle N1N2message transfer", sc.Ref())
	access = sc.Access()

	var ranue RanUe
	var ok bool
	//2. if the UE is in CM_CONNECTED for the given access, forward the message
	//to the state machine of that access
	if ranue, ok = uectx.ranfaces[access]; ok {
		uectx.Info("In CM_CONNECTED, ask RanUe to handle N1N2MessageTransfer")
		if err = ranue.HandleEvent(&common.EventData{
			EvType:  events.SEND_N1N2_TRANSFER,
			Content: job,
		}); err != nil {
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusNotFound,
					Detail: err.Error(),
				},
			}
			job.Done(nil)
			return
		}
	} else {
		uectx.Info("In CM_IDLE, need Paging/Notification")
		//3. else: UE is in CM_IDLE
		//3a. check for the cases to reject transfer
		if req.BinaryDataN2Information != nil &&
			req.JsonData.N2InfoContainer.SmInfo.N2InfoContent.NgapIeType == models.NGAPIETYPE_PDU_RES_REL_CMD {
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusConflict,
					Detail: "Ue in CM_IDLE",
				},
			}
			job.Done(nil)
			return
		}
		/*
			//TODO: need to check the specification again

			if !ue.IsRegistered(models.ACCESSTYPE__3_GPP_ACCESS) {
				//access is non-3gpp but the 3gpp face is no registered
				jobinfo.Ersp = &models.N1N2MessageTransferError{
					ProblemDetails: models.ProblemDetails{
						Status: http.StatusGatewayTimeout,
						Cause:  "Ue not reachable",
					},
				}
				return
			}
		*/
		//3b. if access is 3GPP -> send paging
		//3c. access is Non-3GPP
		//3c1. 3gpp access is connected -> send
		//notification through 3gpp access
		//3c2. else: send paging throuth non-3gpp
		if access == models.ACCESSTYPE__3_GPP_ACCESS {
			//TODO: create new ranue then send paging
			//1. Locate PRAN
			//2. Create RanUe
			//3. Send SEND_PAGING event to RanUe
			uectx.Errorf("Paging not yet implemented")
			jobinfo.Ersp = &models.N1N2MessageTransferError{
				ProblemDetails: models.ProblemDetails{
					Status: http.StatusConflict,
					Detail: "Paging is not implement",
				},
			}
			job.Done(nil)
		} else {
			if ranue, ok = uectx.ranfaces[models.ACCESSTYPE__3_GPP_ACCESS]; ok {
				//send notification to 3GPP face
				uectx.Errorf("Ask RanUe to send Notification")
				if err = ranue.HandleEvent(&common.EventData{
					EvType:  events.SEND_NOTIFICATION,
					Content: job,
				}); err != nil {
					jobinfo.Ersp = &models.N1N2MessageTransferError{
						ProblemDetails: models.ProblemDetails{
							Status: http.StatusNotFound,
							Detail: err.Error(),
						},
					}
					job.Done(nil)
					return
				}

			} else {
				//TODO create a NON-3GPP RanUe then send paging
				//1. Locate PRAN
				//2. Create RanUe
				//3. Send SEND_PAGING event to RanUe
				uectx.Errorf("Paging not yet implemented")
				jobinfo.Ersp = &models.N1N2MessageTransferError{
					ProblemDetails: models.ProblemDetails{
						Status: http.StatusConflict,
						Detail: "Paging is not implement",
					},
				}
				job.Done(nil)
			}
		}
	}

	return
}
