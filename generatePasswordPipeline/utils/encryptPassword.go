package PasswordPipelineUtils

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func EncryptPassword(password string, key string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("invalid key format: %v", err)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	ciphertext := make([]byte, len(password))
	block.Encrypt(ciphertext, []byte(password))

	return hex.EncodeToString(ciphertext), nil
}
