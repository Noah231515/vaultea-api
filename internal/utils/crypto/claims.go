package crypto_utils

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	VaultID  uint   `json: "vaultId"`
	jwt.RegisteredClaims
}
