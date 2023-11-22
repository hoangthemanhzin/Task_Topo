package ran

import (
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"
	"etrib5gc/util/idgen"
	"math"
	"net"
	"sync"

	"github.com/free5gc/ngap/ngapType"
)

type Context interface {
	FindByCuNgapId(int64) *ue.UeContext
	FindByRanNgapId(net.Conn, int64) *ue.UeContext
	GetRanUeList(net.Conn) []*ue.UeContext
	RanNets() []string
}

type RanPool struct {
	ctx      Context
	conn2ran map[net.Conn]*Ran
	id2ran   map[string]*Ran
	ranidgen idgen.IdGenerator
	mutex    sync.Mutex
}

func NewRanPool(ctx Context) *RanPool {
	_initLog()
	return &RanPool{
		ctx:      ctx,
		ranidgen: idgen.NewIdGenerator(1, math.MaxInt16),
		conn2ran: make(map[net.Conn]*Ran),
		id2ran:   make(map[string]*Ran),
	}
}

func (pool *RanPool) ByConn(conn net.Conn) *Ran {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	if r, ok := pool.conn2ran[conn]; ok {
		return r
	}
	return nil
}

func (pool *RanPool) ByRanId(ranid *ngapType.GlobalRANNodeID) *Ran {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	id, _ := ranid2string(ranid)
	if r, ok := pool.id2ran[id]; ok {
		return r
	}
	return nil
}

// add or update ran
func (pool *RanPool) Add(ran *Ran) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	pool.conn2ran[ran.Conn()] = ran
	pool.id2ran[ran.Id()] = ran
}

// remove a Ran (not killing its UeContexts)
func (pool *RanPool) Remove(ran *Ran) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	log.Infof("Remove RAN: %s", ran.name)
	delete(pool.conn2ran, ran.Conn())
	id := ran.Id()
	if len(id) > 0 {
		delete(pool.id2ran, id)
	}
	ran.RemoveUes()
}

// return list of currently connected RAN
func (pool *RanPool) All() (l []*Ran) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	for _, r := range pool.conn2ran {
		l = append(l, r)
	}
	return
}
func (pool *RanPool) NewRan(conn net.Conn) *Ran {
	ran := newRan(conn)
	ran.rannets = pool.ctx.RanNets()
	ran.ctx = pool.ctx
	return ran
}

func ranid2string(id *ngapType.GlobalRANNodeID) (idstr string, access models.AccessType) {
	access = models.ACCESSTYPE__3_GPP_ACCESS
	switch id.Present {
	case ngapType.GlobalRANNodeIDPresentGlobalGNBID:
		idstr = string(id.GlobalGNBID.PLMNIdentity.Value) + string(id.GlobalGNBID.GNBID.GNBID.Bytes)
	case ngapType.GlobalRANNodeIDPresentGlobalNgENBID:
		idstr = string(id.GlobalNgENBID.PLMNIdentity.Value) + string(id.GlobalNgENBID.NgENBID.MacroNgENBID.Bytes)
	case ngapType.GlobalRANNodeIDPresentGlobalN3IWFID:
		idstr = string(id.GlobalN3IWFID.PLMNIdentity.Value) + string(id.GlobalN3IWFID.N3IWFID.N3IWFID.Bytes)
		access = models.ACCESSTYPE_NON_3_GPP_ACCESS
	default:
		idstr = ""
	}
	return
}
