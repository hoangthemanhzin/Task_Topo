package context

import (
	"crypto/sha256"
	"encoding/hex"
	"etrib5gc/logctx"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"strings"

	"etrib5gc/util/sec"
)

type UeContext struct {
	logctx.LogWriter
	id       string //supi or suci
	supi     string
	network  string
	kausf    []byte
	kseaf    []byte
	rand     []byte
	xresstar []byte
	autn     []byte
	authtype models.AuthType
	udmcli   sbi.ConsumerClient
}

func newUeContext(suciORsupi string, snName string) *UeContext {
	return &UeContext{
		LogWriter: log.WithFields(logctx.Fields{
			"ueid": suciORsupi,
		}),
		id:      suciORsupi,
		network: snName,
	}
}
func (ue *UeContext) update(info *models.AuthenticationInfoResult) (err error) {
	ue.authtype = info.AuthType
	var av AuthVector
	if err = av.decode(&info.AuthenticationVector); err != nil {
		return
	}

	if info.AuthType == models.AUTHTYPE__5_G_AKA {
		err = ue.update5gAka(&av)
	} else if info.AuthType == models.AUTHTYPE_EAP_AKA_PRIME {
		err = ue.updateEapAka(&av)
	} else {
		//unknown type
		err = fmt.Errorf("Not supported authentication method %d", info.AuthType)
	}
	if err == nil {
		ue.supi = info.Supi
	}
	return
}

func (ue *UeContext) Udm() sbi.ConsumerClient {
	return ue.udmcli
}

func (ue *UeContext) update5gAka(av *AuthVector) (err error) {
	// Derive Kseaf from Kausf
	var kseaf []byte
	P0 := []byte(ue.network)
	if kseaf, err = sec.SeafKey(av.kausf, P0); err != nil {
		ue.Errorf("GetKDFValue failed: %+v", err)
		return
	}
	//log.Info("KSEAF is generated")
	ue.xresstar = av.xresstar
	ue.kausf = av.kausf
	ue.kseaf = kseaf
	ue.rand = av.rand
	ue.autn = av.autn
	return
}

func (ue *UeContext) updateEapAka(info *AuthVector) (err error) {
	panic("eap-aka is not implemented")
	return
}

func (ue *UeContext) SupiOrSuci() string {
	return ue.id
}
func (ue *UeContext) Supi() string {
	return ue.supi
}

func (ue *UeContext) Kseaf() string {
	return hex.EncodeToString(ue.kseaf)
}

func (ue *UeContext) Rand() string {
	return hex.EncodeToString(ue.rand)
}

func (ue *UeContext) EapId() uint8 {
	return 0
}
func (ue *UeContext) Var5gAuthData() (dat models.UEAuthenticationCtx5gAuthData) {
	dat.Rand = hex.EncodeToString(ue.rand)
	dat.Autn = hex.EncodeToString(ue.autn)
	// Derive HXRES* from XRES*
	h := sha256.Sum256(append(ue.rand, ue.xresstar...))
	dat.HxresStar = hex.EncodeToString(h[16:]) // last 128 bits
	//log.Infof("Send Var5gAuthData to AMF: rand=%s, autn=%s, hash=%s", dat.Rand, dat.Autn, dat.HxresStar)
	return
}

func (ue *UeContext) CheckResStar(resstar string) bool {
	return strings.Compare(resstar, hex.EncodeToString(ue.xresstar)) == 0
}
