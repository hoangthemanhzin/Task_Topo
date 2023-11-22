package config

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/sbi/models"
	"io/ioutil"
)

type SmfConfig struct {
	PlmnId   models.PlmnId    `json:"plmnId"`
	AmfList  []string         `json:"amflist"`
	Slice    models.Snssai    `json:"slice"`
	Dnn      string           `json:"dnn"`
	Mesh     mesh.MeshConfig  `json:"mesh"`
	LogLevel *logctx.LogLevel `json:"loglevel,omitempty"`
}

func LoadConfig(fn string) (cfg SmfConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
