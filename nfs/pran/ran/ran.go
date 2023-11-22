package ran

import (
	"etrib5gc/nfs/pran/ue"
	"etrib5gc/sbi/models"
	"fmt"
	"net"

	"github.com/free5gc/ngap/ngapType"
)

type Ran struct {
	ctx     Context
	access  models.AccessType
	conn    net.Conn
	id      string
	name    string
	rannets []string
	drx     uint64
	tais    []SupportedTai
}

type RanUeId struct {
	Conn      net.Conn
	RanNgapId int64
}

func newRan(conn net.Conn) *Ran {
	ret := &Ran{
		conn:   conn,
		access: models.ACCESSTYPE__3_GPP_ACCESS, //TODO: should be updated from RAN
	}
	return ret
}
func (r *Ran) Access() models.AccessType {
	return r.access
}

func (r *Ran) RanNets() []string {
	return r.rannets
}
func (r *Ran) Conn() net.Conn {
	return r.conn
}

func (r *Ran) Id() string {
	return r.id
}

func (r *Ran) Send(buf []byte) (err error) {
	if len(buf) == 0 {
		err = fmt.Errorf("Empty packet")
		return
	}

	log.Debugf("Send message To Ran")
	var n int
	if n, err = r.conn.Write(buf); err != nil {
		log.Errorf("Send error: %+v", err)
	} else {
		log.Debugf("Write %d bytes", n)
	}
	return
}

// remove all UEs
func (r *Ran) RemoveUes() {
	for _, uectx := range r.ctx.GetRanUeList(r.conn) {
		uectx.Close()
	}
}

func (r *Ran) FindUe(ranngapid *ngapType.RANUENGAPID, cungapid *ngapType.AMFUENGAPID) (uectx *ue.UeContext) {
	if ranngapid != nil {
		if uectx = r.ctx.FindByRanNgapId(r.conn, ranngapid.Value); uectx != nil {
			return
		}
		log.Warnf("Ue for RanNgapId=%d not found", ranngapid.Value)
	}
	if cungapid != nil {
		if uectx = r.ctx.FindByCuNgapId(cungapid.Value); uectx == nil {
			log.Warnf("Ue for CuNgapId=%d not found", cungapid.Value)
		}
	}
	return
}

func (r *Ran) RemoveUe(ranngapid *ngapType.RANUENGAPID, cungapid *ngapType.AMFUENGAPID) {
	if ranngapid != nil {
		if uectx := r.ctx.FindByRanNgapId(r.conn, ranngapid.Value); uectx != nil {
			uectx.Close()
		}
	} else if cungapid != nil {
		if uectx := r.ctx.FindByCuNgapId(cungapid.Value); uectx != nil {
			uectx.Close()
		}
	}
}

func (r *Ran) Setup(id *ngapType.GlobalRANNodeID, name *ngapType.RANNodeName, drx *ngapType.PagingDRX, tailist *ngapType.SupportedTAList) {
	r.id, r.access = ranid2string(id)
	r.name = name.Value
	r.drx = uint64(drx.Value)
	r.tais = convertSupportedTai(tailist)
}
