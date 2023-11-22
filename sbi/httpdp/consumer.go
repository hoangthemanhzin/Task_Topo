package httpdp

import (
	"errors"
	"etrib5gc/fabric"
	"etrib5gc/fabric/common"
	fabricdp "etrib5gc/fabric/httpdp"
	"etrib5gc/sbi"
)

// implement the sbi consumer client abstraction
type requestSender struct {
	fw    fabric.Forwarder
	query common.NfQuery
	addr  common.AgentAddr
}

func NewClient(fw fabric.Forwarder, query common.NfQuery) *requestSender {
	return &requestSender{
		fw:    fw,
		query: query,
	}
}

// ask agent's forwarder to send a request, a connection to a selected producer
// is maintained for all requests
func (s *requestSender) Send(request *sbi.Request) (response *sbi.Response, err error) {
	var commonResponse common.Response
	var ok bool
	if s.query == nil {
		panic(errors.New("No destication to send"))
	}
	if s.addr == nil {
		commonResponse, s.addr, err = s.fw.DiscoveryThenSend(request, s.query)
	} else {
		commonResponse, err = s.fw.DirectSend(request, s.addr)
	}
	if err != nil {
		return
	}
	if response, ok = commonResponse.(*sbi.Response); !ok {
		err = errors.New("the response is not in the sbi format")
	}
	return
}

func (s *requestSender) DecodeResponse(resp *sbi.Response) error {
	return fabricdp.Encoding().DecodeResponse(resp)
}
