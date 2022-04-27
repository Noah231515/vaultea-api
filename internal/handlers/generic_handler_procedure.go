package handlers

type GenericHandlerProcedure interface {
	ValidateData(procData *ProcedureData) bool
	Execute(procData *ProcedureData)
}
