package handlers

type GenericHandlerProcedure interface {
	ValidateRequestMethod(procData *ProcedureData) bool
	CheckPermissions(procData *ProcedureData) bool
	ValidateData(procData *ProcedureData) bool
	Execute(procData *ProcedureData)
}
