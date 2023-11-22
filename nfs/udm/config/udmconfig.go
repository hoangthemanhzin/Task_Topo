package config

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/sbi/models"
	"etrib5gc/util/suci"
	"io/ioutil"
)

type UdmConfig struct {
	PlmnId   models.PlmnId    `json:"plmnid"`
	Profiles []suci.Profile   `json:"profiles"`
	Mesh     mesh.MeshConfig  `json:"mesh"`
	LogLevel *logctx.LogLevel `json:"loglevel,omitempty"`
}

func LoadConfig(fn string) (cfg UdmConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
