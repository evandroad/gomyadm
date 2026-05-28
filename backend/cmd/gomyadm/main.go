package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/evandroad/gomyadm/internal/api"
	"github.com/evandroad/gomyadm/internal/db"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/router"
)

func main() {
	r := chi.NewRouter()
	r.Use(router.CORS)
	r.Use(router.Recovery)
	r.Use(router.Logger)

	manager := db.NewConnectionManager()
	
	connectionHandler := &api.ConnectionHandler{ Connection: manager }
	schemaHandler := &api.SchemaHandler{ Connection: manager }

	r.Get("/health", health)
	r.Post("/api/connection/connect", connectionHandler.Connect)
	r.Post("/api/connection/disconnect", connectionHandler.Disconnect)
	r.Get("/api/connection", connectionHandler.Active)
	r.Post("/api/connection/database/select", connectionHandler.SelectDatabase)
	r.Get("/api/connection/tables", schemaHandler.ListTables)
	r.Get("/api/connection/tables/{table}", schemaHandler.DescribeTable)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) { notFound(w, r) })

	port := ":8181"
	log.Println("server running at http://localhost" + port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	Error(w, http.StatusNotFound, "rota não encontrada", H{ "path": r.URL.Path })
}