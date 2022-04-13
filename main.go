package main

import (
	"log"
	"net/http"
	"time"

	"vaultea/api/internal/database"
	"vaultea/api/internal/environment"
	"vaultea/api/internal/handlers/auth"
	"vaultea/api/internal/handlers/folder"
	crypto_utils "vaultea/api/internal/utils/crypto"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
)

func main() {
	environment.SetEnv()
	router := mux.NewRouter()
	database.ConnectToDatabase()
	database.MakeMigrations()
	validator, err := crypto_utils.InitJwtValidator()

	if err == nil {
		initRoutes(router, validator)
		srv := &http.Server{
			Handler:      router,
			Addr:         "127.0.0.1:8081", // TODO: make configurable
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
	} else {
		log.Fatal(err)
	}

}

func initRoutes(router *mux.Router, validator *validator.Validator) {
	middleware := jwtmiddleware.New(validator.ValidateToken)

	// Auth
	router.HandleFunc("/api/signup", auth.SignUp).Methods("POST")
	router.HandleFunc("/api/login", auth.Login).Methods("POST")

	// Folder
	folderRouter := router.Path("/api/folder").Subrouter()
	folderRouter.Use(middleware.CheckJWT)

	folderRouter.HandleFunc("", folder.Create).Methods(http.MethodPost)
	// TODO: Add middleware func to parse jwt
}
