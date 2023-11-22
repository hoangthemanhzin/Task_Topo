package nas

import (
	"github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

type SmContext interface {
	Id() uint32
	Pti() uint8
	FillPduSessionModificationCommand(*nasMessage.PDUSessionModificationCommand) error
	FillPduSessionEstablishmentAccept(*nasMessage.PDUSessionEstablishmentAccept) error
}

func BuildPduSessionEstablishmentAccept(sm SmContext) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionEstablishmentAccept)
	m.GsmHeader.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	m.PDUSessionEstablishmentAccept = nasMessage.NewPDUSessionEstablishmentAccept(0x0)
	msg := m.PDUSessionEstablishmentAccept

	msg.SetPDUSessionID(uint8(sm.Id()))
	msg.SetMessageType(nas.MsgTypePDUSessionEstablishmentAccept)
	msg.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	msg.SetPTI(sm.Pti())
	if err = sm.FillPduSessionEstablishmentAccept(msg); err != nil {
		return
	}

	pdu, err = m.PlainNasEncode()
	return
}

func BuildPduSessionEstablishmentReject(sm SmContext, cause uint8) ([]byte, error) {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionEstablishmentReject)
	m.GsmHeader.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	m.PDUSessionEstablishmentReject = nasMessage.NewPDUSessionEstablishmentReject(0x0)
	msg := m.PDUSessionEstablishmentReject

	msg.SetMessageType(nas.MsgTypePDUSessionEstablishmentReject)
	msg.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	msg.SetPTI(sm.Pti())
	msg.SetPDUSessionID(uint8(sm.Id()))
	msg.SetCauseValue(cause)

	return m.PlainNasEncode()
}

func BuildPduSessionReleaseCommand(sm SmContext, cause uint8) ([]byte, error) {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionReleaseCommand)
	m.GsmHeader.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	m.PDUSessionReleaseCommand = nasMessage.NewPDUSessionReleaseCommand(0x0)
	msg := m.PDUSessionReleaseCommand

	msg.SetMessageType(nas.MsgTypePDUSessionReleaseCommand)
	msg.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	msg.SetPDUSessionID(uint8(sm.Id()))
	msg.SetPTI(sm.Pti())
	msg.SetCauseValue(cause)

	return m.PlainNasEncode()
}

func BuildPduSessionModificationCommand(sm SmContext) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionModificationCommand)
	m.GsmHeader.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	m.PDUSessionModificationCommand = nasMessage.NewPDUSessionModificationCommand(0x0)
	msg := m.PDUSessionModificationCommand

	msg.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	msg.SetPDUSessionID(uint8(sm.Id()))
	msg.SetPTI(sm.Pti())
	msg.SetMessageType(nas.MsgTypePDUSessionModificationCommand)
	if err = sm.FillPduSessionModificationCommand(msg); err != nil {
		return
	}
	pdu, err = m.PlainNasEncode()
	return
}

//nasMessage.Cause5GSMRequestRejectedUnspecified
func BuildPduSessionReleaseReject(sm SmContext, cause uint8) ([]byte, error) {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionReleaseReject)
	m.GsmHeader.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	m.PDUSessionReleaseReject = nasMessage.NewPDUSessionReleaseReject(0x0)
	pDUSessionReleaseReject := m.PDUSessionReleaseReject

	pDUSessionReleaseReject.SetMessageType(nas.MsgTypePDUSessionReleaseReject)
	pDUSessionReleaseReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)

	pDUSessionReleaseReject.SetPDUSessionID(uint8(sm.Id()))

	pDUSessionReleaseReject.SetPTI(sm.Pti())
	// TODO: fix to real value
	pDUSessionReleaseReject.SetCauseValue(cause)

	return m.PlainNasEncode()
}
