package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSalt() (string, error) {
	const saltLength = 16

	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
