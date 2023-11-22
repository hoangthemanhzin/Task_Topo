package smpc

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{

	{
		Label:   "DeleteSMPolicy",
		Method:  http.MethodPost,
		Path:    "/npcf-smpolicycontrol/v1/sm-policies/:smPolicyId/delete",
		Handler: OnDeleteSMPolicy,
	},
	{
		Label:   "GetSMPolicy",
		Method:  http.MethodGet,
		Path:    "/npcf-smpolicycontrol/v1/sm-policies/:smPolicyId",
		Handler: OnGetSMPolicy,
	},
	{
		Label:   "UpdateSMPolicy",
		Method:  http.MethodPost,
		Path:    "/npcf-smpolicycontrol/v1/sm-policies/:smPolicyId/update",
		Handler: OnUpdateSMPolicy,
	},
	{
		Label:   "CreateSMPolicy",
		Method:  http.MethodPost,
		Path:    "/npcf-smpolicycontrol/v1/sm-policies",
		Handler: OnCreateSMPolicy,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "smpc",
		Routes:  _routes,
		Handler: p,
	}
}
