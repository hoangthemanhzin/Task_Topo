package context

import (
	"etrib5gc/nfs/amf/ranuecontext"
	"etrib5gc/nfs/amf/uecontext"
	"sync"
)

// UePool has all UeContexts which are indexed with suci, supi, tmsi5gs and pei
type UePool struct {
	mux            sync.Mutex
	tmsi5gs2ue     map[string]*uecontext.UeContext
	suci2ue        map[string]*uecontext.UeContext
	supi2ue        map[string]*uecontext.UeContext
	pei2ue         map[string]*uecontext.UeContext
	callback2ranue map[string]*ranuecontext.RanUe
}

func newUePool() (p UePool) {
	p = UePool{}
	p.reset()
	return
}
func (p *UePool) reset() {
	p.tmsi5gs2ue = make(map[string]*uecontext.UeContext)
	p.suci2ue = make(map[string]*uecontext.UeContext)
	p.supi2ue = make(map[string]*uecontext.UeContext)
	p.pei2ue = make(map[string]*uecontext.UeContext)
	p.callback2ranue = make(map[string]*ranuecontext.RanUe)
}

func (p *UePool) findBySuci(suci string) (ue *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	ue, _ = p.suci2ue[suci]
	return
}
func (p *UePool) findByTmsi5gs(tmsi5gs string) (ue *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	ue, _ = p.tmsi5gs2ue[tmsi5gs]
	return
}

func (p *UePool) findBySupi(supi string) (ue *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	ue, _ = p.supi2ue[supi]
	return
}

func (p *UePool) findByPei(pei string) (ue *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	ue, _ = p.pei2ue[pei]
	return
}

func (p *UePool) remove(uectx *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	delete(p.tmsi5gs2ue, uectx.Tmsi5gs())
	delete(p.suci2ue, uectx.Suci())
	delete(p.supi2ue, uectx.Supi())
	delete(p.pei2ue, uectx.Pei())
	uectx.LogWriter.WithFields(_logfields).Info("UeContext is removed from pool")
}

func (p *UePool) add(uectx *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	if suci := uectx.Suci(); len(suci) > 0 {
		p.suci2ue[suci] = uectx
	}
	if supi := uectx.Supi(); len(supi) > 0 {
		p.supi2ue[supi] = uectx
	}

	if pei := uectx.Pei(); len(pei) > 0 {
		p.pei2ue[pei] = uectx
	}

	if tmsi5gs := uectx.Tmsi5gs(); len(tmsi5gs) > 0 {
		p.tmsi5gs2ue[tmsi5gs] = uectx
	}
	uectx.LogWriter.WithFields(_logfields).Info("UeContext is added to pool")
}

func (p *UePool) update(id string, idtype uint8, uectx *uecontext.UeContext) {
	p.mux.Lock()
	defer p.mux.Unlock()
	switch idtype {
	case uecontext.UE_ID_TYPE_SUCI:
		p.suci2ue[id] = uectx
	case uecontext.UE_ID_TYPE_SUPI:
		p.supi2ue[id] = uectx
	case uecontext.UE_ID_TYPE_PEI:
		p.pei2ue[id] = uectx
	case uecontext.UE_ID_TYPE_TMSI5GS:
		p.tmsi5gs2ue[id] = uectx
	case uecontext.UE_ID_TYPE_GUTI:
		//TODO: get TMSI5GS from GUTI
	default:
		//do nothing
	}
}

func (p *UePool) clean() {
	p.mux.Lock()
	defer p.mux.Unlock()
	for _, ue := range p.suci2ue {
		ue.Clean()
	}
	p.reset()
}
