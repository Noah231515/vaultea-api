package main

import (
	"log"
	"net/http"
	"time"

	"vaultea/api/internal/database"
	"vaultea/api/internal/environment"
	"vaultea/api/internal/handlers/auth"

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
	router.HandleFunc("/api/signup", auth.SignUp)
	router.HandleFunc("/api/login", auth.Login)
}
