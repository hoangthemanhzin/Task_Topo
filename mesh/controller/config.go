package controller

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh/models"
	"fmt"
	"io/ioutil"
	"net"
)

const (
	CONTROLLER_PORT int = 8888
)

type CtrlAddress struct {
	Ip   net.IP
	Port int
}

func (addr *CtrlAddress) UnmarshalJSON(b []byte) (err error) {
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
	addr.Port = CONTROLLER_PORT
	if obj.Port != nil {
		addr.Port = *obj.Port
	}
	return
}

type Config struct {
	Addr      *CtrlAddress     `json:"addr"`
	Heartbeat int              `json:"heartbeat"`
	Services  []models.Service `json:"services"`
	LogLevel  *logctx.LogLevel `json:"loglevel,omitempty"`
}

func LoadConfig(fn string) (cfg Config, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
