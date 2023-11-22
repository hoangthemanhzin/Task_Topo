package secmode

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi/utils/nasConvert"
	"etrib5gc/util/fsm"
	"fmt"
	"time"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

const (
	T3560_DURATION = 200 //miliseconds
)

// interface to a completed authentication procedure
type RanUe interface {
	logctx.LogWriter
	UeContext() *uecontext.UeContext //attached UeContext
	Worker() common.Executer
	UpdateSecCtx()
	SecCtx() *nas.SecCtx
	SendSecurityModeCommand(func(*nasMessage.SecurityModeCommand) error) error
}

type SecmodeProc struct {
	logctx.LogWriter
	fsm.State
	ranue RanUe

	eap     string
	hdp     uint8 //horizontal derive parameter (need kamf prime?)
	rinmr   bool  //retransmission of initial NAS message?
	success bool

	container []byte

	t3560 common.UeTimer

	onDone func(*SecmodeProc)
	err    error
}

func New(ranue RanUe, eap string, success bool, hdp uint8, rinmr bool, fn func(*SecmodeProc)) (proc *SecmodeProc) {
	proc = &SecmodeProc{
		LogWriter: ranue.WithFields(logctx.Fields{"mod": "secmod"}),
		ranue:     ranue,
		eap:       eap,
		success:   success,
		hdp:       hdp,
		rinmr:     rinmr,
		onDone:    fn,
		State:     fsm.NewState(SECMODE_IDLE),
	}
	proc.t3560 = common.NewTimer(T3560_DURATION*time.Millisecond, func() {
		//SecurityModeCommand expired
		proc.sendEvent(T3560Event, nil)
	}, nil)

	proc.sendEvent(StartEvent, nil)
	return
}
func (proc *SecmodeProc) NasContainer() []byte {
	return proc.container
}

func (proc *SecmodeProc) GetError() error {
	return proc.err
}

func (proc *SecmodeProc) Handle(msg *libnas.GmmMessage) (err error) {
	switch msg.GetMessageType() {
	case libnas.MsgTypeSecurityModeComplete:
		err = proc.sendEvent(SecmodeCompleteEvent, msg.SecurityModeComplete)
	case libnas.MsgTypeSecurityModeReject:
		err = proc.sendEvent(SecmodeRejectEvent, msg.SecurityModeReject)
	default: //ignore
		err = fmt.Errorf("Unknown Nas message to security mode establishment")
	}
	return
}

func (proc *SecmodeProc) sendEvent(ev fsm.EventType, args interface{}) error {
	return _sm.SendEvent(proc.ranue.Worker(), proc, ev, args)
}

func (proc *SecmodeProc) handleT3560() {
	err := fmt.Errorf("T3560 expired")
	proc.sendEvent(DoneEvent, err)
}

func (proc *SecmodeProc) handleSecurityModeComplete(msg *nasMessage.SecurityModeComplete) {
	proc.Info("Handle Security Mode Complete")
	proc.t3560.Stop()

	// update Kgnb/Kn3iwf
	proc.ranue.UpdateSecCtx()

	//update pei
	if msg.IMEISV != nil {
		pei := nasConvert.PeiToString(msg.IMEISV.Octet[:])
		proc.Infof("Got an IMEISV for the UE: %s", pei)
		proc.ranue.UeContext().UpdatePei(pei)
	}

	//decode nas container
	//a retransmission was requested in the SecurityModeCommand
	if msg.NASMessageContainer != nil {
		proc.Debug("SecurityModeComplete message has a NAS container, decode it")
		proc.container = msg.NASMessageContainer.GetNASMessageContainerContents()

	}
	proc.sendEvent(DoneEvent, nil)
}

func (proc *SecmodeProc) handleSecurityModeReject(msg *nasMessage.SecurityModeReject) {

	proc.Info("Handle Security Mode Reject")
	proc.t3560.Stop()

	cause := msg.Cause5GMM.GetCauseValue()
	proc.Warnf("Reject Cause: %s", nasMessage.Cause5GMMToString(cause))
	proc.sendEvent(DoneEvent, fmt.Errorf("UE reject the security mode command:%s", nasMessage.Cause5GMMToString(cause)))
	return

}
