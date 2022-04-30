package password

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
	return validators.PasswordValidator(GetPassword(proc.Request))
}

func (UpdateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	password := GetPassword(proc.Request)
	passwordId, err := http_utils.GetQueryParamId(proc.Request, "passwordId")

	if err != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(err.Error()))
		return
	}

	password.ID = uint(passwordId)

	db.Model(password).Updates(&password)
	response, err := json.Marshal(&password)
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
