package group

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = sbi.SbiRoutes{

	{
		Label:   "GetNfGroupIDs",
		Method:  http.MethodGet,
		Path:    "/nudr-group-id-map/v1/nf-group-ids",
		Handler: OnGetNfGroupIDs,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "group",
		Routes:  _routes,
		Handler: p,
	}
}
