package context

import (
	"etrib5gc/nfs/smf/sm"
	"sync"
)

type SmList struct {
	mutex sync.Mutex
	list  map[string]*sm.SmContext
}

func newSmList() SmList {
	return SmList{
		list: make(map[string]*sm.SmContext),
	}
}

func (l *SmList) size() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return len(l.list)
}
func (l *SmList) add(smcontext *sm.SmContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.list[smcontext.Ref()] = smcontext
}

func (l *SmList) find(ref string) (smcontext *sm.SmContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	smcontext, _ = l.list[ref]
	return
}

func (l *SmList) remove(smcontext *sm.SmContext) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.list, smcontext.Ref())
}

func (l *SmList) clear() {
	l.mutex.Lock()
	smlist := []*sm.SmContext{}

	for _, smctx := range l.list {
		smlist = append(smlist, smctx)
	}
	l.mutex.Unlock()

	for _, smctx := range smlist {
		smctx.Close()
	}
}
