package addressbook

import (
	"crypto/sha1"
	"encoding/hex"
)

type Hasher interface {
	Hash(string) string
}

type SHA1Hasher struct {
}

func (*SHA1Hasher) Hash(valueToHash string) string {
	h := sha1.New()
	h.Write([]byte(valueToHash))
	sha := h.Sum(nil)
	return hex.EncodeToString(sha)
}
