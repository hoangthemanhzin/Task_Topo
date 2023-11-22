package config

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/sbi/models"
	"fmt"
	"io/ioutil"
	"net"
)

const (
	UPMF_NAME string = "localhost"
	UPMF_PORT int    = 7888
)

type UpmfAddress struct {
	Ip   net.IP
	Port int
}

func (addr *UpmfAddress) String() string {
	return fmt.Sprintf("%s:%d", addr.Ip.String(), addr.Port)
}

func (addr *UpmfAddress) UnmarshalJSON(b []byte) (err error) {
	var obj struct {
		Ip   *string `json:"ip,omitempty"`
		Name *string `json:"name,omitempty"`
		Port *int    `json:"port,omitempty"`
	}

	if err = json.Unmarshal(b, &obj); err != nil {
		return
	}
	if obj.Ip != nil {
		//parse IP
		if addr.Ip = net.ParseIP(*obj.Ip); addr.Ip == nil {
			err = fmt.Errorf("failed to parse ip address [%s]", *obj.Ip)
			return
		}
	} else { //resolve IP address from name
		sname := UPMF_NAME
		if obj.Name != nil {
			sname = *obj.Name
		}
		if addr.Ip, err = name2Ip(sname); err != nil {
			return
		}
	}

	//set port
	addr.Port = UPMF_PORT
	if obj.Port != nil {
		addr.Port = *obj.Port
	}
	return
}

func name2Ip(name string) (ip net.IP, err error) {
	var ips []net.IP
	if ips, err = net.LookupIP(name); err != nil {
		return
	}
	ip = ips[0]
	return
}

func DefaultUpmfAddress() (addr *UpmfAddress, err error) {
	addr = &UpmfAddress{
		Port: UPMF_PORT,
	}
	addr.Ip, err = name2Ip(UPMF_NAME)
	return
}

type DnnInfo struct {
	Dnn  string `json:"dnn"`
	Addr string `json:"addr"`
	Cidr string `json:"cidr"`
}

type IfInfo struct {
	Ip   string `json:"ip"`
	Mtu  uint32 `json:"mtu"`
	Type string `json:"type"` //N3, N9
	Name string `json:"name"` //an, tran
}

type UpfConfig struct {
	PlmnId   models.PlmnId    `json:"plmnid"`
	Upmf     *UpmfAddress     `json:"upmf,omitempty"`
	Id 		 string			  `json:"upfid"`
	Slices    []string		  `json:"slices"`
	DnnList  []DnnInfo        `json:"dnnlist"`
	IfList   []IfInfo         `json:"iflist"`
	Ip       string           `json:"ip,omitempty"`
	SbiPort  int              `json:"sbiport,omitempty"`
	LogLevel *logctx.LogLevel `json:"loglevel,omitempty"`

}

func LoadConfig(fn string) (cfg UpfConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
