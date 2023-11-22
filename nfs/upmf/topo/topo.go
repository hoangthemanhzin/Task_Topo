package topo

import (
	"etrib5gc/logctx"
	"etrib5gc/nfs/upmf/config"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"strings"
)

//NOTE: you should move all code related to Topology management
//(currently in the `context` package into this package

const (
	NET_TYPE_AN   uint8 = 0 //connect to RAN nodes
	NET_TYPE_TRAN uint8 = 1 //between two UPFs
	NET_TYPE_DNN  uint8 = 2 //UPF to DN

	PFCP_DEFAULT_IP = "0.0.0.0"
)

// a link between two network interfaces
type Link struct {
	Inf1 *NetInf
	Inf2 *NetInf
	W    uint16
}

// is this link active and can it serve the given slice?
func (l *Link) IsActive(snssai models.Snssai) bool {
	return l.Inf1.Local.IsActive() && l.Inf2.Local.IsActive() &&
		l.Inf1.Local.Serve(snssai) && l.Inf2.Local.Serve(snssai)
}

// Topo maintains a UPF topology
type Topo struct {
	logctx.LogWriter
	//nodes map[string]UpfNode  //list of current UpfNodes
	Slices      map[string]models.Snssai `json:"slices"`
	Nets        map[string]uint8
	Nodes       map[string]*UpfNode
	Links       []Link
	Pfcpid2node map[string]*UpfNode
	Heartbeat   int
	//Add other topology member attributes if needed
}

func NewTopo(cfg *config.TopoConfig) *Topo {
	topo := &Topo{
		LogWriter: logctx.WithFields(logctx.Fields{
			"mod": "topo",
		}),
		//TODO: initialize member attributes if needed
		Nets:        make(map[string]uint8),
		Nodes:       make(map[string]*UpfNode),
		Pfcpid2node: make(map[string]*UpfNode),
		Slices:      make(map[string]models.Snssai),
	}

	for _, name := range cfg.Networks["access"] {
		topo.Nets[name] = NET_TYPE_AN
	}

	for _, name := range cfg.Networks["transport"] {
		topo.Nets[name] = NET_TYPE_TRAN
	}

	for _, name := range cfg.Networks["dnn"] {
		topo.Nets[name] = NET_TYPE_DNN
	}
	topo.Slices = cfg.Slices
	return topo
	//NOTE: config has definition of slices, and definition of networks. Keep
	//thes information in the Topology for interpreting the RegistrationRequest from UPF
}

func (t *Topo) Start() {
	t.Info("Start topology")
	//Start any go-routines here
}

func (t *Topo) Stop() {
	//When application terminates, it calls this method to terminate all
	//go-routines that was inited for UpfNode

	/*
		for _, node := range t.nodes {
			node.terminate() //implement goroutine termination in this method
		}
	*/
	t.Info("Topology stopped")
}

// Get a node's network interfaces to Access networks
func (topo *Topo) GetNodeAnFaces(node *UpfNode, nets []string) (foundinfs []NetInf) {
	for network, infs := range node.Infs {
		if ntype, ok := topo.Nets[network]; ok && ntype == NET_TYPE_AN {
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
func (topo *Topo) GetNodeDnnFaces(node *UpfNode, dnn string) (foundinfs []NetInf) {
	for network, infs := range node.Infs {
		if ntype, ok := topo.Nets[network]; ok && ntype == NET_TYPE_DNN && strings.Compare(network, dnn) == 0 {
			foundinfs = infs
			break
		}
	}
	return
}

func (t *Topo) CreateUpf(req *n42.RegistrationRequest) (err error) {
	//logic to create a new UPF node and add it to the topogy:
	//return an error if the request is invalid. Follow this logic:
	//1. create a new UpfNode (and its go-routine to send heartbeat):
	//Basically, everytime a new UPF register, you need to create a new
	//UpfNode, initialize a
	//go-routine for the node that will send heartbeats to the remote UPF periodically. When there is no response, you should remove the UPF and terminate this go-routine
	//2. add the UpfNode to the Topo
	// t.nodes[newnode.Id()] = newnode
	nodeConfig := NodeConfig{
		UpfId:  "",
		Slices: make([]string, 0),
		Infs:   make(map[string][]NetInfConfig),
	}

	if nodeConfig.Pfcp == nil {
		nodeConfig.Pfcp = new(PfcpConfig) // Khởi tạo PfcpConfig nếu nó là nil
	}
	nodeConfig.UpfId = req.UpfId
	nodeConfig.Slices = req.Slices
	nodeConfig.Pfcp.Ip = req.Ip
	nodeConfig.Pfcp.Port = req.SbiPort

	for key, value := range req.Infs {
		// Check if the key exists in NodeConfig.Infs
		if _, ok := nodeConfig.Infs[key]; !ok {
			// If the key does not exist, create a new slice
			nodeConfig.Infs[key] = []NetInfConfig{}
		}
		// Add the value value to the corresponding slice
		nodeConfig.Infs[key] = append(nodeConfig.Infs[key], convertNetInf(value))
		// t.LogWriter.Infof(fmt.Sprintf("%s --- %s", key, value.Addr))
	}
	//newnode := NewUpfNode(t, req.UpfId, int(req.Time))
	t.ParseUpfNodes(&nodeConfig)

	for _, value := range t.Nodes {
		if value.Infs == nil {
			//t.LogWriter.Info(value.Infs)
		} else {

			t.LogWriter.Info("data of node infs : ", value.Infs)
		}
	}
	//err = fmt.Errorf("Not implement")
	return
}

func convertNetInf(n42Inf n42.NetInfConfig) NetInfConfig {
	// Perform conversion from n422.NetInfConfig to context.NetInfConfig
	convertedInf := NetInfConfig{
		Addr:    n42Inf.Addr,
		DnnInfo: (*DnnInfoConfig)(n42Inf.DnnInfo),
	}
	return convertedInf
}
