package authentication

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	crypto_utils "vaultea/api/internal/utils/crypto"

	"github.com/golang-jwt/jwt/v4"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validMethods := [1]string{""}
		// parser := jwt.NewParser(jwt.WithValidMethods(validMethods[:]))
		claims := &crypto_utils.Claims{}
		// Validate jwt
		// if valid add to context
		authString := r.Header["Authorization"][0]            // TODO: Handle error here it may panic if this doesn't exist
		tokenString := strings.Split(authString, "Bearer")[1] // TODO: Handle error here, if this is empty less than size 2 then malformed token
		tokenString = strings.Trim(tokenString, " ")
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return "verifyKey", nil
		})
		fmt.Println(claims.VaultID)
		//parser.Parse(tokenString, Test)

		ctx, cancel := context.WithTimeout(r.Context(), time.Duration(60*time.Second))
		defer cancel()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func Test(token *jwt.Token) (interface{}, error) {

	return "", nil
}
