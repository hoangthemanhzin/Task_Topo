package ranuecontext

import (
	"etrib5gc/nas"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n2models"
	rann2 "etrib5gc/sbi/pran/nas"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

func (ranue *RanUe) sendNas(pdu []byte) (err error) {
	msg := n2models.NasDlMsg{
		NasPdu: pdu,
	}
	return rann2.NasDl(ranue.rancli, ranue.ranUeId, msg)
}

func (ranue *RanUe) sendDlN1SmError(sid int32, pdu []byte) (err error) {
	defer ranue.logSendingReport("Downlink N1Sm Error", err)
	err = ranue.sendDlN1Sm(sid, pdu, nasMessage.Cause5GMMPayloadWasNotForwarded)
	return
}
func (ranue *RanUe) sendDlN1Sm(sid int32, n1sm []byte, cause uint8) (err error) {
	var pdu []byte
	ctype := nasMessage.PayloadContainerTypeN1SMInfo
	if pdu, err = nas.BuildDLNASTransport(ranue, ctype, n1sm, sid, cause, nil, 0); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) sendNotification(access models.AccessType) (err error) {
	defer ranue.logSendingReport("NAS Notification", err)
	var pdu []byte
	if pdu, err = nas.BuildNotification(ranue, access); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) sendRegistrationReject(cause uint8) (err error) {
	defer ranue.logSendingReport("NAS RegistrationReject", err)
	var pdu []byte

	eap := ""
	if authctx := ranue.regctx.AuthCtx(); authctx != nil {
		eap = authctx.Eap
	}
	if pdu, err = nas.BuildRegistrationReject(ranue, cause, eap, ranue.fillRegistrationReject); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) fillRegistrationReject(msg *nasMessage.RegistrationReject) (err error) {

	msg.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationRejectT3502ValueType)
	msg.T3502Value.SetLen(1)
	msg.T3502Value.SetGPRSTimer2Value(ranue.amf.T3502())
	return
}
func (ranue *RanUe) sendServiceReject(cause uint8, statuslist *[16]bool) (err error) {
	var pdu []byte
	if pdu, err = nas.BuildServiceReject(ranue, statuslist, cause); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}
func (ranue *RanUe) SendIdentityRequest(idtype uint8) (err error) {
	defer ranue.logSendingReport("NAS IdentityRequest", err)
	var pdu []byte
	if pdu, err = nas.BuildIdentityRequest(ranue, idtype); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) SendAuthenticationRequest(fn func(*nasMessage.AuthenticationRequest) error) (err error) {
	defer ranue.logSendingReport("NAS AuthenticationRequest", err)
	var pdu []byte
	if pdu, err = nas.BuildAuthenticationRequest(ranue, fn); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) SendAuthenticationReject(eap string) (err error) {
	defer ranue.logSendingReport("NAS AuthenticationReject", err)
	var pdu []byte
	if pdu, err = nas.BuildAuthenticationReject(ranue, eap); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) SendSecurityModeCommand(fn func(*nasMessage.SecurityModeCommand) error) (err error) {
	defer ranue.logSendingReport("NAS SecurityModeCommand", err)
	var pdu []byte
	if pdu, err = nas.BuildSecurityModeCommand(ranue, fn); err == nil {
		err = ranue.sendNas(pdu)
	}
	return
}

