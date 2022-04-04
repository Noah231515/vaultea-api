package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
func WriteBadResponse(writer http.ResponseWriter, code int, message string) {
	resp := make(map[string]string)
	resp["message"] = message

	jsonResponse, _ := json.Marshal(resp)

	writer.WriteHeader(code)
	writer.Write(jsonResponse)
}

func IsPost(request *http.Request) bool {
	return request.Method == http.MethodPost
}
