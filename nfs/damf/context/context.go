package context

import (
	"etrib5gc/nfs/damf/config"
	"etrib5gc/nfs/damf/ue"
	"etrib5gc/sbi/models"
	"etrib5gc/util/idgen"
	"math"
	"sync"
	"sync/atomic"
)

const (
	T3502_DURATION = 100 //miliseconds
)

type DamfContext struct {
	id      string
	plmnid  models.PlmnId
	config  *config.DamfConfig
	uelist  UeList
	ueidgen idgen.IdGenerator
	nssf    map[models.Snssai]string

	closech chan *ue.UeContext //to receive UeContext for removing
	done    chan bool
	wg      sync.WaitGroup
	closed  int32
}

func NewDamfContext(cfg *config.DamfConfig) (ctx *DamfContext) {
	_initLog()
	ctx = &DamfContext{
		config:  cfg,
		id:      cfg.Id,
		plmnid:  models.PlmnId(cfg.PlmnId),
		uelist:  newUeList(),
		ueidgen: idgen.NewIdGenerator(0, math.MaxInt64),
		nssf:    make(map[models.Snssai]string),
		closech: make(chan *ue.UeContext),
		done:    make(chan bool),
		closed:  0,
	}
	for _, item := range cfg.AmfMap {
		ctx.nssf[item.Snssai] = item.AmfId
		//log.Infof("Slice[%s] is mapped to AmfId[%s]", item.Snssai.String(), item.AmfId)
	}
	go ctx.loop()
	return
}
func (ctx *DamfContext) IsClosed() bool {
	return atomic.LoadInt32(&ctx.closed) == 1
}

func (ctx *DamfContext) Id() string {
	return ctx.id
}
func (ctx *DamfContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}

func (ctx *DamfContext) GetT3502() uint8 {
	return T3502_DURATION
}
func (ctx *DamfContext) RemoveUe(uectx *ue.UeContext) {
	ctx.closech <- uectx
	ctx.ueidgen.Free(uint64(uectx.AmfUeId()))
}

func (ctx *DamfContext) loop() {
	ctx.wg.Add(1)
	defer ctx.wg.Done()
LOOP:
	for {
		select {
		case <-ctx.done:
			log.Tracef("Receive a signal to close Context's loop")
			break LOOP
		case uectx := <-ctx.closech:
			ctx.uelist.remove(uectx.AmfUeId())
			uectx.Kill()
		}
	}
	//clean up any remaining UeContexts
	for ctx.uelist.size() > 0 {
		log.Trace("Clean remaining UeContext")
		uectx := <-ctx.closech
		ctx.uelist.remove(uectx.AmfUeId())
		uectx.Kill()
	}
}

func (ctx *DamfContext) Clean() {
	log.Tracef("Cleanup UeContexts")
	atomic.StoreInt32(&ctx.closed, 1)
	ctx.uelist.clear()
	close(ctx.done)
	ctx.wg.Wait()
}

func (ctx *DamfContext) HasUe(ranueid int64) bool {
	return ctx.uelist.has(ranueid)
}
func (ctx *DamfContext) FindUe(id int64) *ue.UeContext {
	return ctx.uelist.find(id)
}
func (ctx *DamfContext) AddUe(uectx *ue.UeContext) {
	ctx.uelist.add(uectx)
}

func (ctx *DamfContext) GetUeId() int64 {
	return int64(ctx.ueidgen.Allocate())
}
