package nas

import (
	"fmt"
	"reflect"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/security"
)

type EncodingCtx interface {
	NasSecCtx() *NasSecCtx
	BearerType() uint8
	IsEmergency() bool
}

func Encode(ctx EncodingCtx, msg *libnas.Message) (pdu []byte, err error) {
	//NOTE: panic if ctx is nil or msg is nil

	if ctx == nil {
		pdu, err = msg.PlainNasEncode()
		return
	}

	if secCtx := ctx.NasSecCtx(); secCtx == nil || !secCtx.IsValid() {
		//no valid security context
		pdu, err = msg.PlainNasEncode()
	} else {
		// Security protected NAS Message
		// a security protected NAS message must be integrity protected, and ciphering is optional
		ciphering := false
		switch msg.SecurityHeader.SecurityHeaderType {
		case libnas.SecurityHeaderTypeIntegrityProtected:
		case libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered:
			ciphering = true
		case libnas.SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext:
			secCtx.ulcount.Set(0, 0)
			secCtx.dlcount.Set(0, 0)
		default:
			err = fmt.Errorf("Wrong security header type: 0x%0x", msg.SecurityHeader.SecurityHeaderType)
			return
		}

		// encode plain nas first
		if pdu, err = msg.PlainNasEncode(); err != nil {
			return
		}

		if ciphering {
			if err = security.NASEncrypt(secCtx.encalg, secCtx.enckey, secCtx.dlcount.Get(),
				ctx.BearerType(), security.DirectionDownlink, pdu); err != nil {
				return
			}
		}

		// add sequece number
		pdu = append([]byte{secCtx.dlcount.SQN()}, pdu[:]...)
		var mac32 []byte
		if mac32, err = security.NASMacCalculate(secCtx.intalg, secCtx.intkey, secCtx.dlcount.Get(),
			ctx.BearerType(), security.DirectionDownlink, pdu); err != nil {
			return
		}
		// Add mac value
		pdu = append(mac32, pdu[:]...)

		// Add EPD and Security Type
		msgSecurityHeader := []byte{msg.SecurityHeader.ProtocolDiscriminator, msg.SecurityHeader.SecurityHeaderType}
		pdu = append(msgSecurityHeader, pdu[:]...)

		// Increase DL Count
		secCtx.dlcount.AddOne()
	}

	return
}

func Decode(ctx EncodingCtx, pdu []byte) (msg libnas.Message, err error) {
	msg.SecurityHeaderType = libnas.GetSecurityHeaderType(pdu) & 0x0f
	if msg.SecurityHeaderType == libnas.SecurityHeaderTypePlainNas {

		if err = msg.PlainNasDecode(&pdu); err != nil {
			return
		}

		if ctx != nil {
			secCtx := ctx.NasSecCtx()
			// if having sec-context and not for emergency, check for message type
			if secCtx != nil && !ctx.IsEmergency() {
				if msg.GmmMessage == nil {
					err = fmt.Errorf("Gmm Message is nil")
					return
				}

				// TS 24.501 4.4.4.3: Except the messages listed below, no NAS signalling messages shall be processed
				// by the receiving 5GMM entity in the AMF or forwarded to the 5GSM entity, unless the secure exchange
				// of NAS messages has been established for the NAS signalling connection
				switch msg.GmmHeader.GetMessageType() {
				case libnas.MsgTypeRegistrationRequest:
				case libnas.MsgTypeIdentityResponse:
				case libnas.MsgTypeAuthenticationResponse:
				case libnas.MsgTypeAuthenticationFailure:
				case libnas.MsgTypeSecurityModeReject:
				case libnas.MsgTypeDeregistrationRequestUEOriginatingDeregistration:
				case libnas.MsgTypeDeregistrationAcceptUETerminatedDeregistration:
				default:
					err = fmt.Errorf("UE can not send plain nas for non-emergency service when there is a valid security context")
				}
			}
		}
		return
	} else { // Security protected NAS message
		if ctx == nil {
			err = fmt.Errorf("No valid security context to decrypt the NAS pdu")
			return
		}
		secCtx := ctx.NasSecCtx()
		if secCtx == nil || !secCtx.IsValid() {
			err = fmt.Errorf("No valid security context to decrypt the NAS pdu")
			return
		}
		secHeader := pdu[0:6]
		seqNum := pdu[6]

		rmac32 := secHeader[2:] //receive mac32
		pdu = pdu[6:]

		// a security protected NAS message must be integrity protected, and ciphering is optional
		ciphered := false
		switch msg.SecurityHeaderType {
		case libnas.SecurityHeaderTypeIntegrityProtected:
		case libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered:
			ciphered = true
		case libnas.SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext:
			ciphered = true
			secCtx.ulcount.Set(0, 0)
		default:
			err = fmt.Errorf("Wrong security header type: 0x%0x", msg.SecurityHeader.SecurityHeaderType)
			return
		}

		if secCtx.ulcount.SQN() > seqNum {
			secCtx.ulcount.SetOverflow(secCtx.ulcount.Overflow() + 1)
		}
		secCtx.ulcount.SetSQN(seqNum)
		var mac32 []byte
		if mac32, err = security.NASMacCalculate(secCtx.intalg, secCtx.intkey, secCtx.ulcount.Get(),
			ctx.BearerType(), security.DirectionUplink, pdu); err != nil {
			return
		}

		if !reflect.DeepEqual(mac32, rmac32) {
			err = fmt.Errorf("Integrity check not passed")
			secCtx.invalidate() //make the sec context invalid
			return
		}

		if ciphered {
			// decrypt payload without sequence number (payload[1])
			if err = security.NASEncrypt(secCtx.encalg, secCtx.enckey, secCtx.ulcount.Get(), ctx.BearerType(),
				security.DirectionUplink, pdu[1:]); err != nil {
				return
			}
		}

		// remove sequece Number
		pdu = pdu[1:]
		err = msg.PlainNasDecode(&pdu)
	}
	return
}
