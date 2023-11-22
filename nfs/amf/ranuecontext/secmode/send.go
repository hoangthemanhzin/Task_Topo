package secmode

import (
	"encoding/base64"
	"etrib5gc/nas"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

func (proc *SecmodeProc) sendSecmodeCommand() {
	proc.t3560.Start() //start only if sending is success
	proc.ranue.SendSecurityModeCommand(proc.fillSecurityModeCommand)
}

func (proc *SecmodeProc) fillSecurityModeCommand(cmd *nasMessage.SecurityModeCommand) (err error) {
	uectx := proc.ranue.UeContext()
	secctx := proc.ranue.SecCtx()
	cmd.SelectedNASSecurityAlgorithms.SetTypeOfCipheringAlgorithm(secctx.EncAlg())
	cmd.SelectedNASSecurityAlgorithms.SetTypeOfIntegrityProtectionAlgorithm(secctx.IntAlg())

	cmd.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(*uectx.NgKsi())

	seccap := uectx.SecCap()
	cmd.ReplayedUESecurityCapabilities.SetLen(seccap.GetLen())
	cmd.ReplayedUESecurityCapabilities.Buffer = seccap.Buffer

	if len(uectx.Pei()) > 0 {
		cmd.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		cmd.IMEISVRequest.SetIMEISVRequestValue(nasMessage.IMEISVNotRequested)
	} else {
		cmd.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		cmd.IMEISVRequest.SetIMEISVRequestValue(nasMessage.IMEISVRequested)
	}

	cmd.Additional5GSecurityInformation =
		nasType.NewAdditional5GSecurityInformation(nasMessage.SecurityModeCommandAdditional5GSecurityInformationType)
	cmd.Additional5GSecurityInformation.SetLen(1)

	if proc.rinmr {
		cmd.Additional5GSecurityInformation.SetRINMR(1)
	} else {
		cmd.Additional5GSecurityInformation.SetRINMR(0)
	}

	if proc.hdp != nas.HDP_NONE {
		cmd.Additional5GSecurityInformation.SetHDP(1)
	} else {
		cmd.Additional5GSecurityInformation.SetHDP(0)
	}

	if len(proc.eap) > 0 {
		cmd.EAPMessage = nasType.NewEAPMessage(nasMessage.SecurityModeCommandEAPMessageType)
		var eap []byte
		if eap, err = base64.StdEncoding.DecodeString(proc.eap); err != nil {
			return
		}
		cmd.EAPMessage.SetLen(uint16(len(eap)))
		cmd.EAPMessage.SetEAPMessage(eap)

		if proc.success {
			abba := uectx.Abba()
			cmd.ABBA = nasType.NewABBA(nasMessage.SecurityModeCommandABBAType)
			cmd.ABBA.SetLen(uint8(len(abba)))
			cmd.ABBA.SetABBAContents(abba)
		}
	}
	return
}
