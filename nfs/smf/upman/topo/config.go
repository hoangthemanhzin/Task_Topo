package topo

import "etrib5gc/sbi/models"

type TopoConfig struct {
	Pfcp PfcpConfig `json:"pfcp"`

	Networks NetConfig                `json:"networks"`
	Nodes    map[string]NodeConfig    `json:"nodes"`
	Links    map[string][]LinkConfig  `json:"links"`
	Slices   map[string]models.Snssai `json:"slices"`
}

type NetConfig struct {
	Access    []string `json:"access"`
	Transport []string `json:"transport"`
	Dnn       []string `json:"dnn"`
}

type NodeConfig struct {
	Slices []string                  `json:"slices"`
	Infs   map[string][]NetInfConfig `json:"infs"`
	Pfcp   *PfcpConfig               `json:"pfcp,omitempty"`
}

type PfcpConfig struct {
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	Heartbeat int    `json:"heartbeat,omitempty"`
}

type NetInfConfig struct {
	Addr    string         `json:"addr"`
	DnnInfo *DnnInfoConfig `json:"dnninfo,omitempty"`
}

type DnnInfoConfig struct {
	Cidr string `json:"cidr"`
}

type LinkConfig struct {
	A LinkEndpointConfig `json:"a"`
	B LinkEndpointConfig `json:"b"`
	W *uint16            `json:"w,omitempty"`
}
type LinkEndpointConfig struct {
	Node  string `json:"node"`
	Index *int   `json:"index,omitempty"`
}
