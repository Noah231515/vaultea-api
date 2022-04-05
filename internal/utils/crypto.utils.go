package utils

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"math/rand"
	"time"
	"vaultea/api/internal/environment"
	"vaultea/api/internal/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/pbkdf2"
)

func GetRandomSalt() []byte {
	salt := make([]byte, 8) // Look into seeding
	rand.Read(salt)

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

		return dbPassword == b64.StdEncoding.EncodeToString(hashedPassWithSalt)
		// return bytes.Equal(decodedDbPassword, hashedPassWithSalt) TODO: come back to this. this always passes for some rason

	} else {
		return false
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
} // Break this out into crypto util package

func GetJWT(user models.User) (string, error) {
	expirationDate := time.Now().Add(5 * time.Minute)
	secret := environment.GetEnv()["SECRET_STRING"]
	secretBytes := []byte(secret)

	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedString, err := token.SignedString(secretBytes)

	if err != nil { // TODO: handle error here
		return "", err
	} else {
		return signedString, nil
	}
}
