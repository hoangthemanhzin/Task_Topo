package ranuecontext

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/nfs/amf/ranuecontext/secmode"
	"etrib5gc/nfs/amf/uecontext"

	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/util/fsm"
	"time"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

var _logfields logctx.Fields = logctx.Fields{
	"mod": "ranuecontext",
}

type CommonProc interface {
	Handle(*libnas.GmmMessage) error
}

type AmfContext interface {
	//T3550() time.Duration
	T3550() int
	T3513() int
	T3565() int
	T3522() int
	T3502() uint8
	AllocateAmfUeId() int64
	AddRanUe(*RanUe)
	RemoveRanUe(*RanUe)
	AmfId() string
	SecAlgs() *common.NasSecAlgList
	FillRegistrationAccept(*nasMessage.RegistrationAccept, *RanUe) error
	Callback() models.Callback
}

type RanUe struct {
	logctx.LogWriter
	fsm.State

	rancli sbi.ConsumerClient
	amf    AmfContext
	ue     *uecontext.UeContext //point to the shared UE context (supi, suci, pei etc.)
	secctx *nas.SecCtx

	access   models.AccessType
	rannets  []string
	rrccause string

	drx          uint8 //UE specific DRX parameter
	cap5gmm      nasType.Capability5GMM
	regarea      []models.Tai //regisstration areas
	allowednssai []models.AllowedSnssai

	registered bool //indicating registration status

	activeproc    CommonProc
	regctx        *events.RegistrationContext //registration context
	pendingnaspdu []byte
	contextsent   bool    //has initial context setup request been sent to ran?
	n1n2man       N1N2Man //manage pending N1N2MessageTransfer
	acceptance    *acceptanceReport

	ranUeId int64
	amfUeId int64

	t3522 common.UeTimer //deregistration request
	t3550 common.UeTimer //registration accept
	t3513 common.UeTimer //paging
	t3565 common.UeTimer //notification
}

func NewRanUe(amf AmfContext, ue *uecontext.UeContext, access models.AccessType, rannets []string, callback models.Callback, ranUeId int64) (ranue *RanUe, err error) {
	ranue = &RanUe{
		State:   fsm.NewState(MM_DEREGISTERED),
		amf:     amf,
		ue:      ue,
		rannets: rannets,
		access:  access,
		amfUeId: amf.AllocateAmfUeId(),
		ranUeId: ranUeId,
		n1n2man: newN1N2Man(),
	}

	ranue.LogWriter = ue.LogWriter.WithFields(_logfields).WithFields(logctx.Fields{
		"amfUeId": ranue.amfUeId,
		"ranUeId": ranue.ranUeId,
	})
	if ranue.rancli, err = mesh.ClientWithAddr(string(callback)); err != nil {
		err = common.WrapError("Create Ran consumer failed", err)
		return
	}

	//initialize Security Mode context from UeContext
	//NOTE: it not clear if RanUe should take the current Security Mode context
	//from UeContext or not.
	ranue.secctx = ue.SecCtx()

	ranue.t3550 = common.NewTimer(time.Duration(amf.T3550())*time.Millisecond, func() {
		//IdentityRequest expired
		ranue.sendEvent(T3550Event, nil)
	}, nil)

	ranue.t3513 = common.NewTimer(time.Duration(amf.T3513())*time.Millisecond, func() {
		ranue.Trace("T3513 expired")
		ranue.sendEvent(T3513Event, nil)
	}, nil)
	ranue.t3522 = common.NewTimer(time.Duration(amf.T3522())*time.Millisecond, func() {
		ranue.Trace("T3522 expired")
		ranue.sendEvent(T3522Event, nil)
	}, nil)
	ranue.t3565 = common.NewTimer(time.Duration(amf.T3565())*time.Millisecond, func() {
		ranue.Trace("T3565 expired")
		ranue.sendEvent(T3513Event, nil)
	}, nil)
	amf.AddRanUe(ranue)

	ranue.Infof("RanUe created")
	return
}
func (ranue *RanUe) Worker() common.Executer {
	return ranue.ue.Worker()
}

func (ranue *RanUe) UeContext() *uecontext.UeContext {
	return ranue.ue
}

func (ranue *RanUe) Access() models.AccessType {
	return ranue.access
}

func (ranue *RanUe) RanNets() []string {
	return ranue.rannets
}

