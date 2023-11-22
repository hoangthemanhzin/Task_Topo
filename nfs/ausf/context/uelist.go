package context

import "sync"

//mapping supiORsupi to supi
type IdList struct {
	id2supi map[string]string
	mutex   sync.RWMutex
}

func newIdList() (l IdList) {
	l.id2supi = make(map[string]string)
	return
}

func (l *IdList) add(id string, supi string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.id2supi[id] = supi
}

func (l *IdList) has(id string) (ok bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	_, ok = l.id2supi[id]
	return
}

func (l *IdList) get(id string) (supi string) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	supi, _ = l.id2supi[id]
	return
}
func (l *IdList) del(id string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.id2supi, id)
}

//Ue pool
type UeList struct {
	supi2ue map[string]*UeContext
	mutex   sync.RWMutex
}

func newUeList() (l UeList) {
	l.supi2ue = make(map[string]*UeContext)
	return
}
func (l *UeList) add(ue *UeContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.supi2ue[ue.Supi()] = ue
}

func (l *UeList) del(ue *UeContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.supi2ue, ue.Supi())
}
func (l *UeList) get(supi string) (ue *UeContext) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	ue, _ = l.supi2ue[supi]
	return
}
