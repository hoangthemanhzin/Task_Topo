package authentication

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/nfs/amf/uecontext"
	"etrib5gc/sbi/ausf/uea"
	"etrib5gc/sbi/models"
	"etrib5gc/util/fsm"
	"fmt"
	"time"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

const (
	MAX_SYNC_FAILURE_CNT int = 2
	T3560_DURATION           = 100 //miliseconds
)

type RanUe interface {
	logctx.LogWriter
	Worker() common.Executer
	SendAuthenticationRequest(func(*nasMessage.AuthenticationRequest) error) error
	SendAuthenticationReject(string) error
	UeContext() *uecontext.UeContext
}

type AuthProc struct {
	logctx.LogWriter
	fsm.State
	ranue RanUe

	authtype  models.AuthType
	supi      string
	rand      []byte
	autn      []byte
	hxresstar []byte
	kamf      []byte
	eap       string //temporary eap payload (eap-aka)
	success   bool   //eap-aka

	err        error
	failureCnt int
	onDone     func(*AuthProc)

	t3560 common.UeTimer //authentication request
}

func New(ranue RanUe, fn func(*AuthProc)) (proc *AuthProc) {
	proc = &AuthProc{
		LogWriter: ranue.WithFields(logctx.Fields{"mod": "authentication"}),
		ranue:     ranue,
		onDone:    fn,
		State:     fsm.NewState(AUTH_IDLE),
	}
	proc.t3560 = common.NewTimer(T3560_DURATION*time.Millisecond, func() {
		//AuthenticationRequest expired
		proc.sendEvent(T3560Event, nil)
	}, nil)
	proc.sendEvent(StartEvent, nil)
	return
}

func (proc *AuthProc) GetError() error {
	return proc.err
}

func (proc *AuthProc) AuthInfo() (kamf []byte, supi string, eap string, success bool) {
	kamf = proc.kamf
	supi = proc.supi
	eap = proc.eap
	success = proc.success
	return
}

func (proc *AuthProc) loadCtx(ctx *models.UEAuthenticationCtx) (err error) {
	proc.authtype = ctx.AuthType
	dat := &ctx.Var5gAuthData
	if proc.rand, err = hex.DecodeString(dat.Rand); err != nil {
		return
	}
	if proc.autn, err = hex.DecodeString(dat.Autn); err != nil {
		return
	}
	if proc.hxresstar, err = hex.DecodeString(dat.HxresStar); err != nil {
		return
	}
	return
}

func (proc *AuthProc) Handle(msg *libnas.GmmMessage) (err error) {
	switch msg.GetMessageType() {
	case libnas.MsgTypeAuthenticationResponse:
		err = proc.sendEvent(AuthResponseEvent, msg.AuthenticationResponse)
	case libnas.MsgTypeAuthenticationFailure:
		err = proc.sendEvent(AuthFailureEvent, msg.AuthenticationFailure)
	default:
		err = fmt.Errorf("Unexpected Nas Message during authentication")
	}

	return
}

func (proc *AuthProc) sendEvent(ev fsm.EventType, args interface{}) error {
	return _sm.SendEvent(proc.ranue.Worker(), proc, ev, args)
}

func (proc *AuthProc) challenge() {
	proc.Trace("Start challenging")
	uectx := proc.ranue.UeContext()
	//1. get authentication information from AUSF
	info := &models.AuthenticationInfo{
		SupiOrSuci:         uectx.Suci(), //NOTE: may use supi as well
		ServingNetworkName: uectx.ServingNetwork(),
	}
	//proc.Infof("Send an UeAuthenticationsPost to AUSF [SUCI=%s]", proc.id)
	if ctx, err := uea.UeAuthenticationsPost(uectx.Ausf(), *info); err != nil {
		proc.Errorf("Failed to authenticate with AUSF: %s", err.Error())
		proc.sendEvent(DoneEvent, err)
		return
	} else {
		proc.Info("Receive authentication information from the AUSF")
		if err = proc.loadCtx(ctx); err != nil {
			proc.sendEvent(DoneEvent, err)
			return
		}
		//2.b then send Downlink AuthenticationRequest to UE
		proc.authRequest() //send challening
	}
	return
}

