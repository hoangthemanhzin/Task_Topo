package sm

import (
	"etrib5gc/logctx"
	"net"

	"github.com/free5gc/nas/nasMessage"
)

// utility functions
func toNasIp(ip net.IP, sessiontype uint8) (addr [12]byte, addrlen uint8) {
	log := logctx.WithFields(logctx.Fields{"mod": "sm-util"})
	copy(addr[:], ip)
	switch sessiontype {
	case nasMessage.PDUSessionTypeIPv4:
		log.Trace("convert to nas: type ipv4")
		addrlen = 4 + 1
	case nasMessage.PDUSessionTypeIPv6:
	case nasMessage.PDUSessionTypeIPv4IPv6:
		log.Trace("conver to nas: type ipv6")
		addrlen = 12 + 1
	}
	return
}
