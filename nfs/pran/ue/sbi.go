package ue

import (
	"etrib5gc/common"
	"etrib5gc/sbi/models/n2models"
	"fmt"
	"sync"
	"time"
)

type SbiJob struct {
	info interface{} //job info

	errch chan error
	t     *time.Timer

	onDone func() //called before quit job
}

func NewSbiJob(info interface{}, timeout int) *SbiJob {
	job := &SbiJob{
		info:  info,
		errch: make(chan error, 1),
	}
	if timeout > 0 {
		job.t = time.NewTimer(time.Duration(timeout) * time.Millisecond)
	}
	return job
}
func (job *SbiJob) Wait() (err error) {
	defer func() {
		if job.onDone != nil {
			job.onDone()
		}
	}()
	//blocking (no timeout)
	if job.t == nil {
		err = <-job.errch
		return
	}

	//with timeout
	select {
	case err = <-job.errch:
		job.t.Stop()
	case <-job.t.C:
		//log.Errorf("an sbi job is timeouted")
		err = fmt.Errorf("Timeout")
	}
	return
}

func (job *SbiJob) done(err error) {
	job.errch <- err
}

type PendingJobs struct {
	jobs  map[uint8]*SbiJob
	mutex sync.Mutex
}

func newPendingJobs() PendingJobs {
	return PendingJobs{
		jobs: make(map[uint8]*SbiJob),
	}
}

func (l *PendingJobs) add(t uint8, job *SbiJob) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.jobs[t] = job
}

func (l *PendingJobs) remove(t uint8) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.jobs, t)
}

func (l *PendingJobs) get(t uint8) (job *SbiJob) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	job, _ = l.jobs[t]
	return
}

type InitUeContextStatusJob struct {
	Msg *n2models.InitUeContextStatus
}

type NasDlJob struct {
	Msg *n2models.NasDlMsg
}

type InitCtxSetupReqJob struct {
	Msg  *n2models.InitCtxSetupReq
	Rsp  *n2models.InitCtxSetupRsp
	Ersp *n2models.InitCtxSetupFailure
}
type UeCtxModReqJob struct {
	Msg  *n2models.UeCtxModReq
	Rsp  *n2models.UeCtxModRsp
	Ersp *n2models.UeCtxModFail
}

type PduSessResSetReqJob struct {
	Msg *n2models.PduSessResSetReq
	Rsp *n2models.PduSessResSetRsp
}

type PduSessResModReqJob struct {
	Msg *n2models.PduSessResModReq
	Rsp *n2models.PduSessResModRsp
}

type PduSessResRelCmdJob struct {
	Msg *n2models.PduSessResRelCmd
	Rsp *n2models.PduSessResRelRsp
}

func (uectx *UeContext) HandleSbi(ev *common.EventData) (err error) {
	uectx.sendEvent(SbiEvent, ev)
	return
}

func (uectx *UeContext) addJob(t uint8, job *SbiJob) {
	//remember to remove pending job when it is done
	job.onDone = func() {
		uectx.pendingjobs.remove(t)
	}
	//add job to the pending list
	uectx.pendingjobs.add(t, job)

}
