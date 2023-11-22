package upman

import (
	"etrib5gc/nfs/smf/upman/up"
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/pfcptypes"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"etrib5gc/util/idgen"
	"math"
	"net"
	"strconv"
	"strings"
)

var _seidgen idgen.IdGenerator

func init() {
	_seidgen = idgen.NewIdGenerator(1, math.MaxUint64)
}

func allocateSeid() uint64 {
	return _seidgen.Allocate()
}

func freeSeid(seid uint64) {
	_seidgen.Free(seid)
}

type UpTunnel struct {
	sender   pfcp.PfcpSender
	upmfcli  sbi.ConsumerClient
	path     UpPath
	ranip    net.IP
	ranteid  uint32
	ueip     net.IP
	sessions map[string]*up.PfcpSession //mapping upf id (ip address) to its pfcp session
}

func newTunnel(sender pfcp.PfcpSender, cli sbi.ConsumerClient, path *n43.UpfPath) (tunnel *UpTunnel) {
	tunnel = &UpTunnel{
		sender:   sender,
		ueip:     path.Ip,
		path:     newUpPath(path, sender),
		sessions: make(map[string]*up.PfcpSession),
	}

	//NOTE: later we may need multiple paths in a tunnel with some are
	//activated. In cases of changing activated paths, some nodes stay put so
	//their pfcp sessions should be remained. Only new nodes need to create
	//pfcp sessions

	for _, node := range tunnel.path {
		seid := allocateSeid()
		node.createSessionAndPdrs(seid)

		//NOTE: an UPF node may have multiple sessions, but there can be only
		//one session belonging to a tunnel
		tunnel.sessions[node.upf.Id()] = node.session
	}
	return
}

func (t *UpTunnel) UeIp() net.IP {
	return t.ueip
}

func (t *UpTunnel) AnTeid() uint32 {
	return t.path[0].dl.teid
}

func (t *UpTunnel) AnIp() net.IP {
	return t.path[0].dl.localip
}

func (t *UpTunnel) AnNode() *UpNode {
	return t.path[0]
}

func (t *UpTunnel) AnchorNode() *UpNode {
	return t.path[len(t.path)-1]
}

func (t *UpTunnel) UpdateRanInfo(ip net.IP, teid uint32) {
	log.Infof("Update RanInfo for tunnel gnB ip = %s, teid = %d for data path", ip.String(), teid)

	t.ranip = ip
	t.ranteid = teid

	//node connecting to access network
	annode := t.path[0] //must not be nil
	ulpdr := annode.ul.pdr

	if ulpdr.FAR.ForwardingParameters.OuterHeaderCreation != nil {
		// Old AN tunnel exists
		ulpdr.FAR.ForwardingParameters.SendEndMarker = true
	}

	ulpdr.FAR.ForwardingParameters.OuterHeaderCreation = new(pfcptypes.OuterHeaderCreation)
	headercreator := ulpdr.FAR.ForwardingParameters.OuterHeaderCreation
	headercreator.OuterHeaderCreationDescription = pfcptypes.OuterHeaderCreationGtpUUdpIpv4
	headercreator.Teid = teid
	headercreator.Ipv4Address = ip.To4()
	ulpdr.FAR.State = up.RULE_UPDATE
}

