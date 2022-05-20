package password

import (
	"encoding/json"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
)

type UpdateStarredProcedure struct {
}

func (UpdateStarredProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return true
}

func (UpdateStarredProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	var password models.Password
	id, err := http_utils.GetQueryParamId(proc.Request, "passwordId")

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	db.Where("id = ?", id).Find(&password)
	password.Starred = !password.Starred
	db.Updates(&password)

	json, _ := json.Marshal(password)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(json)
}

func UpdateStarred(writer http.ResponseWriter, request *http.Request) {
	proc := UpdateStarredProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
