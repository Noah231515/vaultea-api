package crypto_utils

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	VaultID  uint   `json: "vaultId"`
	jwt.RegisteredClaims
}

func (claim *Claims) Validate(ctx context.Context) error { // TODO: Implement claim validation
	return nil
}
