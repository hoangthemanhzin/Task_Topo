package identification

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/sbi/utils/nasConvert"
	"etrib5gc/util/fsm"
	"fmt"
	"time"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

const (
	MAX_FAILURE_COUNT int = 2
	T3570_DURATION        = 100 //miliseconds
)

type RanUe interface {
	logctx.LogWriter
	Worker() common.Executer
	SendIdentityRequest(idtype uint8) error
}

type IdProc struct {
	logctx.LogWriter
	fsm.State
	ranue        RanUe
	failurecount int
	idtype       uint8
	id           []byte
	err          error
	t3570        common.UeTimer
	onDone       func(*IdProc)
}

func New(ranue RanUe, idtype uint8, fn func(proc *IdProc)) (proc *IdProc) {
	proc = &IdProc{
		LogWriter: ranue.WithFields(logctx.Fields{"mod": "identification"}),
		ranue:     ranue,
		idtype:    idtype,
		State:     fsm.NewState(PROC_IDLE),
		onDone:    fn,
	}
	proc.t3570 = common.NewTimer(T3570_DURATION*time.Millisecond, func() {
		//IdentityRequest expired
		proc.Infof("IdentificationRequest timer expired")
		proc.sendEvent(T3570Event, nil)
	}, nil)

	proc.sendEvent(StartEvent, nil)
	return
}
func (proc IdProc) GetError() error {
	return proc.err
}

func (proc IdProc) Id() []byte {
	return proc.id
}

func (proc *IdProc) Handle(msg *libnas.GmmMessage) (err error) {
	switch msg.GetMessageType() {
	case libnas.MsgTypeIdentityResponse:
		err = proc.sendEvent(IdResponseEvent, msg.IdentityResponse)
	default:
		err = fmt.Errorf("GmmMessage is not handled:%d", msg.GetMessageType())
	}
	return
}

func (proc *IdProc) sendEvent(ev fsm.EventType, args interface{}) error {
	return _sm.SendEvent(proc.ranue.Worker(), proc, ev, args)
}

// Note: do not handle error, just use T3570 timer
func (proc *IdProc) request() {
	proc.t3570.Start()
	proc.ranue.SendIdentityRequest(proc.idtype)
}

func (proc *IdProc) handleResponse(msg *nasMessage.IdentityResponse) {
	proc.t3570.Stop()
	proc.Info("Handle a NAS Identity Response")
	content := msg.MobileIdentity.GetMobileIdentityContents()
	idtype := nasConvert.GetTypeOfIdentity(content[0])
	if idtype != proc.idtype {
		err := fmt.Errorf("Mistmached identity in Identity response")
		proc.sendEvent(DoneEvent, err)
		return
	}

	proc.id = content
	proc.sendEvent(DoneEvent, nil)
	return
}

func (proc *IdProc) handleT3570() {
	proc.Trace("T3570 expired")
	proc.failurecount++
	if proc.failurecount >= MAX_FAILURE_COUNT {
		err := fmt.Errorf("Too many failed attemps to send IdentityRequest")
		proc.sendEvent(DoneEvent, err)
	} else {
		proc.request()
	}
}
