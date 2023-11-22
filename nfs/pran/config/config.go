package config

import (
	"encoding/json"
	"etrib5gc/common"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/sbi/models"
	"io/ioutil"
)

type NgapConfig struct {
	IpList *[]string `json:"iplist,omitempty"`
	Port   *int      `json:"port,omitempty"`
}

type PRanConfig struct {
	Id       string            `json:"id"`
	RanNets  []string          `json:"rannets"` //GnB User Plane interface name
	Ngap     *NgapConfig       `json:"ngap,omitempty"`
	PlmnList []common.PlmnItem `json:"plmnlist"`
	PlmnId   models.PlmnId     `json:"plmnId"`
	Mesh     mesh.MeshConfig   `json:"mesh"`
	AmfList  []string          `json:"amflist"` //list of amfids
	LogLevel *logctx.LogLevel  `json:"loglevel,omitempty"`
}

func LoadConfig(fn string) (cfg PRanConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
