package context //non3gpp ran state

import (
	"etrib5gc/common"
	"etrib5gc/mesh"
	"etrib5gc/nfs/amf/config"
	"etrib5gc/nfs/amf/ranuecontext"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/nasConvert"
	"etrib5gc/util/idgen"
	"fmt"
	"math"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

type AmfContext struct {
	plmnid models.PlmnId
	amfid  common.AmfId
	//plmnlist map[models.PlmnId][]models.Snssai
	//tailist  []models.Tai
	secalgs *common.NasSecAlgList

	uepool    UePool
	ranuepool RanUePool
	tmsigen   idgen.IdGenerator
	ueidgen   idgen.IdGenerator

	slices []models.Snssai

	t3550 int
	t3522 int
	t3565 int
	t3513 int
	t3502 uint8
	t3512 uint8
}

func NewAmfContext(c *config.AmfConfig) (amf *AmfContext, err error) {
	_initLog()
	amf = &AmfContext{
		slices:    c.Slices,
		uepool:    newUePool(),
		ranuepool: newRanUePool(),
		plmnid:    models.PlmnId(c.PlmnId),
		secalgs:   c.Algs,
		//	plmnlist: make(map[models.PlmnId][]models.Snssai),
		tmsigen: idgen.NewIdGenerator(1, math.MaxUint32),
		ueidgen: idgen.NewIdGenerator(0, math.MaxInt64),
		t3550:   3000, //milliseconds
		t3522:   3000,
		t3565:   3000,
		t3513:   3000,
		t3502:   128,
		t3512:   128,
	}
	if err = amf.amfid.SetHex(c.AmfId); err != nil {
		log.Errorf("Invalid AmfId: %s", err.Error())
		return
	}
	/*
		tacmap := make(map[string]bool)
		for _, tac := range c.Tacs {
			tacmap[tac] = true
		}
		for tac, _ := range tacmap {
			amf.tailist = append(amf.tailist, models.Tai{
				PlmnId: models.PlmnId(c.PlmnId),
				Tac:    tac,
			})
		}
		for _, item := range c.PlmnList {
			amf.plmnlist[models.PlmnId(item.PlmnId)] = item.Slices
		}
	*/
	if amf.secalgs == nil {
		log.Trace("Amf is using the default algorithm settings")
		amf.secalgs = &common.DefaultNasSecAlgs
	}
	return
}

func (ctx *AmfContext) Callback() (callback models.Callback) {
	//	callback.ServiceId = common.AmfServiceName(&ctx.plmnid, ctx.amfid.String())
	//	callback.InstanceId = string(mesh.AgentId())
	return models.Callback(mesh.CallbackAddress())
}

func (ctx *AmfContext) Uri() string {
	return mesh.CallbackAddress()
}

func (ctx *AmfContext) SecAlgs() *common.NasSecAlgList {
	return ctx.secalgs
}

/*
	func (ctx *AmfContext) newGuti() string {
		var tmsi common.Tmsi
		tmsi.Set(ctx.tmsigen.NewId())
		guti := common.Guti{
			Guami: ctx.guami,
			Tmsi:  tmsi,
		}
		return guti.String()
	}
*/
func (ctx *AmfContext) HasGuami(guami *models.Guami) bool {
	//TODO: to be implemented
	return true
}

func (ctx *AmfContext) Clean() {
	//TODO
	//	ctx.uepool.clean()
}

func (ctx *AmfContext) FillRegistrationAccept(msg *nasMessage.RegistrationAccept, ranstate *ranuecontext.RanUe) (err error) {
	// 5gs network feature support
	if ctx.Get5gsNwFeatSuppEnable() {
		msg.NetworkFeatureSupport5GS =
			nasType.NewNetworkFeatureSupport5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
		msg.NetworkFeatureSupport5GS.SetLen(2)
		if ranstate.Access() == models.ACCESSTYPE__3_GPP_ACCESS {
			msg.SetIMSVoPS3GPP(ctx.Get5gsNwFeatSuppImsVoPS())
		} else {
			msg.SetIMSVoPSN3GPP(ctx.Get5gsNwFeatSuppImsVoPS())
		}
		msg.SetEMC(ctx.Get5gsNwFeatSuppEmc())
		msg.SetEMF(ctx.Get5gsNwFeatSuppEmf())
		msg.SetIWKN26(ctx.Get5gsNwFeatSuppIwkN26())
		msg.SetMPSI(ctx.Get5gsNwFeatSuppMpsi())
		msg.SetEMCN(ctx.Get5gsNwFeatSuppEmcN3())
		msg.SetMCSI(ctx.Get5gsNwFeatSuppMcsi())
	}
	//set timer values
	if ranstate.Access() == models.ACCESSTYPE__3_GPP_ACCESS {
		msg.T3512Value = nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
		msg.T3512Value.SetLen(1)
		msg.T3512Value.Octet = ctx.t3512
	}

	if ranstate.Access() == models.ACCESSTYPE_NON_3_GPP_ACCESS {
		msg.Non3GppDeregistrationTimerValue =
			nasType.NewNon3GppDeregistrationTimerValue(nasMessage.RegistrationAcceptNon3GppDeregistrationTimerValueType)
		msg.Non3GppDeregistrationTimerValue.SetLen(1)
		msg.Non3GppDeregistrationTimerValue.SetGPRSTimer2Value(ctx.GetNon3GppDeregistrationTimerValue())
	}

	msg.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationAcceptT3502ValueType)
	msg.T3502Value.SetLen(1)
	msg.T3502Value.SetGPRSTimer2Value(ctx.t3502)
	return
}

