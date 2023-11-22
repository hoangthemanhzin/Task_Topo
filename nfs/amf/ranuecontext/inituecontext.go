package ranuecontext

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"

	"github.com/free5gc/nas/nasMessage"
)

func (ranue *RanUe) handleInitUeContext(dat *events.InitUeContextData) {
	ranue.Infof("handle InitUeContext")
	if ranue.regctx != nil {
		//NOTE: this is a very slim chance
		//TODO: we may need to send a rejection
		return
	}
	msg := dat.GmmMsg
	if regmsg := msg.RegistrationRequest; regmsg != nil {
		ranue.Infof("Process RegistrationRequest content")
		//RegistrationRequest
		//1. Decode the nas container if it is non-nil
		// there must be a security context (from a previous registration)
		if regmsg.NASMessageContainer != nil {
			content := regmsg.NASMessageContainer.GetNASMessageContainerContents()
			if m, err := nas.Decode(ranue, content); err == nil {
				if m.RegistrationRequest != nil {
					regmsg = m.RegistrationRequest
				} else {
					ranue.Errorf("Unexpected message in the NasContainer")
					ranue.rejectRegistration(nasMessage.Cause5GMMInvalidMandatoryInformation)
					return
				}
			} else {
				//NOTE: if integrity check fails the security context is
				//invalidated
				ranue.Warnf("Decode NasContainer in RegistrationRequest failed: %s", err.Error())
				//just try moving on to setup a new security context
			}
		}

		if regmsg.UESecurityCapability == nil {
			ranue.Errorf("No security capability")
			ranue.rejectRegistration(nasMessage.Cause5GMMInvalidMandatoryInformation)
			return
		}

		//re-set the registration type
		regtype := regmsg.NgksiAndRegistrationType5GS.GetRegistrationType5GS()
		switch regtype {
		case nasMessage.RegistrationType5GSInitialRegistration:
			ranue.Debugf("RegistrationType: Initial Registration")
		case nasMessage.RegistrationType5GSMobilityRegistrationUpdating:
			ranue.Debugf("RegistrationType: Mobility Registration Updating")
			if !ranue.registered {
				ranue.Errorf("Invalid registation type")
				//nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState
				ranue.rejectRegistration(nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState)
				return
			}
		case nasMessage.RegistrationType5GSPeriodicRegistrationUpdating:
			ranue.Debugf("RegistrationType: Periodic Registration Updating")
			if !ranue.registered {
				ranue.Errorf("Invalid registation type")
				ranue.rejectRegistration(nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState)
				return
			}

		case nasMessage.RegistrationType5GSEmergencyRegistration:
			ranue.Debugf("RegistrationType: Emergency")
			ranue.Errorf("Emergency Registration not supported")
			ranue.rejectRegistration(nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState)
			return
		case nasMessage.RegistrationType5GSReserved:
			ranue.Debugf("RegistrationType: Reserved")
			regtype = nasMessage.RegistrationType5GSInitialRegistration
		default:
			ranue.Debugf("RegistrationType: %v, change state to InitialRegistration", regtype)
			regtype = nasMessage.RegistrationType5GSInitialRegistration
		}
		ranue.regctx = events.RegCtxFromRegReq(dat.InitUeMsg, regmsg, regtype)
	} else {
		ranue.Infof("Process ServiceRequest content")
		//ServiceRequest
		//service request must be integrity protected
		//and it should be rejected if its verification fails
		//if the service type is just signalling, send an accept, do no more about
		//session management

		//cause := nasMessage.Cause5GMMProtocolErrorUnspecified //default cause for rejection
		servmsg := msg.ServiceRequest
		//2. check paging status
		//TODO:

		//3. decode NasContainer if needs
		if servmsg.NASMessageContainer != nil {
			content := servmsg.NASMessageContainer.GetNASMessageContainerContents()
			if m, err := nas.Decode(ranue, content); err == nil {
				if m.ServiceRequest != nil {
					servmsg = m.ServiceRequest
				} else {
					ranue.Errorf("Unexpected Nas message ServiceRequest")
					ranue.rejectService(nasMessage.Cause5GMMInvalidMandatoryInformation, nil, false)
					return
				}
			} else {
				//NOTE: if integrity check fails the security context is
				//invalidated
				ranue.Warnf("Decode NasContainer in ServiceRequest failed: %s", err.Error())
				//security context is invalidated; move on to setup security
				//context again
			}
		}
		ranue.regctx = events.RegCtxFromServReq(dat.InitUeMsg, servmsg)
	}

	if err := ranue.ue.HandleEvent(&common.EventData{
		EvType:  events.REGISTRATION_REQUEST,
		Content: ranue,
	}); err != nil {
		if ranue.regctx.RegistrationRequest() != nil {
			ranue.rejectRegistration(nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState)
		} else {
			ranue.rejectService(nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState, nil, false)
		}
	}
	return
}
