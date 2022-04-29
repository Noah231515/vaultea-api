package validators

import (
	"vaultea/api/internal/models"
	string_utils "vaultea/api/internal/utils/string"
)

func UserValidator(user models.User) bool {
	if string_utils.IsEmptyString(user.Username) {
		return false
	}

	if string_utils.IsEmptyString(user.Password) {
		return false
	}

	return true
}
