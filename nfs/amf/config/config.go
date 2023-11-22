package config

import (
	"encoding/json"
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/sbi/models"
	"io/ioutil"
)

type AmfConfig struct {
	PlmnId   models.PlmnId         `json:"plmnId"`
	AmfId    string                `json:"amfId"`
	Slices   []models.Snssai       `json:"slices"`
	RanList  []string              `json:"ranlist"`
	Algs     *common.NasSecAlgList `json:"algorithm,omitempty"`
	Mesh     mesh.MeshConfig       `json:"mesh"`
	LogLevel *logctx.LogLevel      `json:"loglevel,omitempty"`
	//	Tacs       []string              `json:"tacs"`
	//	PlmnList   []common.PlmnItem     `json:"plmnlist"`
}

func LoadConfig(fn string) (cfg AmfConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
