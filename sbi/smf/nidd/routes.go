package nidd

import (
	"etrib5gc/sbi"
)

var _routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "nidd",
		Routes:  _routes,
		Handler: p,
	}
}
