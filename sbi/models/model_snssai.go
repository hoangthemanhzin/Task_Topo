package models

import "fmt"

type Snssai struct {
	Sst int32 `json:"sst"`

	Sd string `json:"sd,omitempty"`
}

func (id *Snssai) String() string {
	return fmt.Sprintf("%d-%s", id.Sst, id.Sd)
}
