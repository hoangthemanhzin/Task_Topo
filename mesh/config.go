package mesh

import (
	"encoding/json"
	"etrib5gc/mesh/registry"
	"fmt"
	"net"
)

const (
	SBI_PORT int = 7888
)

type SbiAddress struct {
	Ip   net.IP
	Port int
}

func (addr *SbiAddress) UnmarshalJSON(b []byte) (err error) {
	var obj struct {
		Ip   *string `json:"ip,omitempty"`
		Port *int    `json:"port,omitempty"`
	}

	if err = json.Unmarshal(b, &obj); err != nil {
		return
	}
	var ip string
	if obj.Ip == nil || len(*obj.Ip) == 0 {
		ip = "0.0.0.0"
	} else {
		ip = *obj.Ip
	}

	if addr.Ip = net.ParseIP(ip); addr.Ip == nil {
		err = fmt.Errorf("failed to parse ip address [%s]", ip)
		return
	}
	addr.Port = SBI_PORT
	if obj.Port != nil {
		addr.Port = *obj.Port
	}
	return
}

func DefaultSbiAddress() *SbiAddress {
	return &SbiAddress{
		Port: SBI_PORT,
		Ip:   net.ParseIP("0.0.0.0"),
	}
}

type MeshConfig struct {
	Sbi      *SbiAddress       `json:"sbi,omitempty"`
	Registry registry.Config   `json:"registry"`
	Labels   map[string]string `json:"labels,omitempty"`
}
