package utils

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password string) string {
	salt := make([]byte, 8)
	rand.Read(salt)

	hashedPassword := pbkdf2.Key([]byte(password), salt, 4096, 32, sha512.New)
	hashedPassword = append(salt, hashedPassword...)
	return b64.StdEncoding.EncodeToString([]byte(hashedPassword))
}
