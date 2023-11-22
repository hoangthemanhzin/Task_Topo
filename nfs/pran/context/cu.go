package context

import (
	"etrib5gc/common"
	"etrib5gc/mesh"
	"etrib5gc/nfs/pran/config"
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"
	"etrib5gc/util/idgen"
	"math"
	"net"
	"sync"
	"sync/atomic"
)

const (
	DUMB_AMF_ID   = "100023"
	DUMB_AMF_NAME = "dumbamf"
)

type CuContext struct {
	id       string
	rannets  []string
	amfname  string
	amfid    string
	plmnid   models.PlmnId
	plmnlist map[models.PlmnId][]models.Snssai
	//ngapiplist []net.IP
	relcap int64
	//	ranpool    *RanPool
	uelist    *UeList
	ngapidgen idgen.IdGenerator
	//	nssf      map[models.Snssai]string

	closech chan *ue.UeContext //to receive UeContext for removing
	done    chan bool
	wg      sync.WaitGroup
	closed  int32
}

func NewCuContext(c *config.PRanConfig) (ctx *CuContext, err error) {
	_initLog()
	ctx = &CuContext{
		id:        c.Id,
		rannets:   c.RanNets,
		amfname:   DUMB_AMF_NAME,
		amfid:     DUMB_AMF_ID,
		plmnid:    models.PlmnId(c.PlmnId),
		uelist:    newUeList(),
		plmnlist:  make(map[models.PlmnId][]models.Snssai),
		ngapidgen: idgen.NewIdGenerator(1, math.MaxUint64),
		//		nssf:      make(map[models.Snssai]string),
		closech: make(chan *ue.UeContext),
		done:    make(chan bool),
		closed:  0,
	}

	for _, item := range c.PlmnList {
		//log.Infof("Mnc=%s,Mcc=%s", item.PlmnId.Mnc, item.PlmnId.Mcc)
		ctx.plmnlist[models.PlmnId(item.PlmnId)] = item.Slices
	}

	if err == nil {
		go ctx.loop()
	}

	return
}

func (ctx *CuContext) loop() {
	ctx.wg.Add(1)
	defer ctx.wg.Done()
LOOP:
	for {
		select {
		case <-ctx.done:
			log.Trace("Receive a signal to close Context's loop")
			break LOOP
		case uectx := <-ctx.closech:
			ctx.uelist.remove(uectx)
			uectx.Kill()
		}
	}
	//clean up any remaining UeContexts
	for ctx.uelist.size() > 0 {
		//log.Info("Wait to kill")
		uectx := <-ctx.closech
		ctx.uelist.remove(uectx)
		uectx.Kill()
	}
}

func (ctx *CuContext) RanNets() []string {
	return ctx.rannets
}
func (ctx *CuContext) GetCuNgapId() int64 {
	return int64(ctx.ngapidgen.Allocate())
}

func (ctx *CuContext) IsClosed() bool {
	return atomic.LoadInt32(&ctx.closed) == 1
}
func (ctx *CuContext) AmfId() string {
	return ctx.amfid
}

func (ctx *CuContext) DamfName() string {
	return common.DamfServiceName(&ctx.plmnid, ctx.id)
}

func (ctx *CuContext) Callback() models.Callback {
	return models.Callback(mesh.CallbackAddress())
}

func (ctx *CuContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}

func (ctx *CuContext) Id() string {
	return ctx.id
}

func (ctx *CuContext) Name() string {
	return ctx.amfname
}

func (ctx *CuContext) RelativeCapacity() int64 {
	return ctx.relcap
}

func (ctx *CuContext) AddUe(uectx *ue.UeContext) {
	ctx.uelist.add(uectx)
}

func (ctx *CuContext) RemoveUe(uectx *ue.UeContext) {
	ctx.closech <- uectx
	ctx.ngapidgen.Free(uint64(uectx.CuNgapId()))
}

func (ctx *CuContext) FindByRanNgapId(ran net.Conn, ranNgapId int64) *ue.UeContext {
	return ctx.uelist.findByConn(ran, ranNgapId)
}
func (ctx *CuContext) FindByCuNgapId(cuNgapId int64) *ue.UeContext {
	return ctx.uelist.findByCuNgapId(cuNgapId)
}

// called to remove UeContexts when a Ran is disconnected
func (ctx *CuContext) GetRanUeList(conn net.Conn) (uelist []*ue.UeContext) {
	return ctx.uelist.getRanUeList(conn)
}

func (ctx *CuContext) PlmnList() map[models.PlmnId][]models.Snssai {
	return ctx.plmnlist
}

func (ctx *CuContext) Clean() {
	log.Infof("Clear UeContexts")
	atomic.StoreInt32(&ctx.closed, 1)
	ctx.uelist.clear()
	close(ctx.done)
	ctx.wg.Wait()
}

/*
	func (r *CuContext) ServedGuamiList() []models.Guami {
		return r.config.GuamiList
	}
*/
