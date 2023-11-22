package authentication

import (
	"encoding/hex"
	"etrib5gc/util/sec"
	"fmt"
)

// Kamf Derivation function defined in TS 33.501 Annex A.7
func (proc *AuthProc) createKamf(supi string, kseaf string) (err error) {
	if err = checkSupi(supi); err != nil {
		return
	}
	uectx := proc.ranue.UeContext()
	abba := uectx.Abba()
	var kseafbyte []byte
	if kseafbyte, err = hex.DecodeString(kseaf); err != nil {
		return
	}
	proc.kamf, err = sec.KAMF(kseafbyte, []byte(supi[5:]), abba)
	proc.Tracef("supi=%x, abba=%x", []byte(supi[4:]), abba)
	proc.Tracef("kseaf=%x, kamf=%x", kseafbyte, proc.kamf)
	return
}

// NOTE: Supi should be represented in its strict format (not string), then we
// don't need this check
func checkSupi(supi string) error {
	//TODO: may need to elaborate more
	if len(supi) < 4 {
		return fmt.Errorf("Invalid SUPI")
	}
	return nil
}
