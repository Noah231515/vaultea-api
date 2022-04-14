package folder

import (
	"encoding/json"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	crypto_utils "vaultea/api/internal/utils/crypto"
	http_utils "vaultea/api/internal/utils/http"
	"vaultea/api/internal/validators"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type CreateProcedure struct {
}

func (CreateProcedure) CheckPermissions(procData *handlers.ProcedureData) bool {
	return true
}

func (CreateProcedure) ValidateRequestMethod(procData *handlers.ProcedureData) bool {
	return http_utils.IsPost(procData.Request)
}

func (CreateProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	proc.BodyMap = http_utils.GetRequestBodyMap(proc.Request)
	return validators.FolderValidator(proc.BodyMap)
}

func (CreateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	context := proc.Request.Context()
	claims := context.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).CustomClaims.(*crypto_utils.Claims)
	user := models.User{}
	vault := models.Vault{}

	result := db.Where("username = ?", claims.Username).First(&user)
	if result.Error != nil {

		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte("Error validating JWT."))
		return
	}

	vaultResult := db.Where("user_id = ?", user.ID).First(&vault)
	if vaultResult.Error != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte("Error validating JWT."))
		return
	}

	folder := models.Folder{
		Name:        proc.BodyMap["name"].(string),
		Description: proc.BodyMap["description"].(string),
		VaultID:     vault.ID,
	}

	if proc.BodyMap["folderId"] != nil {
		folderId := proc.BodyMap["folderId"].(uint)
		folder.FolderID = &folderId
	}

	db.Create(&folder)

	json, _ := json.Marshal(folder)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(json)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	proc := CreateProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
