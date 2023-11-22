package topo

import (
	"etrib5gc/common"
	"etrib5gc/nfs/upmf/config"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"etrib5gc/sbi/models/n43"
	"net"
	"testing"
	"time"
)

func Test_Topo(t *testing.T) {
	//create TopoConfig :
	cfg := &config.TopoConfig{
		//TODO: add information :
		Slices: map[string]models.Snssai{
			"slice1": {
				Sd:  "12345",
				Sst: 0,
			},
			"slice2": {
				Sd:  "54321",
				Sst: 1,
			},
		},
		Networks: map[string][]string{
			"access": {
				"an1",
				"an2",
			},
			"transport": {
				"tran",
				"tran2",
				"tran3",
			},
			"dnn": {
				"e1",
				"e2",
				"internet",
			},
		},
	}

	//create empty totpo
	tp := NewTopo(cfg)
	if tp == nil {
		t.Errorf("topo khong co du lieu")
	}

	reqs := []*n42.RegistrationRequest{}

	//create the first UpfNode info:
	//Add UPF1 information :
	reqs = append(reqs, &n42.RegistrationRequest{
		//TODO: add UPF information :
		UpfId:   "upf1",
		Slices:  []string{"slice1"},
		Ip:      common.GetLocalIP().String(),
		SbiPort: 8805,
		Time:    time.Now().Unix(),
		Infs: map[string]n42.NetInfConfig{
			"internet": {
				Addr:    "192.168.10.200",
				DnnInfo: &n42.DnnInfoConfig{Cidr: "10.60.0.0/24"},
			},
			"tran": {
				Addr: "192.168.56.102",
			},
		},
	})

	//TODO: Add more UPF node info to create a topology
	//Add UPF2 information :
	reqs = append(reqs, &n42.RegistrationRequest{
		//TODO: add UPF information :
		UpfId:   "upf2",
		Slices:  []string{"slice1", "slice2"},
		Ip:      common.GetLocalIP().String(),
		SbiPort: 9001,
		Time:    time.Now().Unix(),
		Infs: map[string]n42.NetInfConfig{
			"internet": {
				Addr:    "192.168.10.200",
				DnnInfo: &n42.DnnInfoConfig{Cidr: "10.60.0.0/24"},
			},
			"tran": {
				Addr: "192.168.56.100",
			},
		},
	})

	//Add UPF3 information :
	reqs = append(reqs, &n42.RegistrationRequest{
		//TODO: add UPF information :
		UpfId:   "upf3",
		Slices:  []string{"slice1", "slice2"},
		Ip:      common.GetLocalIP().String(),
		SbiPort: 9002,
		Time:    time.Now().Unix(),
		Infs: map[string]n42.NetInfConfig{
			"an1": {
				Addr: "8.8.8.8",
			},
			"an2": {
				Addr: "10.10.10.10",
			},
			"tran": {
				Addr: "192.168.56.103",
			},
			"tran2": {
				Addr: "192.160.60.103",
			},
		},
	})

	//Add UPF4 information :
	reqs = append(reqs, &n42.RegistrationRequest{
		//TODO: add UPF information :
		UpfId:   "upf4",
		Slices:  []string{"slice1"},
		Ip:      common.GetLocalIP().String(),
		SbiPort: 9003,
		Time:    time.Now().Unix(),
		Infs: map[string]n42.NetInfConfig{
			"an1": {
				Addr: "8.8.8.8",
			},
			"an2": {
				Addr: "10.10.10.10",
			},
			"tran": {
				Addr: "192.168.56.104",
			},
			"tran3": {
				Addr: "192.161.61.104",
			},
		},
	})

	//Add UPF5 information :
	reqs = append(reqs, &n42.RegistrationRequest{
		//TODO: add UPF information :
		UpfId:   "upf5",
		Slices:  []string{"slice1"},
		Ip:      common.GetLocalIP().String(),
		SbiPort: 9004,
		Time:    time.Now().Unix(),
		Infs: map[string]n42.NetInfConfig{
			"e1": {
				Addr:    "12.12.12.12",
				DnnInfo: &n42.DnnInfoConfig{Cidr: "10.10.10.1/24"},
			},
			"tran2": {
				Addr: "192.160.60.105",
			},
		},
	})

	//Add UPF6 information :
	reqs = append(reqs, &n42.RegistrationRequest{
		//TODO: add UPF information :
		UpfId:   "upf6",
		Slices:  []string{"slice1"},
		Ip:      common.GetLocalIP().String(),
		SbiPort: 9005,
		Time:    time.Now().Unix(),
		Infs: map[string]n42.NetInfConfig{
			"e2": {
				Addr:    "10.10.10.10",
				DnnInfo: &n42.DnnInfoConfig{Cidr: "10.10.10.1/24"},
			},
			"tran3": {
				Addr: "192.161.61.106",
			},
		},
	})

	//Now add all the UpfNodes to the Topo
	for _, req := range reqs {
		if err := tp.CreateUpf(req); err != nil {
			t.Errorf("Create Upf failed: %+v", err)
		}
	}

	//create a GetPath querry
	query := &n43.UpfPathQuery{
		//TODO: add query information
		Dnn: "internet",
		Snssai: models.Snssai{
			Sd:  "54321",
			Sst: 1,
		},
		Nets: []string{"an1"},
	}

	//create a GetPath querry 2
	query2 := &n43.UpfPathQuery{
		//TODO: add query information
		Dnn: "e1",
		Snssai: models.Snssai{
			Sd:  "12345",
			Sst: 0,
		},
		Nets: []string{"an2"},
	}
	//create expected return path
	expectedPath := &n43.UpfPath{
		//TODO: add data
		Path: []*n43.PathNode{
			{
				Id:       "upf3",
				UlIp:     net.ParseIP("192.168.56.103"),
				DlIp:     net.ParseIP("8.8.8.8"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9002,
			},
			{
				Id:       "upf2",
				UlIp:     net.ParseIP("192.168.10.200"),
				DlIp:     net.ParseIP("192.168.56.100"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9001,
			},
		},
		Ip: net.ParseIP("10.60.0.1"),
	}

	expectedPath2 := &n43.UpfPath{
		//TODO: add data
		Path: []*n43.PathNode{
			{
				Id:       "upf4",
				UlIp:     net.ParseIP("192.168.56.104"),
				DlIp:     net.ParseIP("8.8.8.8"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9003,
			},
			{
				Id:       "upf2",
				UlIp:     net.ParseIP("192.168.10.200"),
				DlIp:     net.ParseIP("192.168.56.100"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9001,
			},
		},
		Ip: net.ParseIP("10.60.0.1"),
	}

	expectedPathQuery2 := &n43.UpfPath{
		//TODO: add data
		Path: []*n43.PathNode{
			{
				Id:       "upf3",
				UlIp:     net.ParseIP("192.160.60.103"),
				DlIp:     net.ParseIP("10.10.10.10"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9002,
			},
			{
				Id:       "upf5",
				UlIp:     net.ParseIP("12.12.12.12"),
				DlIp:     net.ParseIP("192.160.60.105"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9004,
			},
		},
		Ip: net.ParseIP("10.10.10.1"),
	}

	expectedPathQuery2V2 := &n43.UpfPath{
		//TODO: add data
		Path: []*n43.PathNode{
			{
				Id:       "upf4",
				UlIp:     net.ParseIP("192.168.56.104"),
				DlIp:     net.ParseIP("10.10.10.10"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9003,
			},
			{
				Id:       "upf3",
				UlIp:     net.ParseIP("192.160.60.103"),
				DlIp:     net.ParseIP("192.168.56.103"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9002,
			},
			{
				Id:       "upf5",
				UlIp:     net.ParseIP("12.12.12.12"),
				DlIp:     net.ParseIP("192.160.60.105"),
				PfcpIp:   net.ParseIP("192.168.2.12"),
				PfcpPort: 9004,
			},
		},
		Ip: net.ParseIP("10.10.10.1"),
	}

	if path, err := tp.GetPath(query); err != nil {
		t.Errorf("GetPath failed: %+v", err)
	} else {

		//finally, check if the return path is as you are expected
		if !isMatched(path, expectedPath, t) && !isMatched(path, expectedPath2, t) {
			t.Errorf("Paths not matched")
		}
	}

	if path2, err := tp.GetPath(query2); err != nil {
		t.Errorf("GetPath failed: %+v", err)
	} else {

		//finally, check if the return path is as you are expected
		if !isMatched(path2, expectedPathQuery2, t) && !isMatched(path2, expectedPathQuery2V2, t) {
			t.Errorf("Paths not matched")
		}
	}
}

// func isMatched(p1, p2 *n43.UpfPath) bool {
// 	//TODO: check if two paths are identical
// 	if &p1.Path == &p2.Path {
// 		return true
// 	}
// 	return false
// }

func isMatched(p1, p2 *n43.UpfPath, t *testing.T) bool {
	if p1 == nil || p2 == nil {
		return false
	}

	// Kiểm tra độ dài của Path
	if len(p1.Path) != len(p2.Path) {
		return false
	}

	//Kiểm tra từng phần tử của Path
	for i := range p1.Path {
		if !isPathNodeMatched(p1.Path[i], p2.Path[i], t) {
			return false
		}
	}

	// Kiểm tra IP
	if !p1.Ip.Equal(p2.Ip) {
		return false
	}

	return true
}

func isPathNodeMatched(n1, n2 *n43.PathNode, t *testing.T) bool {
	if n1 == nil || n2 == nil {
		return false
	}

	// Kiểm tra các thuộc tính của PathNode
	if n1.Id != n2.Id || !n1.UlIp.Equal(n2.UlIp) || !n1.DlIp.Equal(n2.DlIp) || !n1.PfcpIp.Equal(n2.PfcpIp) || n1.PfcpPort != n2.PfcpPort {
		t.Errorf("sai %s", n1.DlIp)
		t.Errorf("sai %s", n2.DlIp)
		return false
	}

	return true
}
