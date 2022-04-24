package folder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"

	"github.com/gorilla/mux"
)

type DeleteProcedure struct {
}

func (DeleteProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return true
}

func (DeleteProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folderId, err := strconv.ParseUint(mux.Vars(proc.Request)["folderId"], 10, 64)
	responseMap := make(map[string]string)

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	db.Delete(&models.Folder{}, folderId)

	responseMap["id"] = fmt.Sprint(folderId)
	response, err := json.Marshal(responseMap)

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(response)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	proc := DeleteProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
