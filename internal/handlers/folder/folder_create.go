package folder

import (
	"encoding/json"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	http_utils "vaultea/api/internal/utils/http"
	"vaultea/api/internal/validators"
)

type CreateProcedure struct {
}

func (CreateProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return validators.FolderValidator(GetFolder(proc.Request))
}

func (CreateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	contextfolder := GetFolder(proc.Request)
	vaultID := http_utils.GetVaultId(proc.Writer, proc.Request)

	contextfolder.VaultID = vaultID
	db.Create(&contextfolder)

	json, _ := json.Marshal(contextfolder)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(json)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	proc := CreateProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
