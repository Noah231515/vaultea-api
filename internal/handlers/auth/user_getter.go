package auth

import (
	"net/http"
	"vaultea/api/internal/models"
)

func GetUser(r *http.Request) models.User { // TODO: try and make generic???
	user := r.Context().Value("user").(models.User)
	return user
}
