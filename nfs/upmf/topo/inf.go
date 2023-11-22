package topo

import (
	"etrib5gc/util/ipalloc"
	"net"
)

type InfAddr interface {
	GetIpAddr() net.IP
}

type IpAddr net.IP

type DnnInfo struct {
	IpAddr    `json:"ipaddr"`
	Allocator *ipalloc.IpAllocator `json:"allocator"`
}

func (addr IpAddr) GetIpAddr() net.IP {
	return net.IP(addr)
}

// "infs"
type NetInf struct {
	Id      string     //unique id in a  topo, compose of the node id, the network, and the interface index
	Netname string     //network that this face connects to
	Nettype uint8      //type of network
	Addr    InfAddr    //ipv4 or ipv6
	Local   *UpfNode   //local node attached to this interface
	Remotes []*UpfNode `omitempty` //remote nodes that connect to this network interface
}

func (inf *NetInf) IsAn() bool {
	return inf.Nettype == NET_TYPE_AN
}

func (inf *NetInf) IsDnn() bool {
	return inf.Nettype == NET_TYPE_DNN
}
