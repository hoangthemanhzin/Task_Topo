package up

import (
	"etrib5gc/logctx"
	"etrib5gc/mesh/httpdp"
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/pfcptypes"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/util/idgen"
	"fmt"
	"math"
	"net"
)

var log logctx.LogWriter
var _logfields logctx.Fields = logctx.Fields{
	"mod": "up",
}

func _initLog() {
	if log == nil {
		log = logctx.WithFields(_logfields)
	}
}

type Upf struct {
	pfcp pfcp.PfcpSender //for sending Pfcp Requests
	ip   net.IP
	port int
	cli  sbi.ConsumerClient

	sessions map[uint64]*PfcpSession

	/*
		pdrPool sync.Map
		farPool sync.Map
		barPool sync.Map
		qerPool sync.Map
		// urrPool
		// sync.Map
	*/
	pdridgen  idgen.IdGenerator
	faridgen  idgen.IdGenerator
	teididgen idgen.IdGenerator
}

func NewUpf(pfcp pfcp.PfcpSender, ip net.IP, port int) (upf *Upf) {
	_initLog()
	upf = &Upf{
		pfcp:      pfcp,
		ip:        ip,
		port:      port,
		teididgen: idgen.NewIdGenerator(1, math.MaxUint32),
		pdridgen:  idgen.NewIdGenerator(1, math.MaxUint16),
		faridgen:  idgen.NewIdGenerator(1, math.MaxUint16),
		sessions:  make(map[uint64]*PfcpSession),
	}
	upfaddr := fmt.Sprintf("%s:%d", upf.ip.String(), upf.port)
	fmt.Printf("Create Upf Client: %s", upfaddr)

	upf.cli = httpdp.NewClientWithAddr(upfaddr)
	return
}

func (upf *Upf) Id() string {
	return upf.ip.String()
}

func (upf *Upf) GenerateTeid() uint32 {
	return uint32(upf.teididgen.Allocate())
}

func (upf *Upf) FreeTeid(id uint32) {
	upf.teididgen.Free(uint64(id))
}

func (upf *Upf) CreateSession(localseid uint64) (session *PfcpSession) {
	session = newPfcpSession(localseid, upf)
	upf.sessions[localseid] = session
	return
}
func (upf *Upf) FindSession(localseid uint64) (session *PfcpSession) {
	session, _ = upf.sessions[localseid]
	return
}

func (upf *Upf) createPdr() (pdr *PDR) {
	pdr = &PDR{
		PDRID: uint16(upf.pdridgen.Allocate()),
		FAR:   upf.createFar(),
	}
	return
}

func (upf *Upf) createFar() (far *FAR) {
	far = &FAR{
		FARID: uint32(upf.faridgen.Allocate()),
	}
	return
}

func (upf *Upf) removePdr(pdr *PDR) {
	upf.pdridgen.Free(uint64(pdr.PDRID))
	if pdr.FAR != nil {
		upf.removeFar(pdr.FAR)
	}
}

func (upf *Upf) removeFar(far *FAR) {
	upf.faridgen.Free(uint64(far.FARID))
}
func (upf *Upf) GetQer(authdefqos *models.AuthorizedDefaultQos) (qer *QER, err error) {
	qer = &QER{
		QERID: 1,
		QFI: pfcptypes.QFI{
			QFI: 5,
		},
	}
	return
}
