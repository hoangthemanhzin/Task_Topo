package uepc

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{

	{
		Label:   "DeleteIndividualUEPolicyAssociation",
		Method:  http.MethodDelete,
		Path:    "/npcf-ue-policy-control/v1/policies/:polAssoId",
		Handler: OnDeleteIndividualUEPolicyAssociation,
	},
	{
		Label:   "ReadIndividualUEPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/npcf-ue-policy-control/v1/policies/:polAssoId",
		Handler: OnReadIndividualUEPolicyAssociation,
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualUEPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/npcf-ue-policy-control/v1/policies/:polAssoId/update",
		Handler: OnReportObservedEventTriggersForIndividualUEPolicyAssociation,
	},
	{
		Label:   "CreateIndividualUEPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/npcf-ue-policy-control/v1/policies",
		Handler: OnCreateIndividualUEPolicyAssociation,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "uepc",
		Routes:  _routes,
		Handler: p,
	}
}
