package nas

import (
	"encoding/base64"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/nasConvert"

	"github.com/free5gc/nas"
	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

func BuildIdentityRequest(ctx EncodingCtx, idtype uint8) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeIdentityRequest)

	if ctx.NasSecCtx() != nil {
		m.SecurityHeader = nas.SecurityHeader{
			ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
			SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		}
	}
	req := nasMessage.NewIdentityRequest(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.IdentityRequestMessageIdentity.SetMessageType(nas.MsgTypeIdentityRequest)
	req.SpareHalfOctetAndIdentityType.SetTypeOfIdentity(idtype)

	m.GmmMessage.IdentityRequest = req

	pdu, err = Encode(ctx, m)
	return
}

type AuthenticationRequestCompleter func(*nasMessage.AuthenticationRequest) error

func BuildAuthenticationRequest(ctx EncodingCtx, fn AuthenticationRequestCompleter) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeAuthenticationRequest)

	req := nasMessage.NewAuthenticationRequest(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.AuthenticationRequestMessageIdentity.SetMessageType(nas.MsgTypeAuthenticationRequest)

	if fn != nil {
		if err = fn(req); err != nil {
			return
		}
	}
	m.GmmMessage.AuthenticationRequest = req
	pdu, err = m.PlainNasEncode()
	return
}

func BuildAuthenticationReject(ctx EncodingCtx, eap string) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeAuthenticationReject)

	req := nasMessage.NewAuthenticationReject(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.AuthenticationRejectMessageIdentity.SetMessageType(nas.MsgTypeAuthenticationReject)

	if len(eap) > 0 {
		var eapbytes []byte
		if eapbytes, err = base64.StdEncoding.DecodeString(eap); err != nil {
			return
		}
		req.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRejectEAPMessageType)
		req.EAPMessage.SetLen(uint16(len(eapbytes)))
		req.EAPMessage.SetEAPMessage(eapbytes)
	}

	m.GmmMessage.AuthenticationReject = req

	pdu, err = m.PlainNasEncode()
	return
}

type AuthenticationResultCompleter func(*nasMessage.AuthenticationResult, bool) error

func BuildAuthenticationResult(ctx EncodingCtx, success bool, eap string, fn AuthenticationResultCompleter) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeAuthenticationResult)

	req := nasMessage.NewAuthenticationResult(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.AuthenticationResultMessageIdentity.SetMessageType(nas.MsgTypeAuthenticationResult)

	var eapbytes []byte
	if eapbytes, err = base64.StdEncoding.DecodeString(eap); err != nil {
		return
	}
	req.EAPMessage.SetLen(uint16(len(eapbytes)))
	req.EAPMessage.SetEAPMessage(eapbytes)

	if fn != nil {
		if err = fn(req, success); err != nil {
			return
		}
	}

	m.GmmMessage.AuthenticationResult = req

	pdu, err = m.PlainNasEncode()
	return
}

type RegistrationAcceptCompleter func(*nasMessage.RegistrationAccept) error

func BuildRegistrationAccept(ctx EncodingCtx, statuslist *[16]bool, reactlist *[16]bool, errpdulist []uint8, errcauses []uint8, fn RegistrationAcceptCompleter) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeRegistrationAccept)

	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	req := nasMessage.NewRegistrationAccept(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.RegistrationAcceptMessageIdentity.SetMessageType(nas.MsgTypeRegistrationAccept)

	if statuslist != nil {
		req.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.RegistrationAcceptPDUSessionStatusType)
		req.PDUSessionStatus.SetLen(2)
		req.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*statuslist)
	}

	if reactlist != nil {
		req.PDUSessionReactivationResult =
			nasType.NewPDUSessionReactivationResult(nasMessage.RegistrationAcceptPDUSessionReactivationResultType)
		req.PDUSessionReactivationResult.SetLen(2)
		req.PDUSessionReactivationResult.Buffer = nasConvert.PSIToBuf(*reactlist)
	}

	if len(errpdulist) > 0 {
		req.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(
			nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
		buf := nasConvert.PDUSessionReactivationResultErrorCauseToBuf(errpdulist, errcauses)
		req.PDUSessionReactivationResultErrorCause.SetLen(uint16(len(buf)))
		req.PDUSessionReactivationResultErrorCause.Buffer = buf
	}
	if fn != nil {
		if err = fn(req); err != nil {
			return
		}
	}
	m.GmmMessage.RegistrationAccept = req

	pdu, err = Encode(ctx, m)

	return
}

type RegistrationRejectCompleter func(*nasMessage.RegistrationReject) error

