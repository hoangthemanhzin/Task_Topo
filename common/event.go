package common

import (
	"fmt"
	"time"
)

type EventData struct {
	EvType  uint8
	Content interface{}
}

type AsyncJob struct {
	info interface{} //job info

	errch chan error
	t     *time.Timer

	callback func() //called before quit job
}

func NewAsyncJob(info interface{}, timeout int) *AsyncJob {
	job := &AsyncJob{
		info:  info,
		errch: make(chan error, 1),
	}
	if timeout > 0 {
		job.t = time.NewTimer(time.Duration(timeout) * time.Millisecond)
	}
	return job
}

func (job *AsyncJob) SetCallback(fn func()) {
	job.callback = fn
}

func (job *AsyncJob) Info() interface{} {
	return job.info
}
func (job *AsyncJob) Wait() (err error) {
	defer func() {
		if job.callback != nil {
			job.callback()
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
		log.Errorf("an sbi job is timeouted")
		err = fmt.Errorf("Timeout")
	}
	return
}

func (job *AsyncJob) Done(err error) {
	job.errch <- err
}

/*
type AsyncJobList struct {
	jobs  map[uint8]*AsyncJob
	mutex sync.Mutex
}

func NewAsyncJobList() AsyncJobList {
	return AsyncJobList{
		jobs: make(map[uint8]*AsyncJob),
	}
}

func (l *AsyncJobList) Add(job *AsyncJob) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.jobs[job.info.JobType()] = job
}

func (l *AsyncJobList) Remove(t uint8) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.jobs, t)
}

func (l *AsyncJobList) Get(t uint8) (job *AsyncJob) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	job, _ = l.jobs[t]
	return
}
*/
