package auth_handler

import (
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	"vaultea/api/internal/utils"
	"vaultea/api/internal/validators"

	"gorm.io/gorm"
)

type SignUpProcedure struct {
}

func (SignUpProcedure) CheckPermissions(procData *handlers.ProcedureData) bool {
	return true
}

func (SignUpProcedure) ValidateRequestMethod(procData *handlers.ProcedureData) bool {
	return utils.IsPost(procData.Request)
}

func (SignUpProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	proc.BodyMap = utils.GetRequestBodyMap(proc.Request)
	return validators.SignUpValidator(proc.BodyMap)
}

func (SignUpProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()

	db.Transaction(func(tx *gorm.DB) error {
		hashedPassword := utils.HashPassword(proc.BodyMap["password"].(string))

		user := models.User{
			Key:      proc.BodyMap["key"].(string),
			Username: proc.BodyMap["username"].(string),
			Password: hashedPassword}
		tx.Create(&user)

		vault := models.Vault{
			UserID: user.ID}
		tx.Create(&vault)

		return nil
	})
	// create a private rsa key for the user to use for jwt
	proc.Writer.WriteHeader(200)
}

func SignUp(writer http.ResponseWriter, request *http.Request) {
	proc := SignUpProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
