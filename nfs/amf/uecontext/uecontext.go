package uecontext

import (
	"encoding/hex"
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/nasConvert"
	"etrib5gc/util/fsm"
	"fmt"
	"sync"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

const (
	UE_ID_TYPE_SUCI uint8 = iota
	UE_ID_TYPE_SUPI
	UE_ID_TYPE_PEI
	UE_ID_TYPE_TMSI5GS
	UE_ID_TYPE_GUTI
)

var _logfields logctx.Fields = logctx.Fields{
	"mod": "uecontext",
}

type Ladn struct {
	Dnn     string
	TaiList []models.Tai
}

type RanUe interface {
	Detach()
	Access() models.AccessType
	HandleEvent(*common.EventData) error
	RegistrationContext() *events.RegistrationContext
	IsRegistered() bool
}

type AmfContext interface {
	Uri() string
	UpdateUeId(string, uint8, *UeContext)
	DefaultDnn(models.AccessType) string
}

type UeContext struct {
	logctx.LogWriter
	fsm.State
	worker  common.Executer
	amf     AmfContext
	supi    string
	suci    string
	guti    string
	pei     string
	tmsi5gs string
	tmsi    int32
	gpsi    string

	abba         []byte
	plmnid       models.PlmnId
	ladninfo     []Ladn
	lasttai      *models.Tai //last visited TAI
	ampol        *models.PolicyAssociation
	smfsel       *models.SmfSelectionSubscriptionData
	subdat       *models.SubscribedData
	subnssai     []models.SubscribedSnssai
	tracedat     *models.TraceData
	reachability models.UeReachability
	amdat        *models.AccessAndMobilitySubscriptionData
	sdmsubid     string //Sdm subscription id
	rat          models.RatType
	loc          models.UserLocation
	ueinsmf      *models.UeContextInSmfData

	slicesubchanged bool //has slice subscription been changed?

	secctx *nas.SecCtx //current active security context; the non-current one should be in in the secmode procedure of a RanState
	ngksi  *models.NgKsi
	seccap *nasType.UESecurityCapability

	ranfaces map[models.AccessType]RanUe

	sclist SessionContextList //list of session contex

	mutex sync.Mutex

	ausfcli sbi.ConsumerClient
	udmcli  sbi.ConsumerClient
	pcfcli  sbi.ConsumerClient

	onkill func()
}

func NewUeContextWithId(amf AmfContext, id string, idtype uint8) *UeContext {
	ue := NewUeContext(amf)
	ue.setId(id, idtype)
	return ue
}

func NewUeContext(amf AmfContext) *UeContext {
	ue := &UeContext{
		LogWriter: logctx.WithFields(_logfields),
		State:     fsm.NewState(MM_IDLE),
		worker:    common.NewExecuter(1024),
		amf:       amf,
		abba:      []uint8{0x00, 0x00},
		ranfaces:  make(map[models.AccessType]RanUe), //auto attach
		sclist:    newSessionContextList(),
	}

	return ue
}

func (ue *UeContext) GetUeContextInSmfData() *models.UeContextInSmfData {
	return ue.ueinsmf
}

func (ue *UeContext) Worker() common.Executer {
	return ue.worker
}

func (ue *UeContext) UpdatePei(pei string) {
	ue.pei = pei
	ue.updateLogger(pei, UE_ID_TYPE_PEI)
	ue.amf.UpdateUeId(ue.pei, UE_ID_TYPE_PEI, ue)
}

// content is from a Nas IdentityResponse
// update id for the UE then ask AmfContext to update UePool
func (ue *UeContext) UpdateId(content []byte) (err error) {
	idtype := nasConvert.GetTypeOfIdentity(content[0])
	switch idtype {
	case nasMessage.MobileIdentity5GSTypeNoIdentity:
		ue.Warn("IdentityResponse without an identity")

	case nasMessage.MobileIdentity5GSTypeSuci:
		ue.Tracef("IdentityResponse with a suci: %x", content)
		//suci is requested only if the pending RegistrationRequest does not have one
		if err = ue.SetPlmnId(content[1:4]); err == nil {
			ue.suci, _ = nasConvert.SuciToString(content)
			ue.updateLogger(ue.suci, UE_ID_TYPE_SUCI)
			ue.amf.UpdateUeId(ue.suci, UE_ID_TYPE_SUCI, ue)
		}

	case nasMessage.MobileIdentity5GSType5gGuti:
		//there is no use case yet where GUTI is requested by AMF
		ue.Warn("Guti identity is not handled")

	case nasMessage.MobileIdentity5GSTypeImei:
		ue.Trace("IMEI")
		ue.pei = nasConvert.PeiToString(content)
		ue.updateLogger(ue.pei, UE_ID_TYPE_PEI)
		ue.amf.UpdateUeId(ue.pei, UE_ID_TYPE_PEI, ue)

	case nasMessage.MobileIdentity5GSTypeImeisv:
		ue.Trace("IMEISV")
		ue.pei = nasConvert.PeiToString(content)
		ue.updateLogger(ue.pei, UE_ID_TYPE_PEI)
		ue.amf.UpdateUeId(ue.pei, UE_ID_TYPE_PEI, ue)
	}
	return
}

func (ue *UeContext) updateLogger(id string, idtype uint8) {
	switch idtype {
	case UE_ID_TYPE_SUCI:
		ue.LogWriter = ue.WithFields(
			logctx.Fields{
				"ue-suci": id,
			})
	case UE_ID_TYPE_SUPI:
		ue.LogWriter = ue.WithFields(
			logctx.Fields{
				"ue-supi": id,
			})

	case UE_ID_TYPE_PEI:
		ue.LogWriter = ue.WithFields(
			logctx.Fields{
				"ue-pei": id,
			})

	default:
		//do nothing
	}

}

func (ue *UeContext) setId(id string, idtype uint8) {
	switch idtype {
	case UE_ID_TYPE_SUCI:
		ue.suci = id
	case UE_ID_TYPE_SUPI:
		ue.supi = id

	case UE_ID_TYPE_PEI:
		ue.pei = id

	case UE_ID_TYPE_TMSI5GS:
		ue.tmsi5gs = id
	case UE_ID_TYPE_GUTI:
		//TODO: get TMSI5GS from GUTI
	default:
		//do nothing
	}
	ue.updateLogger(id, idtype)
}

func (ue *UeContext) sendEvent(ev fsm.EventType, args interface{}) error {
	ue.Tracef("in ranstate Send event %d to statemachine:", ev)
	return _sm.SendEvent(ue.worker, ue, ev, args)
}

func (ue *UeContext) update(regctx *events.RegistrationContext) {
	ue.Trace("Update UE with information from InitialUeMsg")
	authctx := regctx.AuthCtx()
	// UE has been authenticated (from DAMF)
	if len(authctx.Supi) > 0 {
		ue.supi = authctx.Supi
		ue.amf.UpdateUeId(ue.supi, UE_ID_TYPE_SUPI, ue)
	}

	if msg := regctx.RegistrationRequest(); msg != nil {
		//RegistrationRequest
		var ngksi models.NgKsi
		// NgKsi: TS 24.501 9.11.3.32
		switch msg.NgksiAndRegistrationType5GS.GetTSC() {
		case nasMessage.TypeOfSecurityContextFlagNative:
			//ue.Info("sctype_native")
			ngksi.Tsc = models.SCTYPE_NATIVE
		case nasMessage.TypeOfSecurityContextFlagMapped:
			//ue.Info("sctype_mapped")
			ngksi.Tsc = models.SCTYPE_MAPPED
		}
		ngksi.Ksi = int32(msg.NgksiAndRegistrationType5GS.GetNasKeySetIdentifiler())
		if ngksi.Tsc == models.SCTYPE_NATIVE && ngksi.Ksi != 7 {
		} else {
			ngksi.Tsc = models.SCTYPE_NATIVE
			ngksi.Ksi = 0
		}
		ue.ngksi = &ngksi
		if msg.UESecurityCapability != nil {
			ue.Trace("UE has Security Capability")
			ue.seccap = msg.UESecurityCapability
		}
	} else {
		//TODO: update UE with information from the ServiceRequest
	}
}

// check if ue is in allowed service area for pdu session establishment
func (ue *UeContext) IsReEstablishPduSessionAllowed() bool {
	//TODO: check if pdu session re-estabishment is allowed
	/*
		if ue.AmPolicyAssociation != nil && ue.AmPolicyAssociation.ServAreaRes != nil {
			switch ue.AmPolicyAssociation.ServAreaRes.RestrictionType {
			case models.RestrictionType_ALLOWED_AREAS:
				allowReEstablishPduSession = context.TacInAreas(ue.Tai.Tac, ue.AmPolicyAssociation.ServAreaRes.Areas)
			case models.RestrictionType_NOT_ALLOWED_AREAS:
				allowReEstablishPduSession = !context.TacInAreas(ue.Tai.Tac, ue.AmPolicyAssociation.ServAreaRes.Areas)
			}
		}
	*/

	return true
}

// check if an Snssai is allowed in an access
func (ue *UeContext) IsSnssaiAllowed(snssai models.Snssai, access models.AccessType) bool {
	//TODO: to be implemente
	return true
}

func (ue *UeContext) cmConnected(access models.AccessType) (ok bool) {
	_, ok = ue.ranfaces[access]
	return
}

func (ue *UeContext) IsRegistered(access models.AccessType) (ok bool) {
	if ranue, ok := ue.ranfaces[access]; ok {
		ok = ranue.IsRegistered()
	}
	return
}
func (ue *UeContext) DefaultAllowedSnssai(access models.AccessType) *models.Snssai {
	/* TODO:
	if allowedNssai, ok := ue.AllowedNssai[anType]; ok {
				snssai = *allowedNssai[0].AllowedSnssai
			} else {
				return errors.New("Ue doesn't have allowedNssai")
			}
	*/
	return nil
}
func (ue *UeContext) AllowedNssai(access models.AccessType) []models.AllowedSnssai {
	//TODO: now it is just a dummy one
	return []models.AllowedSnssai{
		models.AllowedSnssai{
			AllowedSnssai: models.Snssai{
				Sst: 10,
				Sd:  "ff",
			},
		},
	}
}

// Ue Radio Capability
func (ue *UeContext) RadCap() string {
	//TODO
	return ""
}

func (ue *UeContext) SecCtx() *nas.SecCtx {
	return ue.secctx
}

func (ue *UeContext) AmfContext() AmfContext {
	return ue.amf
}

func (ue *UeContext) SetLocation(loc *models.UserLocation) {
	if loc != nil {
		ue.loc = *loc
	}
}

func (ue *UeContext) AmPolUri() string {
	return fmt.Sprintf("%s:/namf-callback/v1/am-policy/", ue.amf.Uri())
}
func (ue *UeContext) Supi() string {
	return ue.supi
}
func (ue *UeContext) RatType() models.RatType {
	return ue.rat
}
func (ue *UeContext) Abba() []byte {
	return ue.abba
}
func (ue *UeContext) Pei() string {
	return ue.pei
}
func (ue *UeContext) PlmnId() models.PlmnId {
	return ue.plmnid
}

func (ue *UeContext) Gpsi() string {
	return ue.gpsi
}

func (ue *UeContext) Suci() string {
	return ue.suci
}
func (ue *UeContext) Tmsi5gs() string {
	return ue.tmsi5gs
}
func (ue *UeContext) Guti() string {
	return ue.guti
}

func (ue *UeContext) LadnInfo() []Ladn {
	return ue.ladninfo
}
func (ue *UeContext) NgKsi() *models.NgKsi {
	return ue.ngksi
}

func (ue *UeContext) SliceSubChanged() bool {
	return ue.slicesubchanged
}

func (ue *UeContext) SetSliceSubChanged(f bool) {
	ue.slicesubchanged = f
}

func (ue *UeContext) SecCap() *nasType.UESecurityCapability {
	return ue.seccap
}

func (ue *UeContext) SetPlmnId(buf []byte) (err error) {
	if id, err := models.Bytes2PlmnId(buf); err == nil {
		ue.plmnid = *id
	}
	return
}

func (ue *UeContext) AmPol() *models.PolicyAssociation {
	return ue.ampol
}
func (ue *UeContext) SetAmPol(ampol *models.PolicyAssociation) {
	ue.ampol = ampol
}

func (ue *UeContext) AmData() *models.AccessAndMobilitySubscriptionData {
	return ue.amdat
}
func (ue *UeContext) SmfSelData() *models.SmfSelectionSubscriptionData {
	return ue.smfsel
}

/*
func (ue *UeContext) RanState(access models.AccessType) RanState {
	return ue.ranfaces[access]
}
*/
// check if the Ue has a SUPI or SUCI

func (ue *UeContext) ServingNetwork() string {
	str := common.ServingNetworkName(&ue.plmnid)
	ue.Tracef("Serving network:%s", str)
	return str
}

/*
func (ue *UeContext) IsRegistered(access models.AccessType) bool {
	return ue.ranfaces[access].IsRegistered()
}
*/
// get registration status (for preparing a RegistrationRequest)
func (ue *UeContext) RegStatus(access models.AccessType) (status uint8) {
	status = 0
	if access == models.ACCESSTYPE__3_GPP_ACCESS {
		status |= nasMessage.AccessType3GPP
		if ranstate, ok := ue.ranfaces[models.ACCESSTYPE_NON_3_GPP_ACCESS]; ok {
			if ranstate.IsRegistered() {
				status |= nasMessage.AccessTypeNon3GPP
			}
		}
	} else {
		status |= nasMessage.AccessTypeNon3GPP
		if ranstate, ok := ue.ranfaces[models.ACCESSTYPE__3_GPP_ACCESS]; ok {
			if ranstate.IsRegistered() {
				status |= nasMessage.AccessType3GPP
			}
		}
	}
	return
}

func (ue *UeContext) ResetUeRadioCap() {
	//TODO: needs more investigation
	//ue.UeRadioCapability = ""
	//ue.UeRadioCapabilityForPaging = nil
}

func (ue *UeContext) Clean() {
	/*
		for _, f := range ue.ranfaces {
			f.Clean()
		}
	*/
}

func (ue *UeContext) FindSessionContext(id int32) (smctx SessionContext) {
	return ue.sclist.find(id)
}

func (ue *UeContext) StoreSessionContext(sc SessionContext) {
	ue.Infof("Store SessionContext [sid=%d]", sc.Id())
	ue.sclist.add(sc)
}

/*
func (ue *UeContext) SessionContextList() []SessionContext {
	return ue.sclist.list()
}
*/
// find Dnn value for a Session from subscription data of the Ue
func (ue *UeContext) FindDnn(sc SessionContext) (dnn string) {
	snssai := sc.Snssai()
	//convert ssnssai to string
	id := fmt.Sprintf("%02x%s", snssai.Sst, snssai.Sd)
	if info, ok := ue.smfsel.SubscribedSnssaiInfos[id]; ok {
		for _, dnninfo := range info.DnnInfos {
			if dnninfo.DefaultDnnIndicator {
				dnn = dnninfo.Dnn
				return
			}
		}
	}
	if len(dnn) == 0 {
		dnn = ue.amf.DefaultDnn(sc.Access()) //NOTE: AMF should always has a default DNN
	}

	return
}

func (ue *UeContext) SetAmPolRequest(req *models.PolicyAssociationRequest) {
	plmnid := ue.plmnid //should be Amf' plmnid
	req.NotificationUri = ue.AmPolUri()
	req.Supi = ue.Supi()
	req.Pei = ue.Pei()
	req.Gpsi = ue.Gpsi()
	req.ServingPlmn = models.PlmnIdNid{
		Mcc: plmnid.Mcc,
		Mnc: plmnid.Mnc,
	}
	//Guami:       &amfSelf.ServedGuamiList[0],

	/*
		if amdat := ue.AmData(); amdat != nil {
			req.Rfsp = amdat.RfspIndex
		}
	*/
}

func (ue *UeContext) Pcf() sbi.ConsumerClient {
	var err error
	if ue.pcfcli == nil {
		sid := common.PcfServiceName(&ue.plmnid)
		if ue.pcfcli, err = mesh.Consumer(meshmodels.ServiceName(sid), nil, false); err != nil {
			ue.Errorf("Fail to create a consumer to %s: %s", sid, err.Error())
		}
	}

	return ue.pcfcli
}
func (ue *UeContext) Ausf() sbi.ConsumerClient {
	var err error
	if ue.ausfcli == nil {
		sid := common.AusfServiceName(&ue.plmnid)
		if ue.ausfcli, err = mesh.Consumer(meshmodels.ServiceName(sid), nil, false); err != nil {
			ue.Errorf("Fail to create a consumer to %s: %s", sid, err.Error())
		}
	}
	return ue.ausfcli
}

func (ue *UeContext) Udm() sbi.ConsumerClient {
	var err error
	if ue.udmcli == nil {
		sid := common.UdmServiceName(&ue.plmnid)
		if ue.udmcli, err = mesh.Consumer(meshmodels.ServiceName(sid), nil, false); err != nil {
			ue.Errorf("Fail to create a consumer to %s: %s", sid, err.Error())
		}
	}

	return ue.udmcli
}

func (ue *UeContext) AttachRanUe(ranstate RanUe) (err error) {
	//TODO
	access := ranstate.Access()
	if old, ok := ue.ranfaces[access]; ok {
		old.Detach()
		delete(ue.ranfaces, access)
	}
	ue.ranfaces[access] = ranstate
	return
}

func (ue *UeContext) Kill() {
	ue.Infof("Kill %s", ue.suci)
	ue.worker.Terminate()
	if ue.onkill != nil {
		ue.onkill()
	}
}

/*
	func (ue *UeContext) SendAuthenticationPost() (*models.UEAuthenticationCtx, error) {
		return nil, nil
	}
*/
func (ue *UeContext) SetLastTai(nastai *nasType.LastVisitedRegisteredTAI) {
	if nastai == nil {
		return
	}
	plmnid := nasConvert.PlmnIDToString(nastai.Octet[1:4])
	nastac := nastai.GetTAC()
	tac := hex.EncodeToString(nastac[:])
	ue.lasttai = &models.Tai{
		PlmnId: models.PlmnId{
			Mcc: plmnid[:3],
			Mnc: plmnid[3:],
		},
		Tac: tac,
	}

}
