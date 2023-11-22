package nas

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{
	{
		Label:   "InitUeContextStatus",
		Method:  http.MethodPost,
		Path:    "status/:ueId",
		Handler: OnInitUeContextStatus,
	},

	{
		Label:   "DownlinkNas",
		Method:  http.MethodPost,
		Path:    "dl/:ueId",
		Handler: OnNasDl,
	},
	{
		Label:   "InitCtxSetup",
		Method:  http.MethodPost,
		Path:    "uectx/:ueId/set",
		Handler: OnInitCtxSet,
	},
	{
		Label:   "UeCtxMod",
		Method:  http.MethodPost,
		Path:    "uectx/:ueId/mod",
		Handler: OnUeCtxMod,
	},
	{
		Label:   "UeCtxRel",
		Method:  http.MethodPost,
		Path:    "uectx/:ueId/rel",
		Handler: OnUeCtxRel,
	},
	{
		Label:   "PduSessResSet",
		Method:  http.MethodPost,
		Path:    "pdu/:ueId/set",
		Handler: OnPduSessResSet,
	},
	{
		Label:   "PduSessResMod",
		Method:  http.MethodPost,
		Path:    "pdu/:ueId/mod",
		Handler: OnPduSessResMod,
	},
	{
		Label:   "PduSessResRel",
		Method:  http.MethodPost,
		Path:    "pdu/:ueId/rel",
		Handler: OnPduSessResRel,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "amf",
		Routes:  _routes,
		Handler: p,
	}
}
