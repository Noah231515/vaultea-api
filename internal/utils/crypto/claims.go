package crypto_utils

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	VaultID  uint   `json: "vaultId"`
	jwt.RegisteredClaims
}

func (claim *Claims) Validate(ctx context.Context) error {
	return nil
}
