package ueau

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{
	{
		Label:   "GenerateAuthData",
		Method:  http.MethodPost,
		Path:    "/nudm-ueau/v1/:ueid/security-information/generate-auth-data",
		Handler: OnGenerateAuthData,
	},
	{
		Label:   "ConfirmAuth",
		Method:  http.MethodPost,
		Path:    "/nudm-ueau/v1/:ueid/auth-events",
		Handler: OnConfirmAuth,
	},
	{
		Label:   "DeleteAuth",
		Method:  http.MethodPut,
		Path:    "/nudm-ueau/v1/:ueid/auth-events/:authEventId",
		Handler: OnDeleteAuth,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "ueau",
		Routes:  _routes,
		Handler: p,
	}
}
