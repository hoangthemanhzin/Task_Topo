package ngapConvert

import (
	"etrib5gc/sbi/models"

	"github.com/free5gc/ngap/ngapType"
)

func AllowedNssaiToNgap(allowedNssaiModels []models.AllowedSnssai) (allowedNssaiNgap ngapType.AllowedNSSAI) {
	/*
		//TungTQ: TODO
		for _, allowedSnssai := range allowedNssaiModels {
			item := ngapType.AllowedNSSAIItem{
				SNSSAI: SNssaiToNgap(*allowedSnssai.AllowedSnssai),
			}
			allowedNssaiNgap.List = append(allowedNssaiNgap.List, item)
		}
	*/
	return
}

func AllowedNssaiToModels(allowedNssaiNgap ngapType.AllowedNSSAI) (allowedNssaiModels []models.AllowedSnssai) {
	/*
		//TungTQ: TODO
		for _, item := range allowedNssaiNgap.List {
			snssai := SNssaiToModels(item.SNSSAI)
			allowedSnssai := models.AllowedSnssai{
				AllowedSnssai: &snssai,
			}
			allowedNssaiModels = append(allowedNssaiModels, allowedSnssai)
		}
	*/
	return
}
