package PasswordPipelineUtils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateEncryptionKey(masterPassword string, salt string) (string, error) {
	combined := []byte(masterPassword + salt)

	hash := sha256.Sum256(combined)

	return hex.EncodeToString(hash[:]), nil
}