func (t *UpTunnel) FillPdrs(srule *models.SessionRule, precedence uint32, dnn string) (err error) {
	//log.Info("filling pdr for the tunnel")
	authdefqos := srule.AuthDefQos
	var fqer *up.QER
	plen := len(t.path)
	for i, node := range t.path {
		//TODO: check adding QER
		if fqer, err = node.upf.GetQer(&authdefqos); err != nil {
			return
		}
		fqer.QFI.QFI = uint8(authdefqos.Var5qi)
		fqer.GateStatus = &pfcptypes.GateStatus{
			ULGate: pfcptypes.GateOpen,
			DLGate: pfcptypes.GateOpen,
		}
		fqer.MBR = &pfcptypes.MBR{
			ULMBR: bitrate2kbps(srule.AuthSessAmbr.Uplink),
			DLMBR: bitrate2kbps(srule.AuthSessAmbr.Downlink),
		}

		//downlink to previous node (closer to RAN)
		//so the rules are defined for uplink traffic going into this UPF
		dlpdr := node.dl.pdr
		dlpdr.QER = append(dlpdr.QER, fqer)
		dlpdr.Precedence = precedence
		dlpdr.PDI = up.PDI{
			SourceInterface: pfcptypes.SourceInterface{InterfaceValue: pfcptypes.SourceInterfaceAccess},
			LocalFTeid: &pfcptypes.FTEID{
				V4:          true,
				Ipv4Address: node.dl.localip,
				Teid:        node.dl.teid,
			},
			NetworkInstance: &pfcptypes.NetworkInstance{NetworkInstance: dnn},
			UEIPAddress: &pfcptypes.UEIPAddress{
				V4:          true,
				Ipv4Address: t.ueip.To4(),
			},
		}
		dlpdr.OuterHeaderRemoval = &pfcptypes.OuterHeaderRemoval{
			OuterHeaderRemovalDescription: pfcptypes.OuterHeaderRemovalGtpUUdpIpv4,
		}

		dlfar := dlpdr.FAR
		dlfar.ApplyAction = pfcptypes.ApplyAction{
			Buff: false,
			Drop: false,
			Dupl: false,
			Forw: true,
			Nocp: false,
		}

		facevalue := pfcptypes.DestinationInterfaceCore
		if i == plen-1 { //node is an anchor upf
			facevalue = pfcptypes.DestinationInterfaceSgiLanN6Lan
		}
		dlfar.ForwardingParameters = &up.ForwardingParameters{
			DestinationInterface: pfcptypes.DestinationInterface{
				InterfaceValue: facevalue,
			},
			NetworkInstance: &pfcptypes.NetworkInstance{NetworkInstance: dnn},
		}
		if node.ul.remote != nil { //or i != plen-1; node is not an anchor upf
			dlfar.ForwardingParameters.OuterHeaderCreation = &pfcptypes.OuterHeaderCreation{
				OuterHeaderCreationDescription: pfcptypes.OuterHeaderCreationGtpUUdpIpv4,
				Ipv4Address:                    node.ul.remote.dl.localip,
				Teid:                           node.ul.remote.dl.teid,
			}
		}

		//uplink to the next node (closer to DN)
		//so the rules are defined for downlink traffic going into this UPF
		ulpdr := node.ul.pdr
		ulpdr.QER = append(ulpdr.QER, fqer)
		ulpdr.Precedence = precedence
		if i == plen-1 { //node is an anchor upf
			ulpdr.PDI = up.PDI{
				SourceInterface: pfcptypes.SourceInterface{InterfaceValue: pfcptypes.SourceInterfaceSgiLanN6Lan},
				NetworkInstance: &pfcptypes.NetworkInstance{NetworkInstance: dnn},
				UEIPAddress: &pfcptypes.UEIPAddress{
					V4:          true,
					Sd:          true,
					Ipv4Address: t.ueip.To4(),
				},
			}
		} else {
			ulpdr.OuterHeaderRemoval = &pfcptypes.OuterHeaderRemoval{
				OuterHeaderRemovalDescription: pfcptypes.OuterHeaderRemovalGtpUUdpIpv4,
			}
			ulpdr.PDI = up.PDI{
				SourceInterface: pfcptypes.SourceInterface{InterfaceValue: pfcptypes.SourceInterfaceCore},
				LocalFTeid: &pfcptypes.FTEID{
					V4:          true,
					Ipv4Address: node.ul.localip,
					Teid:        node.ul.teid,
				},
				UEIPAddress: &pfcptypes.UEIPAddress{
					V4:          true,
					Ipv4Address: t.ueip.To4(),
				},
			}
		}
		ulfar := ulpdr.FAR
		if i == 0 { //node is connecting to RAN (node.dl.remote is nil)
			ulfar.ForwardingParameters = &up.ForwardingParameters{
				DestinationInterface: pfcptypes.DestinationInterface{
					InterfaceValue: pfcptypes.DestinationInterfaceAccess,
				},
				NetworkInstance: &pfcptypes.NetworkInstance{
					NetworkInstance: dnn,
				},
			}
			ulfar.ForwardingParameters.OuterHeaderCreation = &pfcptypes.OuterHeaderCreation{
				OuterHeaderCreationDescription: pfcptypes.OuterHeaderCreationGtpUUdpIpv4,
				Ipv4Address:                    t.ranip.To4(),
				Teid:                           t.ranteid,
			}
		} else { //node.dl.remote != nil; not connecting to RAN
			ulfar.ApplyAction = pfcptypes.ApplyAction{
				Buff: false,
				Drop: false,
				Dupl: false,
				Forw: true,
				Nocp: false,
			}
			ulfar.ForwardingParameters = &up.ForwardingParameters{
				DestinationInterface: pfcptypes.DestinationInterface{InterfaceValue: pfcptypes.DestinationInterfaceAccess},
				OuterHeaderCreation: &pfcptypes.OuterHeaderCreation{
					OuterHeaderCreationDescription: pfcptypes.OuterHeaderCreationGtpUUdpIpv4,
					Ipv4Address:                    node.dl.remote.ul.localip,
					Teid:                           node.dl.remote.ul.teid,
				},
			}
		}
	}
	return
}

