package ue

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/util/fsm"
	"net"
	"time"

	"github.com/free5gc/ngap/ngapType"
)

const (
	UE_LIFETIME = 3000 //miliseconds //How long UeContex can stay in Idle state
)

type Context interface {
	DamfName() string //create default Amf service name
	Callback() models.Callback
	GetCuNgapId() int64  //generate a new CuNgapId
	AddUe(*UeContext)    //add UeContext to pool
	RemoveUe(*UeContext) //ask to remove and terminate UeContext
}

type Ran interface {
	Conn() net.Conn
	Access() models.AccessType
	Send([]byte) error //send encoded Ngap message to Ran
}

type HandoverInfo struct {
	HandoverType ngapType.HandoverType
}

type UeAuth struct {
	suci      string
	ngksi     models.NgKsi
	abba      []byte
	rand      []byte
	autn      []byte
	kamf      []byte
	hxresstar []byte
	eap       string
	authtype  models.AuthType
}

type UeContext struct {
	logctx.LogWriter
	fsm.State

	cu  Context
	ran Ran

	worker    common.Executer //thread to handle related events in all procedures
	ranNgapId int64
	cuNgapId  int64
	amfUeId   int64
	//	fivegstmsi   *ngapType.FiveGSTMSI
	//	amfsetid     *ngapType.AMFSetID
	//	allowednssai *ngapType.AllowedNSSAI

	tai          models.Tai
	routingId    string
	trsr         string
	handoverInfo *HandoverInfo

	amfcli sbi.ConsumerClient
	amfid  string //will be set by default amf

	uea  UeAuth
	supi string

	slices  []models.Snssai
	initmsg *n2models.InitUeContextRequest

	alivetimer  common.UeTimer
	pendingjobs PendingJobs //pending SBI request processing jobs (waiting for Ngap's response)
}

func NewUeContext(ran Ran, ctx Context, ranNgapId int64, fivegstmsi *ngapType.FiveGSTMSI, amfsetid *ngapType.AMFSetID, allowednssai *ngapType.AllowedNSSAI) *UeContext {
	uectx := &UeContext{
		State:     fsm.NewState(CM_IDLE),
		ran:       ran,
		cu:        ctx,
		worker:    common.NewExecuter(1024),
		ranNgapId: ranNgapId,
		cuNgapId:  ctx.GetCuNgapId(),
		uea: UeAuth{
			abba: []uint8{0x00, 0x00},
		},
		pendingjobs: newPendingJobs(),
	}
	uectx.setAmfCli(fivegstmsi, amfsetid, allowednssai)
	ctx.AddUe(uectx)
	uectx.alivetimer = common.NewTimer(UE_LIFETIME*time.Millisecond, func() {
		uectx.Trace("UeContext reaches its end-of-life")
		uectx.sendEvent(EndOfLifeEvent, nil)
	}, nil)
	uectx.alivetimer.Start()
	uectx.LogWriter = logctx.WithFields(logctx.Fields{
		"ranNgapId": uectx.ranNgapId,
		"cuNgapId":  uectx.cuNgapId,
	})
	uectx.Info("UeContext created")
	return uectx
}
func (uectx *UeContext) setAmfCli(fivegstmsi *ngapType.FiveGSTMSI, amfsetid *ngapType.AMFSetID, allowednssai *ngapType.AllowedNSSAI) {
	//TODO: create non-default AmfCli
	if fivegstmsi != nil {
		uectx.Info("UE has TMSI")
	}

	if amfsetid != nil {
		uectx.Info("UE has amfsetid")
	}

	if allowednssai != nil {
		uectx.Info("UE has allowedNssai")
	}
}

// send to RAN
func (uectx *UeContext) Send(buf []byte) error {
	return uectx.ran.Send(buf)
}

func (uectx *UeContext) Close() {
	uectx.sendEvent(CloseEvent, nil)
}

func (uectx *UeContext) Kill() {
	uectx.worker.Terminate()
	uectx.Infof("UeContext terminated")
}

func (uectx *UeContext) sendEvent(ev fsm.EventType, args interface{}) error {
	return _sm.SendEvent(uectx.worker, uectx, ev, args)
}

func (uectx *UeContext) Access() models.AccessType {
	return uectx.ran.Access()
}

func (uectx *UeContext) getAmf() (err error) {
	amf := common.AmfServiceName(&uectx.initmsg.UeCtx.PlmnId, uectx.amfid)
	uectx.amfcli, err = mesh.Consumer(meshmodels.ServiceName(amf), nil, false)
	return
}

func (uectx *UeContext) getDefaultAmf() (err error) {
	damf := uectx.cu.DamfName()
	uectx.amfcli, err = mesh.Consumer(meshmodels.ServiceName(damf), nil, false)
	return
}

func (uectx *UeContext) Amf() sbi.ConsumerClient {
	if uectx.amfcli == nil {
		uectx.Errorf("AMF not found")
	}
	return uectx.amfcli
}

func (uectx *UeContext) RanConn() net.Conn {
	return uectx.ran.Conn()
}
func (uectx *UeContext) CuNgapId() int64 {
	return int64(uectx.cuNgapId)
}

func (uectx *UeContext) AmfUeId() int64 {
	return int64(uectx.amfUeId)
}

func (uectx *UeContext) RanNgapId() int64 {
	return int64(uectx.ranNgapId)
}

func (uectx *UeContext) HandoverInfo() *HandoverInfo {
	return uectx.handoverInfo
}
func (uectx *UeContext) RoutingId() string {
	return uectx.routingId
}
func (uectx *UeContext) Trsr() string {
	return uectx.trsr
}
func (uectx *UeContext) SetTrsr(trsr string) {
	uectx.trsr = trsr
}

func (uectx *UeContext) SetRoutingId(id string) {
	uectx.routingId = id
}

func (uectx *UeContext) UeCtxRelReq(msg *n2models.UeCtxRelReq) {
	//TODO
}

func (uectx *UeContext) UeCtxRelCmpl(msg *n2models.UeCtxRelCmpl) {
	//TODO
}

/*
func (uectx *UeContext) RrcInactTranRep(msg *n2models.RrcInactTranRep) {
	//TODO
}

func (uectx *UeContext) PduSessResSetRsp(msg *n2models.PduSessResSetRsp) {
	//TODO
}

func (uectx *UeContext) PduSessResModRsp(msg *n2models.PduSessResModRsp) {
	//TODO
}
func (uectx *UeContext) PduSessResRelRsp(msg *n2models.PduSessResRelRsp) {
	//TODO
}

func (uectx *UeContext) PduSessResNot(msg *n2models.PduSessResNot) {
	//TODO
}

func (uectx *UeContext) PduSessResModInd(msg *n2models.PduSessResModInd) {
	//TODO
}
*/
