package registry

import (
	"encoding/json"
	"etrib5gc/mesh/controller"
	"fmt"
	"net"
)

const (
	AGENT_PORT      int    = 8889
	CONTROLLER_NAME string = "b5gc-ctrl"
)

type AgentAddress struct {
	Ip   net.IP
	Port int
}

func (addr *AgentAddress) UnmarshalJSON(b []byte) (err error) {
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
	}

	//set port
	addr.Port = AGENT_PORT
	if obj.Port != nil {
		addr.Port = *obj.Port
	}
	return
}

func DefaultAgentAddress() *AgentAddress {
	return &AgentAddress{
		Port: AGENT_PORT,
		Ip:   net.ParseIP("0.0.0.0"),
	}
}

type ControllerAddress struct {
	Ip   net.IP
	Port int
}

func (addr *ControllerAddress) String() string {
	return fmt.Sprintf("%s:%d", addr.Ip.String(), addr.Port)
}
func (addr *ControllerAddress) UnmarshalJSON(b []byte) (err error) {
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
		sname := CONTROLLER_NAME
		if obj.Name != nil {
			sname = *obj.Name
		}
		if addr.Ip, err = name2Ip(sname); err != nil {
			return
		}
	}

	//set port
	addr.Port = controller.CONTROLLER_PORT
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

func DefaultControllerAddress() (addr *ControllerAddress, err error) {
	addr = &ControllerAddress{
		Port: controller.CONTROLLER_PORT,
	}
	addr.Ip, err = name2Ip(CONTROLLER_NAME)
	return
}

type Config struct {
	Controller *ControllerAddress `json:"controller,omitempty"`
	Agent      *AgentAddress      `json:"agent,omitempty"`
}
