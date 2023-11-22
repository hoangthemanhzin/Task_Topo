package common

import (
	"fmt"
	"testing"
)

func Test_PlmnId(t *testing.T) {
	fmt.Printf("Test PlmnId\n")
	var plmnid _PlmnId
	var mcc uint16 = 1000
	var mnc uint16 = 500
	if err := plmnid.Set(mcc, mnc); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		emcc := plmnid.Mcc()
		emnc := plmnid.Mnc()
		if mcc != emcc || mnc != emnc {
			t.Error("Test PlmnId failed")
		}
	}
	if err := plmnid.SetHex("001400"); err != nil {
		t.Errorf("Test PlmnId failed: %s", err.Error())
	} else {
		if plmnid.Mcc() != 1 || plmnid.Mnc() != 1024 {
			t.Error("Test PlmnId (loaded from hex) failed")
		}
	}
}

func Test_AmfId(t *testing.T) {
	fmt.Printf("Test AmfId\n")
	var reg uint8 = 12
	var set uint16 = 135
	var pointer uint8 = 50
	var amfid _AmfId
	if err := amfid.Set(reg, set, pointer); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		ereg := amfid.AmfRegion()
		eset := amfid.AmfSet()
		epointer := amfid.AmfPointer()
		if reg != ereg || set != eset || pointer != epointer {
			t.Error("Test AmfId failed")
		}
	}
}
