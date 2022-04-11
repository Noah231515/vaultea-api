package main

import (
	"log"
	"net/http"
	"time"

	"vaultea/api/internal/database"
	"vaultea/api/internal/environment"
	"vaultea/api/internal/handlers/auth"
	"vaultea/api/internal/handlers/folder"
	"vaultea/api/internal/middleware/authentication"

	"github.com/gorilla/mux"
)

func main() {
	environment.SetEnv()
	router := mux.NewRouter()
	database.ConnectToDatabase()
	database.MakeMigrations()
	initRoutes(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8081", // TODO: make configurable
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func initRoutes(router *mux.Router) {
	// Auth
	router.HandleFunc("/api/signup", auth.SignUp).Methods("POST")
	router.HandleFunc("/api/login", auth.Login).Methods("POST")

	// Folder
	folderRouter := router.Path("/api/folder").Subrouter()
	folderRouter.Use(authentication.LoggingMiddleware)
	folderRouter.HandleFunc("", folder.Create).Methods(http.MethodPost)
	// TODO: Add middleware func to parse jwt
}