// deactivate all nodes and send pfcp session release request to UPFs
func (t *UpTunnel) Release() (err error) {
	plen := len(t.path)
	tasks := make([]func() error, plen)
	for i, node := range t.path {
		node.deactivate()
		tasks[i] = node.releaseFn(t.sender) //session release
	}
	//release ue ip
	//TODO: we may need to request upmf to release the UE's allocated IP
	// req := n43.ReleaseIpRequest{}
	// upmf2smf.ReleaseIp(t.upmfcli, req)
	//execute pfcp session release in a batch pararelly
	return batchExecution(tasks)

}

// send Pfcp session establishment/modification
func (t *UpTunnel) Update() (err error) {
	plen := len(t.path)
	tasks := make([]func() error, plen)
	for i, node := range t.path {
		tasks[i] = node.updateFn(t.sender) //session establishment/modification
	}
	//execute pfcp session release in a batch pararelly
	return batchExecution(tasks)

}

func (t *UpTunnel) Path() UpPath {
	return t.path
}

func (t *UpTunnel) DnsIpv4Addr() net.IP {
	//TODO: get from tunnel
	return []byte{}
}
func (t *UpTunnel) DnsIpv6Addr() net.IP {
	//TODO: get from tunnel
	return []byte{}
}
func (t *UpTunnel) PcscfIpv4Addr() net.IP {
	//TODO: get from tunnel
	return []byte{}
}

func (t *UpTunnel) Ipv4LinkMtu() uint16 {
	return 1400
}

func batchExecution(tasks []func() error) (err error) {
	len := len(tasks)
	errch := make(chan error, len)
	for _, task := range tasks {
		go func(ch chan error) {
			ch <- task()
		}(errch)
	}
	for i := 0; i < len; i++ {
		if err = <-errch; err != nil {
			//return immediately at the first error detected
			return
		}
	}
	return
}
func bitrate2kbps(bitrate string) (kbps uint64) {
	s := strings.Split(bitrate, " ")

	var value int

	if n, err := strconv.Atoi(s[0]); err != nil {
		return
	} else {
		value = n
	}

	switch s[1] {
	case "bps":
		kbps = uint64(value / 1000)
	case "Kbps":
		kbps = uint64(value * 1)
	case "Mbps":
		kbps = uint64(value * 1000)
	case "Gbps":
		kbps = uint64(value * 1000000)
	case "Tbps":
		kbps = uint64(value * 1000000000)
	}
	return
}