func (ctx *AmfContext) GetNon3GppDeregistrationTimerValue() uint8 {
	//TODO: get from Amf context
	return 200
}

func (ctx *AmfContext) Get5gsNwFeatSuppEnable() bool {
	return true
}

func (ctx *AmfContext) Get5gsNwFeatSuppImsVoPS() uint8 {
	return 1
}
func (ctx *AmfContext) Get5gsNwFeatSuppEmc() uint8 {
	return 1
}
func (ctx *AmfContext) Get5gsNwFeatSuppEmf() uint8 {
	return 1
}
func (ctx *AmfContext) Get5gsNwFeatSuppIwkN26() uint8 {
	return 1
}
func (ctx *AmfContext) Get5gsNwFeatSuppMpsi() uint8 {
	return 1
}
func (ctx *AmfContext) Get5gsNwFeatSuppEmcN3() uint8 {
	return 1
}
func (ctx *AmfContext) Get5gsNwFeatSuppMcsi() uint8 {
	return 1
}

func (ctx *AmfContext) AmfId() string {
	return ctx.amfid.String()
}

func (ctx *AmfContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}

func (ctx *AmfContext) DefaultDnn(access models.AccessType) string {
	//TODO: select a default one from configured list
	return "internet"
}

/*
func (ctx *AmfContext) hasTac(tac string) bool {
	for _, tai := range ctx.tailist {
		if strings.Compare(tai.Tac, tac) == 0 {
			return true
		}
	}
	return false
}
*/

func (ctx *AmfContext) SearchUeContext(msg *libnas.GmmMessage) (uectx *uecontext.UeContext, err error) {
	switch msg.GetMessageType() {
	case libnas.MsgTypeRegistrationRequest:
		uectx, err = ctx.searchUeWithRegistrationRequest(msg.RegistrationRequest)

	case libnas.MsgTypeServiceRequest:
		//TODO: check for errors
		tmsi5gs, _, _ := msg.ServiceRequest.TMSI5GS.Get5GSTMSI()
		if uectx = ctx.FindUeByTmsi(tmsi5gs); uectx == nil {
			err = fmt.Errorf("UEontext not found [tmsi5gs=%s]", tmsi5gs)
		}

	//case libnas.MsgTypeDeregistrationRequest:

	default:
		err = fmt.Errorf("Unexpected Nas message")
	}
	return
}

