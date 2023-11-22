package sdm

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{

	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/nudm-sdm/v2/{supi}/am-data",
		Handler: OnGetAmData,
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/nudm-sdm/v2/{ueId}/id-translation-result",
		Handler: OnGetSupiOrGpsi,
	},
	{
		Label:   "GetSmfSelData",
		Method:  http.MethodGet,
		Path:    "/nudm-sdm/v2/{supi}/smf-select-data",
		Handler: OnGetSmfSelData,
	},
	{
		Label:   "GetSmData",
		Method:  http.MethodGet,
		Path:    "/nudm-sdm/v2/{supi}/sm-data",
		Handler: OnGetSmData,
	}, {
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/nudm-sdm/v2/{supi}/ue-context-in-amf-data",
		Handler: OnGetUeCtxInAmfData,
	},
	{
		Label:   "GetUeCtxInSmfData",
		Method:  http.MethodGet,
		Path:    "/nudm-sdm/v2/{supi}/ue-context-in-smf-data",
		Handler: OnGetUeCtxInSmfData,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "sdm",
		Routes:  _routes,
		Handler: p,
	}
}
