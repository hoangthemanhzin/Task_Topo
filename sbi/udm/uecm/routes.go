package uecm

import (
	"etrib5gc/sbi"
)

var _routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "uecm",
		Routes:  _routes,
		Handler: p,
	}
}
