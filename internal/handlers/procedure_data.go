package handlers

import "net/http"

type ProcedureData struct {
	Writer  http.ResponseWriter
	Request *http.Request
	BodyMap map[string]interface{}
}