func (ranue *RanUe) fillAuthenticationResult(msg *nasMessage.AuthenticationResult, success bool) {
	uectx := ranue.ue
	if success {
		abba := uectx.Abba()
		msg.ABBA = nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
		msg.ABBA.SetLen(uint8(len(abba)))
		msg.ABBA.SetABBAContents(abba)
		ngksi := uectx.NgKsi()
		msg.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(*ngksi)

	}
	return
}
func (ranue *RanUe) fillRegistrationAccept(msg *nasMessage.RegistrationAccept) (err error) {
	uectx := ranue.ue
	amf := ranue.amf
	//set registration result
	msg.RegistrationResult5GS.SetLen(1)
	msg.RegistrationResult5GS.SetRegistrationResultValue5GS(uectx.RegStatus(ranue.Access()))

	//set guti
	if guti := uectx.Guti(); len(guti) > 0 {
		gutiNas := nasConvert.GutiToNas(guti) //TODO: there can be an error if guti is malformed
		msg.GUTI5G = &gutiNas
		msg.GUTI5G.SetIei(nasMessage.RegistrationAcceptGUTI5GType)

	}
	/*
		// TODO: set smsAllowed value of RegistrationResult5GS if need


		amfSelf := context.AMF_Self()
		if len(amfSelf.PlmnSupportList) > 1 {
			msg.EquivalentPlmns = nasType.NewEquivalentPlmns(nasMessage.RegistrationAcceptEquivalentPlmnsType)
			var buf []uint8
			for _, plmnSupportItem := range amfSelf.PlmnSupportList {
				buf = append(buf, nasConvert.PlmnIDToNas(*plmnSupportItem.PlmnId)...)
			}
			msg.EquivalentPlmns.SetLen(uint8(len(buf)))
			copy(msg.EquivalentPlmns.Octet[:], buf)
		}
	*/
	if len(ranue.regarea) > 0 {
		msg.TAIList = nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
		taiListNas := nasConvert.TaiListToNas(ranue.regarea)
		msg.TAIList.SetLen(uint8(len(taiListNas)))
		msg.TAIList.SetPartialTrackingAreaIdentityList(taiListNas)
	}
	if len(ranue.allowednssai) > 0 {
		msg.AllowedNSSAI = nasType.NewAllowedNSSAI(nasMessage.RegistrationAcceptAllowedNSSAIType)
		var buf []uint8
		for _, nssai := range ranue.allowednssai {
			buf = append(buf, nasConvert.SnssaiToNas(nssai.AllowedSnssai)...)
		}
		msg.AllowedNSSAI.SetLen(uint8(len(buf)))
		msg.AllowedNSSAI.SetSNSSAIValue(buf)
	}
	if err = amf.FillRegistrationAccept(msg, ranue); err != nil {
		return
	}
	/*
		if ue.NetworkSliceInfo != nil {
			if len(ue.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 || len(ue.NetworkSliceInfo.RejectedNssaiInTa) != 0 {
				rejectedNssaiNas := nasConvert.RejectedNssaiToNas(
					ue.NetworkSliceInfo.RejectedNssaiInPlmn, ue.NetworkSliceInfo.RejectedNssaiInTa)
				msg.RejectedNSSAI = &rejectedNssaiNas
				msg.RejectedNSSAI.SetIei(nasMessage.RegistrationAcceptRejectedNSSAIType)
			}
		}

		if includeConfiguredNssaiCheck(ue) {
			msg.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(nasMessage.RegistrationAcceptConfiguredNSSAIType)
			var buf []uint8
			for _, snssai := range ue.ConfiguredNssai {
				buf = append(buf, nasConvert.SnssaiToNas(*snssai.ConfiguredSnssai)...)
			}
			msg.ConfiguredNSSAI.SetLen(uint8(len(buf)))
			msg.ConfiguredNSSAI.SetSNSSAIValue(buf)
		}
	*/
	if ladninfo := uectx.LadnInfo(); len(ladninfo) > 0 {
		msg.LADNInformation = nasType.NewLADNInformation(nasMessage.RegistrationAcceptLADNInformationType)
		buf := make([]uint8, 0)
		for _, ladn := range ladninfo {
			ladnNas := nasConvert.LadnToNas(ladn.Dnn, ladn.TaiList)
			buf = append(buf, ladnNas...)
		}
		msg.LADNInformation.SetLen(uint16(len(buf)))
		msg.LADNInformation.SetLADND(buf)
	}

	if uectx.SliceSubChanged() {
		msg.NetworkSlicingIndication =
			nasType.NewNetworkSlicingIndication(nasMessage.RegistrationAcceptNetworkSlicingIndicationType)
		msg.NetworkSlicingIndication.SetNSSCI(1)
		msg.NetworkSlicingIndication.SetDCNI(0)
		ranue.ue.SetSliceSubChanged(false) // reset the value
	}

	/*
		//NOTE: PolicyAssociation does not have ServAreRes anymore
			if ampol := ranue.ue.AmPol(); ranue.access == models.ACCESSTYPE__3_GPP_ACCESS && ampol != nil &&
				ampol.ServAreaRes != nil {
				msg.ServiceAreaList = nasType.NewServiceAreaList(nasMessage.RegistrationAcceptServiceAreaListType)
				partialServiceAreaList := nasConvert.PartialServiceAreaListToNas(ranue.ue.ModelPlmnId(), *ampol.ServAreaRes)
				msg.ServiceAreaList.SetLen(uint8(len(partialServiceAreaList)))
				msg.ServiceAreaList.SetPartialServiceAreaList(partialServiceAreaList)
			}
	*/
	//set ue specific drx parameter
	if ranue.drx != nasMessage.DRXValueNotSpecified {
		msg.NegotiatedDRXParameters =
			nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
		msg.NegotiatedDRXParameters.SetLen(1)
		msg.NegotiatedDRXParameters.SetDRXValue(ranue.drx)
	}

	return
}

