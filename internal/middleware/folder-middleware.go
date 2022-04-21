package middleware

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
)

func FolderDataMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var folderData models.Folder
		rawData, err := http_utils.GetBodyData(r, reflect.TypeOf(models.Folder{}))

		folderData = rawData.(models.Folder)
		fmt.Println(folderData.Name)
		if (err) != nil {
			panic(err)
		}

		// put folder data into request context
		next.ServeHTTP(w, r.WithContext(context.Background()))
	})
}
