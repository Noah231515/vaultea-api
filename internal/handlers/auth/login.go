package auth

import (
	"encoding/json"
	"net/http"
	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	crypto_utils "vaultea/api/internal/utils/crypto"
	http_utils "vaultea/api/internal/utils/http"
	"vaultea/api/internal/validators"

	"gorm.io/gorm"
)

type LoginProcedure struct {
}

func (LoginProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return validators.UserValidator(GetUser(proc.Request))
}

func (LoginProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	contextUser := GetUser(proc.Request)
	invalidMessage := "Invalid username or password"

	user := models.User{}
	vault := models.Vault{}
	folders := []models.Folder{}
	passwords := []models.Password{}

	result := db.Where("username = ?", contextUser.Username).First(&user)
	vaultResult := db.Where("user_id = ?", user.ID).First(&vault)
	foldersResult := db.Where("vault_id = ?", vault.ID).Find(&folders)
	passwordsResult := db.Where("vault_id = ?", vault.ID).Find(&passwords)

	if result.Error == nil && vaultResult.Error == nil && foldersResult.Error == nil && passwordsResult.Error == nil {
		if crypto_utils.ComparePassword(user.Password, contextUser.Password) {
			resp := make(map[string]interface{})
			jwt, _ := crypto_utils.GetJWT(user)

			resp["id"] = user.ID
			resp["username"] = user.Username
			resp["key"] = user.Key
			resp["accessToken"] = jwt
			resp["folders"] = folders
			resp["passwords"] = passwords

			jsonResponse, _ := json.Marshal(resp)

			proc.Writer.WriteHeader(http.StatusOK)
			proc.Writer.Write(jsonResponse)
			return
		} else {
			http_utils.WriteBadResponse(proc.Writer, 500, invalidMessage)
			return
		}
	} else if result.Error == gorm.ErrRecordNotFound {
		http_utils.WriteBadResponse(proc.Writer, 500, invalidMessage)
		return
	}

	proc.Writer.WriteHeader(200)
}

func Login(writer http.ResponseWriter, request *http.Request) {
	proc := LoginProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
