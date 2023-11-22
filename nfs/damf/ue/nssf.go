package ue

import (
	"etrib5gc/sbi/models"
	"fmt"
)

func (uectx *UeContext) findAmf() (err error) {
	if len(uectx.aslices) == 0 {
		uectx.Trace("Get subscribed slices for UE")
		if err = uectx.getSubscribedSlices(); err != nil {
			return
		}

		uectx.Trace("Determine allowed slices for UE")
		if err = uectx.getAllowedSlices(); err != nil {
			return
		}
	}

	uectx.Infof("Subscribed Snssai = %s", uectx.aslices[0].String())
	uectx.amfid = uectx.ctx.FindAmfId(uectx.aslices)
	if len(uectx.amfid) == 0 {
		err = fmt.Errorf("Amf not found for slice %s", uectx.aslices[0].String())
		return
	} else {
		uectx.Infof("AMF[amfid=%s] found for slice %s", uectx.amfid, uectx.aslices[0].String())
	}
	/*
		sid := common.AmfServiceName(&uectx.plmnid, amfid)
		if uectx.amfcli, err = uectx.ctx.Agent().Sender(types.ServiceId(sid), nil); err == nil {
			log.Infof("A consumer to AMF[%s] is created for Ue[SUPI=%s]", sid, uectx.supi)
		} else {
			log.Errorf("No AMF to handle slice", uectx.slices[0].String())
		}
	*/
	return
}

func (uectx *UeContext) getSubscribedSlices() (err error) {
	//TODO: get the information from UDM
	uectx.sslices = []models.Snssai{
		models.Snssai{
			Sst: 1,
			Sd:  "010203",
		},
	}
	return
}

func (uectx *UeContext) getAllowedSlices() (err error) {
	//TODO: implement the logic for selecting allowed slices
	uectx.aslices = make([]models.Snssai, len(uectx.sslices))
	for i, s := range uectx.sslices {
		uectx.aslices[i] = s
	}
	return
}
