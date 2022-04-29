package password

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
	return validators.PasswordValidator(GetPassword(proc.Request))
}

func (UpdateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	password := GetPassword(proc.Request)
	passwordId, err := strconv.ParseUint(mux.Vars(proc.Request)["passwordId"], 10, 64)

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
