package topo

import (
	"etrib5gc/util/ipalloc"
	"net"
)

type InfAddr interface {
	IpAddr() net.IP
}

type ipAddr net.IP

func (addr ipAddr) IpAddr() net.IP {
	return net.IP(addr)
}

type dnnInfo struct {
	ipAddr
	allocator *ipalloc.IpAllocator
}

type NetInf struct {
	id      string      //unique id in a  topo, compose of the node id, the network, and the interface index
	netname string      //network that this face connects to
	nettype uint8       //type of network
	addr    InfAddr     //ipv4 or ipv6
	local   *topoNode   //local node attached to this interface
	remotes []*topoNode //remote nodes that connect to this network interface
}

func (inf *NetInf) isAn() bool {
	return inf.nettype == NET_TYPE_AN
}

func (inf *NetInf) isDnn() bool {
	return inf.nettype == NET_TYPE_DNN
}
