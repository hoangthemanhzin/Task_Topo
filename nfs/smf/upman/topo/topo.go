package topo

import (
	"etrib5gc/nfs/smf/upman/upmodels"
	"etrib5gc/pfcp"
	"etrib5gc/sbi/models"
	"etrib5gc/util/dijkstra"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

const (
	NET_TYPE_AN   uint8 = 0 //connect to RAN nodes
	NET_TYPE_TRAN uint8 = 1 //between two UPFs
	NET_TYPE_DNN  uint8 = 2 //UPF to DN

	PFCP_DEFAULT_IP = "0.0.0.0"
)

// a link between two network interfaces
type Link struct {
	inf1 *NetInf
	inf2 *NetInf
	w    uint16
}

// is this link active and can it serve the given slice?
func (l *Link) isActive(snssai models.Snssai) bool {
	return l.inf1.local.isActive() && l.inf2.local.isActive() &&
		l.inf1.local.serve(snssai) && l.inf2.local.serve(snssai)
}

type UpfTopo struct {
	nets        map[string]uint8
	nodes       map[string]*topoNode
	links       []Link
	pfcpid2node map[string]*topoNode

	heartbeat int
}

func NewTopo(config *TopoConfig) (topo *UpfTopo) {
	_initLog()
	topo = newTopo()
	topo.load(config)
	return
}

func newTopo() *UpfTopo {
	return &UpfTopo{
		nets:        make(map[string]uint8),
		nodes:       make(map[string]*topoNode),
		pfcpid2node: make(map[string]*topoNode),
	}
}

func (topo *UpfTopo) Start(pfcp *pfcp.Pfcp) (err error) {
	log.Infof("Start to try connecting to UPFs")
	//create UPFs for all the nodes
	for _, node := range topo.nodes {
		node.start(pfcp)
	}
	return
}

func (topo *UpfTopo) Stop() {
	for _, node := range topo.nodes {
		node.stop()
	}
	//topo.pfcp.Stop()
}

// Get a node's network interfaces to Access networks
func (topo *UpfTopo) getNodeAnFaces(node *topoNode, nets []string) (foundinfs []NetInf) {
	for network, infs := range node.infs {
		if ntype, ok := topo.nets[network]; ok && ntype == NET_TYPE_AN {
			for _, netname := range nets {
				if strings.Compare(netname, network) == 0 {
					foundinfs = append(foundinfs, infs...)
					break
				}
			}
		}
	}
	return
}

// Get a node's network interfaces to Dnn
func (topo *UpfTopo) getNodeDnnFaces(node *topoNode, dnn string) (foundinfs []NetInf) {
	for network, infs := range node.infs {
		if ntype, ok := topo.nets[network]; ok && ntype == NET_TYPE_DNN && strings.Compare(network, dnn) == 0 {
			foundinfs = infs
			break
		}
	}
	return
}
func (topo *UpfTopo) FindPath(query *upmodels.PathQuery) (datapath *upmodels.DataPath) {

	//find all anchors and source nodes for searching(at the same time)
	dnnfaces := []NetInf{} //Net interfaces to Dnn
	srcfaces := []NetInf{} //nodes for start searching
	for _, node := range topo.nodes {
		if node.isActive() && node.serve(query.Snssai) {
			if infs := topo.getNodeDnnFaces(node, query.Dnn); len(infs) > 0 {
				dnnfaces = append(dnnfaces, infs...)
			}

			if infs := topo.getNodeAnFaces(node, query.Nets); len(infs) > 0 { //a starting node
				srcfaces = append(srcfaces, infs...)
			}
		}
	}

	//select an dnn face and allocate ip for UE
	var dnnface *NetInf //selected dnn face
	var ip net.IP
	//	1. first shuffling anchors
	for i := range dnnfaces {
		j := rand.Intn(i + 1)
		dnnfaces[i], dnnfaces[j] = dnnfaces[j], dnnfaces[i]
	}
	//	2. then pick the first one that can allocate an Ip address
	for _, face := range dnnfaces {
		dnninfo := face.addr.(*dnnInfo) //must not panic
		if ip = dnninfo.allocator.Allocate(); ip != nil {
			dnnface = &face
			break
		}
	}
	if dnnface == nil {
		log.Errorf("can't select an anchor to allocate Ue's IP")
		return
	}
	log.Tracef("UE's IP = %s(%d) on Dnn=%s", ip.String(), len(ip), dnnface.netname)
	//build a graph of active links then find the shortest paths from source to
	//destination
	edges := []dijkstra.EdgeInfo{} //edges to build the grap

	//a structure to keep the endpoint's ip addresses of a link
	type edgesig struct {
		ip1 net.IP
		ip2 net.IP
	}

	ipmap := make(map[string]edgesig) //map edge name to a tuple of its endpoint's ip addresses

	for _, l := range topo.links {
		if l.isActive(query.Snssai) { //only pick active links
			edges = append(edges, dijkstra.EdgeInfo{
				A: l.inf1.local.id,
				B: l.inf2.local.id,
				W: int64(l.w),
			})
			log.Tracef("add link %s-%s", l.inf1.local.id, l.inf2.local.id)
			//keep the ip addresses of the edges for later use
			ipmap[fmt.Sprintf("%s-%s", l.inf1.local.id, l.inf2.local.id)] = edgesig{
				ip1: l.inf1.addr.IpAddr(),
				ip2: l.inf2.addr.IpAddr(),
			}
			ipmap[fmt.Sprintf("%s-%s", l.inf2.local.id, l.inf1.local.id)] = edgesig{
				ip1: l.inf2.addr.IpAddr(),
				ip2: l.inf1.addr.IpAddr(),
			}

		}
	}
	graph := dijkstra.New(edges)
	//	1. shuffle sources
	for i := range srcfaces {
		j := rand.Intn(i + 1)
		srcfaces[i], srcfaces[j] = srcfaces[j], srcfaces[i]
	}
	for _, srcface := range srcfaces {
		log.Tracef("Search path from %s to %s", srcface.local.id, dnnface.local.id)
		if _, paths := graph.ShortestPath(srcface.local.id, dnnface.local.id); len(paths) > 0 {
			path := paths[0] //pick the first path

			//build the path with ip address of the faces
			plen := len(path)
			pathnodes := make([]*upmodels.PathNode, plen)
			for i, id := range path {
				pfcpinfo := topo.nodes[id].pfcpinfo
				pathnodes[i] = &upmodels.PathNode{
					Id:       id,
					PfcpIp:   pfcpinfo.Ip,
					PfcpPort: pfcpinfo.Port,
				}
			}
			//set ip addresses for the An face and Dnn face of the path
			pathnodes[0].DlIp = srcface.addr.IpAddr()
			pathnodes[plen-1].UlIp = dnnface.addr.IpAddr()
			//set ip addresses for remaining faces on the path
			for i := 0; i < plen-1; i++ {
				info := ipmap[fmt.Sprintf("%s-%s", path[i], path[i+1])]
				pathnodes[i].UlIp = info.ip1
				pathnodes[i+1].DlIp = info.ip2
			}
			datapath = &upmodels.DataPath{
				Path:        pathnodes,
				Ip:          ip,
				Deallocator: dnnface.addr.(*dnnInfo).allocator.Release,
			}
			break
		}
	}
	return
}
