package n43

import (
	"encoding/json"
	"net"
	"strings"
	"testing"
)

type IpStruct struct {
	Ip1 net.IP
	Ip2 net.IP
}

func (obj *IpStruct) Compare(other *IpStruct) bool {
	b := strings.Compare(obj.Ip1.String(), other.Ip1.String()) == 0
	b = b && strings.Compare(obj.Ip2.String(), other.Ip2.String()) == 0
	return b
}
func Test_IpAddress(t *testing.T) {
	obj := IpStruct{
		Ip1: net.ParseIP("192.168.0.3"),
		//Ip2: net.ParseIP("192.168.0.4"),
	}
	var err error
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		t.Errorf("Marshal failed: %+v", err)
	} else {
		var newobj IpStruct
		if err = json.Unmarshal(buf, &newobj); err != nil {
			t.Errorf("Unmarshal failed: %+v", err)
		} else if !obj.Compare(&newobj) {
			t.Errorf("Not match")
		}
	}
}