func (proc *AuthProc) handleT3560() {
	//TODO: should resend an authentication request
	proc.Trace("t3560 expired")
	proc.sendEvent(DoneEvent, fmt.Errorf("T3560 expired"))
}

// handle an authentication response from the UE. The handling is dependent on
// the authentication method (5g-aka or eap-aka-prime).
func (proc *AuthProc) handleAuthenticationResponse(msg *nasMessage.AuthenticationResponse) {
	proc.Info("Receive a NAS AuthenticationResponse from UE")
	proc.t3560.Stop()
	uectx := proc.ranue.UeContext()
	var err error
	switch proc.authtype {
	case models.AUTHTYPE__5_G_AKA:
		resstar := msg.AuthenticationResponseParameter.GetRES()
		hx := sha256.Sum256(append(proc.rand, resstar[:]...))
		hresstar := hx[16:]
		proc.Tracef("rand=%x,resstar=%x,hash=%x", proc.rand, resstar[:], hresstar)
		//check resstar
		if !bytes.Equal(hresstar, proc.hxresstar) {
			proc.Errorf("Mismatch hxresstar %x vs %x", hresstar, proc.hxresstar)
			proc.ranue.SendAuthenticationReject("")
			proc.sendEvent(DoneEvent, fmt.Errorf("Mismatched ResStar"))
			return
		}
		proc.Infof("Send a UeAuthenticationsAuthCtxId5gAkaConfirmationPut to AUSF for UE[SUCI=%s]", uectx.Suci())
		body := &models.ConfirmationData{
			ResStar: hex.EncodeToString(resstar[:]),
		}
		var resp *models.ConfirmationDataResponse
		if resp, err = uea.UeAuthenticationsAuthCtxId5gAkaConfirmationPut(uectx.Ausf(), uectx.Suci(), body); err != nil {
			proc.Error(err.Error())
			proc.ranue.SendAuthenticationReject("")
			proc.sendEvent(DoneEvent, err)
			return
		}
		switch resp.AuthResult {
		case models.AUTHRESULT_SUCCESS:
			proc.Infof("Authentication is success; receive a SUPI=%s", resp.Supi)
			if err = proc.createKamf(resp.Supi, resp.Kseaf); err != nil {
				proc.ranue.SendAuthenticationReject("")
				proc.sendEvent(DoneEvent, err)
				return
			}
			proc.Info("A secret key (KAMF) is created for the AMF")
			proc.supi = resp.Supi
			proc.sendEvent(DoneEvent, nil)
		case models.AUTHRESULT_FAILURE:
			proc.ranue.SendAuthenticationReject("")
			proc.sendEvent(DoneEvent, fmt.Errorf("An authentication failure from AUSF"))
		}
	case models.AUTHTYPE_EAP_AKA_PRIME:
		/*
			//TODO: handle eap_aka authentication

			var resp *models.EapSession
			eapmsg := base64.StdEncoding.EncodeToString(msg.EAPMessage.GetEAPMessage())
			body := &models.EapSession{
				EapPayload: eapmsg,
			}
			if resp, err = uea.EapAuthMethod(ue.Ausf(), proc.id, body); err != nil {
				proc.sendEvent(RejectEvent, err)
				return
			}
			switch resp.AuthResult {
			case models.AUTHRESULT_SUCCESS:
				if err = proc.createKamf(resp.Supi, resp.KSeaf); err != nil {
					proc.sendEvent(RejectEvent, err)
					return
				}
				proc.proc.eap = resp.EapPayload
				proc.proc.supi = resp.Supi

			case models.AUTHRESULT_FAILURE:
				proc.sendEvent(RejectEvent, err)

			case models.AUTHRESULT_ONGOING:
				proc.proc.eap = resp.EapPayload
				proc.authRequest()
			}
		*/
	}
}

