package config

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/sbi/models"
	"io/ioutil"
)

// definition of slices and networks in the topology
type TopoConfig struct {
	Slices   map[string]models.Snssai `json:"slices"`
	Networks map[string][]string      `json:"networks"`
}

type UpmfConfig struct {
	PlmnId   models.PlmnId    `json:"plmnid"`
	Mesh     mesh.MeshConfig  `json:"mesh"`
	LogLevel *logctx.LogLevel `json:"loglevel,omitempty"`
	Topo     TopoConfig       `json:"topo"`
}

func LoadConfig(fn string) (cfg UpmfConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
