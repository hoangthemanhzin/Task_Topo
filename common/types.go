package common

import (
	"encoding/hex"
	"encoding/json"
	"etrib5gc/sbi/models"
	"fmt"
	"net"

	"github.com/free5gc/nas/security"
)

type Endpoint struct {
	Ip   net.IP `json:"ip"`
	Port int    `json:"port"`
}

func (o *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", o.Ip.String(), o.Port)
}

func (o *Endpoint) MarshalJSON() (b []byte, err error) {
	obj := make(map[string]interface{})
	obj["port"] = o.Port
	obj["ip"] = o.Ip.String()
	b, err = json.Marshal(obj)
	return
}

func (o *Endpoint) UnmarshalJSON(b []byte) (err error) {
	var obj struct {
		Ip   string
		Port int
	}

	if err = json.Unmarshal(b, &obj); err != nil {
		return
	}
	if len(obj.Ip) == 0 {
		obj.Ip = "0.0.0.0"
	}
	if o.Ip = net.ParseIP(obj.Ip); o.Ip == nil {
		err = fmt.Errorf("failed to parse ip address [%s], set to empty", obj.Ip)
	}
	o.Port = obj.Port
	return
}

type Pfcp Endpoint

//type Sbi Endpoint

type NasSecAlgList struct {
	IntegrityOrder []byte `json:"intorder"`
	CipheringOrder []byte `json:"ciporder"`
}

var DefaultNasSecAlgs NasSecAlgList = NasSecAlgList{
	IntegrityOrder: []byte{
		//security.AlgIntegrity128NIA0,
		security.AlgIntegrity128NIA1,
		security.AlgIntegrity128NIA2,
		security.AlgIntegrity128NIA3,
	},
	CipheringOrder: []byte{
		//security.AlgCiphering128NEA0,
		security.AlgCiphering128NEA1,
		security.AlgCiphering128NEA2,
		security.AlgCiphering128NEA3,
	},
}

type PlmnItem struct {
	PlmnId models.PlmnId   `json:"plmnId"`
	Slices []models.Snssai `json:"slices"`
}

type AmfId struct {
	Region  uint8
	Set     uint16
	Pointer uint8
}

func (id *AmfId) SetHex(str string) (err error) {
	var buf [3]byte
	if err = loadHex(buf[:], str); err != nil {
		return
	}
	id.Region = buf[0]
	id.Set = uint16(buf[1])<<2 + (uint16(buf[2])&0x00c0)>>6
	id.Pointer = buf[2] & 0x3f
	return
}

func (id *AmfId) Bytes() (b []byte, err error) {
	var buf [3]byte
	if id.Set >= MAX_10_BITS {
		err = fmt.Errorf("AmfSet must be a 10-bit number")
		return
	}

	if id.Pointer >= MAX_6_BITS {
		err = fmt.Errorf("AmfPointer must be a 6-bit number")
		return
	}
	buf[0] = id.Region
	buf[1] = uint8(id.Set>>2) & 0xff
	buf[2] = uint8(id.Set&0x03) + id.Pointer&0x3f
	b = buf[:]
	return
}
func (id *AmfId) String() string {
	if b, err := id.Bytes(); err == nil {
		return hex.EncodeToString(b)
	}
	return ""
}

/*
type Guami struct {
	PlmnId
	AmfId
}

func (id Guami) Bytes() (b []byte, err error) {
		var amfid _AmfId
		if err = amfid.Set(id.AmfId.Region, id.AmfId.Set, id.AmfId.Pointer); err != nil {
			return
		}
	var plmnid _PlmnId
	if err = plmnid.Set(id.PlmnId.Mcc, id.PlmnId.Mnc); err != nil {
		return
	}
	var guami _Guami
	//guami.Set(plmnid, amfid)
	b = guami[:]
	return
}

type Snssai struct {
	Sst uint8
	Sd  uint32
}

func (id Snssai) SetHex(str string) (err error) {
	var _id _Snssai
	if err = _id.SetHex(str); err != nil {
		return
	}
	id.Sst = _id.Sst()
	id.Sd = _id.Sd()
	return
}

func (id Snssai) Bytes() []byte {
	var _id _Snssai
	_id.Set(id.Sst, id.Sd)
	return _id[:]
}

type Guti struct {
	Guami
	Tmsi
}

func (id Guti) Bytes() (b []byte, err error) {
	if guami, err := id.Guami.Bytes(); err == nil {
		b = make([]byte, 10)
		copy(b, guami[:])
		copy(b[len(guami):], id.Tmsi[:])
	}
	return
}

func (id Guti) String() (str string) {
	if b, err := id.Bytes(); err == nil {
		str = hex.EncodeToString(b[:])
	}
	return
}
*/
