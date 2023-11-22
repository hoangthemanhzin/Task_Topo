package context

import (
	"etrib5gc/nfs/damf/ue"
	"sync"
)

type UeList struct {
	uelist     map[int64]*ue.UeContext
	ranueid2ue map[int64]*ue.UeContext
	mutex      sync.RWMutex
}

func newUeList() UeList {
	return UeList{
		uelist:     make(map[int64]*ue.UeContext),
		ranueid2ue: make(map[int64]*ue.UeContext),
	}
}

func (l *UeList) size() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.uelist)
}
func (l *UeList) add(uectx *ue.UeContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	uectx.WithFields(_logfields).Info("UeContext is added")
	l.uelist[uectx.AmfUeId()] = uectx
	l.ranueid2ue[uectx.RanUeId()] = uectx
}

func (l *UeList) has(ranueid int64) (ok bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	_, ok = l.ranueid2ue[ranueid]
	return
}
func (l *UeList) remove(id int64) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if uectx, ok := l.uelist[id]; ok {
		uectx.WithFields(_logfields).Info("UeContext is removed")
		delete(l.uelist, id)
		delete(l.ranueid2ue, uectx.RanUeId())
	}
}

func (l *UeList) find(id int64) (uectx *ue.UeContext) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	uectx, _ = l.uelist[id]
	return
}

func (l *UeList) clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	for _, uectx := range l.uelist {
		uectx.Close()
	}
}