func (ue *RanUe) fillConfigurationUpdateCommand(req *nasMessage.ConfigurationUpdateCommand, access models.AccessType) (err error) {
	/*
		if ue.ConfigurationUpdateIndication.Octet != 0 {
			configurationUpdateCommand.ConfigurationUpdateIndication =
				nasType.NewConfigurationUpdateIndication(nasMessage.ConfigurationUpdateCommandConfigurationUpdateIndicationType)
			configurationUpdateCommand.ConfigurationUpdateIndication = &ue.ConfigurationUpdateIndication
		}

			if ue.Guti != "" {
			gutiNas := nasConvert.GutiToNas(ue.Guti)
			configurationUpdateCommand.GUTI5G = &gutiNas
			configurationUpdateCommand.GUTI5G.SetIei(nasMessage.ConfigurationUpdateCommandGUTI5GType)
		}

		if len(ue.RegistrationArea[anType]) > 0 {
			configurationUpdateCommand.TAIList = nasType.NewTAIList(nasMessage.ConfigurationUpdateCommandTAIListType)
			taiListNas := nasConvert.TaiListToNas(ue.RegistrationArea[anType])
			configurationUpdateCommand.TAIList.SetLen(uint8(len(taiListNas)))
			configurationUpdateCommand.TAIList.SetPartialTrackingAreaIdentityList(taiListNas)
		}

		if len(ue.AllowedNssai[anType]) > 0 {
			configurationUpdateCommand.AllowedNSSAI =
				nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
			var buf []uint8
			for _, allowedSnssai := range ue.AllowedNssai[anType] {
				buf = append(buf, nasConvert.SnssaiToNas(*allowedSnssai.AllowedSnssai)...)
			}
			configurationUpdateCommand.AllowedNSSAI.SetLen(uint8(len(buf)))
			configurationUpdateCommand.AllowedNSSAI.SetSNSSAIValue(buf)
		}

		if len(ue.ConfiguredNssai) > 0 {
			configurationUpdateCommand.ConfiguredNSSAI =
				nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
			var buf []uint8
			for _, snssai := range ue.ConfiguredNssai {
				buf = append(buf, nasConvert.SnssaiToNas(*snssai.ConfiguredSnssai)...)
			}
			configurationUpdateCommand.ConfiguredNSSAI.SetLen(uint8(len(buf)))
			configurationUpdateCommand.ConfiguredNSSAI.SetSNSSAIValue(buf)
		}

		if ue.NetworkSliceInfo != nil {
			if len(ue.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 || len(ue.NetworkSliceInfo.RejectedNssaiInTa) != 0 {
				rejectedNssaiNas := nasConvert.RejectedNssaiToNas(
					ue.NetworkSliceInfo.RejectedNssaiInPlmn, ue.NetworkSliceInfo.RejectedNssaiInTa)
				configurationUpdateCommand.RejectedNSSAI = &rejectedNssaiNas
				configurationUpdateCommand.RejectedNSSAI.SetIei(nasMessage.ConfigurationUpdateCommandRejectedNSSAIType)
			}
		}

		// TODO: UniversalTimeAndLocalTimeZone
		if anType == models.AccessType__3_GPP_ACCESS && ue.AmPolicyAssociation != nil &&
			ue.AmPolicyAssociation.ServAreaRes != nil {
			configurationUpdateCommand.ServiceAreaList =
				nasType.NewServiceAreaList(nasMessage.ConfigurationUpdateCommandServiceAreaListType)
			partialServiceAreaList := nasConvert.PartialServiceAreaListToNas(ue.PlmnId, *ue.AmPolicyAssociation.ServAreaRes)
			configurationUpdateCommand.ServiceAreaList.SetLen(uint8(len(partialServiceAreaList)))
			configurationUpdateCommand.ServiceAreaList.SetPartialServiceAreaList(partialServiceAreaList)
		}

		amfSelf := context.AMF_Self()
		if amfSelf.NetworkName.Full != "" {
			fullNetworkName := nasConvert.FullNetworkNameToNas(amfSelf.NetworkName.Full)
			configurationUpdateCommand.FullNameForNetwork = &fullNetworkName
			configurationUpdateCommand.FullNameForNetwork.SetIei(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
		}

		if amfSelf.NetworkName.Short != "" {
			shortNetworkName := nasConvert.ShortNetworkNameToNas(amfSelf.NetworkName.Short)
			configurationUpdateCommand.ShortNameForNetwork = &shortNetworkName
			configurationUpdateCommand.ShortNameForNetwork.SetIei(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
		}

		if ue.TimeZone != "" {
			localTimeZone := nasConvert.LocalTimeZoneToNas(ue.TimeZone)
			localTimeZone.SetIei(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
			configurationUpdateCommand.LocalTimeZone =
				nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
			configurationUpdateCommand.LocalTimeZone = &localTimeZone
		}

		if ue.TimeZone != "" {
			daylightSavingTime := nasConvert.DaylightSavingTimeToNas(ue.TimeZone)
			daylightSavingTime.SetIei(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
			configurationUpdateCommand.NetworkDaylightSavingTime =
				nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
			configurationUpdateCommand.NetworkDaylightSavingTime = &daylightSavingTime
		}

		if len(ue.LadnInfo) > 0 {
			configurationUpdateCommand.LADNInformation =
				nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
			var buf []uint8
			for _, ladn := range ue.LadnInfo {
				ladnNas := nasConvert.LadnToNas(ladn.Dnn, ladn.TaiLists)
				buf = append(buf, ladnNas...)
			}
			configurationUpdateCommand.LADNInformation.SetLen(uint16(len(buf)))
			configurationUpdateCommand.LADNInformation.SetLADND(buf)
		}
	*/
	return
}
