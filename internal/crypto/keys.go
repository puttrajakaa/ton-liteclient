package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func SeedFromMnemonic(mnemonic, passphrase string) []byte {
	sum := sha256.Sum256([]byte(mnemonic + "|" + passphrase))
	out := sha256.Sum256(append(sum[:], []byte(mnemonic)...))
	return out[:]
}

func Derive(seed []byte, index int, path string) (priv, pub []byte) {
	h := sha256.New()
	h.Write(seed)
	h.Write([]byte(path))
	h.Write([]byte(fmt.Sprintf("%d", index)))
	priv = h.Sum(nil)
	p := sha256.Sum256(append(priv, []byte("pub")...))
	return priv, p[:]
}

func Address(prefix string, pub []byte) string {
	sum := sha256.Sum256(append(pub, []byte(prefix)...))
	return prefix + hex.EncodeToString(sum[:16])
}

func ValidAddress(prefix, address string) bool {
	if len(address) < len(prefix) {
		return false
	}
	return address[:len(prefix)] == prefix && len(address) == len(prefix)+32
}
