package uea

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{

	{
		Label:   "Delete5gAkaAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/nausf-auth/v1/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: OnDelete5gAkaAuthenticationResult,
	},
	{
		Label:   "DeleteEapAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/nausf-auth/v1/ue-authentications/:authCtxId/eap-session",
		Handler: OnDeleteEapAuthenticationResult,
	},
	{
		Label:   "EapAuthMethod",
		Method:  http.MethodPost,
		Path:    "/nausf-auth/v1/ue-authentications/:authCtxId/eap-session",
		Handler: OnEapAuthMethod,
	},
	{
		Label:   "RgAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/nausf-auth/v1/rg-authentications",
		Handler: OnRgAuthenticationsPost,
	},
	{
		Label:   "UeAuthenticationsAuthCtxId5gAkaConfirmationPut",
		Method:  http.MethodPut,
		Path:    "/nausf-auth/v1/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: OnUeAuthenticationsAuthCtxId5gAkaConfirmationPut,
	},
	{
		Label:   "UeAuthenticationsDeregisterPost",
		Method:  http.MethodPost,
		Path:    "/nausf-auth/v1/ue-authentications/deregister",
		Handler: OnUeAuthenticationsDeregisterPost,
	},
	{
		Label:   "UeAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/nausf-auth/v1/ue-authentications",
		Handler: OnUeAuthenticationsPost,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "uea",
		Routes:  _routes,
		Handler: p,
	}
}
