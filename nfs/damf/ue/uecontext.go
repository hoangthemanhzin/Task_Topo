package ue

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/nas"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/sbi/utils/nasConvert"
	"etrib5gc/util/fsm"
	"fmt"
	"time"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

const (
	UE_LIFETIME    = 3000 //miliseconds //How long UeContex can stay in DAMF
	T3560_DURATION = 100  //miliseconds
	MAX_T3560_CNT  = 3
)

type Context interface {
	GetT3502() uint8
	RemoveUe(*UeContext)
	FindAmfId([]models.Snssai) string
	GetUeId() int64
}

type UeContext struct {
	logctx.LogWriter
	fsm.State
	ctx       Context
	plmnid    models.PlmnId
	suci      string
	supi      string
	ngksi     models.NgKsi
	abba      []byte
	rand      []byte
	autn      []byte
	kamf      []byte
	hxresstar []byte
	eap       string
	authtype  models.AuthType

	sslices []models.Snssai //subscribed slices (from UDM)
	aslices []models.Snssai //allowed slices, either received from UE or being determined from sslices

	worker     common.Executer //thread to handle related events in all procedures
	ranueid    int64           //id at RAN
	amfueid    int64           //id at AMF
	t3560      common.UeTimer
	alivetimer common.UeTimer
	amfid      string

	//tai          models.Tai
	//routingId    string
	//trsr         string
	ausfcli sbi.ConsumerClient
	udmcli  sbi.ConsumerClient
	rancli  sbi.ConsumerClient
	msg     *nasMessage.RegistrationRequest
}

func CreateUeContext(ctx Context, callback models.Callback, msg *n2models.InitUeContextRequest) (uectx *UeContext, err error) {
	uectx = &UeContext{
		LogWriter: logctx.WithFields(logctx.Fields{
			"ranUeId": msg.RanUeId,
		}),
		State:   fsm.NewState(UE_IDLE),
		ctx:     ctx,
		ranueid: msg.RanUeId,
		abba:    []uint8{0x00, 0x00},
	}
	var nasMsg libnas.Message
	if nasMsg, err = nas.Decode(nil, msg.NasPdu); err != nil {
		return
	} else {
		if uectx.msg = nasMsg.RegistrationRequest; uectx.msg == nil {
			err = fmt.Errorf("Unexpected Nas message")
			return
		}
	}

	//process RegistrationRequest message
	if err = uectx.init(); err != nil {
		return
	}

	if uectx.rancli, err = mesh.ClientWithAddr(string(callback)); err != nil {
		err = fmt.Errorf("Create PRan consumer failed: %s", err.Error())
		return
	}

	uectx.amfueid = ctx.GetUeId()
	uectx.worker = common.NewExecuter(1024)
	uectx.t3560 = common.NewTimer(T3560_DURATION*time.Millisecond, func() {
		//AuthenticationRequest expired
		uectx.sendEvent(T3560Event, nil)
	}, nil)
	uectx.alivetimer = common.NewTimer(UE_LIFETIME*time.Millisecond, func() {
		uectx.Tracef("UeContext reaches its end-of-life")
		uectx.sendEvent(CloseEvent, nil)
	}, nil)
	uectx.alivetimer.Start()

	//update log fields for UeContext
	uectx.LogWriter = uectx.WithFields(logctx.Fields{
		"ue-suci": uectx.suci,
	})
	return
}

func (uectx *UeContext) init() (err error) {
	uectx.Info("Init UeContext with RegistrationRequest")
	content := uectx.msg.MobileIdentity5GS.GetMobileIdentity5GSContents()
	idtype := nasConvert.GetTypeOfIdentity(content[0])
	if idtype != nasMessage.MobileIdentity5GSTypeSuci {
		err = fmt.Errorf("SUCI missing")
		return
	}
	var id *models.PlmnId
	if uectx.suci, _ = nasConvert.SuciToString(content); len(uectx.suci) == 0 {
		err = fmt.Errorf("Decode Suci failed %x", content)
		return
	}

	if id, err = models.Bytes2PlmnId(content[1:4]); err == nil {
		uectx.plmnid = *id
		uectx.Infof("suci=%s; plmnid=%s", uectx.suci, uectx.plmnid.String())
		sid := common.AusfServiceName(&uectx.plmnid)
		if uectx.ausfcli, err = mesh.Consumer(meshmodels.ServiceName(sid), nil, false); err != nil {
			err = fmt.Errorf("Create Ausf consumer failed: %s", err.Error())
			return
		}
		sid = common.UdmServiceName(&uectx.plmnid)
		if uectx.udmcli, err = mesh.Consumer(meshmodels.ServiceName(sid), nil, false); err != nil {
			err = fmt.Errorf("Create Udm consumer failed: %s", err.Error())
			return
		}

	} else {
		err = fmt.Errorf("Decode PlmnId failed: %s", err.Error())
		return
	}
	var ngksi models.NgKsi
	// NgKsi: TS 24.501 9.11.3.32
	switch uectx.msg.NgksiAndRegistrationType5GS.GetTSC() {
	case nasMessage.TypeOfSecurityContextFlagNative:
		//log.Info("sctype_native")
		ngksi.Tsc = models.SCTYPE_NATIVE
	case nasMessage.TypeOfSecurityContextFlagMapped:
		//log.Info("sctype_mapped")
		ngksi.Tsc = models.SCTYPE_MAPPED
	}
	ngksi.Ksi = int32(uectx.msg.NgksiAndRegistrationType5GS.GetNasKeySetIdentifiler())
	if ngksi.Tsc == models.SCTYPE_NATIVE && ngksi.Ksi != 7 {
	} else {
		ngksi.Tsc = models.SCTYPE_NATIVE
		ngksi.Ksi = 0
	}

	uectx.ngksi = ngksi
	return
}

func (uectx *UeContext) AmfUeId() int64 {
	return uectx.amfueid
}
func (uectx *UeContext) RanUeId() int64 {
	return uectx.ranueid
}

func (uectx *UeContext) Close() {
	uectx.sendEvent(CloseEvent, nil)
}
func (uectx *UeContext) Kill() {
	uectx.worker.Terminate()
	uectx.Tracef("UeContext terminated")
}

func (uectx *UeContext) sendEvent(ev fsm.EventType, args interface{}) error {
	return _sm.SendEvent(uectx.worker, uectx, ev, args)
}
func (uectx *UeContext) ServingNetwork() string {
	str := common.ServingNetworkName(&uectx.plmnid)
	uectx.Tracef("Serving network:%s", str)
	return str
}
