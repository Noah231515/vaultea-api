package handlers

import (
	"net/http"
	"vaultea/api/internal/utils"
)

func ExecuteHandler(proc GenericHandlerProcedure, procData *ProcedureData) {
	if !proc.ValidateRequestMethod(procData) {
		utils.WriteBadResponse(procData.Writer, http.StatusBadRequest, "Incorrect request method")
	}
	if !proc.CheckPermissions(procData) {
		utils.WriteBadResponse(procData.Writer, http.StatusForbidden, "No access to this resource")
	}
	if !proc.ValidateData(procData) {
		utils.WriteBadResponse(procData.Writer, http.StatusBadRequest, "Invalid data")
	}
	proc.Execute(procData)
}