func (proc *AuthProc) handleAuthenticationFailure(msg *nasMessage.AuthenticationFailure) {
	proc.Info("Receive a NAS AuthenticationFailure")
	proc.t3560.Stop()
	uectx := proc.ranue.UeContext()
	cause := msg.Cause5GMM.GetCauseValue()

	if proc.authtype == models.AUTHTYPE__5_G_AKA {
		switch cause {
		case nasMessage.Cause5GMMMACFailure:
			proc.Warnln("Authentication Failure Cause: Mac Failure")
			proc.sendEvent(DoneEvent, fmt.Errorf("Mac Failure"))
		case nasMessage.Cause5GMMNon5GAuthenticationUnacceptable:
			proc.Warnln("Authentication Failure Cause: Non-5G Authentication Unacceptable")
			proc.sendEvent(DoneEvent, fmt.Errorf("Non-5G Authentication Unacceptable"))
		case nasMessage.Cause5GMMngKSIAlreadyInUse:
			proc.Warnln("Authentication Failure Cause: NgKSI Already In Use")
			ngksi := uectx.NgKsi()
			proc.failureCnt = 0
			proc.Warn("Select new NgKsi")
			// select new ngksi
			if ngksi.Ksi < 6 { // ksi is range from 0 to 6
				ngksi.Ksi += 1
			} else {
				ngksi.Ksi = 0
			}
			proc.authRequest()
		case nasMessage.Cause5GMMSynchFailure: // TS 24.501 5.4.1.3.7 case f
			proc.Warn("Authentication Failure 5GMM Cause: Synch Failure")
			proc.failureCnt++
			if proc.failureCnt >= MAX_SYNC_FAILURE_CNT {
				proc.Warn("max consecutive Synch Failure, terminate authentication procedure")
				proc.sendEvent(DoneEvent, fmt.Errorf("Max sync failure"))
				return
			}

			auts := msg.AuthenticationFailureParameter.GetAuthenticationFailureParameter()
			proc.Tracef("Resync with auts[%d]", len(auts))
			info := models.AuthenticationInfo{
				SupiOrSuci:         uectx.Suci(), //NOTE: may use supi as well
				ServingNetworkName: uectx.ServingNetwork(),
				ResynchronizationInfo: &models.ResynchronizationInfo{
					Auts: hex.EncodeToString(auts[:]),
					Rand: hex.EncodeToString(proc.rand),
				}}

			if ctx, err := uea.UeAuthenticationsPost(uectx.Ausf(), info); err != nil {

				proc.ranue.SendAuthenticationReject("")
				proc.sendEvent(DoneEvent, err)
			} else {
				if err = proc.loadCtx(ctx); err != nil {

					proc.ranue.SendAuthenticationReject("")
					proc.sendEvent(DoneEvent, err)
				} else {
					proc.authRequest()
				}
			}
		default:
			proc.Warn("Unknown authentication failure cause")
			proc.sendEvent(DoneEvent, fmt.Errorf("Unknown failure cause"))
		}
	} else if proc.authtype == models.AUTHTYPE_EAP_AKA_PRIME {
		proc.Trace("Handle authentication failure: eap")
		proc.sendEvent(DoneEvent, fmt.Errorf("Eap-aka is not supported"))
		/*
			switch cause {
			case nasMessage.Cause5GMMngKSIAlreadyInUse:
				proc.Warn("Authentication Failure 5GMM Cause: NgKSI Already In Use")
				if ue.NgKsi.Ksi < 6 { // ksi is range from 0 to 6
					ue.NgKsi.Ksi += 1
				} else {
					ue.NgKsi.Ksi = 0
				}
				if err = nas.SendAuthenticationRequest(sender, proc.ranue); err != nil {
					fn(OP_STATUS_ONGOING)
				}
			}
		*/
	} else {
		proc.Warnf("Handle authentication failure: unknown authentication type")
		proc.sendEvent(DoneEvent, fmt.Errorf("Unknown failure type"))
	}
	proc.Trace("Handle authentication failure is complete")
	return
}
