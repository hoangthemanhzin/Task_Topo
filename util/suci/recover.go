package suci

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/curve25519"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "suci"})
}

// suci-0(SUPI type)-mcc-mnc-routingIndentifier-protectionScheme-homeNetworkPublicKeyIdentifier-schemeOutput.
const (
	SUPI_TYPE_PLACE = 1
	MCC_PLACE       = 2
	MNC_PLACE       = 3
	SCHEME_PLACE    = 5
	KEY_INDEX_PLACE = 6
)

const (
	TYPE_IMSI_STR = "0"
	IMSI_PREFIX   = "imsi-"
	NULL_SCHEME   = "0"
	A_SCHEME      = "1"
	B_SCHEME      = "2"
)

type Profile struct {
	ProtectionScheme string `json:"scheme,omitempty"`
	PrivateKey       string `json:"prvkey,omitempty"`
	PublicKey        string `json:"pubkey,omitempty"`
}

// profile A.
const (
	A_MAC_K_LEN = 32 // octets
	A_ENC_K_LEN = 16 // octets
	A_ICB_LEN   = 16 // octets
	A_MAC_LEN   = 8  // octets
	A_HASH_LEN  = 32 // octets
)

// profile B.
const (
	B_MAC_K_LEN = 32 // octets
	B_ENC_K_LEN = 16 // octets
	B_ICB_LEN   = 16 // octets
	B_MAC_LEN   = 8  // octets
	B_HASH_LEN  = 32 // octets
)

func CompressKey(uncompressed []byte, y *big.Int) []byte {
	compressed := uncompressed[0:33]
	if y.Bit(0) == 1 { // 0x03
		compressed[0] = 0x03
	} else { // 0x02
		compressed[0] = 0x02
	}
	return compressed
}

// modified from https://stackoverflow.com/questions/46283760/
// how-to-uncompress-a-single-x9-62-compressed-point-on-an-ecdh-p256-curve-in-go.
func uncompressKey(compressedBytes []byte, priv []byte) (*big.Int, *big.Int) {
	// Split the sign byte from the rest
	signByte := uint(compressedBytes[0])
	xBytes := compressedBytes[1:]

	x := new(big.Int).SetBytes(xBytes)
	three := big.NewInt(3)

	// The params for P256
	c := elliptic.P256().Params()

	// The equation is y^2 = x^3 - 3x + b
	// x^3, mod P
	xCubed := new(big.Int).Exp(x, three, c.P)

	// 3x, mod P
	threeX := new(big.Int).Mul(x, three)
	threeX.Mod(threeX, c.P)

	// x^3 - 3x + b mod P
	ySquared := new(big.Int).Sub(xCubed, threeX)
	ySquared.Add(ySquared, c.B)
	ySquared.Mod(ySquared, c.P)

	// find the square root mod P
	y := new(big.Int).ModSqrt(ySquared, c.P)
	if y == nil {
		// If this happens then you're dealing with an invalid point.
		log.Error("Uncompressed key with invalid point")
		return nil, nil
	}

	// Finally, check if you have the correct root. If not you want -y mod P
	if y.Bit(0) != signByte&1 {
		y.Neg(y)
		y.Mod(y, c.P)
	}
	return x, y
}

func HmacSha256(input, mackey []byte, maclen int) (tag []byte, err error) {
	h := hmac.New(sha256.New, mackey)
	if _, err = h.Write(input); err != nil {
		log.Errorf("HMAC SHA256 error %+v", err)
		return
	}
	mac := h.Sum(nil)
	tag = mac[:maclen]
	return
}

func Aes128ctr(input, enckey, icb []byte) (output []byte, err error) {
	output = make([]byte, len(input))
	var block cipher.Block
	if block, err = aes.NewCipher(enckey); err != nil {
		log.Errorf("AES128 CTR error %+v", err)
		return
	}
	stream := cipher.NewCTR(block, icb)
	stream.XORKeyStream(output, input)
	return
}

func AnsiX963KDF(sharedkey, pubkey []byte, enckeylen, mackeylen, hashlen int) (kdfkey []byte) {
	var counter uint32 = 0x00000001
	numrounds := int(math.Ceil(float64(enckeylen+mackeylen) / float64(hashlen)))
	for i := 1; i <= numrounds; i++ {
		counterbytes := make([]byte, 4)
		binary.BigEndian.PutUint32(counterbytes, counter)
		tmpk := sha256.Sum256(append(append(sharedkey, counterbytes...), pubkey...))
		slicek := tmpk[:]
		kdfkey = append(kdfkey, slicek...)
		counter++
	}
	return
}

func swapnibbles(input []byte) []byte {
	output := make([]byte, len(input))
	for i, b := range input {
		output[i] = bits.RotateLeft8(b, 4)
	}
	return output
}

func decompose(input []byte, klen int, maclen int) (mac []byte, pubkey []byte, ciphertext []byte, err error) {
	if len(input) < klen+maclen {
		log.Error("len of input data is too short!")
		err = fmt.Errorf("suci input too short\n")
		return
	}
	pubkey = input[:klen]
	mac = input[len(input)-maclen:]
	ciphertext = input[klen : len(input)-maclen]
	return
}

