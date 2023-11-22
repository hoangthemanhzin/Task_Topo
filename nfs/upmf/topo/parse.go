package topo

import (
	"etrib5gc/sbi/models"
	"etrib5gc/util/ipalloc"
	"fmt"
	"net"
	"strings"
)

func (topo *Topo) Load(config *TopoConfig) {
	topo.Heartbeat = config.Pfcp.Heartbeat

	for _, name := range config.Networks.Access {
		topo.Nets[name] = NET_TYPE_AN
	}

	for _, name := range config.Networks.Transport {
		topo.Nets[name] = NET_TYPE_TRAN
	}

	for _, name := range config.Networks.Dnn {
		topo.Nets[name] = NET_TYPE_DNN
	}

	topo.ParseNodes(config)
	topo.parseLinks(config)
}

// TODO: check for duplicated IP address at each node config
func (topo *Topo) ParseNodes(config *TopoConfig) {
	for nodeid, nodeinfo := range config.Nodes {
		node := NewUpfNode(topo, nodeid, topo.Heartbeat, nodeinfo.Pfcp, true)
		//add slices
		for _, slice := range nodeinfo.Slices {
			if snssai, ok := config.Slices[slice]; ok {
				node.Slices = append(node.Slices, snssai)
			} else {
				//log.Warnf("slice '%s' is not defined", slice)
			}
		}
		//add interfaces
		var infs []NetInf
		for netname, inf := range nodeinfo.Infs {
			if nettype, ok := topo.Nets[netname]; ok {
				infs = []NetInf{}
				if nettype == NET_TYPE_DNN {
					dnninfolist := parseDnnInfoList(inf)
					for i, addr := range dnninfolist {
						infs = append(infs, NetInf{
							Id:      fmt.Sprintf("%s:%s:%d", node.Id, netname, i),
							Netname: netname,
							Nettype: nettype,
							Addr:    addr,
							Local:   node,
						})
						//log.Infof(fmt.Sprintf("%s:%s:%d", node.id, netname, i))
					}
				} else {
					ipaddrlist := parseIpAddrList(inf)
					for i, addr := range ipaddrlist {
						infs = append(infs, NetInf{
							Id:      fmt.Sprintf("%s:%s:%d", node.Id, netname, i),
							Netname: netname,
							Nettype: nettype,
							Addr:    addr,
							Local:   node,
						})
						//topo.LogWriter.Infof(fmt.Sprintf("%s:%s:%d", node.Id, netname, i))
					}
				}
				node.Infs[netname] = infs
			} else {
				//topo.LogWriter.Warnf("network '%s' is not defined", netname)
			}
		}
		topo.Nodes[nodeid] = node
		if node.hasPfcpIp() {
			topo.Pfcpid2node[node.Pfcpinfo.NodeId()] = node
		}
	}

}

func (topo *Topo) parseLinks(config *TopoConfig) {
	var (
		a, b                 *UpfNode
		inf1, inf2           *NetInf
		ok                   bool
		ntype                uint8
		aindex, bindex       int
		w                    uint16
		linkname1, linkname2 string
		err                  error
	)

	existedlinks := make(map[string]uint16) //mark parsed links

	//loop through all networks to parse links
	for netname, linklist := range config.Links {
		//check if network is defined and it is a transport network
		if ntype, ok = topo.Nets[netname]; !ok || ntype != NET_TYPE_TRAN {
			//log.Warnf("'%s' either not exist or not a transport network", netname)
			continue
		}
		//parse all links in this network
		for _, linkconfig := range linklist {
			if a, ok = topo.Nodes[linkconfig.A.Node]; !ok {
				//topo.LogWriter.Warnf("'%s' not exist", linkconfig.A.Node)
				continue
			}
			if b, ok = topo.Nodes[linkconfig.B.Node]; !ok {
				//topo.LogWriter.Warnf("'%s' not exist", linkconfig.B.Node)
				continue
			}
			//make sure a and b are different nodes :
			if strings.Compare(a.Id, b.Id) == 0 {
				//topo.LogWriter.Warnf("link with duplicated endpoint '%s'", a.Id)
				continue
			}
			//get interface indexes :
			aindex = 0
			if linkconfig.A.Index != nil {
				aindex = *linkconfig.A.Index
			}
			bindex = 0
			if linkconfig.B.Index != nil {
				bindex = *linkconfig.B.Index
			}
			//get link weight
			w = 1 //default link weight
			if linkconfig.W != nil {
				w = *linkconfig.W
			}

			if inf1, err = a.getInf(netname, aindex); err != nil {
				//topo.LogWriter.Warnf("get face from node '%s' fails: %s", a.Id, err.Error())
				continue
			}
			if inf2, err = b.getInf(netname, bindex); err != nil {
				//topo.LogWriter.Warnf("get face from node '%s' fails: %s", b.Id, err.Error())
				continue
			}
			//check if link existed already
			linkname1 = fmt.Sprintf("%s:%s:%d-%s:%s:%d", a.Id, netname, aindex, b.Id, netname, bindex)
			linkname2 = fmt.Sprintf("%s:%s:%d-%s:%s:%d", b.Id, netname, bindex, a.Id, netname, aindex)
			if _, ok := existedlinks[linkname1]; ok {
				//topo.LogWriter.Warnf("link '%s' existed", linkname1)
				continue
			}
			//mark link as existed
			existedlinks[linkname1] = w
			existedlinks[linkname2] = w
			//log.Infof("add links: [%s] (alias[%s]) with weight=%d", linkname1, linkname2, w)
			inf1.Remotes = append(inf1.Remotes, b)
			inf2.Remotes = append(inf2.Remotes, a)
			topo.Links = append(topo.Links, Link{
				Inf1: inf1,
				Inf2: inf2,
				W:    w,
			})
		}
	}
}

