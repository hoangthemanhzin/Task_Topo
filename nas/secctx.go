package nas

import (
	"encoding/binary"
	"encoding/hex"
	"etrib5gc/sbi/models"
	"etrib5gc/util/sec"

	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/nas/security"
)

const (
	HDP_NONE uint8 = iota
	HDP_HANDOVER
	HDP_MOBILITY_UPDATE
)

type NasSecCtx struct {
	dlcount security.Count
	ulcount security.Count
	encalg  uint8
	intalg  uint8
	intkey  [16]uint8
	enckey  [16]uint8
	valid   bool
}

func (ctx *NasSecCtx) EncAlg() uint8 {
	return ctx.encalg
}

func (ctx *NasSecCtx) IntAlg() uint8 {
	return ctx.intalg
}

func (ctx *NasSecCtx) UlCount() uint32 {
	return ctx.ulcount.Get()
}

func (ctx *NasSecCtx) DlCount() uint32 {
	return ctx.dlcount.Get()
}
func (ctx *NasSecCtx) SelectAlg(intorder, encorder []uint8, seccap *nasType.UESecurityCapability) {
	ctx.encalg = security.AlgCiphering128NEA1
	ctx.intalg = security.AlgIntegrity128NIA1

	supported := uint8(0)
	for _, alg := range intorder {
		switch alg {
		case security.AlgIntegrity128NIA0:
			supported = seccap.GetIA0_5G()
		case security.AlgIntegrity128NIA1:
			supported = seccap.GetIA1_128_5G()
		case security.AlgIntegrity128NIA2:
			supported = seccap.GetIA2_128_5G()
		case security.AlgIntegrity128NIA3:
			supported = seccap.GetIA3_128_5G()
		}
		if supported == 1 {
			ctx.intalg = alg
			break
		}
	}

	supported = uint8(0)
	for _, alg := range encorder {
		switch alg {
		case security.AlgCiphering128NEA0:
			supported = seccap.GetEA0_5G()
		case security.AlgCiphering128NEA1:
			supported = seccap.GetEA1_128_5G()
		case security.AlgCiphering128NEA2:
			supported = seccap.GetEA2_128_5G()
		case security.AlgCiphering128NEA3:
			supported = seccap.GetEA3_128_5G()
		}
		if supported == 1 {
			ctx.encalg = alg
			break
		}
	}
	//log.Infof("Selected IntegrityAlg[%d], EncryptionAlg[%d]", ctx.intalg, ctx.encalg)
}
func (ctx *NasSecCtx) IsValid() bool {
	return ctx.valid
}

func (ctx *NasSecCtx) invalidate() {
	ctx.valid = false
}

type SecCtx struct {
	NasSecCtx
	kamf   []byte
	kgnb   []uint8 //gnb key
	kn3iwf []uint8 //n3iwf key
	nh     []uint8 //next hop parameter (for AS security context)
	ncc    uint8   // next chain counter (for AS security context) used by ngap (Handover, PathSwitch)
}

func NewSecCtx(kamf []byte) *SecCtx {
	ctx := &SecCtx{}
	ctx.valid = false //invalid util keys are derived
	ctx.kamf = make([]byte, len(kamf))
	copy(ctx.kamf, kamf)
	return ctx
}

func (ctx *SecCtx) Kgnb() []byte {
	return ctx.kgnb
}
func (ctx *SecCtx) Kn3iwf() []byte {
	return ctx.kn3iwf
}

// Algorithm key Derivation function defined in TS 33.501 Annex A.9
func (ctx *SecCtx) DeriveAlgKeys(hdp uint8) (err error) {
	var p0 []byte
	var p1 [4]byte
	var kamf []byte
	switch hdp {
	case HDP_HANDOVER:
		//derive kamf prime
		p0, _ = hex.DecodeString("01")
		//NOTE: p1 = dlcount is only applied for 3gpp access (not sure about
		//non-3gpp (see TS 133.501 A.13)
		binary.BigEndian.PutUint32(p1[:], ctx.dlcount.Get())
		kamf, err = sec.KamfPrime(ctx.kamf, p0, p1[:])
	case HDP_MOBILITY_UPDATE:
		//derive kamf prime
		p0, _ = hex.DecodeString("00")
		binary.BigEndian.PutUint32(p1[:], ctx.ulcount.Get())
		kamf, err = sec.KamfPrime(ctx.kamf, p0, p1[:])
	default:
	}
	if err != nil {
		return
	}
	// Security Key
	P0 := []byte{security.NNASEncAlg}
	P1 := []byte{ctx.encalg}

	var kenc, kint []byte
	if kenc, err = sec.AlgKey(kamf, P0, P1); err != nil {
		return
	}

	// Integrity Key
	P0 = []byte{security.NNASIntAlg}
	P1 = []byte{ctx.intalg}

	if kint, err = sec.AlgKey(kamf, P0, P1); err != nil {
		return
	}
	copy(ctx.enckey[:], kenc[16:32])
	copy(ctx.intkey[:], kint[16:32])
	//log.Infof("kamf=%x, kamf-prime=%x, kenc=%x, kint=%x", ctx.kamf, kamf, ctx.enckey, ctx.intkey)
	ctx.kamf = kamf
	ctx.ulcount.Set(0, 0)
	ctx.dlcount.Set(0, 0)
	ctx.valid = true
	return
}

// Access Network key Derivation function defined in TS 33.501 Annex A.9
func (ctx *SecCtx) CreateAnKey(access models.AccessType) (err error) {

	nasaccess := security.AccessType3GPP
	if access == models.ACCESSTYPE_NON_3_GPP_ACCESS {
		nasaccess = security.AccessTypeNon3GPP
	}

	P0 := make([]byte, 4)
	binary.BigEndian.PutUint32(P0, ctx.ulcount.Get())
	P1 := []byte{nasaccess}

	switch nasaccess {
	case security.AccessType3GPP:
		ctx.kgnb, err = sec.RanKey(ctx.kamf, P0, P1)
	case security.AccessTypeNon3GPP:
		ctx.kn3iwf, err = sec.RanKey(ctx.kamf, P0, P1)
	}
	return
}

// NH Derivation function defined in TS 33.501 Annex A.10
func (ctx *SecCtx) CreateNh(syncinput []byte) (err error) {
	ctx.nh, err = sec.NhKey(ctx.kamf, syncinput)
	return
}

func (ctx *SecCtx) Update(access models.AccessType) (err error) {
	if err = ctx.CreateAnKey(access); err != nil {
		return
	}

	switch access {
	case models.ACCESSTYPE__3_GPP_ACCESS:
		ctx.CreateNh(ctx.kgnb)
	case models.ACCESSTYPE_NON_3_GPP_ACCESS:
		ctx.CreateNh(ctx.kn3iwf)
	}
	ctx.ncc = 1
	return
}

func (ctx *SecCtx) UpdateNh() error {
	ctx.ncc++
	return ctx.CreateNh(ctx.nh)
}
