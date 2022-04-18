package handlers

import (
	"net/http"
	http_utils "vaultea/api/internal/utils/http"
)

func ExecuteHandler(proc GenericHandlerProcedure, procData *ProcedureData) {
	if !proc.ValidateData(procData) {
		http_utils.WriteBadResponse(procData.Writer, http.StatusBadRequest, "Invalid data")
		return
	}
	proc.Execute(procData)
}
