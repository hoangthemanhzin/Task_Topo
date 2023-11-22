package context

import (
	"etrib5gc/nfs/smf/config"
	"etrib5gc/nfs/smf/sm"
	"etrib5gc/sbi/models"
	"sync"
	"sync/atomic"

	"github.com/free5gc/nas/nasMessage"
)

//QFI = QoS Flow Identifier
//it is a 6-bit scalar value assgined by the SMF and applied to the UE, gNB,
//and UPF during a QoS Flow setup. QFI is present in Gtp header  so that UPF
//knows how to process the packet and maps it to the corresponding QoS Flow.

//QER: QoS Enforcement Rule: SMF sends UPF the QFI/QER mapping and they are
//kept by the UPF in a mapping table.

//RQI: Reflective QoS Indicator: a bit that indicates to the UE that it can
//copy the DL's QoS settings for the corresponding UL traffic to simplify QoS
//signalling for the UE in the 5GS

//PDR: Packet Detection Rule: is sent from SMF to UPF during QoS Flow setup.
//UPF uses 5-tuple to match PDR so that it map DL packet to a corresponding QoS
//Flow.

//Forwarding Action Rule (FAR)
// Usage Report Rule (URR)
//Buffering Action Rule (BAR)

//(non)-GBR (Guanranteed Bit Rate)

//SDF Service Data Flow

//When a PCF sends a PCC Rule to the SMF, the SMF formulates different QoS
//constructs for other entities along the QoS Flow as follows:
// - SDF template to UPF (PFCP)
// - QoS Profile to GnB (N2)
// - QoS Rule to UE (N1)

//Reading SMF source code:

/*
UEDefaultPaths holds a list of anchoring UPFs and a pool of data paths. The pool is indexed with anchoring UPFs' names.
The UEDefaultPaths is created from the configuration file.

A DataPath is a linked list consisting of DataPathNode items. So, it has a start node, and some other supporting attributes such as destination information (anchor), branching, etc

A DataPathNode has: a reference to UPF, uplink and downlink GtpTunnel which are also pointers to next/prev nodes.
*/
type SmfContext struct {
	plmnid      models.PlmnId
	dnn         string
	snssai      models.Snssai
	sessiontype string //defaul pdu session type (in case UE does not specify)
	smlist      SmList

	closech chan *sm.SmContext //to receive SmContext for removing
	done    chan bool
	wg      sync.WaitGroup
	closed  int32
}

func New(cfg *config.SmfConfig) *SmfContext {
	_initLog()
	ret := &SmfContext{
		plmnid:      models.PlmnId(cfg.PlmnId),
		dnn:         cfg.Dnn,
		snssai:      cfg.Slice,
		sessiontype: "IPv4",
		smlist:      newSmList(),
		closech:     make(chan *sm.SmContext),
		done:        make(chan bool),
		closed:      0,
	}

	go ret.loop()
	return ret
}

func (ctx *SmfContext) loop() {
	ctx.wg.Add(1)
	defer ctx.wg.Done()
LOOP:
	for {
		select {
		case <-ctx.done:
			log.Trace("Receive a signal to close Context's loop")
			break LOOP
		case smctx := <-ctx.closech:
			ctx.smlist.remove(smctx)
			smctx.Kill() //terminate worker
		}
	}
	//clean up any remaining SmContexts
	for ctx.smlist.size() > 0 {
		smctx := <-ctx.closech
		log.Info("Kill it")
		ctx.smlist.remove(smctx)
		smctx.Kill()
	}
}

func (ctx *SmfContext) RemoveSmContext(smctx *sm.SmContext) {
	ctx.closech <- smctx
}

func (ctx *SmfContext) Clean() {
	log.Infof("Clear SmContexts")
	atomic.StoreInt32(&ctx.closed, 1) //do not accept new SmContext
	ctx.smlist.clear()                //add all SmContexts to closing list
	close(ctx.done)
	ctx.wg.Wait()
}

func (ctx *SmfContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}

func (ctx *SmfContext) Dnn() string {
	return ctx.dnn
}

func (ctx *SmfContext) Snssai() *models.Snssai {
	return &ctx.snssai
}
func (ctx *SmfContext) DefaultPduSessionType() (t uint8) {
	switch ctx.sessiontype {
	case "IPv4":
		t = nasMessage.PDUSessionTypeIPv4
	case "IPv6":
		t = nasMessage.PDUSessionTypeIPv6
	case "IPv4v6":
		t = nasMessage.PDUSessionTypeIPv4IPv6
	case "Ethernet":
		t = nasMessage.PDUSessionTypeEthernet
	default:
		t = nasMessage.PDUSessionTypeIPv4
	}
	return
}
func (ctx *SmfContext) IsClosed() bool {
	return atomic.LoadInt32(&ctx.closed) == 1
}

func (ctx *SmfContext) IsIpv4SessionSupported() bool {
	//TODO: use configued parameters
	return false
}

func (ctx *SmfContext) IsIpv6SessionSupported() bool {
	//TODO: use configued parameters
	return false
}

func (ctx *SmfContext) IsEthernetSessionSupported() bool {
	//TODO: use configued parameters
	return false
}

func (ctx *SmfContext) FindSmContext(ref string) *sm.SmContext {
	return ctx.smlist.find(ref)
}

func (ctx *SmfContext) AddSmContext(smctx *sm.SmContext) {
	ctx.smlist.add(smctx)
}
