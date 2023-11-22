package models

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func PlmnId2Bytes(id *PlmnId) (buf []uint8, err error) {
	if len(id.Mcc) != 3 {
		err = fmt.Errorf("Mcc len must be 3: %s", id.Mcc)
		return
	}
	if len(id.Mnc) != 2 && len(id.Mnc) != 3 {
		err = fmt.Errorf("Mnc len must be 2 or 3: %s", id.Mnc)
		return
	}

	var (
		mcc [3]uint8
		mnc [3]uint8
		tmp int
	)

	mnc[2] = 0x0f

	for i := 0; i < 3; i++ {
		if tmp, err = strconv.Atoi(string(id.Mcc[i])); err != nil {
			return
		}
		mcc[i] = uint8(tmp)
	}
	for i := 0; i < len(id.Mnc); i++ {
		if tmp, err = strconv.Atoi(string(id.Mnc[i])); err != nil {
			return
		}
		mnc[i] = uint8(tmp)
	}

	buf = []uint8{
		(mcc[1] << 4) | mcc[0],
		(mnc[2] << 4) | mcc[2],
		(mnc[1] << 4) | mnc[0],
	}
	return
}

func Bytes2PlmnId(buf []byte) (id *PlmnId, err error) {
	if len(buf) != 3 {
		err = fmt.Errorf("plmnid must be 3-byte length")
		return
	}
	var mcc [3]byte
	var mnc [3]byte
	mcc[0] = buf[0] & 0x0f
	mcc[1] = (buf[0] & 0xf0) >> 4
	mcc[2] = (buf[1] & 0x0f)

	mnc[0] = (buf[2] & 0x0f)
	mnc[1] = (buf[2] & 0xf0) >> 4
	mnc[2] = (buf[1] & 0xf0) >> 4

	tmp := []byte{(mcc[0] << 4) | mcc[1], (mcc[2] << 4) | mnc[0], (mnc[1] << 4) | mnc[2]}

	str := hex.EncodeToString(tmp)
	plmnid := PlmnId{
		Mcc: str[:3],
	}
	if str[5] == 'f' {
		plmnid.Mnc = str[3:5] //discard the last letter
	} else {
		plmnid.Mnc = str[3:]
	}
	id = &plmnid
	return
}
