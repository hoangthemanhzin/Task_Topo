package n43

import (
	"encoding/json"
	"fmt"
	"net"
)

type IpAddress struct {
	net.IP
}

func (ip *IpAddress) MarshalJSON() (buf []byte, err error) {
	if len(ip.IP) > 0 {
		buf, err = json.Marshal(ip.String())
	} else {
		buf = []byte{}
	}
	return
}

func (ip *IpAddress) UnmarshalJSON(buf []byte) (err error) {
	var ipstr string
	if err = json.Unmarshal(buf, &ipstr); err != nil {
		return
	}
	if len(ipstr) == 0 {
		return
	}
	if parsedIp := net.ParseIP(ipstr); parsedIp == nil {
		err = fmt.Errorf("Invalid IP: %s[%d]", ipstr, len(ipstr))
		return
	} else {
		ip.IP = parsedIp
	}
	return
}
