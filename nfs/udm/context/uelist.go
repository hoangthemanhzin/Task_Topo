package context

import "sync"

type UeList struct {
	mutex sync.RWMutex
	list  map[string]*UeContext
}

func newUeList() UeList {
	return UeList{
		list: make(map[string]*UeContext),
	}
}
func (l *UeList) find(supi string) *UeContext {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if ue, ok := l.list[supi]; ok {
		return ue
	}
	return nil
}

func (l *UeList) add(ue *UeContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.list[ue.Supi()] = ue
	return
}

func (l *UeList) del(supi string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.list, supi)
}
