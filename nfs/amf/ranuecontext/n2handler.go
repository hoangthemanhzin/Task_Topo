package ranuecontext

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"fmt"
)

type smContextUpdateTask struct {
	sid      int32
	n2type   models.N2SmInfoType
	transfer []byte
}

type n1SmMsg struct {
	sid int32
	pdu []byte
}
type n2Info struct {
	n2type  models.N2SmInfoType
	session n2models.DlPduSessionResourceInfo
}

type smContextUpdateStatus struct {
	n1msglist  []n1SmMsg
	n2infolist []n2Info
}

func (status *smContextUpdateStatus) add(sid int32, rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse) {
	if ersp != nil {
	} else if rsp != nil {
		n1msg := rsp.BinaryDataN1SmMessage
		if len(rsp.BinaryDataN2SmInformation) > 0 {
			status.n2infolist = append(status.n2infolist, n2Info{
				n2type: rsp.JsonData.N2SmInfoType,
				session: n2models.DlPduSessionResourceInfo{
					Id:       int64(sid),
					Transfer: rsp.BinaryDataN2SmInformation,
					NasPdu:   n1msg,
				},
			})
		} else if len(n1msg) > 0 {
			status.n1msglist = append(status.n1msglist, n1SmMsg{
				sid: sid,
				pdu: n1msg,
			})
		}
	}
}

func (ranue *RanUe) sendN1N2Dl(status *smContextUpdateStatus) (err error) {
	var errors []string
	//send N1 messages
	for _, n1msg := range status.n1msglist {
		//TODO: sending in parallel
		if err = ranue.sendDlN1Sm(n1msg.sid, n1msg.pdu, 0); err != nil {
			ranue.logSendingReport("N1Sm Downlink", err)
			errors = append(errors, err.Error())
		}
	}
	sessmap := make(map[models.N2SmInfoType][]n2models.DlPduSessionResourceInfo)
	for _, n2info := range status.n2infolist {
		switch n2info.n2type {
		case models.N2SMINFOTYPE_PDU_RES_SETUP_REQ:
			fallthrough
		case models.N2SMINFOTYPE_PDU_RES_REL_CMD:
			fallthrough
		case models.N2SMINFOTYPE_PDU_RES_MOD_REQ:
			fallthrough
		case models.N2SMINFOTYPE_PDU_RES_MOD_CFM:
			if list, ok := sessmap[n2info.n2type]; ok {
				sessmap[n2info.n2type] = append(list, n2info.session)
			} else {
				sessmap[n2info.n2type] = []n2models.DlPduSessionResourceInfo{
					n2info.session,
				}
			}
		}
	}
	for n2type, sessions := range sessmap {
		//TODO: send in parallel
		switch n2type {
		case models.N2SMINFOTYPE_PDU_RES_SETUP_REQ:
			if err = ranue.sendPduSessionResourceSetupRequest(sessions, nil); err != nil {
				errors = append(errors, err.Error())
			}
		case models.N2SMINFOTYPE_PDU_RES_REL_CMD:
			if err = ranue.sendPduSessionResourceReleaseCommand(sessions, nil); err != nil {
				errors = append(errors, err.Error())
			}
		case models.N2SMINFOTYPE_PDU_RES_MOD_REQ:
			if err = ranue.sendPduSessionResourceModifyRequest(sessions, nil); err != nil {
				errors = append(errors, err.Error())
			}
		case models.N2SMINFOTYPE_PDU_RES_MOD_CFM:
			ranue.Warnf("send PduSessionModifyConfirm not implemented")
		}
	}

	//N2SMINFOTYPE_PDU_RES_NTY             N2SmInfoType = "PDU_RES_NTY"
	//N2SMINFOTYPE_PDU_RES_NTY_REL         N2SmInfoType = "PDU_RES_NTY_REL"
	//N2SMINFOTYPE_PDU_RES_MOD_IND         N2SmInfoType = "PDU_RES_MOD_IND"
	//N2SMINFOTYPE_PATH_SWITCH_REQ         N2SmInfoType = "PATH_SWITCH_REQ"
	//N2SMINFOTYPE_PATH_SWITCH_SETUP_FAIL  N2SmInfoType = "PATH_SWITCH_SETUP_FAIL"
	//N2SMINFOTYPE_PATH_SWITCH_REQ_ACK     N2SmInfoType = "PATH_SWITCH_REQ_ACK"
	//N2SMINFOTYPE_PATH_SWITCH_REQ_FAIL    N2SmInfoType = "PATH_SWITCH_REQ_FAIL"
	//N2SMINFOTYPE_HANDOVER_REQUIRED       N2SmInfoType = "HANDOVER_REQUIRED"
	//N2SMINFOTYPE_HANDOVER_CMD            N2SmInfoType = "HANDOVER_CMD"
	//N2SMINFOTYPE_HANDOVER_PREP_FAIL      N2SmInfoType = "HANDOVER_PREP_FAIL"
	//N2SMINFOTYPE_HANDOVER_REQ_ACK        N2SmInfoType = "HANDOVER_REQ_ACK"
	//N2SMINFOTYPE_HANDOVER_RES_ALLOC_FAIL N2SmInfoType = "HANDOVER_RES_ALLOC_FAIL"
	//N2SMINFOTYPE_SECONDARY_RAT_USAGE     N2SmInfoType = "SECONDARY_RAT_USAGE"
	//N2SMINFOTYPE_PDU_RES_MOD_IND_FAIL    N2SmInfoType = "PDU_RES_MOD_IND_FAIL"

	return
}

