package uecontext

func (uectx *UeContext) UpdateWithUdm() (err error) {
	uectx.Warnf("Communicate to UDM not implemented")
	//TODO: make all the requests in a parallel fashion
	/*
		cli := ue.Udm()
		if err = cli.UeCmRegister(access, true); err != nil {
			return
		}
		if ue.amdat, err = cli.GetAmData(); err != nil {
			return
		}
		if ue.smfsel, err = cli.GetSmfSelectData(); err != nil {
			return
		}
		if ue.ueinsmf, err = cli.GetUeContextInSmfData(); err != nil {
			return
		}

		if ue.sdmsubid, err = cli.Subscribe(); err != nil {
			return
		}
	*/
	return
}

// retrieve subscribed nssai from UDM
// the context must have a supi
func (uectx *UeContext) GetSubscribedNssai() (err error) {
	uectx.Warnf("Get subscribed Nssai from UDM not implmented")
	/*
		var nssai *models.Nssai
		//NOTE: perhaps we may need to select a different UDM based on the UE's
		//supi (which is different from the previous one which was selected with
		//the Ue's suci)
		if nssai, err = ue.udmcli.GetNssai(ue.supi, ue.ModelPlmnId()); err != nil {
			return
		}
		//extra
		for _, snssai := range nssai.DefaultSingleNssais {
			ue.subnssai = append(ue.subnssai, models.SubscribedSnssai{
				SubscribedSnssai: &models.Snssai{
					Sst: snssai.Sst,
					Sd:  snssai.Sd,
				},
				DefaultIndication: true,
			})
		}
		for _, snssai := range nssai.SingleNssais {
			ue.subnssai = append(ue.subnssai, models.SubscribedSnssai{
				SubscribedSnssai: &models.Snssai{
					Sst: snssai.Sst,
					Sd:  snssai.Sd,
				},
				DefaultIndication: false,
			})
		}
	*/
	return
}

func (uectx *UeContext) HandleRequestedNssai() (err error) {
	uectx.Warnf("Handle requested Nssai not implmented")
	//TODO: to be implemented
	return
}
