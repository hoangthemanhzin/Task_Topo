package events

import (
	"etrib5gc/sbi/models/n2models"
	"fmt"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

type InitUeContextData struct {
	InitUeMsg *n2models.InitUeContextRequest
	GmmMsg    *libnas.GmmMessage
}

type RegistrationContext struct {
	initmsg *n2models.InitUeContextRequest
	regmsg  *nasMessage.RegistrationRequest
	servmsg *nasMessage.ServiceRequest
	regtype uint8 //registration type: initial or update

	auth n2models.AuthUeCtx
}

func newRegistrationContext(initmsg *n2models.InitUeContextRequest) *RegistrationContext {
	regctx := &RegistrationContext{
		initmsg: initmsg,
	}
	if initmsg.UeCtx != nil {
		//TODO: make a deep copy instead
		regctx.auth = *initmsg.UeCtx
	}

	return regctx
}

func RegCtxFromRegReq(initmsg *n2models.InitUeContextRequest, regmsg *nasMessage.RegistrationRequest, regtype uint8) (regctx *RegistrationContext) {
	regctx = newRegistrationContext(initmsg)
	regctx.regmsg = regmsg
	regctx.regtype = regtype
	return
}
func RegCtxFromServReq(initmsg *n2models.InitUeContextRequest, servmsg *nasMessage.ServiceRequest) (regctx *RegistrationContext) {
	regctx = newRegistrationContext(initmsg)
	regctx.servmsg = servmsg
	return
}

func (r *RegistrationContext) AuthCtx() *n2models.AuthUeCtx {
	return &r.auth
}

func (r *RegistrationContext) RegType() uint8 {
	return r.regtype
}

func (r *RegistrationContext) IsRequestContext() bool {
	return r.initmsg.ContextRequest
}
func (r *RegistrationContext) RegistrationRequest() *nasMessage.RegistrationRequest {
	return r.regmsg
}
func (r *RegistrationContext) ServiceRequest() *nasMessage.ServiceRequest {
	return r.servmsg
}

// NasMessage is re-transmissed in a NasContainer of the SecurityModeComplete
func (r *RegistrationContext) UpdateMsg(msg *libnas.GmmMessage) (err error) {
	if r.regmsg != nil {
		if msg.RegistrationRequest != nil {
			r.regmsg = msg.RegistrationRequest
			return
		}
	} else {
		if msg.ServiceRequest != nil {
			r.servmsg = msg.ServiceRequest
			return
		}

	}
	err = fmt.Errorf("Invalid re-tranmissioned NasMessage")
	return
}
