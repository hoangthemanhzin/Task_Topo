package sessioncontext

import (
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/nasConvert"
	"fmt"

	"github.com/free5gc/nas/nasMessage"
)

var _logfields logctx.Fields = logctx.Fields{
	"mod": "sessioncontext",
}

type RanUe interface {
	Access() models.AccessType
	RanNets() []string
	UeContext() *uecontext.UeContext
}

type SessionContext struct {
	logctx.LogWriter
	id        int32
	ref       string
	ue        *uecontext.UeContext
	rannets   []string
	smfcli    sbi.ConsumerClient
	snssai    models.Snssai
	loc       models.UserLocation
	dnn       string
	access    models.AccessType
	plmnid    models.PlmnId
	smfid     string
	statusUri string
	// slice  string
}

func CreateSessionContext(ranue RanUe, msg *nasMessage.ULNASTransport, preexist bool) (sc *SessionContext, err error) {
	ue := ranue.UeContext()
	access := ranue.Access()
	log := ue.LogWriter.WithFields(_logfields)
	defer func() {
		if err == nil {
			sc.Infof("Create a SMF consumer %s", sc.smfid)
			routematch := meshmodels.NewRouteMatch()
			routematch.Set("usertype", "tester")
			if sc.smfcli, err = mesh.Consumer(meshmodels.ServiceName(sc.smfid), routematch, false); err != nil {
				err = common.WrapError("Create Smf consumer failed", err)
			}
		}
	}()

	sid := int32(msg.PduSessionID2Value.GetPduSessionID2Value())

	if preexist {
		log.Infof("Load smcontext[id=%d] from data", sid)
		if sc = loadSessionContext(sid, ue); sc != nil {
			sc.access = access
		} else {
			err = fmt.Errorf("Load smcontext from data failed")
		}
		return
	}

	var snssai models.Snssai
	// If the S-NSSAI IE is not included and the user's
	// subscription context obtained from UDM.  AMF shall select a
	// default snssai
	if msg.SNSSAI != nil {
		snssai = nasConvert.SnssaiToModels(msg.SNSSAI)
		//TODO: check for conversion error
		log.Infof("Requested session[%d] in slice %s", sid, snssai.String())
	} else {
		if dsnssai := ue.DefaultAllowedSnssai(access); dsnssai != nil {
			snssai = *dsnssai
		} else {
			err = fmt.Errorf("Ue has no allowed Snssai")
			return
		}
	}
	sc = newSessionContext(ue, sid, snssai)
	sc.access = access
	sc.rannets = ranue.RanNets()

	if msg.DNN != nil {
		sc.dnn = msg.DNN.GetDNN()
	} else {
		// if user's subscription context obtained from UDM does
		// not contain the default DNN for the S-NSSAI, the AMF
		// shall use a locally configured DNN as the DNN
		sc.dnn = ue.FindDnn(sc)
	}
	sc.smfid = common.SmfServiceName(&sc.plmnid, &sc.snssai)
	return
}

func newSessionContext(ue *uecontext.UeContext, id int32, snssai models.Snssai) *SessionContext {
	sc := &SessionContext{
		LogWriter: ue.LogWriter.WithFields(logctx.Fields{"pid": id}),
		id:        id,
		ue:        ue,
		plmnid:    ue.PlmnId(),
		snssai:    snssai,
	}
	return sc
}

// create a SessionContext from Ue context data in Smf received from Udm
func loadSessionContext(id int32, ue *uecontext.UeContext) (sc *SessionContext) {
	dat := ue.GetUeContextInSmfData()
	if dat == nil {
		return
	}
	if session, ok := dat.PduSessions[fmt.Sprintf("%d", id)]; ok {
		sc = newSessionContext(ue, id, session.SingleNssai)
		sc.dnn = session.Dnn
		sc.plmnid = session.PlmnId
		sc.smfid = session.SmfInstanceId
	}
	return
}

func (sc *SessionContext) SetRef(ref string) {
	sc.ref = ref
}

func (sc *SessionContext) Ref() string {
	return sc.ref
}
func (sc *SessionContext) SmfId() string {
	return sc.smfid
}

func (sc *SessionContext) Id() int32 {
	return sc.id
}

func (sc *SessionContext) Snssai() models.Snssai {
	return sc.snssai
}
func (sc *SessionContext) Access() models.AccessType {
	return sc.access
}

func (sc *SessionContext) Update(access models.AccessType) {
	sc.access = access
	//sm.loc=sm.ue.loc
}

func (sc *SessionContext) fillSessionContextCreateData(dat *models.SmContextCreateData) {
	//sc.Tracef("Filling smcontext for supi=%s", sc.ue.supi)
	dat.Supi = sc.ue.Supi()
	//	dat.UnauthenticatedSupi 	=	ue.UnauthenticatedSupi
	dat.Pei = sc.ue.Pei()
	//dat.Gpsi =sc.ue.gpsi
	dat.PduSessionId = sc.id
	dat.SNssai = sc.snssai
	dat.Dnn = sc.dnn
	dat.RanNets = sc.rannets
	//dat.ServingNfId =context.NfId
	//dat.Guami =&context.ServedGuamiList[0]
	//dat.ServingNetwork = context.ServedGuamiList[0].PlmnId
	//if requestType!= nil
	//{
	//		dat.RequestType = *requestType
	//	}
	dat.AnType = sc.access
	//if ue.RatType  != ""
	//{
	//		dat.RatType = 	ue.RatType
	//	}
	// TODO: location
	// is
	// used
	// in
	// roaming
	// scenerio

	// if ue.Location != nil {
	// smContextCreateData.UeLocation  = ue.Location
	// }
	//	dat.UeTimeZone 	=/	ue.TimeZone
	dat.SmContextStatusUri = sc.statusUri
}
