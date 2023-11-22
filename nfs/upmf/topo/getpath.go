package topo

import (
	"etrib5gc/sbi/models/n43"
	"etrib5gc/util/dijkstra"
	"fmt"
	"math/rand"
	"net"
)

func (t *Topo) GetPath(query *n43.UpfPathQuery) (datapath *n43.UpfPath, err error) {
	//find all anchors and source nodes for searching(at the same time)
	dnnfaces := []NetInf{} //Net interfaces to Dnn
	srcfaces := []NetInf{} //nodes for start searching
	for _, node := range t.Nodes {
		if node.IsActive() && node.Serve(query.Snssai) {
			if infs := t.GetNodeDnnFaces(node, query.Dnn); len(infs) > 0 {
				dnnfaces = append(dnnfaces, infs...)
			}

			if infs := t.GetNodeAnFaces(node, query.Nets); len(infs) > 0 { //a starting node
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
		dnninfo := face.Addr.(*DnnInfo) //must not panic
		if ip = dnninfo.Allocator.Allocate(); ip != nil {
			dnnface = &face
			break
		}
	}
	if dnnface == nil {
		//t.Errorf("can't select an anchor to allocate Ue's IP")
		return
	}
	//t.LogWriter.Tracef("UE's IP = %s(%d) on Dnn=%s", ip.String(), len(ip), dnnface.Netname)
	//build a graph of active links then find the shortest paths from source to
	//destination
	edges := []dijkstra.EdgeInfo{} //edges to build the grap

	//a structure to keep the endpoint's ip addresses of a link
	type edgesig struct {
		ip1 net.IP
		ip2 net.IP
	}

	ipmap := make(map[string]edgesig) //map edge name to a tuple of its endpoint's ip addresses

	for _, l := range t.Links {
		if l.IsActive(query.Snssai) { //only pick active links
			edges = append(edges, dijkstra.EdgeInfo{
				A: l.Inf1.Local.Id,
				B: l.Inf2.Local.Id,
				W: int64(l.W),
			})
			//t.LogWriter.Tracef("add link %s-%s", l.Inf1.Local.Id, l.Inf2.Local.Id)
			//keep the ip addresses of the edges for later use
			ipmap[fmt.Sprintf("%s-%s", l.Inf1.Local.Id, l.Inf2.Local.Id)] = edgesig{
				ip1: l.Inf1.Addr.GetIpAddr(),
				ip2: l.Inf2.Addr.GetIpAddr(),
			}
			ipmap[fmt.Sprintf("%s-%s", l.Inf2.Local.Id, l.Inf1.Local.Id)] = edgesig{
				ip1: l.Inf2.Addr.GetIpAddr(),
				ip2: l.Inf1.Addr.GetIpAddr(),
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
		//t.LogWriter.Tracef("Search path from %s to %s", srcface.Local.Id, dnnface.Local.Id)
		if _, paths := graph.ShortestPath(srcface.Local.Id, dnnface.Local.Id); len(paths) > 0 {
			path := paths[0] //pick the first path

			//build the path with ip address of the faces
			plen := len(path)
			pathnodes := make([]*n43.PathNode, plen)
			for i, id := range path {
				pfcpinfo := t.Nodes[id].Pfcpinfo
				pathnodes[i] = &n43.PathNode{
					Id:       id,
					PfcpIp:   pfcpinfo.Ip,
					PfcpPort: pfcpinfo.Port,
				}
			}
			//set ip addresses for the An face and Dnn face of the path
			pathnodes[0].DlIp = srcface.Addr.GetIpAddr()
			pathnodes[plen-1].UlIp = dnnface.Addr.GetIpAddr()
			//set ip addresses for remaining faces on the path
			for i := 0; i < plen-1; i++ {
				info := ipmap[fmt.Sprintf("%s-%s", path[i], path[i+1])]
				pathnodes[i].UlIp = info.ip1
				pathnodes[i+1].DlIp = info.ip2
			}
			// Sua lai o day :
			datapath = &n43.UpfPath{
				Path: pathnodes,
				Ip:   ip,
				// Deallocator: dnnface.addr.(*dnnInfo).allocator.Release,
			}
			break
		}
	}
	//err = fmt.Errorf("Not implement")
	return
}
