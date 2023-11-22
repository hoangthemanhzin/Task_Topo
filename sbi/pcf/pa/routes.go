package pa

import (
	"etrib5gc/sbi"
)

var _routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "pa",
		Routes:  _routes,
		Handler: p,
	}
}