func (ranue *RanUe) smContextBatchUpdate(tasks []smContextUpdateTask, status *smContextUpdateStatus) {
	errch := make(chan error)
	update := func(task *smContextUpdateTask, ch chan error) {
		if sc := findSession(ranue.ue, task.sid); sc != nil {
			if rsp, ersp, err := sc.SendSmContextN2Info(task.n2type, task.transfer); err == nil {
				status.add(task.sid, rsp, ersp)
				ch <- nil
			} else {
				ch <- err
			}
		}
		ch <- fmt.Errorf("SmContext not found for session id %d", task.sid)
	}

	for _, task := range tasks {
		go update(&task, errch)

	}
	//wait for all tasks completed
	var err error
	for i := 0; i < len(tasks); i++ {
		if err = <-errch; err != nil {
			ranue.Errorf(err.Error())
		}
	}
}

func (ranue *RanUe) handleInitCtxSetupRsp(msg *n2models.InitCtxSetupRsp) (err error) {
	ranue.Info("Receive InitialContextSetupResponse")
	//TODO: some more handling
	return
}

func (ranue *RanUe) handleInitCtxSetupFail(msg *n2models.InitCtxSetupFailure) (err error) {
	ranue.Info("Receive InitialContextSetupFailure")
	//TODO: some more handling
	return
}

func (ranue *RanUe) handlePduSessResSetRsp(msg *n2models.PduSessResSetRsp) error {
	ranue.Info("Receive PduSessionResourceSetupResponse")
	status := &smContextUpdateStatus{}
	tasks := []smContextUpdateTask{}
	for _, session := range msg.SuccessList {
		ranue.Infof("Success session %v", session)
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_SETUP_RSP,
			transfer: session.Transfer,
		})
	}
	for _, session := range msg.FailedList {
		ranue.Infof("Failed session %v", session)
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_SETUP_FAIL,
			transfer: session.Transfer,
		})
	}

	ranue.smContextBatchUpdate(tasks, status)
	return ranue.sendN1N2Dl(status)
}

func (ranue *RanUe) handlePduSessResModRsp(msg *n2models.PduSessResModRsp) (err error) {
	ranue.Info("Receive PduSessionResourceModifyResponse")
	//TODO: update User Location
	status := &smContextUpdateStatus{}
	tasks := []smContextUpdateTask{}
	for _, session := range msg.SuccessList {
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_MOD_RSP,
			transfer: session.Transfer,
		})
	}
	for _, session := range msg.FailedList {
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_MOD_FAIL,
			transfer: session.Transfer,
		})
	}
	ranue.smContextBatchUpdate(tasks, status)
	return ranue.sendN1N2Dl(status)
}

func (ranue *RanUe) handlePduSessResRelRsp(msg *n2models.PduSessResRelRsp) (err error) {
	ranue.Info("Receive PduSessionResourceReleaseResponse")
	//TODO: update User Location
	status := &smContextUpdateStatus{}
	tasks := []smContextUpdateTask{}
	for _, session := range msg.List {
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_REL_RSP,
			transfer: session.Transfer,
		})
	}

	ranue.smContextBatchUpdate(tasks, status)
	return ranue.sendN1N2Dl(status)
}
func (ranue *RanUe) handlePduSessResNot(msg *n2models.PduSessResNot) (err error) {
	ranue.Info("Receive PduSessionResourceNotify")
	//TODO: update User Location
	status := &smContextUpdateStatus{}
	tasks := []smContextUpdateTask{}
	for _, session := range msg.NotifyList {
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_NTY,
			transfer: session.Transfer,
		})
	}
	for _, session := range msg.ReleasedList {
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_NTY_REL,
			transfer: session.Transfer,
		})
	}
	ranue.smContextBatchUpdate(tasks, status)
	return ranue.sendN1N2Dl(status)
}

func (ranue *RanUe) handlePduSessResModInd(msg *n2models.PduSessResModInd) (err error) {
	ranue.Info("Receive PduSessionResourceModifyIndication")
	status := &smContextUpdateStatus{}
	tasks := []smContextUpdateTask{}
	for _, session := range msg.ModifyList {
		tasks = append(tasks, smContextUpdateTask{
			sid:      int32(session.Id),
			n2type:   models.N2SMINFOTYPE_PDU_RES_MOD_IND,
			transfer: session.Transfer,
		})
	}
	ranue.smContextBatchUpdate(tasks, status)
	return ranue.sendN1N2Dl(status)
}
