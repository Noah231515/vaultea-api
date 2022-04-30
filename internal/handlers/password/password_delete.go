package password

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
)

type DeleteProcedure struct {
}

func (DeleteProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return true
}

func (DeleteProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	passwordId, err := http_utils.GetQueryParamId(proc.Request, "passwordId")
	responseMap := make(map[string]string)

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	deleteErr := db.Where("id = ?", passwordId).Delete(&models.Password{}).Error

	if deleteErr != nil {
		http_utils.WriteErrorResponse(proc.Writer, deleteErr)
		return
	}

	responseMap["id"] = fmt.Sprint(passwordId)
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
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
