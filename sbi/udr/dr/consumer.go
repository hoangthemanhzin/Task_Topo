package dr

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

const (
	SERVICE_PATH = "{apiRoot}/nudr-dr/v2"
)

//mock api
func GetUeSub(client sbi.ConsumerClient, supi string) (sub *models.AuthenticationSubscription, err error) {
	sub = &models.AuthenticationSubscription{}

	sub.AuthenticationMethod = models.AUTHMETHOD__5_G_AKA
	sub.AuthenticationManagementField = "8000"
	sub.EncPermanentKey = "8baf473f2f8fd09487cccbd7097c6862"
	sub.EncOpcKey = "8e27b6af0e692e750f32667a3b14605d"
	sub.Supi = "imsi-208930000000003"
	sub.SequenceNumber = models.SequenceNumber{
		Sqn: "0000000000af",
	}
	return
}
