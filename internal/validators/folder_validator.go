package validators

import (
	string_utils "vaultea/api/internal/utils/string"
)

func FolderValidator(bodyData map[string]interface{}) bool {
	if string_utils.IsEmptyString(bodyData["name"].(string)) {
		return false
	}

	return true
}
