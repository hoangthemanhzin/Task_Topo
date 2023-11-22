package sm

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/nfs/smf/upman"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/pcf/smpc"
	"etrib5gc/sbi/utils/nasConvert"
	"etrib5gc/util/fsm"
	"fmt"
	"time"

	"github.com/free5gc/nas/nasMessage"
)

const (
	SM_CONTEXT_ACTIVATING_TIMEOUT   time.Duration = 5000 //milliseconds
	SM_CONTEXT_DEACTIVATING_TIMEOUT time.Duration = 5000 //milliseconds
)

type SmfContext interface {
	PlmnId() *models.PlmnId
	DefaultPduSessionType() uint8
	IsIpv4SessionSupported() bool
	IsEthernetSessionSupported() bool
	AddSmContext(*SmContext)
	RemoveSmContext(*SmContext)
}

type SmContext struct {
	logctx.LogWriter
	ctx SmfContext

	supi        string
	pei         string
	sid         uint32 //pdu session id
	ref         string //smctx reference
	pti         uint8
	sessiontype uint8
	dnn         string
	snssai      *models.Snssai
	access      models.AccessType
	rat         models.RatType
	pco         *Pco //protocol configuration options
	dnnconf     models.DnnConfiguration

	estacceptcause uint8

	maxul models.MaxIntegrityProtectedDataRate //max data rate per ue for Uplink UP integrity protection (string value)
	maxdl models.MaxIntegrityProtectedDataRate //max data rate per ue for downlink UP integrity protection (string value)

	smpolid     string //SM policy identity
	smpol       *models.SmPolicyDecision
	activesmpol string //activated SessionRule
	upmanager   *upman.UpManager
	tunnel      *upman.UpTunnel
	//pfcpsessions map[string]*up.PfcpSession //mapping upf id (ip address) to its pfcp session

	fsm.State
	worker     common.Executer
	acttimer   common.UeTimer
	deacttimer common.UeTimer

	amfcli  sbi.ConsumerClient
	pcfcli  sbi.ConsumerClient
	upmfcli sbi.ConsumerClient

	onkill func()
}

func CreateSmContext(ctx SmfContext, upmanager *upman.UpManager, callback models.Callback, info *models.SmContextCreateData) (smctx *SmContext, err error) {
	smctx = &SmContext{
		LogWriter: logctx.WithFields(logctx.Fields{
			"mod":       "sm",
			"supi":      info.Supi,
			"sessionId": info.PduSessionId,
		}),
		worker:    common.NewExecuter(1024),
		State:     fsm.NewState(SM_INACTIVE),
		supi:      info.Supi,
		sid:       uint32(info.PduSessionId),
		snssai:    &info.SNssai,
		dnn:       info.Dnn,
		ctx:       ctx,
		upmanager: upmanager,
	}

	smctx.Info("Create SMContext")
	smctx.ref = common.SmContextRef(smctx.supi, smctx.sid)

	pcfid := common.PcfServiceName(ctx.PlmnId())
	smctx.Infof("Create PCF consumer [%s]", pcfid)
	if smctx.pcfcli, err = mesh.Consumer(meshmodels.ServiceName(pcfid), nil, false); err != nil {
		smctx.Errorf("Create Pcf consumer failed: %+v", err)
		return
	}

	upmfid := common.UpmfServiceName(ctx.PlmnId())
	smctx.Infof("Create Upmf consumer [%s]", upmfid)
	if smctx.upmfcli, err = mesh.Consumer(meshmodels.ServiceName(upmfid), nil, false); err != nil {
		smctx.Errorf("Create Upmf consumer failed: %+v", err)
		return
	}

	smctx.Infof("Create callback AMF consumer [%s]", callback)
	if smctx.amfcli, err = mesh.ClientWithAddr(string(callback)); err != nil {
		smctx.Errorf("Create AMF consumer failed: %+v", err)
		return
	}

	smctx.acttimer = common.NewTimer(SM_CONTEXT_ACTIVATING_TIMEOUT*time.Millisecond, func() {
		smctx.sendEvent(ActTimeoutEvent, nil)
	}, nil)
	smctx.deacttimer = common.NewTimer(SM_CONTEXT_DEACTIVATING_TIMEOUT*time.Millisecond, func() {
		smctx.sendEvent(DeactTimeoutEvent, nil)
	}, nil)

	smctx.Info("SmContext created")
	return
}

func (smctx *SmContext) sendEvent(ev fsm.EventType, args interface{}) error {
	return _sm.SendEvent(smctx.worker, smctx, ev, args)
}

func (smctx *SmContext) PduSessionType() models.PduSessionType {
	//TODO: convert sessiontype from uint8 to models.PduSessionType
	return models.PDUSESSIONTYPE_IPV4
}

func (smctx *SmContext) CreateSmPol() (err error) {
	req := models.SmPolicyContextData{}
	req.Supi = smctx.supi
	req.Dnn = smctx.dnn
	req.PduSessionId = int32(smctx.sid)
	req.AccessType = smctx.access
	req.PduSessionType = smctx.PduSessionType()
	req.Pei = smctx.pei
	req.SliceInfo = *smctx.snssai
	//should be from subscribed data
	req.SubsSessAmbr = models.Ambr{
		Uplink:   "100 Mbps",
		Downlink: "100 Mbps",
	}
	//should be from subscribed data
	req.SubsDefQos = models.SubscribedDefaultQos{
		Var5qi: 5,
		Arp: models.Arp{
			PriorityLevel: 10,
			PreemptCap:    models.PREEMPTIONCAPABILITY_NOT_PREEMPT,
			PreemptVuln:   models.PREEMPTIONVULNERABILITY_NOT_PREEMPTABLE,
		},
		PriorityLevel: 10,
	}
	smctx.Infof("Send CreateSMPolicy to PCF")
	if smctx.smpol, err = smpc.CreateSMPolicy(smctx.pcfcli, req); err != nil {
		smctx.Errorf("Send CreateSMPolicy returned error: %+v", err)
	}
	return
}

