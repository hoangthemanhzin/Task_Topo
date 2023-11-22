package context

import (
	"etrib5gc/nfs/udm/config"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/udr/dr"
	"etrib5gc/util/suci"
	"fmt"
)

type UdmContext struct {
	config *config.UdmConfig
	uelist UeList
	udrcli sbi.ConsumerClient
	plmnid models.PlmnId
}

func New(cfg *config.UdmConfig) *UdmContext {
	_initLog()
	ret := &UdmContext{
		config: cfg,
		uelist: newUeList(),
		plmnid: models.PlmnId(cfg.PlmnId),
	}
	return ret
}

func (ctx *UdmContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}

func (ctx *UdmContext) GetUeContext(supiOrSuci string) (ue *UeContext, err error) {
	var supi string
	if supi, err = suci.RecoverSupi(supiOrSuci, ctx.config.Profiles); err != nil {
		err = fmt.Errorf("Recovering SUPI failed: %s", err.Error())
		return
	}
	log.Tracef("Recover SUPI[%s] from SUCI[%s]", supi, supiOrSuci)
	if ue = ctx.uelist.find(supi); ue == nil {
		log.Tracef("No UeContext for %s, create new one from subscription data", supiOrSuci)
		ue, err = ctx.createUeContext(supi)
	} else {
		ue.Info("UeContext found")
	}
	return
}

func (ctx *UdmContext) createUeContext(supi string) (ue *UeContext, err error) {
	//talk to UDR to build a new UeContext
	var sub *models.AuthenticationSubscription
	if sub, err = dr.GetUeSub(ctx.udrcli, supi); err != nil {
		err = fmt.Errorf(" Get subscription data failed: %s", err.Error())
		return
	}
	if ue, err = newUeContext(supi, sub); err != nil {
		err = fmt.Errorf("Create UeContext failed: %s", err.Error())
		return
	}
	ue.Info("UeContext created")
	ctx.uelist.add(ue)
	return
}
