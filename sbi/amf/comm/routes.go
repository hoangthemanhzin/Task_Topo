package comm

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{
	{
		Label:   "CreateUEContext",
		Method:  http.MethodPut,
		Path:    "/namf-comm/v1/ue-contexts/:ueContextId",
		Handler: OnCreateUEContext,
	},
	{
		Label:   "UEContextTransfer",
		Method:  http.MethodPost,
		Path:    "/namf-comm/v1/ue-contexts/:ueContextId/transfer",
		Handler: OnUEContextTransfer,
	},
	{
		Label:   "N1N2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/namf-comm/v1/ue-contexts/:ueContextId/n1-n2-messages",
		Handler: OnN1N2MessageTransfer,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "comm",
		Routes:  _routes,
		Handler: p,
	}
}
