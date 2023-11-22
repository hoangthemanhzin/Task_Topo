package pdu

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{
	{
		Label:   "ReleaseSmContext",
		Method:  http.MethodPost,
		Path:    "/nsmf-pdusession/v1/sm-contexts/:smContextRef/release",
		Handler: OnReleaseSmContext,
	},
	{
		Label:   "UpdateSmContext",
		Method:  http.MethodPost,
		Path:    "/nsmf-pdusession/v1/sm-contexts/:smContextRef/modify",
		Handler: OnUpdateSmContext,
	},
	{
		Label:   "PostSmContexts",
		Method:  http.MethodPost,
		Path:    "/nsmf-pdusession/v1/sm-contexts",
		Handler: OnPostSmContexts,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "pdu",
		Routes:  _routes,
		Handler: p,
	}
}
