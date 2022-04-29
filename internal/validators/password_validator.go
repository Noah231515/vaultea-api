package validators

import (
	"vaultea/api/internal/models"
	string_utils "vaultea/api/internal/utils/string"
)

func PasswordValidator(password models.Password) bool {
	if string_utils.IsEmptyString(password.Name) {
		return false
	}

	if string_utils.IsEmptyString(password.Username) {
		return false
	}

	if string_utils.IsEmptyString(password.Password) {
		return false
	}

	return true
}
