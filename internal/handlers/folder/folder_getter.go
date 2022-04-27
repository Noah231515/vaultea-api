package folder

import (
	"net/http"
	"vaultea/api/internal/models"
)

func GetFolder(r *http.Request) models.Folder {
	folder := r.Context().Value("folder").(models.Folder)
	return folder
}
