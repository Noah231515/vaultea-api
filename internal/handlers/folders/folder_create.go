package folder

import (
	"encoding/json"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
	"vaultea/api/internal/validators"
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
	return validators.SignUpValidator(proc.BodyMap)
}

func (CreateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folderId := proc.BodyMap["folderId"].(uint)

	folder := models.Folder{
		Name:        proc.BodyMap["name"].(string),
		Description: proc.BodyMap["description"].(string),
		FolderID:    &folderId,
	}

	db.Create(&folder)

	json, _ := json.Marshal(folder)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(json)
}

func SignUp(writer http.ResponseWriter, request *http.Request) {
	proc := CreateProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
