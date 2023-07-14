package pass

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

const saltSize = 16
const keyLen = 64
const iter = 4096

func generateSalt() (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	return base64.StdEncoding.EncodeToString(salt), err
}

func hashPassword(password string, salt string) (string, error) {
	passwordBytes := []byte(password)
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	hashedPassword := pbkdf2.Key(passwordBytes, saltBytes, iter, keyLen, sha256.New)
	return base64.StdEncoding.EncodeToString(hashedPassword), err
}

func CreatePassHash(password string) (string, string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", "", err
	}
	hashedPassword, err := hashPassword(password, salt)
	return hashedPassword, salt, err
}

func PasswordChecks(password string, storedHash string, salt string) (bool, error) {
	expectedHash, err := hashPassword(password, salt)
	if err != nil {
		return false, err
	}
	return storedHash == expectedHash, nil
}
