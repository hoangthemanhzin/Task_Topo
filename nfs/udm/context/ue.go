package context

import (
	"bytes"
	"encoding/hex"
	"etrib5gc/logctx"
	"etrib5gc/sbi/models"
	"etrib5gc/util/sec"
	"fmt"
)

type UeContext struct {
	logctx.LogWriter
	supi     string
	authtype models.AuthType
	sub      *models.AuthenticationSubscription
	milenage *sec.Milenage
	amf      [2]byte
	sqn      [6]byte
}

// create a UEContext from subscription data received from UDR
func newUeContext(supi string, sub *models.AuthenticationSubscription) (ue *UeContext, err error) {
	ue = &UeContext{
		LogWriter: log.WithFields(logctx.Fields{"ue-supi": supi}),
		supi:      supi,
	}
	ue.Infof("Create UE with fields=%s", supi)
	if supi != sub.Supi {
		err = fmt.Errorf("Mismatch supi: %s vs %s", supi, sub.Supi)
		return
	}
	err = ue.loadSub(sub)
	return
}

func (ue *UeContext) loadSub(sub *models.AuthenticationSubscription) (err error) {
	var k, opopc, amf, sqn []byte
	if k, err = hex.DecodeString(sub.EncPermanentKey); err != nil {
		return
	}
	if amf, err = hex.DecodeString(sub.AuthenticationManagementField); err != nil {
		return
	}

	if sqn, err = hex.DecodeString(sub.SequenceNumber.Sqn); err != nil {
		return
	}
	var isopc bool = false
	if len(sub.EncOpcKey) > 0 {
		if opopc, err = hex.DecodeString(sub.EncOpcKey); err != nil {
			return
		}
		isopc = true
	} else if len(sub.EncTopcKey) > 0 {
		if opopc, err = hex.DecodeString(sub.EncTopcKey); err != nil {
			return
		}
	} else {
		err = fmt.Errorf("OP/OPC is missing from subscription data")
		return
	}
	if ue.milenage, err = sec.NewMilenage(k, opopc, isopc); err != nil {
		return
	}

	if len(sqn) != 6 || len(amf) != 2 {
		err = fmt.Errorf("Input size for SQN or AMF is incorrect")
		return
	}
	copy(ue.sqn[:], sqn)
	copy(ue.amf[:], amf)
	switch sub.AuthenticationMethod {
	case models.AUTHMETHOD__5_G_AKA:
		ue.authtype = models.AUTHTYPE__5_G_AKA
	case models.AUTHMETHOD_EAP_AKA_PRIME:
		ue.authtype = models.AUTHTYPE_EAP_AKA_PRIME
	case models.AUTHMETHOD_EAP_TLS:
		ue.authtype = models.AUTHTYPE_EAP_TLS
	default:
		err = fmt.Errorf("Unsupported Auth method")
		return
	}
	ue.Infof("Subscription data loaded")
	return
}

func (ue *UeContext) Supi() string {
	return ue.supi
}

func (ue *UeContext) AuthType() models.AuthType {
	return ue.authtype
}

func (ue *UeContext) Resync(auststr string, randstr string) (err error) {
	defer func() {
		if err != nil {
			ue.Errorf("Resync failed: %s", err.Error())
		}
	}()

	var auts, uerand []byte
	if auts, err = hex.DecodeString(auststr); err != nil {
		return
	}

	if uerand, err = hex.DecodeString(randstr); err != nil {
		return
	}
	if len(auts) != 14 || len(uerand) != 16 {
		err = fmt.Errorf("Wrong input size:auts[%d],rand[%d]", len(auts), len(uerand))
		return
	}

	var sqn, macs []byte
	if sqn, macs, err = ue.milenage.CheckSqn(auts[:6], uerand); err != nil {
		return
	}
	if bytes.Compare(macs, auts[6:]) != 0 {
		err = fmt.Errorf("Resycn MAC failed [%x vs %x]", ue.supi, macs, auts[6:])
		return
	}
	copy(ue.sqn[:], sqn)
	return
}

// TODO: move this procedure into util/sec package
func (ue *UeContext) BuildAuthenticationVector(servingnet string) (v models.AuthenticationVector, err error) {
	defer func() {
		if err != nil {
			ue.Errorf("Build authentication vector failed: %s", err.Error())
		}
	}()
	//make sure the sequence number is updated successfully before building the
	//vector
	ue.milenage.Refresh() //update RAND
	maca, _, _ := ue.milenage.F1(ue.sqn[:], ue.amf[:])
	res, ak := ue.milenage.F2F5()
	ck := ue.milenage.F3()
	ik := ue.milenage.F4()
	//akstar := ue.milenage.F5star()
	sqnXORak := make([]byte, 6)
	for i := 0; i < 6; i++ {
		sqnXORak[i] = ue.sqn[i] ^ ak[i]
	}
	autn := append(append(sqnXORak, ue.amf[:]...), maca...)

	v.Rand = hex.EncodeToString(ue.milenage.GetRand())
	v.Autn = hex.EncodeToString(autn)

	key := append(ck, ik...)
	var xresstar []byte
	if ue.authtype == models.AUTHTYPE__5_G_AKA {
		if _, xresstar, err = sec.ResstarXresstar(key, []byte(servingnet), ue.milenage.GetRand(), res); err != nil {
			return
		}
		var kausf []byte
		if kausf, err = sec.KAUSF(key, []byte(servingnet), sqnXORak); err != nil {
			return
		}

		v.XresStar = hex.EncodeToString(xresstar)
		v.Kausf = hex.EncodeToString(kausf)
		v.AvType = models.AVTYPE__5_G_HE_AKA
	} else if ue.authtype == models.AUTHTYPE_EAP_AKA_PRIME {
		var ckprime, ikprime []byte
		if ckprime, ikprime, err = sec.CkPrimeIkPrime(key, []byte(servingnet), sqnXORak); err != nil {
			return
		}
		v.CkPrime = hex.EncodeToString(ckprime)
		v.IkPrime = hex.EncodeToString(ikprime)
		v.AvType = models.AVTYPE_EAP_AKA_PRIME

	} else {
		err = fmt.Errorf("Unsupported authentication type")
		//should never happen
	}
	ue.Info("Authentication vector built")
	return
}
