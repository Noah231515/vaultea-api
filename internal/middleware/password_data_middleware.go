package middleware

import (
	"context"
	"net/http"
	"reflect"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
)

func PasswordDataMiddleware(next http.Handler) http.Handler { // TODO: Work into generic middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			next.ServeHTTP(w, r)
			return
		}

		rawData, err := http_utils.GetBodyData(r, reflect.TypeOf(models.Password{}))

		if (err) != nil {
			panic(err)
		}

		ctx := context.WithValue(r.Context(), "password", rawData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