func BuildRegistrationReject(ctx EncodingCtx, cause uint8, eap string, fn RegistrationRejectCompleter) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeRegistrationReject)

	req := nasMessage.NewRegistrationReject(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.RegistrationRejectMessageIdentity.SetMessageType(nas.MsgTypeRegistrationReject)
	req.Cause5GMM.SetCauseValue(cause)
	if len(eap) > 0 {
		var eapbytes []byte
		req.EAPMessage = nasType.NewEAPMessage(nasMessage.RegistrationRejectEAPMessageType)
		if eapbytes, err = base64.StdEncoding.DecodeString(eap); err != nil {
			return
		}
		req.EAPMessage.SetLen(uint16(len(eapbytes)))
		req.EAPMessage.SetEAPMessage(eapbytes)
	}
	if fn != nil {
		if err = fn(req); err != nil {
			return
		}
	}
	m.GmmMessage.RegistrationReject = req

	pdu, err = m.PlainNasEncode()
	return
}

/*
	type DLNASTParams struct {
		ContainerType    uint8
		Pdu              []byte
		PduSessionId     int32
		Cause            uint8
		BackoffTimerUint *uint8
		BackoffTimer     uint8
	}
*/
func BuildDLNASTransport(ctx EncodingCtx, ctype uint8, content []byte, pduid int32, cause uint8, timeru *uint8, timerv uint8) (pdu []byte, err error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeDLNASTransport)

	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	req := nasMessage.NewDLNASTransport(0)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SetMessageType(nas.MsgTypeDLNASTransport)
	req.SpareHalfOctetAndPayloadContainerType.SetPayloadContainerType(ctype)
	req.PayloadContainer.SetLen(uint16(len(content)))
	req.PayloadContainer.SetPayloadContainerContents(content)

	if pduid != 0 {
		req.PduSessionID2Value = new(nasType.PduSessionID2Value)
		req.PduSessionID2Value.SetIei(nasMessage.DLNASTransportPduSessionID2ValueType)
		req.PduSessionID2Value.SetPduSessionID2Value(uint8(pduid))
	}
	if cause != 0 {
		req.Cause5GMM = new(nasType.Cause5GMM)
		req.Cause5GMM.SetIei(nasMessage.DLNASTransportCause5GMMType)
		req.Cause5GMM.SetCauseValue(cause)
	}
	if timeru != nil {
		req.BackoffTimerValue = new(nasType.BackoffTimerValue)
		req.BackoffTimerValue.SetIei(nasMessage.DLNASTransportBackoffTimerValueType)
		req.BackoffTimerValue.SetLen(1)
		req.BackoffTimerValue.SetUnitTimerValue(*timeru)
		req.BackoffTimerValue.SetTimerValue(timerv)
	}

	m.GmmMessage.DLNASTransport = req

	pdu, err = Encode(ctx, m)
	return
}

func BuildNotification(ctx EncodingCtx, access models.AccessType) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeNotification)

	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	notification := nasMessage.NewNotification(0)
	notification.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	notification.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	notification.SetMessageType(nas.MsgTypeNotification)

	if access == models.ACCESSTYPE__3_GPP_ACCESS {
		notification.SetAccessType(nasMessage.AccessType3GPP)
	} else {
		notification.SetAccessType(nasMessage.AccessTypeNon3GPP)
	}

	m.GmmMessage.Notification = notification

	pdu, err = Encode(ctx, m)
	return
}

func BuildServiceAccept(ctx EncodingCtx, statuslist *[16]bool, reactlist *[16]bool, errpdulist []uint8, errcauses []uint8) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeServiceAccept)

	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	req := nasMessage.NewServiceAccept(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SetMessageType(nas.MsgTypeServiceAccept)
	if statuslist != nil {
		req.PDUSessionStatus = new(nasType.PDUSessionStatus)
		req.PDUSessionStatus.SetIei(nasMessage.ServiceAcceptPDUSessionStatusType)
		req.PDUSessionStatus.SetLen(2)
		req.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*statuslist)
	}
	if reactlist != nil {
		req.PDUSessionReactivationResult = new(nasType.PDUSessionReactivationResult)
		req.PDUSessionReactivationResult.SetIei(nasMessage.ServiceAcceptPDUSessionReactivationResultType)
		req.PDUSessionReactivationResult.SetLen(2)
		req.PDUSessionReactivationResult.Buffer = nasConvert.PSIToBuf(*reactlist)
	}
	if len(errpdulist) > 0 {
		req.PDUSessionReactivationResultErrorCause = new(nasType.PDUSessionReactivationResultErrorCause)
		req.PDUSessionReactivationResultErrorCause.SetIei(
			nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType)
		buf := nasConvert.PDUSessionReactivationResultErrorCauseToBuf(errpdulist, errcauses)
		req.PDUSessionReactivationResultErrorCause.SetLen(uint16(len(buf)))
		req.PDUSessionReactivationResultErrorCause.Buffer = buf
	}
	m.GmmMessage.ServiceAccept = req

	pdu, err = Encode(ctx, m)
	return
}

