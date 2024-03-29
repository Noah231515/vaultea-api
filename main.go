package main

import (
	"log"
	"net/http"
	"time"

	"vaultea/api/internal/database"
	"vaultea/api/internal/environment"
	"vaultea/api/internal/handlers/auth"
	"vaultea/api/internal/handlers/folder"
	"vaultea/api/internal/handlers/password"
	"vaultea/api/internal/handlers/user_preferences"
	"vaultea/api/internal/middleware"
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
	jwtMiddleware := jwtmiddleware.New(validator.ValidateToken)

	// Auth
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.UserDataMiddleware)

	authRouter.HandleFunc("/signup", auth.SignUp).Methods("POST")
	authRouter.HandleFunc("/login", auth.Login).Methods("POST")

	// Folder
	folderRouter := router.PathPrefix("/api/folder").Subrouter()
	folderRouter.Use(jwtMiddleware.CheckJWT, middleware.FolderDataMiddleware)

	folderRouter.HandleFunc("", folder.Create).Methods(http.MethodPost)
	folderRouter.HandleFunc("/{folderId:[0-9]+}", folder.Update).Methods(http.MethodPut)
	folderRouter.HandleFunc("/{folderId:[0-9]+/updateStarred}", folder.UpdateStarred).Methods(http.MethodPut)
	folderRouter.HandleFunc("/{folderId:[0-9]+}", folder.Delete).Methods(http.MethodDelete)

	// Password
	passwordRouter := router.PathPrefix("/api/password").Subrouter()
	passwordRouter.Use(jwtMiddleware.CheckJWT, middleware.PasswordDataMiddleware)

	passwordRouter.HandleFunc("", password.Create).Methods(http.MethodPost)
	passwordRouter.HandleFunc("/{passwordId:[0-9]+}", password.Update).Methods(http.MethodPut)
	passwordRouter.HandleFunc("/{passwordId:[0-9]+/updateStarred}", password.UpdateStarred).Methods(http.MethodPut)
	passwordRouter.HandleFunc("/{passwordId:[0-9]+}", password.Delete).Methods(http.MethodDelete)

	// User Preferences
	userPreferencesRouter := router.PathPrefix("/api/userPreferences").Subrouter()
	userPreferencesRouter.Use(jwtMiddleware.CheckJWT)
	userPreferencesRouter.HandleFunc("/toggleVaultView", user_preferences.ToggleVaultView).Methods(http.MethodPut)
}
