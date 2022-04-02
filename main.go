package main

import (
	"log"
	"net/http"
	"time"

	"vaultea/api/internal/database"
	"vaultea/api/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	database.ConnectToDatabase()
	database.MakeMigrations()
	initRoutes(router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/api/signup", handlers.SignUp)
}
