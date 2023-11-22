package common

import (
	"etrib5gc/sbi/models"
	"fmt"
	"net"
	"strings"

	"github.com/free5gc/nas/security"
)

func WrapError(msg string, err error) error {
	return fmt.Errorf("%s: %s", msg, err.Error())
}

func BearerType(access models.AccessType) uint8 {
	if access == models.ACCESSTYPE__3_GPP_ACCESS {
		return security.Bearer3GPP
	} else if access == models.ACCESSTYPE_NON_3_GPP_ACCESS {
		return security.BearerNon3GPP
	} else {
		return security.OnlyOneBearer
	}
}

func IsSliceEqual(s1, s2 *models.Snssai) bool {
	if s1 == nil && s2 == nil {
		return true
	} else if s1 == nil || s2 == nil {
		return false
	}

	return s1.Sst == s2.Sst && strings.Compare(s1.Sd, s2.Sd) == 0
}

func IsPlmnIdEqual(id1, id2 *models.PlmnId) bool {
	if id1 == nil && id2 == nil {
		return true
	} else if id1 == nil || id2 == nil {
		return false
	}

	return strings.Compare(id1.Mnc, id2.Mnc) == 0 && strings.Compare(id1.Mcc, id2.Mcc) == 0
}
func ServingNetworkName(id *models.PlmnId) string {
	//return fmt.Sprintf("5G:mnc%03x.mcc%03x.3gppnetwork.org", id.Mnc, id.Mcc)
	if len(id.Mnc) == 2 {
		return fmt.Sprintf("5G:mnc0%s.mcc%s.3gppnetwork.org", id.Mnc, id.Mcc)
	} else {
		return fmt.Sprintf("5G:mnc%s.mcc%s.3gppnetwork.org", id.Mnc, id.Mcc)
	}
}

func Prob2Err(prob *models.ProblemDetails) error {
	return fmt.Errorf("%d: %s", prob.Status, prob.Detail)
}

func SmContextRef(supi string, sid uint32) string {
	return fmt.Sprintf("supi=%s-sid=%d", supi, sid)
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
