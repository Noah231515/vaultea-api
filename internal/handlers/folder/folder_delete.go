package folder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"
	"vaultea/api/internal/models"
	http_utils "vaultea/api/internal/utils/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type DeleteProcedure struct {
}

var foldersToDelete []models.Folder

func (DeleteProcedure) ValidateData(proc *handlers.ProcedureData) bool {
	return true
}

func (DeleteProcedure) Execute(proc *handlers.ProcedureData) {
	db := database.GetDb()
	folderId, err := strconv.ParseUint(mux.Vars(proc.Request)["folderId"], 10, 64)
	responseMap := make(map[string]string)
	foldersToDelete = make([]models.Folder, 1)

	var mainFolder models.Folder
	var children []models.Folder

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	db.Model(models.Folder{}).Where("id = ?", folderId).First(&mainFolder)
	db.Where("folder_id = ?", mainFolder.ID).Find(&children)
	for _, child := range children {
		traverse(child, &db)
	}

	db.Where("folder_id", mainFolder.ID).Delete(&models.Password{})

	for i := len(foldersToDelete) - 1; i >= 0; i-- {
		db.Where("id = ?", foldersToDelete[i].ID).Delete(&models.Folder{})
	}

	db.Delete(&mainFolder)

	// IMplement recursive delete

	responseMap["id"] = fmt.Sprint(folderId)
	response, err := json.Marshal(responseMap)

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(response)
}

func traverse(folder models.Folder, db *gorm.DB) {
	var children []models.Folder
	foldersToDelete = append(foldersToDelete, folder)

	db.Where("folder_id = ?", folder.ID).Delete(&models.Password{})
	db.Where("folder_id = ?", folder.ID).Find(&children)
	for _, child := range children {
		traverse(child, db)
	}

	// get passwords for child
	// get childs children
	// traverse on those
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	proc := DeleteProcedure{}
	procData := handlers.ProcedureData{writer, request, make(map[string]interface{})}
	handlers.ExecuteHandler(proc, &procData)
}
