package folder

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/validators"

	"github.com/gorilla/mux"
)

type UpdateProcedure struct {
}

func (UpdateProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return validators.FolderValidator(GetFolder(proc.Request))
}

func (UpdateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folder := GetFolder(proc.Request)
	folderId, err := strconv.ParseUint(mux.Vars(proc.Request)["folderId"], 10, 64)

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
