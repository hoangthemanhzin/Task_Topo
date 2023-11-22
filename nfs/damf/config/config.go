package config

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/sbi/models"
	"io/ioutil"
)

type Slice2AmfId struct {
	Snssai models.Snssai `json:"snssai"`
	AmfId  string        `json:"amfid"`
}

type DamfConfig struct {
	Id       string           `json:"id"`
	PlmnId   models.PlmnId    `json:"plmnid"`
	Mesh     mesh.MeshConfig  `json:"mesh"`
	AmfMap   []Slice2AmfId    `json:"nssf"`
	LogLevel *logctx.LogLevel `json:"loglevel,omitempty"`
}

func LoadConfig(fn string) (cfg DamfConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
