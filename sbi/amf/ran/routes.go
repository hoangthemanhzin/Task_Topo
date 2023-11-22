package ran

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{
	{
		Label:   "UplinkNasTransport",
		Method:  http.MethodPut,
		Path:    "/:ueId/ul",
		Handler: OnUlNasTransport,
	},
	{
		Label:   "InitUeContext",
		Method:  http.MethodPost,
		Path:    "/init",
		Handler: OnInitUeContext,
	},
	{
		Label:   "NasNonDeliveryIndication",
		Method:  http.MethodPost,
		Path:    "/:ueId/naserr",
		Handler: OnNasNonDeliveryIndication,
	},
	{
		Label:   "UeCtxRelReq",
		Method:  http.MethodPost,
		Path:    "/:ueId/not/ctxrel",
		Handler: OnUeCtxRelReq,
	},
	{
		Label:   "PduSessResNot",
		Method:  http.MethodPost,
		Path:    "/:ueId/rep/pdu",
		Handler: OnPduSessResNot,
	},
	{
		Label:   "PduSessResModInd",
		Method:  http.MethodPost,
		Path:    "/:ueId/rep/modind",
		Handler: OnPduSessResModInd,
	},
	{
		Label:   "RrcInactTranRep",
		Method:  http.MethodPost,
		Path:    "/:ueId/rep/rrc",
		Handler: OnRrcInactTranRep,
	},
}
var _damfroutes = sbi.SbiRoutes{
	{
		Label:   "UplinkNasTransport",
		Method:  http.MethodPut,
		Path:    "/:ueId/ul",
		Handler: OnUlNasTransport,
	},
	{
		Label:   "InitUeContext",
		Method:  http.MethodPost,
		Path:    "/init",
		Handler: OnInitUeContext,
	},
	{
		Label:   "NasNonDeliveryIndication",
		Method:  http.MethodPost,
		Path:    "/:ueId/naserr",
		Handler: OnNasNonDeliveryIndication,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "ran",
		Routes:  _routes,
		Handler: p,
	}
}

func DamfService(p DProducer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "ran",
		Routes:  _damfroutes,
		Handler: p,
	}
}
