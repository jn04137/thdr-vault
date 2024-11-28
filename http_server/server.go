package http_server

import (
	"net/http"
	"database/sql"

	"thdr_vault_go/http_server/handlers"
)

type CustomHttpServer struct {
	dbConn *sql.DB
}

func NewCustomHttpServer(dbConn *sql.DB) *CustomHttpServer {
	return &CustomHttpServer {
		dbConn: dbConn,
	}
}

func (server *CustomHttpServer) GetMuxHandler() *http.ServeMux {
	mux := http.NewServeMux()

	handler := handlers.NewCustomHandler(server.dbConn)

	mux.HandleFunc("GET /", handler.ServeHomePage)
	mux.HandleFunc("GET /vault_page", handler.ServeVaultPage)
	mux.HandleFunc("POST /add_secret", handler.PutKeyValue)

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return mux
}

