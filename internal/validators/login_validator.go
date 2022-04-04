package validators

import "vaultea/api/internal/utils"

func LoginValidator(bodyData map[string]interface{}) bool {
	if utils.IsEmptyString(bodyData["username"].(string)) {
		return false
	}

	if utils.IsEmptyString(bodyData["password"].(string)) {
		return false
	}

	return true
}
