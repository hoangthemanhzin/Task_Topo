package sm

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"fmt"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

func (smctx *SmContext) handlePostSmContexts(jobinfo *PostSmContextsJob) {
	var err error
	req := jobinfo.Req
	//1. decode Nas Sm message
	m := libnas.NewMessage()
	n1cause := nasMessage.Cause5GSMProtocolErrorUnspecified
	if err = m.GsmMessageDecode(&req.BinaryDataN1SmMessage); err == nil {
		if m.GsmHeader.GetMessageType() != libnas.MsgTypePDUSessionEstablishmentRequest {
			err = fmt.Errorf("Expecting PduSessionEstablishmentRequest")
		}
	}

	if err != nil {
		jobinfo.setN1Error(err, n1cause, smctx)
		return
	}
	//handle Nas Sm message (PduSessionEstablishmentRequest)
	if err, n1cause = smctx.handlePduSessionEstablishmentRequest(m.PDUSessionEstablishmentRequest); err != nil {
		smctx.Errorf("Fail to handle PduSessionEstablishmentRequest: %s", err.Error())
		jobinfo.setN1Error(err, n1cause, smctx)
		return
	}

	//2.  create SmPolicy (request PCF)
	if err = smctx.CreateSmPol(); err != nil {
		//TODO: set n1cause
		jobinfo.setN1Error(err, n1cause, smctx)
		return
	}

	//3. select anchoring UPF and allocate Ue's IP
	//TODO: upf query shoud be composed with a combination of information from
	//AMF, PCF (AF) and UE
	//query := &upmodels.PathQuery{
	query := &n43.UpfPathQuery{
		Dnn:    smctx.dnn,
		Snssai: *smctx.snssai,
		//Nets:   string{"an1"},
		Nets: req.JsonData.RanNets,
	}

	smctx.Infof("Search UPF path to Dnn=%s", query.Dnn)
	//select upf, allocate Ue's ip then find a path
	if smctx.tunnel, err = smctx.upmanager.CreateTunnel(query, smctx.upmfcli); err != nil {
		smctx.Errorf("Create tunnel fails: %s", err.Error())
		//TODO: set n1cause
		jobinfo.setN1Error(err, n1cause, smctx)
		return
	}

	smctx.Infof("Setup tunnel")
	//4. create PDRs for all links in the path
	if err = smctx.setupTunnel(); err != nil {
		//TODO: set n1cause
		smctx.Errorf("Setup tunnel failed: %s", err.Error())
		//TODO: does the created tunnel need to be clean up?
		jobinfo.setN1Error(err, n1cause, smctx)
		return
	}

	smctx.sendEvent(SessActEvent, nil) //must never fail

	jobinfo.Rsp = &models.PostSmContextsResponse{
		JsonData: smctx.BuildSmContextCreatedData(),
	}
	smctx.ctx.AddSmContext(smctx)
	smctx.Infof("SmContext added")
	return
}

