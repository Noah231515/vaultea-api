package middleware

import (
	"context"
	"net/http"
	"reflect"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
)

func FolderDataMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawData, err := http_utils.GetBodyData(r, reflect.TypeOf(models.Folder{}))

		if (err) != nil {
			panic(err)
		}

		ctx := context.WithValue(r.Context(), "folder", rawData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
