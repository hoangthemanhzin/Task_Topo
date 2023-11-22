package topo

import (
	"etrib5gc/util/ipalloc"
	"fmt"
	"net"
	"strings"
)

func (topo *UpfTopo) load(config *TopoConfig) {
	topo.heartbeat = config.Pfcp.Heartbeat

	for _, name := range config.Networks.Access {
		topo.nets[name] = NET_TYPE_AN
	}

	for _, name := range config.Networks.Transport {
		topo.nets[name] = NET_TYPE_TRAN
	}

	for _, name := range config.Networks.Dnn {
		topo.nets[name] = NET_TYPE_DNN
	}

	topo.parseNodes(config)
	topo.parseLinks(config)
}

// TODO: check for duplicated IP address at each node config
func (topo *UpfTopo) parseNodes(config *TopoConfig) {
	for nodeid, nodeinfo := range config.Nodes {
		node := newNode(nodeid, topo.heartbeat, nodeinfo.Pfcp, true)
		//add slices
		for _, slice := range nodeinfo.Slices {
			if snssai, ok := config.Slices[slice]; ok {
				node.slices = append(node.slices, snssai)
			} else {
				log.Warnf("slice '%s' is not defined", slice)
			}
		}
		//add interfaces
		var infs []NetInf
		for netname, inf := range nodeinfo.Infs {
			if nettype, ok := topo.nets[netname]; ok {
				infs = []NetInf{}
				if nettype == NET_TYPE_DNN {
					dnninfolist := parseDnnInfoList(inf)
					for i, addr := range dnninfolist {
						infs = append(infs, NetInf{
							id:      fmt.Sprintf("%s:%s:%d", node.id, netname, i),
							netname: netname,
							nettype: nettype,
							addr:    addr,
							local:   node,
						})
						//log.Infof(fmt.Sprintf("%s:%s:%d", node.id, netname, i))
					}
				} else {
					ipaddrlist := parseIpAddrList(inf)
					for i, addr := range ipaddrlist {
						infs = append(infs, NetInf{
							id:      fmt.Sprintf("%s:%s:%d", node.id, netname, i),
							netname: netname,
							nettype: nettype,
							addr:    addr,
							local:   node,
						})
						//log.Infof(fmt.Sprintf("%s:%s:%d", node.id, netname, i))
					}
				}
				node.infs[netname] = infs
			} else {
				log.Warnf("network '%s' is not defined", netname)
			}
		}
		topo.nodes[nodeid] = node
		if node.hasPfcpIp() {
			topo.pfcpid2node[node.pfcpinfo.NodeId()] = node
		}
	}

}

func (topo *UpfTopo) parseLinks(config *TopoConfig) {
	var (
		a, b                 *topoNode
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
		if ntype, ok = topo.nets[netname]; !ok || ntype != NET_TYPE_TRAN {
			log.Warnf("'%s' either not exist or not a transport network", netname)
			continue
		}
		//parse all links in this network
		for _, linkconfig := range linklist {
			if a, ok = topo.nodes[linkconfig.A.Node]; !ok {
				log.Warnf("'%s' not exist", linkconfig.A.Node)
				continue
			}
			if b, ok = topo.nodes[linkconfig.B.Node]; !ok {
				log.Warnf("'%s' not exist", linkconfig.B.Node)
				continue
			}
			//make sure a and b are different nodes
			if strings.Compare(a.id, b.id) == 0 {
				log.Warnf("link with duplicated endpoint '%s'", a.id)
				continue
			}
			//get interface indexes
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
				log.Warnf("get face from node '%s' fails: %s", a.id, err.Error())
				continue
			}
			if inf2, err = b.getInf(netname, bindex); err != nil {
				log.Warnf("get face from node '%s' fails: %s", b.id, err.Error())
				continue
			}
			//check if link existed already
			linkname1 = fmt.Sprintf("%s:%s:%d-%s:%s:%d", a.id, netname, aindex, b.id, netname, bindex)
			linkname2 = fmt.Sprintf("%s:%s:%d-%s:%s:%d", b.id, netname, bindex, a.id, netname, aindex)
			if _, ok := existedlinks[linkname1]; ok {
				log.Warnf("link '%s' existed", linkname1)
				continue
			}
			//mark link as existed
			existedlinks[linkname1] = w
			existedlinks[linkname2] = w
			//log.Infof("add links: [%s] (alias[%s]) with weight=%d", linkname1, linkname2, w)
			inf1.remotes = append(inf1.remotes, b)
			inf2.remotes = append(inf2.remotes, a)
			topo.links = append(topo.links, Link{
				inf1: inf1,
				inf2: inf2,
				w:    w,
			})
		}
	}
}

func parseIpAddrList(infs []NetInfConfig) (addrlist []ipAddr) {
	for _, info := range infs {
		if ip := net.ParseIP(info.Addr); ip != nil {
			addrlist = append(addrlist, ipAddr(ip))
		} else {
			log.Warnf("parse IP fails '%s'", info.Addr)
		}
	}
	return
}

func parseDnnInfoList(infs []NetInfConfig) (addrlist []*dnnInfo) {
	for _, info := range infs {
		if info.DnnInfo != nil {
			if ip := net.ParseIP(info.Addr); ip != nil {
				if _, ipnet, err := net.ParseCIDR(info.DnnInfo.Cidr); err == nil && ipnet != nil {
					addrlist = append(addrlist, &dnnInfo{
						ipAddr:    ipAddr(ip),
						allocator: ipalloc.New(ipnet),
					})
				} else if err != nil {
					log.Warnf("Parse CIDR %s returns error: %s", info.DnnInfo.Cidr, err.Error())
				}
			} else {
				log.Warnf("parse IP fails '%s'", info.Addr)
			}
		} else {
			log.Warnf("Dnn at %s has empty cidr", info.Addr)
		}
	}

	return
}
