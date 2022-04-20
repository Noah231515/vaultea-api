package folder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"
	"vaultea/api/internal/validators"
)

type UpdateProcedure struct {
}

func (UpdateProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	proc.BodyMap = http_utils.GetRequestBodyMap(proc.Request)
	return validators.FolderValidator(proc.BodyMap)
}

func (UpdateProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folder := models.Folder{}

	b, err := ioutil.ReadAll(proc.Request.Body) // We want to put body in context

	if err != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(err.Error()))
		return
	}
	json.Unmarshal(b, &folder)

	updatedFolder := db.Model(&folder).Where("folder_id = ?", folder.ID).Updates(models.Folder{FolderID: folder.FolderID, Name: folder.Name, Description: folder.Description}) // TODO: Verify that the user can edit this
	if updatedFolder.Error != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(updatedFolder.Error.Error()))
		return
	}

	fmt.Println(folder.Name)
	response, err := json.Marshal(&folder)
	if err != nil {
		proc.Writer.WriteHeader(500)
		proc.Writer.Write([]byte(err.Error()))
		return
	}

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(response)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	proc := UpdateProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
