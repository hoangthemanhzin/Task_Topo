package sm

import (
	"encoding/hex"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

//implementation of Nas SM message builder
func (smctx *SmContext) FillPduSessionEstablishmentAccept(msg *nasMessage.PDUSessionEstablishmentAccept) (err error) {
	smctx.Tracef("Fill PduSessionEstablishmentAccept")
	srule := smctx.getActivatedSessionRule()
	authdefqos := srule.AuthDefQos

	if smctx.estacceptcause != 0 {
		msg.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionEstablishmentAcceptCause5GSMType)
		msg.Cause5GSM.SetCauseValue(smctx.estacceptcause)
	}
	msg.SetPDUSessionType(smctx.sessiontype)

	msg.SetSSCMode(1)
	msg.SessionAMBR = nasConvert.ModelsToSessionAMBR(&srule.AuthSessAmbr)
	msg.SessionAMBR.SetLen(uint8(len(msg.SessionAMBR.Octet)))

	qrule := QoSRules{
		QoSRule{
			Identifier:    0x01,
			DQR:           0x01,
			OperationCode: OperationCodeCreateNewQoSRule,
			Precedence:    0xff,
			QFI:           uint8(authdefqos.Var5qi),
			PacketFilterList: []PacketFilter{
				{
					Identifier:    0x01,
					Direction:     PacketFilterDirectionBidirectional,
					ComponentType: PacketFilterComponentTypeMatchAll,
				},
			},
		},
	}

	var qrulebytes []byte
	if qrulebytes, err = qrule.MarshalBinary(); err != nil {
		return
	}

	msg.AuthorizedQosRules.SetLen(uint16(len(qrulebytes)))
	msg.AuthorizedQosRules.SetQosRule(qrulebytes)

	ueip := smctx.tunnel.UeIp()
	if len(ueip) > 0 {
		addr, addrlen := toNasIp(ueip, smctx.sessiontype)
		smctx.Tracef("Set ip: %v - %d", addr, addrlen)
		msg.PDUAddress =
			nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
		msg.PDUAddress.SetLen(addrlen)
		msg.PDUAddress.SetPDUSessionTypeValue(smctx.sessiontype)
		msg.PDUAddress.SetPDUAddressInformation(addr)
	}

	msg.AuthorizedQosFlowDescriptions =
		nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	msg.AuthorizedQosFlowDescriptions.SetLen(6)
	msg.SetQoSFlowDescriptions([]uint8{uint8(authdefqos.Var5qi), 0x20, 0x41, 0x01, 0x01, 0x09})

	var sd [3]uint8
	var buf []byte

	if buf, err = hex.DecodeString(smctx.snssai.Sd); err != nil {
		return
	} else {
		copy(sd[:], buf)
	}

	msg.SNSSAI = nasType.NewSNSSAI(nasMessage.ULNASTransportSNSSAIType)
	msg.SNSSAI.SetLen(4)
	msg.SNSSAI.SetSST(uint8(smctx.snssai.Sst))
	msg.SNSSAI.SetSD(sd)

	msg.DNN = nasType.NewDNN(nasMessage.ULNASTransportDNNType)
	msg.DNN.SetDNN(smctx.dnn)
	/*
		if smctx.pco != nil {
			smctx.Info("pco")
			if naspco := smctx.pco.toNas(smctx); len(naspco) > 0 {
				msg.ExtendedProtocolConfigurationOptions =
					nasType.NewExtendedProtocolConfigurationOptions(
						nasMessage.PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType,
					)
				msg.ExtendedProtocolConfigurationOptions.SetLen(uint16(len(naspco)))
				msg.ExtendedProtocolConfigurationOptions.SetExtendedProtocolConfigurationOptionsContents(naspco)
			}
			smctx.Info("end pco")
		}
	*/
	return
}

func (smctx *SmContext) FillPduSessionModificationCommand(msg *nasMessage.PDUSessionModificationCommand) (err error) {
	//TODO: dynamic filling with SmContext content
	//msg.SetQosRule()
	//msg.AuthorizedQosRules.SetLen()
	//msg.SessionAMBR.SetSessionAMBRForDownlink([2]uint8{0x11, 0x11})
	//msg.SessionAMBR.SetSessionAMBRForUplink([2]uint8{0x11, 0x11})
	//msg.SessionAMBR.SetUnitForSessionAMBRForDownlink(10)
	//msg.SessionAMBR.SetUnitForSessionAMBRForUplink(10)
	//msg.SessionAMBR.SetLen(uint8(len(msg.SessionAMBR.Octet)))

	return
}