func BuildServiceReject(ctx EncodingCtx, pduSessionStatus *[16]bool, cause uint8) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeServiceReject)

	req := nasMessage.NewServiceReject(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SetMessageType(nas.MsgTypeServiceReject)
	req.SetCauseValue(cause)
	if pduSessionStatus != nil {
		req.PDUSessionStatus = new(nasType.PDUSessionStatus)
		req.PDUSessionStatus.SetIei(nasMessage.ServiceAcceptPDUSessionStatusType)
		req.PDUSessionStatus.SetLen(2)
		req.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pduSessionStatus)
	}

	m.GmmMessage.ServiceReject = req

	pdu, err = m.PlainNasEncode()
	return
}

type ConfigurationUpdateCommandCompleter func(*nasMessage.ConfigurationUpdateCommand) error

func BuildConfigurationUpdateCommand(ctx EncodingCtx, slicing *nasType.NetworkSlicingIndication, fn ConfigurationUpdateCommandCompleter) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeConfigurationUpdateCommand)

	req := nasMessage.NewConfigurationUpdateCommand(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.SetMessageType(nas.MsgTypeConfigurationUpdateCommand)

	if slicing != nil {
		req.NetworkSlicingIndication =
			nasType.NewNetworkSlicingIndication(nasMessage.ConfigurationUpdateCommandNetworkSlicingIndicationType)
		req.NetworkSlicingIndication = slicing
	}
	if fn != nil {
		if err = fn(req); err != nil {
			return
		}
	}

	m.GmmMessage.ConfigurationUpdateCommand = req

	pdu, err = m.PlainNasEncode()
	return
}

func BuildStatus5GMM(ctx EncodingCtx, cause uint8) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeStatus5GMM)

	req := nasMessage.NewStatus5GMM(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SetMessageType(nas.MsgTypeStatus5GMM)
	req.SetCauseValue(cause)

	m.GmmMessage.Status5GMM = req

	pdu, err = m.PlainNasEncode()
	return
}
func BuildDeregistrationRequest(ctx EncodingCtx, antype uint8, rereg bool, cause uint8) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeDeregistrationRequestUETerminatedDeregistration)

	req := nasMessage.NewDeregistrationRequestUETerminatedDeregistration(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.SetMessageType(nas.MsgTypeDeregistrationRequestUETerminatedDeregistration)

	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}
	req.SetAccessType(antype)
	req.SetSwitchOff(0)
	if rereg {
		req.SetReRegistrationRequired(nasMessage.ReRegistrationRequired)
	} else {
		req.SetReRegistrationRequired(nasMessage.ReRegistrationNotRequired)
	}

	if cause != 0 {
		req.Cause5GMM = nasType.NewCause5GMM(
			nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
		req.Cause5GMM.SetCauseValue(cause)
	}
	m.GmmMessage.DeregistrationRequestUETerminatedDeregistration = req

	pdu, err = Encode(ctx, m)
	return
}

func BuildDeregistrationAccept(ctx EncodingCtx) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeDeregistrationAcceptUEOriginatingDeregistration)

	req := nasMessage.NewDeregistrationAcceptUEOriginatingDeregistration(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.SetMessageType(nas.MsgTypeDeregistrationAcceptUEOriginatingDeregistration)

	m.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration = req

	pdu, err = m.PlainNasEncode()
	return
}

type SecurityModeCommandCompleter func(*nasMessage.SecurityModeCommand) error

func BuildSecurityModeCommand(ctx EncodingCtx, fn SecurityModeCommandCompleter) (pdu []byte, err error) {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeSecurityModeCommand)

	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nas.SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext,
	}

	req := nasMessage.NewSecurityModeCommand(0)
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.SecurityModeCommandMessageIdentity.SetMessageType(nas.MsgTypeSecurityModeCommand)

	if fn != nil {
		if err = fn(req); err != nil {
			return
		}
	}

	m.GmmMessage.SecurityModeCommand = req
	pdu, err = Encode(ctx, m)
	return
}