/*

//keep this old implementation just as a reference

func (gsm *GSM) HandlePostSmContext(req *models.PostSmContextsRequest, callback *models.Callback) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse) {
	smctx.Infof("Receive a PostSmContexts from AMF [callback = %s]", callback.String())
	var err error
	//1. decode Nas Sm message
	m := libnas.NewMessage()
	if err = m.GsmMessageDecode(&req.BinaryDataN1SmMessage); err == nil {
		// Check has PDU Session Establishment Request
		if m.GsmHeader.GetMessageType() != libnas.MsgTypePDUSessionEstablishmentRequest {
			err = fmt.Errorf("Expecting PduSessionEstablishmentRequest")
		}
	}
	if err != nil {
		ersp = &models.PostSmContextsErrorResponse{
			JsonData: models.SmContextCreateError{
				Error: &models.ExtProblemDetails{
					Status: http.StatusForbidden,
					Detail: err.Error(),
					Type:   "N1SmError",
				},
			},
		}
		return
	}

	supi := req.JsonData.Supi
	sid := uint32(req.JsonData.PduSessionId)
	//pei := req.pei

	var smcontext *SmContext
	var n1cause uint8

	n1cause = nasMessage.Cause5GSMProtocolErrorUnspecified

	//2. release existing sm context (if any)
	smcontext = gsm.findSmContext(supi, sid)
	if smcontext != nil {
		//existing smcontext
		if err, n1cause = gsm.releaseSmContext(smcontext); err != nil {
			ersp = postSmCtxErrRsp(smcontext, err, n1cause)
			return
		}
	} else {
		//3. create a new smcontext from the request
		if smcontext, err, n1cause = gsm.createSmContext(&req.JsonData, m.PDUSessionEstablishmentRequest); err != nil {
			ersp = postSmCtxErrRsp(smcontext, err, n1cause)
			return
		}
	}

	if err = smcontext.CreateSmPol(); err != nil {
		n1cause = nasMessage.Cause5GSMProtocolErrorUnspecified
		ersp = postSmCtxErrRsp(smcontext, err, n1cause)
		return
	}

	//create AMF consumer
	instanceid := types.InstanceId(callback.InstanceId)
	if smcontext.amfcli, err = gsm.ctx.Agent().Sender(types.ServiceId(callback.ServiceId), &instanceid); err != nil {
		smctx.Errorf("Fail to create amf consumer for %s : %s : %s", callback.ServiceId, callback.InstanceId, err.Error())
		return
	} else {
		smctx.Infof("Amf consumer is created: %s/%s", callback.ServiceId, callback.InstanceId)
	}

	//4. select anchoring UPF and allocate Ue's IP
	//TODO: upf query shoud be composed with a combination of information from
	//AMF, PCF (AF) and UE
	smctx.Infof("Start activating the session %d [SUPI=%s]", smcontext.sid, smcontext.supi)
	query := &upmodels.PathQuery{
		Dnn:    smcontext.dnn,
		Snssai: *smcontext.snssai,
		//	Dnn: "internet",
		//	Snssai: models.Snssai{
		//		Sd:  "12345",
		//		Sst: 0,
		//	},
		Nets: []string{"an1"},
	}
	smctx.Infof("Looking for a data path (UPFs) in %s slice that reaches to Dnn=%s", query.Snssai.String(), query.Dnn)
	//select upf, allocate Ue's ip then find a path
	if smcontext.tunnel, err = gsm.upman.CreateTunnel(query); err != nil {
		smctx.Errorf("Create tunnel fails: %s", err.Error())
		n1cause = nasMessage.Cause5GSMProtocolErrorUnspecified
		ersp = postSmCtxErrRsp(smcontext, err, n1cause)
		return
	}
	//create PDRs for all links in the path
	if err = smcontext.setupTunnel(); err != nil {
		smctx.Errorf("Setup tunnel fails: %s", err.Error())
		n1cause = nasMessage.Cause5GSMProtocolErrorUnspecified
		ersp = postSmCtxErrRsp(smcontext, err, n1cause)
		return
	}
	smctx.Infof("A data path has been setup for the session %d [SUPI=%s]", smcontext.sid, smcontext.supi)
	//NOTE: definitely we can move to activting event abit early (before
	//processing the PDUSessionEstablishment and setting up the tunnel. So
	//these procedures will be performed in the SM_ACTIVATING state
	if err = smcontext.sendEvent(SessActEvent, nil); err != nil {
		n1cause = nasMessage.Cause5GSMProtocolErrorUnspecified
		ersp = postSmCtxErrRsp(smcontext, err, n1cause)
		return
	}

	rsp = &models.PostSmContextsResponse{
		JsonData: smcontext.BuildCreatedData(),
	}

	//add SmContext to the managed list
	gsm.smlist.add(smcontext)
	return
}


*/
