package mt

import (
	"etrib5gc/sbi"
)

var Routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "loc",
		Routes:  Routes,
		Handler: p,
	}
}
