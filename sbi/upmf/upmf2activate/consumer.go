package upmf2activate

import (
	"etrib5gc/sbi"
	"net/http"
)

const SERVICE_PATH string = "upmf2activate"

var _routes = sbi.SbiRoutes{
	{
		Label:   "UpfsActivate",
		Method:  http.MethodPost,
		Path:    "activate",
		Handler: UpfsActivate,
	},{
		Label: 	 "UpfsDeActivate",
		Method:  http.MethodPost,
		Path: 	 "deactivate",
		Handler: UpfsDeActivate,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upmf2activate",
		Routes:  _routes,
		Handler: p,
	}
}