func parseIpAddrList(infs []NetInfConfig) (addrlist []IpAddr) {
	for _, info := range infs {
		if ip := net.ParseIP(info.Addr); ip != nil {
			addrlist = append(addrlist, IpAddr(ip))
		} else {
			//.Warnf("parse IP fails '%s'", info.Addr)
		}
	}
	return
}

func parseDnnInfoList(infs []NetInfConfig) (addrlist []*DnnInfo) {
	for _, info := range infs {
		if info.DnnInfo != nil {
			if ip := net.ParseIP(info.Addr); ip != nil {
				if _, ipnet, err := net.ParseCIDR(info.DnnInfo.Cidr); err == nil && ipnet != nil {
					addrlist = append(addrlist, &DnnInfo{
						IpAddr:    IpAddr(ip),
						Allocator: ipalloc.New(ipnet),
					})
				} else if err != nil {
					//log.Warnf("Parse CIDR %s returns error: %s", info.DnnInfo.Cidr, err.Error())
				}
			} else {
				//log.Warnf("parse IP fails '%s'", info.Addr)
			}
		} else {
			//log.Warnf("Dnn at %s has empty cidr", info.Addr)
		}
	}

	return
}

// TODO: check for duplicated IP address at each node config :
func (topo *Topo) ParseUpfNodes(upfNode *NodeConfig) (node *UpfNode) {
	node = NewUpfNode(topo, upfNode.UpfId, upfNode.Pfcp.Heartbeat, upfNode.Pfcp, true)
	//add slices :
	for _, slice := range upfNode.Slices {
		if snssai, ok := topo.Slices[slice]; ok {
			node.Slices = append(node.Slices, snssai)
		} else {
			topo.LogWriter.Warnf("slice '%s' is not defined", slice)
		}
	}
	var infs []NetInf
	for netname, inf := range upfNode.Infs {
		if nettype, ok := topo.Nets[netname]; ok {
			infs = []NetInf{}
			if nettype == NET_TYPE_DNN {
				dnninfolist := parseDnnInfoList(inf)
				for i, addr := range dnninfolist {
					infs = append(infs, NetInf{
						Id:      fmt.Sprintf("%s:%s:%d", node.Id, netname, i),
						Netname: netname,
						Nettype: nettype,
						Addr:    addr,
						Local:   node,
					})
					//topo.LogWriter.Infof(fmt.Sprintf("-------------------%s:%s:%d", node.Id, netname, i))
				}
			} else {
				ipaddrlist := parseIpAddrList(inf)
				for i, addr := range ipaddrlist {
					addInf := NetInf{
						Id:      fmt.Sprintf("%s:%s:%d", node.Id, netname, i),
						Netname: netname,
						Nettype: nettype,
						Addr:    addr,
						Local:   node,
					}
					infs = append(infs, addInf)
					//topo.LogWriter.Infof(fmt.Sprintf("====================%s:%s:%d", node.Id, netname, i))
					topo.GenLink(*node, &addInf)
				}
			}
			node.Infs[netname] = infs
		} else {
			topo.LogWriter.Warnf(fmt.Sprintf("network '%s' is not defined", netname))
		}
	}
	topo.Nodes[upfNode.UpfId] = node
	if node.hasPfcpIp() {
		topo.Pfcpid2node[node.Pfcpinfo.NodeId()] = node
	}
	topo.LogWriter.Infof(fmt.Sprintf("Added node %s", node.Id))
	return
}

func (topo *Topo) GenLink(nodeA UpfNode, netInf *NetInf) {
	if netInf.Nettype != NET_TYPE_TRAN {
		return
	}
	for _, nodeB := range topo.Nodes {
		if nodeA.Id == nodeB.Id {
			continue
		}

		for _, infs := range nodeB.Infs {
			for _, inf := range infs {
				if inf.Nettype != NET_TYPE_TRAN {
					continue
				} else if inf.Netname != netInf.Netname {
					// !checkSameGateWay(netInf.Addr.GetIpAddr(), inf.Addr.GetIpAddr())
					continue
				}
				inf.Remotes = append(inf.Remotes, &nodeA)
				netInf.Remotes = append(netInf.Remotes, nodeB)
				topo.Links = append(topo.Links, Link{
					Inf1: netInf,
					Inf2: &inf,
					W:    1,
				})
				//if slice != nil {
				// topo.Links = append(topo.Links, Link{
				// 	Inf1: netInf,
				// 	Inf2: &inf,
				// 	W:    1,
				// })
				//}
			}
		}
	}
}

func checkSameSlice(slice1 []models.Snssai, slice2 []models.Snssai) *models.Snssai {
	for _, snssai1 := range slice1 {
		for _, snssai2 := range slice2 {
			if snssai1 == snssai2 {
				return &snssai1
			}
		}
	}
	return nil
}

func checkSameGateWay(ip1 net.IP, ip2 net.IP) bool {
	mask1 := ip1.DefaultMask()
	mask2 := ip2.DefaultMask()
	return mask1.String() == mask2.String()
}

func (upf *UpfNode) RemoteUpfNode(upfId string, topo *Topo) {
	delete(topo.Nodes, upfId)

	for sbi, node := range topo.Pfcpid2node {
		if node.Id == upfId {
			delete(topo.Pfcpid2node, sbi)
		}
	}
}
