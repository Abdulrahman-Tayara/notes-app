package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(length int) string {
	b := make([]byte, length/2)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