func RecoverSupi(suci string, profiles []Profile) (supi string, err error) {
	parts := strings.Split(suci, "-")

	prefix := parts[0]
	if prefix == "imsi" || prefix == "nai" {
		log.Infof("Got supi\n")
		supi = suci
		return
	} else if prefix == "suci" {
		if parts[SUPI_TYPE_PLACE] != TYPE_IMSI_STR {
			err = fmt.Errorf("Unsupport type *%s) of supi", parts[SUPI_TYPE_PLACE])
			return
		}
		//supi types IMSI
		if len(parts) < 6 {
			err = fmt.Errorf("Suci with wrong format\n")
			return
		}
	} else {
		err = fmt.Errorf("Unknown succi prefix [%s]", prefix)
		return
	}

	//it is a imsi typed suci, let's process

	scheme := parts[SCHEME_PLACE]
	plmn := parts[MCC_PLACE] + parts[MNC_PLACE]

	supiprefix := IMSI_PREFIX

	if scheme == NULL_SCHEME { // NULL scheme
		supi = supiprefix + plmn + parts[len(parts)-1]
		return
	}

	var keyindex int
	if keyindex, err = strconv.Atoi(parts[KEY_INDEX_PLACE]); err != nil {
		return
	}
	if keyindex > len(profiles) {
		err = fmt.Errorf("keyIndex(%d) out of range(%d)", keyindex, len(profiles))
		return
	}

	profile := profiles[keyindex-1]

	if scheme != profile.ProtectionScheme {
		err = fmt.Errorf("Protect Scheme mismatch [%s:%s]", scheme, profile.ProtectionScheme)
		return
	}

	var prvkey []byte
	if prvkey, err = hex.DecodeString(profile.PrivateKey); err != nil {
		log.Errorf("Decode private key error: %+v", err)
		return
	}
	var schemeoutput []byte

	if schemeoutput, err = hex.DecodeString(parts[len(parts)-1]); err != nil {
		log.Errorf("Decode scheme output error: %+v", err)
		return
	}

	//get key length (scheme dependent)
	klen := 32 //A_SCHEME
	maclen := A_MAC_LEN
	uncompressed := false

	if scheme == B_SCHEME {
		maclen = B_MAC_LEN
		if schemeoutput[0] == 0x02 || schemeoutput[0] == 0x03 {
			klen = 33 // ceil(log(2, q)/8) + 1 = 33
			uncompressed = false
		} else if schemeoutput[0] == 0x04 {
			klen = 65 // 2*ceil(log(2, q)/8) + 1 = 65
			uncompressed = true
		} else {
			log.Error("input error")
			err = fmt.Errorf("suci input error\n")
			return
		}
	} else if scheme != A_SCHEME {
		err = fmt.Errorf("Unknown scheme")
		return
	}
	var mac, pubkey, ciphertext []byte
	if mac, pubkey, ciphertext, err = decompose(schemeoutput, klen, maclen); err != nil {
		return
	}

	var sharedkey, kdfkey, enckey, icb, mackey, mactag []byte
	if scheme == A_SCHEME {
		if sharedkey, err = curve25519.X25519(prvkey, pubkey); err != nil {
			log.Errorf("X25519 error: %+v", err)
			return
		}
		kdfkey = AnsiX963KDF(sharedkey, pubkey, A_ENC_K_LEN, A_MAC_K_LEN, A_HASH_LEN)
		enckey = kdfkey[:A_ENC_K_LEN]
		icb = kdfkey[A_ENC_K_LEN : A_ENC_K_LEN+A_ICB_LEN]
		mackey = kdfkey[len(kdfkey)-A_MAC_K_LEN:]
		mactag, err = HmacSha256(ciphertext, mackey, A_MAC_LEN)

	} else {
		var x, y *big.Int
		if uncompressed {
			x = new(big.Int).SetBytes(pubkey[1:(klen/2 + 1)])
			y = new(big.Int).SetBytes(pubkey[(klen/2 + 1):])
		} else {
			x, y = uncompressKey(pubkey, prvkey)
			if x == nil || y == nil {
				log.Error("Uncompressed key has invalid point")
				err = fmt.Errorf("Key uncompression error\n")
				return
			}
		}

		// x-coordinate is the shared key
		tmp, _ := elliptic.P256().ScalarMult(x, y, prvkey)
		sharedkey = tmp.Bytes()
		if uncompressed {
			pubkey = CompressKey(pubkey, y)
		}
		kdfkey = AnsiX963KDF(sharedkey, pubkey, B_ENC_K_LEN, B_MAC_K_LEN, B_HASH_LEN)
		enckey = kdfkey[:B_ENC_K_LEN]
		icb = kdfkey[B_ENC_K_LEN : B_ENC_K_LEN+B_ICB_LEN]
		mackey = kdfkey[len(kdfkey)-B_MAC_K_LEN:]
		mactag, err = HmacSha256(ciphertext, mackey, B_MAC_LEN)
	}

	if !bytes.Equal(mactag, mac) {
		log.Error("MAC unmatches")
		err = fmt.Errorf("MAC failed\n")
		return
	}
	var plaintext []byte
	if plaintext, err = Aes128ctr(ciphertext, enckey, icb); err != nil {
		return
	}
	var tmpstr string

	tmpstr = hex.EncodeToString(swapnibbles(plaintext))
	if tmpstr[len(tmpstr)-1] == 'f' {
		tmpstr = tmpstr[:len(tmpstr)-1]
	}

	supi = supiprefix + plmn + tmpstr
	return
}
