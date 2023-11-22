package ee

import (
	"etrib5gc/sbi"
)

var Routes = sbi.SbiRoutes{}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "ee",
		Routes:  Routes,
		Handler: p,
	}
}
