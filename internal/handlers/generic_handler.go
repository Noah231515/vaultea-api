package handlers

import (
	"net/http"
	"vaultea/api/internal/utils"
)

func ExecuteHandler(proc GenericHandlerProcedure, procData *ProcedureData) {
	if !proc.ValidateRequestMethod(procData) {
		utils.WriteBadResponse(procData.Writer, http.StatusBadRequest, "Incorrect request method")
		return
	}
	if !proc.CheckPermissions(procData) {
		utils.WriteBadResponse(procData.Writer, http.StatusForbidden, "No access to this resource")
		return
	}
	if !proc.ValidateData(procData) {
		utils.WriteBadResponse(procData.Writer, http.StatusBadRequest, "Invalid data")
		return
	}
	proc.Execute(procData)
}
