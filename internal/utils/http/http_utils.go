package http_utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"vaultea/api/internal/database"
	"vaultea/api/internal/models"
	crypto_utils "vaultea/api/internal/utils/crypto"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func GetRequestBodyMap(request *http.Request) map[string]interface{} {
	bodyMap := make(map[string]interface{})

	b, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &bodyMap)

	return bodyMap
}

func GetBodyData(request *http.Request, dataType reflect.Type) (interface{}, error) {
	var jsonErr error
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	switch dataType.Name() {
	case "Folder":
		var folder models.Folder
		jsonErr = json.Unmarshal(b, &folder)
		return folder, jsonErr
	default:
		panic("Unknown data type")
	}
}

func WriteBadResponse(writer http.ResponseWriter, code int, message string) {
	resp := make(map[string]string)
	resp["message"] = message

	jsonResponse, _ := json.Marshal(resp)

	writer.WriteHeader(code)
	writer.Write(jsonResponse)
}

func WriteErrorResponse(writer http.ResponseWriter, err error) {
	writer.WriteHeader(500)
	writer.Write([]byte(err.Error()))
}

func IsPost(request *http.Request) bool {
	return request.Method == http.MethodPost
}

func GetVaultId(writer http.ResponseWriter, request *http.Request) uint { // TODO: Clean up
	db := database.GetDb()
	user := models.User{}
	vault := models.Vault{}
	context := request.Context()

	claims := context.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).CustomClaims.(*crypto_utils.Claims)

	result := db.Where("username = ?", claims.Username).First(&user)
	if result.Error != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("Error validating JWT."))
		return 0
	}

	vaultResult := db.Where("user_id = ?", user.ID).First(&vault)
	if vaultResult.Error != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("Error validating JWT."))
		return 0
	}

	return vault.ID
}
