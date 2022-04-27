package validators

import (
	"vaultea/api/internal/models"
	string_utils "vaultea/api/internal/utils/string"
)

func FolderValidator(folder models.Folder) bool {
	if string_utils.IsEmptyString(folder.Name) {
		return false
	}

	return true
}
