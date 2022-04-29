package password

import (
	"net/http"
	"vaultea/api/internal/models"
)

func GetPassword(r *http.Request) models.Password { // TODO: try and make generic???
	password := r.Context().Value("password").(models.Password)
	return password
}
