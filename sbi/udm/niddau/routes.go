package niddau

import (
	"etrib5gc/sbi"
)

var _routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "niddau",
		Routes:  _routes,
		Handler: p,
	}
}
