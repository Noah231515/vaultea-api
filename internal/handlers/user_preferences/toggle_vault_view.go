package user_preferences

import (
	"encoding/json"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
)

type ToggleVaultViewProcedure struct {
}

func (ToggleVaultViewProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return true
}

func (ToggleVaultViewProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	userID := http_utils.GetUserID(proc.Writer, proc.Request)
	userPreferences := models.UserPreferences{}

	queryError := db.Model(models.UserPreferences{}).Where("user_id = ?", userID).First(&userPreferences)
	if queryError.Error != nil {
		http_utils.WriteErrorResponse(proc.Writer, queryError.Error)
		return
	}

	if userPreferences.VaultView == 0 {
		userPreferences.VaultView = 1
	} else {
		userPreferences.VaultView = 0
	}

	updateError := db.Model(userPreferences).Select("vault_view").Updates(&userPreferences)
	if updateError.Error != nil {
		http_utils.WriteErrorResponse(proc.Writer, updateError.Error)
	}

	response, err := json.Marshal(&userPreferences)
	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	// TODO: Get error catching as it should and fix updating and only return vaultView: value instead of the whole object
	proc.Writer.WriteHeader(200)
	proc.Writer.Write(response)
}

func ToggleVaultView(writer http.ResponseWriter, request *http.Request) {
	proc := ToggleVaultViewProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
