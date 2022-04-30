package folder

import (
	"encoding/json"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	http_utils "vaultea/api/internal/utils/http"
	"vaultea/api/internal/validators"
)

type UpdateProcedure struct {
}

func (UpdateProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return validators.FolderValidator(GetFolder(proc.Request))
}

func (UpdateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folder := GetFolder(proc.Request)
	folderId, err := http_utils.GetQueryParamId(proc.Request, "folderId")

	if err != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(err.Error()))
		return
	}

	folder.ID = uint(folderId)

	db.Model(folder).Updates(&folder)
	response, err := json.Marshal(&folder)
	if err != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(err.Error()))
		return
	}

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(response)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	proc := UpdateProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
