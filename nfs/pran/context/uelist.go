package context

import (
	"etrib5gc/nfs/pran/ue"
	"net"
	"sync"
)

type UeMap map[int64]*ue.UeContext

func newUeMap() UeMap {
	return make(map[int64]*ue.UeContext)
}

type UeList struct {
	cuNgapId2ue UeMap
	conn2ue     map[net.Conn]UeMap
	mutex       sync.RWMutex
}

func newUeList() *UeList {
	l := &UeList{}
	l.init()
	return l
}

func (pool *UeList) init() {
	pool.cuNgapId2ue = newUeMap()
	pool.conn2ue = make(map[net.Conn]UeMap)

}

func (pool *UeList) findByCuNgapId(cuNgapId int64) (uectx *ue.UeContext) {
	pool.mutex.RLock()
	defer pool.mutex.RUnlock()
	uectx, _ = pool.cuNgapId2ue[cuNgapId]
	return
}

func (pool *UeList) findByConn(conn net.Conn, ranNgapId int64) (uectx *ue.UeContext) {
	pool.mutex.RLock()
	defer pool.mutex.RUnlock()
	if r, ok := pool.conn2ue[conn]; ok {
		uectx, _ = r[ranNgapId]
	}
	return
}
func (pool *UeList) getRanUeList(conn net.Conn) (list []*ue.UeContext) {
	pool.mutex.RLock()
	defer pool.mutex.RUnlock()
	if r, ok := pool.conn2ue[conn]; ok {
		for _, uectx := range r {
			list = append(list, uectx)
		}
	}
	return
}

// add a new UeContext to the pool. The UeContext must have
// (RanNgapId[RAN], CuNgapId[PRAN])
func (pool *UeList) add(uectx *ue.UeContext) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	log.Infof("Add UeContext [CuNgapId=%d]", uectx.CuNgapId())
	pool.cuNgapId2ue[uectx.CuNgapId()] = uectx
	conn := uectx.RanConn()
	if r, ok := pool.conn2ue[conn]; ok {
		r[uectx.RanNgapId()] = uectx
	} else {
		pool.conn2ue[conn] = UeMap{
			uectx.RanNgapId(): uectx,
		}
	}
}

func (pool *UeList) size() int {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	return len(pool.cuNgapId2ue)
}

func (pool *UeList) remove(uectx *ue.UeContext) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	delete(pool.cuNgapId2ue, uectx.CuNgapId())
	//	delete(pool.amfUeId2ue, uectx.AmfUeId())
	conn := uectx.RanConn()
	if r, ok := pool.conn2ue[conn]; ok {
		delete(r, uectx.RanNgapId())
	}

	log.Infof("Remove UeContext CuNgapId=%d", uectx.CuNgapId())
}

func (pool *UeList) all() (list []*ue.UeContext) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	for _, uectx := range pool.cuNgapId2ue {
		list = append(list, uectx)
	}
	return
}

func (pool *UeList) clear() {
	pool.mutex.Lock()
	uelist := []*ue.UeContext{}
	for _, uectx := range pool.cuNgapId2ue {
		uelist = append(uelist, uectx)
	}
	pool.mutex.Unlock()

	for _, uectx := range uelist {
		uectx.Close()
	}
}
