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

	// Recursive tree DFS to traverse, and delete folders
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(models.Folder{}).Where("id = ?", folderId).First(&mainFolder).Error; err != nil {
			return err
		}

		if err := tx.Where("folder_id = ?", mainFolder.ID).Find(&children).Error; err != nil {
			return err
		}

		for _, child := range children {
			traverse(child, tx)
		}

		if err := tx.Where("folder_id", mainFolder.ID).Delete(&models.Password{}).Error; err != nil {
			return err
		}

		for i := len(foldersToDelete) - 1; i >= 0; i-- {
			if err := tx.Where("id = ?", foldersToDelete[i].ID).Delete(&models.Folder{}).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(&mainFolder).Error; err != nil {
			return err
		}

		return nil
	})

	responseMap["id"] = fmt.Sprint(folderId)
	response, err := json.Marshal(responseMap)

	if err != nil {
		http_utils.WriteErrorResponse(proc.Writer, err)
		return
	}

	proc.Writer.WriteHeader(200)
	proc.Writer.Write(response)
}

func traverse(folder models.Folder, tx *gorm.DB) error {
	var children []models.Folder
	foldersToDelete = append(foldersToDelete, folder)

	if err := tx.Where("folder_id = ?", folder.ID).Delete(&models.Password{}).Error; err != nil {
		return err
	}
	if err := tx.Where("folder_id = ?", folder.ID).Find(&children).Error; err != nil {
		return err
	}
	for _, child := range children {
		traverse(child, tx)
	}

	return nil
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	proc := DeleteProcedure{}
	procData := handlers.ProcedureData{writer, request}
	handlers.ExecuteHandler(proc, &procData)
}
