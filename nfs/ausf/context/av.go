package context

import (
	"encoding/hex"
	"etrib5gc/sbi/models"
)

type AuthVector struct {
	rand, xres, autn, ckprime, ikprime, xresstar, kausf []byte
}

func (av *AuthVector) decode(info *models.AuthenticationVector) (err error) {
	if len(info.Rand) > 0 {
		if av.rand, err = hex.DecodeString(info.Rand); err != nil {
			return
		}
	}
	if len(info.Xres) > 0 {
		if av.xres, err = hex.DecodeString(info.Xres); err != nil {
			return
		}
	}

	if len(info.Autn) > 0 {
		if av.autn, err = hex.DecodeString(info.Autn); err != nil {
			return
		}
	}

	if len(info.CkPrime) > 0 {
		if av.ckprime, err = hex.DecodeString(info.CkPrime); err != nil {
			return
		}
	}

	if len(info.IkPrime) > 0 {
		if av.ikprime, err = hex.DecodeString(info.IkPrime); err != nil {
			return
		}
	}

	if len(info.XresStar) > 0 {
		if av.xresstar, err = hex.DecodeString(info.XresStar); err != nil {
			return
		}
	}

	if len(info.Kausf) > 0 {
		if av.kausf, err = hex.DecodeString(info.Kausf); err != nil {
			return
		}
	}
	return
}
