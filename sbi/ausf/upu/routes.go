package upu

import (
	"etrib5gc/sbi"
)

var _routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upu",
		Routes:  _routes,
		Handler: p,
	}
}
