package folder

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
	var folder models.Folder
	id, err := http_utils.GetQueryParamId(proc.Request, "folderId")

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	db.Where("id = ?", id).Find(&folder)
	folder.Starred = !folder.Starred
	db.Model(&folder).Select("starred").Updates(&folder)

	json, _ := json.Marshal(folder)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(json)
}

func UpdateStarred(writer http.ResponseWriter, request *http.Request) {
	proc := UpdateStarredProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