func (ranue *RanUe) AmfUeId() int64 {
	return ranue.amfUeId
}

func (ranue *RanUe) RanUeId() int64 {
	return ranue.ranUeId
}

func (ranue *RanUe) RegistrationContext() *events.RegistrationContext {
	return ranue.regctx
}
func (ranue *RanUe) sendEvent(ev fsm.EventType, args interface{}) error {
	ranue.Tracef("in ranue Send event %d to statemachine:", ev)
	return _sm.SendEvent(ranue.ue.Worker(), ranue, ev, args)
}
func (ranue *RanUe) BearerType() uint8 {
	return common.BearerType(ranue.access)
}

func (ranue *RanUe) IsEmergency() bool {
	return ranue.rrccause == "0"
}

func (ranue *RanUe) SecCtx() (ctx *nas.SecCtx) {
	return ranue.secctx
}

// security mode establishment complete
func (ranue *RanUe) UpdateSecCtx() {
	ranue.secctx.Update(ranue.access)
}

// return a valid sec context, otherwise return nil
func (ranue *RanUe) NasSecCtx() (ctx *nas.NasSecCtx) {
	if ranue.secctx != nil {
		//ranue.Trace("Has a non-current Nas Security Context from current security mode establishment procedure")
		ctx = &ranue.secctx.NasSecCtx
	} /* else if secctx := ranue.ue.SecCtx(); secctx != nil {
		if seccx.IsValid() {
			ranue.Trace("Has a current Nas Security Context")
			ctx = &secctx.NasSecCtx
		}
	}
	*/
	return
}

// NOTE: need to use atomic reading
func (ranue *RanUe) IsRegistered() bool {
	return ranue.registered
}
func (ranue *RanUe) Detach() {
	//TODO: cleanup
	ranue.Infof("Detach UeContext from RanUe:%d", ranue.ranUeId)
	ranue.ue = nil //detach from UeContext
	ranue.amf.RemoveRanUe(ranue)
}

func (ranue *RanUe) isAuthenticated() bool {
	kamf := ranue.regctx.AuthCtx().Kamf
	return len(kamf) > 0
}

func (ranue *RanUe) hasId() bool {
	return len(ranue.ue.Suci()) > 0
}

func (ranue *RanUe) hasValidSecmode() bool {
	return ranue.secctx != nil && ranue.secctx.IsValid()
}

func (ranue *RanUe) startSecmode() {
	ranue.Infof("Start security mode establishment procedure")

	//create the security mode context
	authctx := ranue.regctx.AuthCtx()
	secctx := nas.NewSecCtx(authctx.Kamf)
	algs := ranue.amf.SecAlgs()
	secctx.SelectAlg(algs.IntegrityOrder, algs.CipheringOrder, ranue.ue.SecCap())
	hdp := nas.HDP_MOBILITY_UPDATE //mobility/registration update
	if err := secctx.DeriveAlgKeys(hdp); err != nil {
		//send a reject
		ranue.Errorf(err.Error())
		ranue.rejectRegistration(nasMessage.Cause5GMMSecurityModeRejectedUnspecified)
		return
	}
	ranue.secctx = secctx
	rinmr := true //retranmission initial NAS message
	//start the security mode context establishment procedure with the UE
	ranue.activeproc = secmode.New(ranue, authctx.Eap, authctx.Success, hdp, rinmr, func(proc *secmode.SecmodeProc) {
		ranue.sendEvent(SecmodeCmplEvent, proc)
	})
}

func (ranue *RanUe) rejectRegistration(cause uint8) (err error) {
	//send a registration reject
	err = ranue.sendRegistrationReject(cause)
	ranue.sendEvent(FailEvent, err)
	return
}

func (ranue *RanUe) rejectService(cause uint8, statuslist *[16]bool, release bool) (err error) {
	if err = ranue.sendServiceReject(cause, statuslist); err == nil {
		if release {
			//ranue.conn.rancli.SendUeCtxRelCmd()
		}
	}

	ranue.sendEvent(FailEvent, nil)
	return
}

func (ranue *RanUe) RanKey() (k []byte) {
	if ranue.access == models.ACCESSTYPE__3_GPP_ACCESS {
		k = ranue.secctx.Kgnb()
	} else {
		k = ranue.secctx.Kn3iwf()
	}
	return
}
