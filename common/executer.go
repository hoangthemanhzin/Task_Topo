package common

import (
	"sync"
)

//a generic function to be executed
type ExecuteFn func()

//hold a stack of functions to be executed
type fnStack struct {
	list   []ExecuteFn
	maxlen int
	cur    int
	mux    sync.Mutex
}

func newFnStack(maxlen int) fnStack {
	return fnStack{
		maxlen: maxlen,
		list:   make([]ExecuteFn, maxlen, maxlen),
		cur:    -1,
	}
}

// false if the queue is full
func (s *fnStack) Add(fn ExecuteFn) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.cur >= s.maxlen-1 { //the stack is full
		return false
	}
	s.list[s.cur+1] = fn
	s.cur++
	return true
}

func (s *fnStack) Pop() (fn ExecuteFn) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.cur >= 0 {
		fn = s.list[s.cur]
		s.cur--
	}
	return
}

//go routine worker to execute functions
type Executer interface {
	Terminate()
	AddJob(ExecuteFn) bool
}
type executer struct {
	jobs fnStack
	quit chan struct{}
}

//create then run
//maxjobsize is the max size of job queue
func NewExecuter(maxjobsize int) Executer {
	e := &executer{
		jobs: newFnStack(maxjobsize),
		quit: make(chan struct{}),
	}
	go e.loop()
	return e
}

//terminate once, don't be stupid
func (e *executer) Terminate() {
	e.quit <- struct{}{}
}

//add job to the executer
//return false of the job queue is full
func (e *executer) AddJob(fn ExecuteFn) bool {
	return e.jobs.Add(fn)
}

//inner loop, exits when Terminate method is called
func (e *executer) loop() {

LOOP:
	for {
		select {
		case <-e.quit:
			break LOOP
		default:
			if fn := e.jobs.Pop(); fn != nil {
				fn()
			}
		}
	}
}
