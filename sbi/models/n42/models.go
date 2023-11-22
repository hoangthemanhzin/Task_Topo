package n42

import (
	"encoding/json"
	"etrib5gc/pfcp/pfcpmsg"

	"github.com/free5gc/tlv"
)

//ALL data models between UPF and UPMF are defined in this package

type HeartbeatRequest struct {
	Nonce int64 //for matching request/response
	Msg   pfcpmsg.HeartbeatRequest
}

func (m *HeartbeatRequest) MarshalJSON() (buf []byte, err error) {
	//encode pfcp message first
	var msgbuf []byte
	if msgbuf, err = tlv.Marshal(&m.Msg); err != nil {
		return
	}
	//compose a json obj for json marshalling
	tmp := struct {
		Nonce int64
		Buf   []byte
	}{
		Nonce: m.Nonce,
		Buf:   msgbuf,
	}
	buf, err = json.Marshal(&tmp)
	return
}

func (m *HeartbeatRequest) UnmarshalJSON(buf []byte) (err error) {
	//decode to a json struct
	var tmp struct {
		Nonce int64
		Buf   []byte
	}

	if err = json.Unmarshal(buf, &tmp); err != nil {
		return
	}

	//then decode the pfcp message
	if err = tlv.Unmarshal(tmp.Buf, &m.Msg); err != nil {
		return
	}

	return
}

type HeartbeatResponse struct {
	Nonce int64 //must be equal the value in request
	Msg   pfcpmsg.HeartbeatResponse
}

func (m *HeartbeatResponse) MarshalJSON() (buf []byte, err error) {
	//encode pfcp message first
	var msgbuf []byte
	if msgbuf, err = tlv.Marshal(&m.Msg); err != nil {
		return
	}
	//compose a json obj for json marshalling
	tmp := struct {
		Nonce int64
		Buf   []byte
	}{
		Nonce: m.Nonce,
		Buf:   msgbuf,
	}
	buf, err = json.Marshal(&tmp)
	return
}

func (m *HeartbeatResponse) UnmarshalJSON(buf []byte) (err error) {
	//decode to a json struct
	var tmp struct {
		Nonce int64
		Buf   []byte
	}

	if err = json.Unmarshal(buf, &tmp); err != nil {
		return
	}

	//then decode the pfcp message
	if err = tlv.Unmarshal(tmp.Buf, &m.Msg); err != nil {
		return
	}

	return
}