// try to extract ue identity then find its context, otherwise create a new one
func (ctx *AmfContext) searchUeWithRegistrationRequest(msg *nasMessage.RegistrationRequest) (uectx *uecontext.UeContext, err error) {

	content := msg.MobileIdentity5GS.GetMobileIdentity5GSContents()
	idtype := nasConvert.GetTypeOfIdentity(content[0])
	newue := true
	switch idtype {
	case nasMessage.MobileIdentity5GSTypeNoIdentity:
		//create an empty context
		log.Info("RegistrationRequest without UE identity")
		uectx = uecontext.NewUeContext(ctx)

	case nasMessage.MobileIdentity5GSTypeSuci:
		//TODO: reimplement identity conversion
		suci, _ := nasConvert.SuciToString(content)
		if uectx = ctx.FindUeBySuci(suci); uectx == nil {
			log.Infof("UeContext not found [suci=%s], create UeContext", suci)
			uectx = uecontext.NewUeContextWithId(ctx, suci, uecontext.UE_ID_TYPE_SUCI)
			if err = uectx.SetPlmnId(content[1:4]); err != nil {
				err = common.WrapError("Decode PlmnId failed", err)
			}
		} else {
			newue = false
			log.Infof("UeContext found [suci=%s]", suci)
		}
		/*
			//NOTE: not expect to receive a Guti or Tmsi because in etrib5gc we
			//enforce that this is the right AMF to handle the UE, no AMF
			//relocation is needed
				case nasMessage.MobileIdentity5GSType5gGuti:
					log.Info("RegistrationRequest with a Guti")
					guami, guti := nasConvert.GutiToString(content)
					if ctx.HasGuami(&guami) {
						//this guami is for this AMF
						if ue = ctx.FindUe(guti, context.UE_ID_TYPE_TMSI5GS); ue == nil {
							err = fmt.Errorf("No uecontext is found for guti=%s", guti)
						}
					} else {
						//handover (not the AMF's guami)
						//TODO: fetch UeContext from old AMF
						log.Warnf("UE context feching has not been implemented")
						err = fmt.Errorf("Changing amf is not supported")
					}

				case nasMessage.MobileIdentity5GSType5gSTmsi:
					log.Info("RegistrationRequest with a tmsi")
					tmsi5gs := hex.EncodeToString(content[1:])
					if ue = ctx.FindUe(tmsi5gs, context.UE_ID_TYPE_TMSI5GS); ue == nil {
						err = fmt.Errorf("No ue context is found for tmsi5gs=%s", tmsi5gs)
					}
		*/
	case nasMessage.MobileIdentity5GSTypeImei:
		log.Info("RegistrationRequest with an IMEI")
		imei := nasConvert.PeiToString(content)

		uectx = uecontext.NewUeContextWithId(ctx, imei, uecontext.UE_ID_TYPE_PEI)

	case nasMessage.MobileIdentity5GSTypeImeisv:
		log.Info("RegistrationRequest with an IMEISV")
		imeisv := nasConvert.PeiToString(content)

		uectx = uecontext.NewUeContextWithId(ctx, imeisv, uecontext.UE_ID_TYPE_PEI)
	default:
		//create an empty context anyway
		log.Info("RegistrationRequest with unknown UE identity")
		uectx = uecontext.NewUeContext(ctx)
	}
	if newue {
		//add the UeContext to the pool
		ctx.uepool.add(uectx)
	}
	return
}

func (ctx *AmfContext) T3550() int {
	return ctx.t3550
}
func (ctx *AmfContext) T3513() int {
	return ctx.t3513
}
func (ctx *AmfContext) T3565() int {
	return ctx.t3565
}
func (ctx *AmfContext) T3522() int {
	return ctx.t3522
}

func (ctx *AmfContext) T3502() uint8 {
	return ctx.t3502
}

func (ctx *AmfContext) T3512() uint8 {
	return ctx.t3512
}

func (ctx *AmfContext) AllocateAmfUeId() int64 {
	return int64(ctx.ueidgen.Allocate())
}

func (ctx *AmfContext) UpdateUeId(id string, idtype uint8, uectx *uecontext.UeContext) {
	ctx.uepool.update(id, idtype, uectx)
}

func (ctx *AmfContext) FindUeBySupi(supi string) (uectx *uecontext.UeContext) {
	return ctx.uepool.findBySupi(supi)
}

func (ctx *AmfContext) FindUeBySuci(suci string) (uectx *uecontext.UeContext) {
	return ctx.uepool.findBySuci(suci)
}

func (ctx *AmfContext) FindUeByPei(pei string) (uectx *uecontext.UeContext) {
	return ctx.uepool.findByPei(pei)
}

func (ctx *AmfContext) FindUeByTmsi(tmsi5gs string) (uectx *uecontext.UeContext) {
	return ctx.uepool.findByTmsi5gs(tmsi5gs)
}

func (ctx *AmfContext) AddRanUe(ranue *ranuecontext.RanUe) {
	ctx.ranuepool.add(ranue)
}
func (ctx *AmfContext) RemoveRanUe(ranue *ranuecontext.RanUe) {
	ctx.ranuepool.remove(ranue)
}

func (ctx *AmfContext) FindRanUe(ueid int64) *ranuecontext.RanUe {
	return ctx.ranuepool.findById(ueid)
}

func (ctx *AmfContext) FindRanUeByIdAtRan(cuNgapId int64) *ranuecontext.RanUe {
	return ctx.ranuepool.findByIdAtRan(cuNgapId)
}
