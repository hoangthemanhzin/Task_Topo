package up

import (
	"etrib5gc/sbi"
	"net"
)

type PfcpSession struct {
	upf        *Upf   //the session is on this upf
	localseid  uint64 // session identification in PFCP received messages
	remoteseid uint64 //session identification in PFCP sending messages
	pdrs       []*PDR
}

func newPfcpSession(localseid uint64, upf *Upf) *PfcpSession {
	return &PfcpSession{
		localseid: localseid,
		upf:       upf,
	}
}

func (s *PfcpSession) UdpAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   s.upf.ip,
		Port: int(s.upf.port),
	}
}

func (s *PfcpSession) UpfCli() sbi.ConsumerClient {
	return s.upf.cli
}

func (s *PfcpSession) LocalSeid() uint64 {
	return s.localseid
}
func (s *PfcpSession) SetRemoteSeid(seid uint64) {
	s.remoteseid = seid
}

func (s *PfcpSession) RemoteSeid() uint64 {
	return s.remoteseid
}

func (s *PfcpSession) CreatePdr() (pdr *PDR) {
	pdr = s.upf.createPdr()
	s.pdrs = append(s.pdrs, pdr)
	return
}

func (s *PfcpSession) RemovePdrs() {
	for _, pdr := range s.pdrs {
		s.upf.removePdr(pdr)
	}
	s.pdrs = []*PDR{}
}

/*
//send Session Release
func (s *PfcpSession) Release() (err error) {
	return
}

//send Session Establishment/Modification (if remoteseid exists)
func (s *PfcpSession) Update() (err error) {
	return
}
*/
