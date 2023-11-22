package n41

import (
	"encoding/json"
	"etrib5gc/pfcp/pfcpmsg"
	"fmt"

	"github.com/free5gc/tlv"
)

type SessionEstablishmentRequest struct {
	Msg *pfcpmsg.PFCPSessionEstablishmentRequest
}

func (m *SessionEstablishmentRequest) MarshalJSON() (buf []byte, err error) {
	//encode pfcp message first
	var msgbuf []byte
	if msgbuf, err = tlv.Marshal(m.Msg); err != nil {
		err = fmt.Errorf("TLV failed: %+v", err)
		return
	}
	//compose a json obj for json marshalling
	tmp := struct {
		Buf []byte
	}{
		Buf: msgbuf,
	}
	buf, err = json.Marshal(&tmp)
	return
}

func (m *SessionEstablishmentRequest) UnmarshalJSON(buf []byte) (err error) {
	//decode to a json struct
	var tmp struct {
		Buf []byte
	}

	if err = json.Unmarshal(buf, &tmp); err != nil {
		return
	}

	//then decode the pfcp message
	m.Msg = &pfcpmsg.PFCPSessionEstablishmentRequest{}
	if err = tlv.Unmarshal(tmp.Buf, m.Msg); err != nil {
		return
	}
	return
}

type SessionEstablishmentResponse struct {
	Seid uint64
	Msg  *pfcpmsg.PFCPSessionEstablishmentResponse
}

func (m *SessionEstablishmentResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = tlv.Marshal(m.Msg)
	return
}

func (m *SessionEstablishmentResponse) UnmarshalJSON(buf []byte) (err error) {
	var msg pfcpmsg.PFCPSessionEstablishmentResponse
	m.Msg = &msg
	err = tlv.Unmarshal(buf, &msg)
	return
}

type SessionModificationRequest struct {
	Msg *pfcpmsg.PFCPSessionModificationRequest
}

func (m *SessionModificationRequest) MarshalJSON() (buf []byte, err error) {
	buf, err = tlv.Marshal(m.Msg)
	return
}

func (m *SessionModificationRequest) UnmarshalJSON(buf []byte) (err error) {
	var msg pfcpmsg.PFCPSessionModificationRequest
	m.Msg = &msg
	err = tlv.Unmarshal(buf, &msg)
	return
}

type SessionModificationResponse struct {
	Seid uint64
	Msg  *pfcpmsg.PFCPSessionModificationResponse
}

func (m *SessionModificationResponse) MarshalJSON() (buf []byte, err error) {
	buf, err = tlv.Marshal(m.Msg)
	return
}

func (m *SessionModificationResponse) UnmarshalJSON(buf []byte) (err error) {
	var msg pfcpmsg.PFCPSessionModificationResponse
	m.Msg = &msg
	err = tlv.Unmarshal(buf, &msg)
	return
}

type SessionDeletionRequest struct {
	Msg *pfcpmsg.PFCPSessionDeletionRequest
}

type SessionDeletionResponse struct {
	Seid uint64
	Msg  *pfcpmsg.PFCPSessionDeletionResponse
}
