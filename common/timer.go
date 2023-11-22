package common

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "common"})
}

type UeTimer interface {
	Stop()
	Start() //Start or Reset timer
}

type ueTimer struct {
	d     time.Duration
	e     Executer
	abort chan bool
	_t    *time.Timer
	fn    func() //callback
	wg    sync.WaitGroup
}

// create an idling timer
func NewTimer(d time.Duration, fn func(), e Executer) UeTimer {
	return &ueTimer{
		d:  d,
		fn: fn,
		e:  e,
	}
}

func (t *ueTimer) Start() {
	//log.Info("Start a timer")
	t.Stop()

	t.abort = make(chan bool)
	t._t = time.NewTimer(t.d) //always create a new one

	go func() {
		t.wg.Add(1)
		defer t.wg.Done()
		select {
		case <-t.abort:
			//just exit
		case <-t._t.C:
			if t.e != nil {
				//run callback in Executer
				t.e.AddJob(t.fn)
			} else {
				t.fn() //or run directly
			}
		}
		//t.abort = nil
		//t._t = nil
	}()
}

func (t *ueTimer) Stop() {
	if t._t != nil && t.abort != nil {
		//log.Info("Stop a timer")
		close(t.abort)
		t._t.Stop()
		t.wg.Wait() //make sure the running goroutine is terminated
		//note: t._t and t.abort will be reset with the termination
		t.abort = nil
		t._t = nil
	}
}
