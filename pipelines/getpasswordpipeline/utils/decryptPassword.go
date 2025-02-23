package utils

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func DecryptPassword(encryptedHex string, key string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("invalid key format: %v", err)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", fmt.Errorf("invalid ciphertext format: %v", err)
	}

	plaintext := make([]byte, len(ciphertext))
	block.Decrypt(plaintext, ciphertext)

	return string(plaintext), nil
}
