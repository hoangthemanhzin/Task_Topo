package config

import (
	"encoding/json"
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"io/ioutil"
	//	"strconv"
	//	"time"
)

type UdrConfig struct {
	UdrName  string
	Mesh     mesh.MeshConfig  `json:"mesh"`
	LogLevel *logctx.LogLevel `json:"loglevel,omitempty"`
}

func LoadConfig(fn string) (cfg UdrConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}
	return
}
