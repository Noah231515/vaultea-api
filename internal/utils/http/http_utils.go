package http_utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"vaultea/api/internal/models"
	crypto_utils "vaultea/api/internal/utils/crypto"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
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
	resp := make(map[string]string)
	resp["message"] = err.Error()

	jsonResponse, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	writer.WriteHeader(500)
	writer.Write(jsonResponse)
}

func GetVaultId(writer http.ResponseWriter, request *http.Request) uint {
	context := request.Context()
	claims := context.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).CustomClaims.(*crypto_utils.Claims)
	return claims.VaultID
}

func GetQueryParamId(request *http.Request, paramName string) (uint, error) {
	paramString := mux.Vars(request)[paramName]
	idPart := strings.Split(paramString, "/")[0]
	id, err := strconv.ParseUint(idPart, 10, 64)
	return uint(id), err
}
