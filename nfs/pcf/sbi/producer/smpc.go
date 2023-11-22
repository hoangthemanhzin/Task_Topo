package producer

import (
	"etrib5gc/sbi/models"
	"fmt"
)

func (p *Producer) SMPC_HandleDeleteSMPolicy(smPolicyId string, body models.SmPolicyDeleteData) (prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SMPC_HandleDeleteSMPolicy has not been implemented")
	return
}
func (p *Producer) SMPC_HandleGetSMPolicy(smPolicyId string) (result models.SmPolicyControl, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SMPC_HandleGetSMPolicy has not been implemented")
	return
}
func (p *Producer) SMPC_HandleUpdateSMPolicy(smPolicyId string, body models.SmPolicyUpdateContextData) (result models.SmPolicyDecision, prob *models.ProblemDetails) {
	//TODO: to be implemented
	panic("SMPC_HandleUpdateSMPolicy has not been implemented")
	return
}
func (p *Producer) SMPC_HandleCreateSMPolicy(body models.SmPolicyContextData) (decision models.SmPolicyDecision, prob *models.ProblemDetails) {
	//TODO: to be implemented
	log.Infof("Receive CreateSmPolicy: pduSessionId=%d, SUPI=%s, return a dummpy PolicyDecision", body.PduSessionId, body.Supi)
	//supi := body.Supi
	sid := body.PduSessionId

	//var smdata models.SmPolicyData
	//smpolid := fmt.Sprintf("%s-%d", supi, sid)
	//smPolicyData = ue.NewUeSmPolicyData(smPolicyID, request, &smData)

	decision = models.SmPolicyDecision{
		SessRules: make(map[string]models.SessionRule),
	}
	srid := fmt.Sprintf("srid-%d", sid) //session rule id
	srule := models.SessionRule{
		AuthSessAmbr: body.SubsSessAmbr,
		SessRuleId:   srid,
		// RefUmData
		// RefCondData
	}
	defqos := &body.SubsDefQos
	//just authorized the requested Qos no matter what
	srule.AuthDefQos = models.AuthorizedDefaultQos{
		Var5qi:        defqos.Var5qi,
		Arp:           defqos.Arp,
		PriorityLevel: defqos.PriorityLevel,
		// AverWindow
		// MaxDataBurstVol
	}

	decision.SessRules[srid] = srule

	dnndata := models.SmPolicyDnnData{ //default one (should be from subscribed data for given Supi/SliceInfo/Dnn
		Online:    true,
		Offline:   false,
		Ipv4Index: 0,
		Ipv6Index: 0,
	}

	decision.Online = dnndata.Online
	decision.Offline = dnndata.Offline
	decision.Ipv4Index = dnndata.Ipv4Index
	decision.Ipv6Index = dnndata.Ipv6Index

	flowinfos := []models.FlowInformation{
		{
			FlowDescription:   "permit out ip from any to assigned",
			FlowDirection:     models.FLOWDIRECTIONRM_DOWNLINK,
			PacketFilterUsage: true,
			PackFiltId:        "PackFiltId-0",
		},
		{
			FlowDescription:   "permit out ip from any to assigned",
			FlowDirection:     models.FLOWDIRECTIONRM_DOWNLINK,
			PacketFilterUsage: true,
			PackFiltId:        "PackFiltId-1",
		},
	}
	precedence := int32(10)
	pccrule := models.PccRule{
		AppId:      "",
		FlowInfos:  flowinfos,
		PccRuleId:  fmt.Sprintf("pccrule-id-%d", sid),
		Precedence: precedence,
	}

	qosdata := models.QosData{ //should be created with flow information from UDR
		QosId:   fmt.Sprintf("qos-id-%d", 10),
		GbrUl:   "100 Mbps",
		GbrDl:   "100 Mbps",
		MaxbrUl: "1000 Mbps",
		MaxbrDl: "1000 Mbps",
		Qnc:     false,
		Var5qi:  int32(5),
	}
	decision.QosDecs = make(map[string]models.QosData)
	decision.QosDecs[qosdata.QosId] = qosdata
	pccrule.RefQosData = []string{qosdata.QosId}
	decision.PccRules = make(map[string]models.PccRule)
	decision.PccRules[pccrule.PccRuleId] = pccrule

	decision.SuppFeat = "dummy"
	decision.QosFlowUsage = body.QosFlowUsage
	//decision.PolicyCtrlReqTriggers = util.PolicyControlReqTrigToArray(0x40780f)
	return
}