func (smctx *SmContext) Ref() string {
	return smctx.ref
}
func (smctx *SmContext) Id() uint32 {
	return smctx.sid
}
func (smctx *SmContext) Pti() uint8 {
	return smctx.pti
}

func (smctx *SmContext) Supi() string {
	return smctx.supi
}
func (smctx *SmContext) Kill() {
	smctx.Infof("Kill SmContext")
	smctx.worker.Terminate()
	if smctx.onkill != nil {
		smctx.onkill()
	}
}

func (smctx *SmContext) checkPduSessionType(reqtype uint8) (err error) {
	//TODO: just for testing
	smctx.sessiontype = nasMessage.PDUSessionTypeIPv4
	//smctx.Infof("session type is %v", smctx.sessiontype)
	return nil

	v4 := false
	v6 := false
	eth := false

	for _, allowedtype := range smctx.dnnconf.PduSessionTypes.AllowedSessionTypes {
		switch allowedtype {
		case models.PDUSESSIONTYPE_IPV4:
			v4 = true
		case models.PDUSESSIONTYPE_IPV6:
			v6 = true
		case models.PDUSESSIONTYPE_IPV4_V6:
			v4 = true
			v6 = true
		case models.PDUSESSIONTYPE_ETHERNET:
			eth = true
		}
	}
	if v4 {
		v4 = smctx.ctx.IsIpv4SessionSupported()
	}
	if v6 {
		v6 = smctx.ctx.IsIpv4SessionSupported()
	}
	if eth {
		eth = smctx.ctx.IsEthernetSessionSupported()
	}

	smctx.estacceptcause = 0
	switch nasConvert.PDUSessionTypeToModels(reqtype) {
	case models.PDUSESSIONTYPE_IPV4:
		if v4 {
			smctx.sessiontype = nasConvert.ModelsToPDUSessionType(models.PDUSESSIONTYPE_IPV4)
		} else {
			return fmt.Errorf("PduSessionType_IPV4 is not allowed in DNN[%s] configuration", smctx.dnn)
		}
	case models.PDUSESSIONTYPE_IPV6:
		if v6 {
			smctx.sessiontype = nasConvert.ModelsToPDUSessionType(models.PDUSESSIONTYPE_IPV6)
		} else {
			return fmt.Errorf("PduSessionType_IPV6 is not allowed in DNN[%s] configuration", smctx.dnn)
		}
	case models.PDUSESSIONTYPE_IPV4_V6:
		if v4 && v6 {
			smctx.sessiontype = nasConvert.ModelsToPDUSessionType(models.PDUSESSIONTYPE_IPV4_V6)
		} else if v4 {
			smctx.sessiontype = nasConvert.ModelsToPDUSessionType(models.PDUSESSIONTYPE_IPV4)
			smctx.estacceptcause = nasMessage.Cause5GSMPDUSessionTypeIPv4OnlyAllowed
		} else if v6 {
			smctx.sessiontype = nasConvert.ModelsToPDUSessionType(models.PDUSESSIONTYPE_IPV6)
			smctx.estacceptcause = nasMessage.Cause5GSMPDUSessionTypeIPv6OnlyAllowed
		} else {
			return fmt.Errorf("PduSessionType_IPV4_V6 is not allowed in DNN[%s] configuration", smctx.dnn)
		}
	case models.PDUSESSIONTYPE_ETHERNET:
		if eth {
			smctx.sessiontype = nasConvert.ModelsToPDUSessionType(models.PDUSESSIONTYPE_ETHERNET)
		} else {
			return fmt.Errorf("PduSessionType_ETHERNET is not allowed in DNN[%s] configuration", smctx.dnn)
		}
	default:
		return fmt.Errorf("Requested PDU Sesstion type[%d] is not supported", reqtype)
	}
	return nil
}

func (smctx *SmContext) BuildSmContextCreatedData() (dat models.SmContextCreatedData) {
	dat = models.SmContextCreatedData{
		SmContextRef: smctx.ref,
	}
	return
}

func (smctx *SmContext) setupTunnel() (err error) {
	//populate the PDRs with an activated session rule (and pcc rules?) from smctx
	srule := smctx.getActivatedSessionRule()
	err = smctx.tunnel.FillPdrs(srule, 256, smctx.dnn)
	return
}

func (smctx *SmContext) getActivatedSessionRule() (srule *models.SessionRule) {
	if smctx.smpol != nil && len(smctx.smpol.SessRules) > 0 {
		if len(smctx.activesmpol) > 0 {
			rule, _ := smctx.smpol.SessRules[smctx.activesmpol]
			srule = &rule
		} else {
			for id, rule := range smctx.smpol.SessRules {
				smctx.activesmpol = id
				srule = &rule
				//just take the first rule as an activated rule
				//TODO: needs to known when a session rule is activated
				break
			}
		}
	}
	return
}

func (smctx *SmContext) Close() {
	smctx.sendEvent(CloseEvent, nil)
}
func (smctx *SmContext) notifyAmf() (err error) {
	smctx.Warn("Notify AMF not implemented")
	return
}

func (smctx *SmContext) terminatePolAsso() (err error) {
	smctx.Warn("Terminate Policy Association not implemented")
	return
}
