package auth_handler

import (
	"encoding/json"
	"net/http"
	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	"vaultea/api/internal/utils"
	"vaultea/api/internal/validators"

	"gorm.io/gorm"
)

type LoginProcedure struct {
}

func (LoginProcedure) CheckPermissions(procData *handlers.ProcedureData) bool {
	return true
}

func (LoginProcedure) ValidateRequestMethod(procData *handlers.ProcedureData) bool {
	return utils.IsPost(procData.Request)
}

func (LoginProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	proc.BodyMap = utils.GetRequestBodyMap(proc.Request)
	return validators.LoginValidator(proc.BodyMap)
}

func (LoginProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	invalidMessage := "Invalid username or password"

	user := models.User{}
	result := db.Where("username = ?", proc.BodyMap["username"]).First(&user)

	if result.Error == nil {
		if utils.ComparePassword(user.Password, proc.BodyMap["password"].(string)) {
			resp := make(map[string]interface{})
			jwt, _ := utils.GetJWT(user)

			resp["id"] = user.ID
			resp["username"] = user.Username
			resp["key"] = user.Key
			resp["accessToken"] = jwt
			resp["vaultId"] = 1
			resp["folders"] = make([]models.Folder, 0)
			resp["passwords"] = make([]models.Password, 0)

			jsonResponse, _ := json.Marshal(resp)

			proc.Writer.WriteHeader(http.StatusOK)
			proc.Writer.Write(jsonResponse)
			return
		} else {
			utils.WriteBadResponse(proc.Writer, 500, invalidMessage)
			return
		}
	} else if result.Error == gorm.ErrRecordNotFound {
		utils.WriteBadResponse(proc.Writer, 500, invalidMessage)
		return
	}

	proc.Writer.WriteHeader(200)
}

func Login(writer http.ResponseWriter, request *http.Request) {
	proc := LoginProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
