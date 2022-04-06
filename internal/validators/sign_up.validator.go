package validators

import (
	string_utils "vaultea/api/internal/utils/string"
)

func SignUpValidator(bodyData map[string]interface{}) bool {
	if string_utils.IsEmptyString(bodyData["key"].(string)) {
		return false
	}

	if string_utils.IsEmptyString(bodyData["username"].(string)) {
		return false
	}

	if string_utils.IsEmptyString(bodyData["password"].(string)) {
		return false
	}

	return true
}
