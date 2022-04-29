package http_utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"vaultea/api/internal/models"
	crypto_utils "vaultea/api/internal/utils/crypto"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

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
	case "Password":
		var password models.Password
		jsonErr = json.Unmarshal(b, &password)
		return password, jsonErr
	case "User":
		var user models.User
		jsonErr = json.Unmarshal(b, &user)
		return user, jsonErr
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

func GetVaultId(writer http.ResponseWriter, request *http.Request) uint {
	context := request.Context()
	claims := context.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).CustomClaims.(*crypto_utils.Claims)
	return claims.VaultID
}
