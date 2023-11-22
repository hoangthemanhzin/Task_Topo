package models

import (
	"encoding/json"
	"fmt"
)

type PlmnId struct {
	Mcc string `json:"mcc"`

	Mnc string `json:"mnc"`
}

func (id *PlmnId) String() string {
	return fmt.Sprintf("%s-%s", id.Mcc, id.Mnc)
}

func (id *PlmnId) UnmarshalJSON(b []byte) (err error) {
	tmpid := struct {
		Mcc string
		Mnc string
	}{}

	if err = json.Unmarshal(b, &tmpid); err != nil {
		return
	}
	id.Mnc = tmpid.Mnc
	id.Mcc = tmpid.Mcc

	if _, err = PlmnId2Bytes(id); err != nil {
		return
	}
	return
}
