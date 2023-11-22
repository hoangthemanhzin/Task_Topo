package producer

import (
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"net/http"
)

// var topo_test context.UpfTopo
//var topoConfig context.TopoConfig //NOTE (tungtq): you will not need a
//TopoConfig

//var nodeConfig context.NodeConfig

func (prod *Producer) HandleRegistration(req *n42.RegistrationRequest) (rsp *n42.RegistrationResponse, prob *models.ProblemDetails) {
	prod.Infof("Receive RegistrationRequest from UPF")

	var err error

	//create a new UpfNode then add to topology
	if err = prod.topology.CreateUpf(req); err != nil {
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}
	} else {
		//Create a response
		rsp = &n42.RegistrationResponse{
			//TODO: add attributes:
			Time: req.Time,
			Status: http.StatusCreated,
		}
	}

	/*
			nodeConfig := context.NodeConfig{
				UpfId:  "",
				Slices: make([]string, 0),
				Infs:   make(map[string][]context.NetInfConfig),
			}

			topo_test := context.UpfTopo{
				Nets:        make(map[string]uint8),
				Nodes:       make(map[string]*context.TopoNode),
				Links:       make([]context.Link, 0),
				Pfcpid2node: make(map[string]*context.TopoNode),
				Heartbeat:   0,
			}
			if nodeConfig.Pfcp == nil {
				nodeConfig.Pfcp = new(context.PfcpConfig) // Khởi tạo PfcpConfig nếu nó là nil
			}

		topo_test.Load(&topoConfig)
		prod.Infof("Register upf : %s", nodeConfig.UpfId)
		prod.Infof("du lieu cua topo : %s", prod.topo)
		// @ManhHT sua lai o day :
		prod.topo = topo_test
		prod.Infof("du lieu cua topo da sua lai: %s", prod.topo.Nodes)

		for _, value := range prod.topo.Nodes {
			prod.Info("Log du lieu node topo : %s", value.GetNetInf())
		}
		rsp = &n42.RegistrationResponse{
			Time: time.Now().UnixNano(),
		}
			var jsonData []byte
			if jsonData, err = ioutil.ReadFile("nfs/upmf/context/topoconfig.json"); err != nil {
				// log
				return
			}
			// Chuyển đổi dữ liệu JSON thành biến cấu trúc TopoConfig
			if err := json.Unmarshal(jsonData, &topoConfig); err != nil {
				return nil, prob
			}

			//var Infs map[string][]context.NetInfConfig
			//var topo_test context.UpfTopo
			nodeConfig.UpfId = req.UpfId
			nodeConfig.Slices = req.Slices
			nodeConfig.Pfcp.Ip = req.Ip
			nodeConfig.Pfcp.Port = req.PfcpPort
			prod.Infof("Xu lys duoc vong lap")
			// req.Infs chứa dữ liệu bạn muốn thêm vào NodeConfig.Infs
			for key, value := range req.Infs {
				// Kiểm tra xem key đã tồn tại trong NodeConfig.Infs chưa
				if _, ok := nodeConfig.Infs[key]; !ok {
					// Nếu key chưa tồn tại, tạo một slice mới
					nodeConfig.Infs[key] = []context.NetInfConfig{}
				}
				//convertNetInf(value)
				// Thêm giá trị value vào slice tương ứng

				nodeConfig.Infs[key] = append(nodeConfig.Infs[key], convertNetInf(value))
			}
			topoConfig.Nodes = make(map[string]context.NodeConfig)
			topoConfig.Nodes[nodeConfig.UpfId] = nodeConfig

			topo_test.Load(&topoConfig)
			prod.Infof("Register upf : %s", nodeConfig.UpfId)
			prod.Infof("du lieu cua topo : %s", prod.topo)
			// @ManhHT sua lai o day :
			prod.topo = topo_test
			prod.Infof("du lieu cua topo da sua lai: %s", prod.topo)
			rsp = &n42.RegistrationResponse{
				Time: time.Now().UnixNano(),
			}
	*/
	//NOTE: Following part for testing communication and message encoding/decoding
	//We send a heartbeat to the UPF using Pfcp message embeded in an SBI
	//message. The response message is an SBI message that embeds a PFCP
	//response message
	/*
		upfaddr := fmt.Sprintf("%s:%d", req.Ip, req.SbiPort)
		prod.Infof("Send Heartbeat to %s", upfaddr)

		upfcli := httpdp.NewClientWithAddr(upfaddr)
		msg := n42.HeartbeatRequest{
			Nonce: rand.Int63(),
			Msg: pfcpmsg.HeartbeatRequest{
				RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
					RecoveryTimeStamp: time.Now(),
				},
			},
		}
		if hbrsp, err := upf2upmf.Heartbeat(upfcli, msg); err != nil {
			prod.Errorf("Send Heartbeat to %s failed: %+v", upfaddr, err)
		} else {
			prod.Infof("Nonce=%d, Time=%s", hbrsp.Nonce, hbrsp.Msg.RecoveryTimeStamp.RecoveryTimeStamp.String())
		}
	*/
	return
}

/*
func convertNetInf(n42Inf n42.NetInfConfig) context.NetInfConfig {
	// Thực hiện chuyển đổi từ n422.NetInfConfig sang context.NetInfConfig
	convertedInf := context.NetInfConfig{
		Addr:    n42Inf.Addr,
		DnnInfo: (*context.DnnInfoConfig)(n42Inf.DnnInfo),
	}
	return convertedInf
}
*/
