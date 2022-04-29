package auth

import (
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	crypto_utils "vaultea/api/internal/utils/crypto"
	"vaultea/api/internal/validators"

	"gorm.io/gorm"
)

type SignUpProcedure struct {
}

func (SignUpProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return validators.SignUpValidator(GetUser(proc.Request))
}

func (SignUpProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	contextUser := GetUser(proc.Request)

	db.Transaction(func(tx *gorm.DB) error {
		hashedPassword := crypto_utils.HashPassword(contextUser.Password)
		contextUser.Password = hashedPassword

		tx.Create(&contextUser)

		vault := models.Vault{
			UserID: contextUser.ID}
		tx.Create(&vault)

		return nil
	})
	// do nothing for now
	proc.Writer.WriteHeader(200)
}

func SignUp(writer http.ResponseWriter, request *http.Request) {
	proc := SignUpProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
