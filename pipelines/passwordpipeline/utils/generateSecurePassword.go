package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
)

func GenerateSecurePassword() (string, error) {
	length := 32

	var password string
	var allChars = lowerCharSet + upperCharSet + specialCharSet + numberSet

	for _, charSet := range []string{lowerCharSet, upperCharSet, specialCharSet, numberSet} {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		password += string(charSet[n.Int64()])
	}

	for i := len(password); i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password += string(allChars[n.Int64()])
	}

	return password, nil
}
