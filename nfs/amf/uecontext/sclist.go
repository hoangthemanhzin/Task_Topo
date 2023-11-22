package uecontext

import (
	"etrib5gc/sbi/models"
	"sync"
)

type SessionContext interface {
	Access() models.AccessType
	Id() int32
	SmfId() string
	Snssai() models.Snssai
}

type SessionContextList struct {
	id2ctx map[int32]SessionContext
	mutex  sync.RWMutex
}

func newSessionContextList() (l SessionContextList) {
	l.id2ctx = make(map[int32]SessionContext)
	return
}

func (l *SessionContextList) add(sm SessionContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	//TODO:
	//check
	//existence
	l.id2ctx[sm.Id()] = sm
}

func (l *SessionContextList) find(id int32) (sm SessionContext) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	sm, _ = l.id2ctx[id]
	return
}

func (l *SessionContextList) list() (out []SessionContext) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	out = make([]SessionContext, len(l.id2ctx))
	i := 0
	for _, sm := range l.id2ctx {
		out[i] = sm
		i++
	}
	return
}
