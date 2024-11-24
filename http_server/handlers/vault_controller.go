package handlers

import (
	"log"
	"net/http"
	"html/template"
	"database/sql"

	"thdr_vault_go/repositories"
)

type CustomHandler struct {
	SecretRepo *repositories.SecretRepo
}

func NewCustomHandler(DBConn *sql.DB) *CustomHandler {
	return &CustomHandler{
		SecretRepo: repositories.CreateSecretRepo(DBConn),
	}
}

func (config *CustomHandler) ServeHomePage(w http.ResponseWriter, r *http.Request) {
	filepath := "templates/home.html"
	t, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Couldn't parse the following file: %s", err.Error())
	}
	err = t.ExecuteTemplate(w, "main", nil)
	if err != nil {
		log.Printf("Exception executing template: %s", err.Error());
	}
}

type ServeVaultPageData struct {
	Secrets []repositories.Secret
}

func (handler *CustomHandler) ServeVaultPage(w http.ResponseWriter, r *http.Request) {
	repo := handler.SecretRepo

	secrets, repoErr := repo.GetSecrets()
	if repoErr != nil {
		log.Printf("There was an error getting secrets from repo: %v", repoErr.Error())
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	pageData := ServeVaultPageData{
		Secrets: secrets,
	}

	filepath := "templates/vault_page.html"
	t, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Couldn't parse the following file: %s", err.Error())
	}
	err = t.ExecuteTemplate(w, "vault_page", pageData)
	if err != nil {
		log.Printf("Exception executing template: %s", err.Error());
	}
}

func (config *CustomHandler) PutKeyValue(w http.ResponseWriter, r *http.Request) {
	repo := config.SecretRepo

	key := r.FormValue("key")
	secret := r.FormValue("secret")
	err := repo.PutSecret(key, secret)
	if err != nil {
		log.Printf("Error occurred: %v", err.Error())
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (config *CustomHandler) GetKeyValues(w http.ResponseWriter, r *http.Request) {

}


