package folder

import (
	"net/http"
	"strconv"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"

	"github.com/gorilla/mux"
)

type DeleteProcedure struct {
}

func (DeleteProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return true
}

func (DeleteProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folderId, err := strconv.ParseUint(mux.Vars(proc.Request)["folderId"], 10, 1)

	if err != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(err.Error()))
	}

	db.Model(models.Folder{}).Delete(folderId)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write([]byte("wow"))
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	proc := DeleteProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
