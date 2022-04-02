package handlers

import (
	"net/http"

	"vaultea/api/database"
	"vaultea/api/models"
	"vaultea/api/utils"
	"vaultea/api/validators"

	"gorm.io/gorm"
)

func (ProcedureData) ValidateRequestMethod(procData *ProcedureData) bool {
	return utils.IsPost(procData.Request)
}

func (ProcedureData) CheckPermissions(procData *ProcedureData) bool {
	return true
}

func (ProcedureData) ValidateData(proc *ProcedureData) bool {
	proc.BodyMap = utils.GetRequestBodyMap(proc.Request)
	return validators.SignUpValidator(proc.BodyMap)
}

func (ProcedureData) Execute(proc *ProcedureData) {
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
	procData := ProcedureData{writer, request, make(map[string]interface{})}
	ExecuteHandler(procData, &procData)
}

func Login(writer http.ResponseWriter, request *http.Request) {
	// authenticate
	// return user's data
}
