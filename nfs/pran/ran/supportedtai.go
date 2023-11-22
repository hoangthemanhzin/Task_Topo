package ran

import (
	"encoding/hex"
	"etrib5gc/common"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/ngapConvert"

	"github.com/free5gc/ngap/ngapType"
)

type SupportedTai struct {
	Tac   string
	Plmns []common.PlmnItem
}

func convertSupportedTai(tailist *ngapType.SupportedTAList) (results []SupportedTai) {
	results = make([]SupportedTai, len(tailist.List))
	for i, item := range tailist.List {
		newitem := &results[i]
		newitem.Tac = hex.EncodeToString(item.TAC.Value)
		log.Tracef("Supported tac=%s", newitem.Tac)
		newitem.Plmns = make([]common.PlmnItem, len(item.BroadcastPLMNList.List))
		for j, plmn := range item.BroadcastPLMNList.List {
			newitem.Plmns[j] = common.PlmnItem{
				PlmnId: models.PlmnId(ngapConvert.PlmnIdToModels(plmn.PLMNIdentity)),
				Slices: make([]models.Snssai, len(plmn.TAISliceSupportList.List)),
			}
			log.Tracef("Plmnid: Mcc=%s, Mnc=%s", newitem.Plmns[j].PlmnId.Mcc, newitem.Plmns[j].PlmnId.Mnc)
			for k, snssai := range plmn.TAISliceSupportList.List {
				newitem.Plmns[j].Slices[k] = ngapConvert.SNssaiToModels(snssai.SNSSAI)
				log.Tracef("Supported slice: Sst=%d, Sd=%s", newitem.Plmns[j].Slices[k].Sst, newitem.Plmns[j].Slices[k].Sd)
			}
		}
	}
	return
}
