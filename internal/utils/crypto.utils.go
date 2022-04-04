package utils

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

func GetRandomSalt() []byte {
	salt := make([]byte, 8)
	rand.Read(salt)

	fmt.Println(string(salt))
	return salt
}

func PbKdf2(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 4096, 32, sha512.New)
}

func HashPassword(password string) string {
	salt := GetRandomSalt()

	hashedPassword := PbKdf2(password, salt)
	hashedPassword = append(salt, hashedPassword...)
	return b64.StdEncoding.EncodeToString([]byte(hashedPassword))
}

func ComparePassword(dbPassword string, clientPassword string) bool {
	decodedDbPassword, err := b64.StdEncoding.DecodeString(dbPassword)
	if err == nil {
		salt := decodedDbPassword[0:8]

		hashedClientPassword := PbKdf2(clientPassword, salt)
		hashedPassWithSalt := append(salt, hashedClientPassword...)

		fmt.Println(dbPassword)
		fmt.Println(b64.StdEncoding.EncodeToString([]byte(hashedPassWithSalt)))

		return dbPassword == b64.StdEncoding.EncodeToString(hashedPassWithSalt)
		// return bytes.Equal(decodedDbPassword, hashedPassWithSalt) TODO: come back to this. this always passes for some rason

	} else {
		return false
	}
}
