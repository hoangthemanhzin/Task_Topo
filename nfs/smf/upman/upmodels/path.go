package upmodels

import (
	"etrib5gc/sbi/models"
	"net"
	"strings"
)

// this package contains data models for communication between a Upf topology
// manager (server) and its clients (Smf)

// represent a query to search for a data path to serve a given pdu session
type PathQuery struct {
	Dnn    string        //to locate anchor UPF
	Snssai models.Snssai //all UPFs must serve this snssai
	Nets   []string      //names of UP networks that connect to RAN UP (to find the first RAN-connected UPF)
}

type PathNode struct {
	Id       string //Upf Identity (read from its configuration)
	UlIp     net.IP //ip of the uplink face
	DlIp     net.IP //ip of the downlink face
	PfcpIp   net.IP
	PfcpPort int
}

func (n *PathNode) String() string {
	//return fmt.Sprintf("%s_%s_%s", n.dlip, n.id, n.ulip)
	return n.Id
}

//represent a path consisting of Nodes in user plane from access network to
//data network

type DataPath struct {
	Path []*PathNode //UP nodes in the path
	Ip   net.IP      //Ue IP allocated by the anchor node
	//Deallocator func(net.IP) //for release Ue IP
}

func (p *DataPath) String() string {
	nodes := []string{}
	for _, n := range p.Path {
		nodes = append(nodes, n.String())
	}
	return strings.Join(nodes, " ")
}

// func (p *DataPath) ReleaseIp() {
// 	p.Deallocator(p.Ip)
// }
