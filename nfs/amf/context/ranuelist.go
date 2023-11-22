package context

import (
	"etrib5gc/nfs/amf/ranuecontext"
	"sync"
)

// RanUePool has all UeContexts which are indexed with suci, supi, tmsi5gs and pei
type RanUePool struct {
	mux           sync.Mutex
	idatran2ranue map[int64]*ranuecontext.RanUe
	id2ranue      map[int64]*ranuecontext.RanUe
}

func newRanUePool() (p RanUePool) {
	p = RanUePool{}
	p.reset()
	return
}
func (p *RanUePool) reset() {
	p.idatran2ranue = make(map[int64]*ranuecontext.RanUe)
	p.id2ranue = make(map[int64]*ranuecontext.RanUe)
}

func (p *RanUePool) findById(id int64) (ranue *ranuecontext.RanUe) {
	p.mux.Lock()
	defer p.mux.Unlock()
	ranue, _ = p.id2ranue[id]
	return
}
func (p *RanUePool) findByIdAtRan(idatran int64) (ranue *ranuecontext.RanUe) {
	p.mux.Lock()
	defer p.mux.Unlock()
	ranue, _ = p.idatran2ranue[idatran]
	return
}

func (p *RanUePool) add(ranue *ranuecontext.RanUe) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.idatran2ranue[ranue.RanUeId()] = ranue
	p.id2ranue[ranue.AmfUeId()] = ranue
}

func (p *RanUePool) remove(ranue *ranuecontext.RanUe) {
	p.mux.Lock()
	defer p.mux.Unlock()
	delete(p.idatran2ranue, ranue.RanUeId())
	delete(p.id2ranue, ranue.AmfUeId())
}

func (p *RanUePool) clean() {
	p.mux.Lock()
	defer p.mux.Unlock()
	/*
		for _, ranue := range p.id2ranue {
			//TODO
			//ranue.Clean()
		}
	*/
	p.reset()
}
