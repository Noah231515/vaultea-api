package password

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
	return validators.PasswordValidator(GetPassword(proc.Request))
}

func (CreateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	contextPassword := GetPassword(proc.Request)
	vaultID := http_utils.GetVaultId(proc.Writer, proc.Request)

	contextPassword.VaultID = vaultID
	db.Create(&contextPassword)

	json, _ := json.Marshal(contextPassword)

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(json)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	proc := CreateProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
