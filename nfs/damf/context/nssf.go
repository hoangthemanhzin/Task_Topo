package context

import "etrib5gc/sbi/models"

func (ctx *DamfContext) FindAmfId(slices []models.Snssai) (amfid string) {
	amfid, _ = ctx.nssf[slices[0]]
	return //"112233"
}
