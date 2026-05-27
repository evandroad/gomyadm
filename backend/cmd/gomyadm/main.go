package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/evandroad/gomyadm/internal/api"
	"github.com/evandroad/gomyadm/internal/db"
)

func main() {
	r := chi.NewRouter()

	manager := db.NewConnectionManager()
	
	connectionHandler := &api.ConnectionHandler{
		Connections: manager,
	}
	schemaHandler := &api.SchemaHandler{
		Connections: manager,
	}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Post("/api/connections/connect", connectionHandler.Connect)
	r.Post("/api/connections/{id}/disconnect", connectionHandler.Disconnect)
	r.Get("/api/connections/{id}/tables", schemaHandler.ListTables)

	port := ":8181"
	log.Println("server running at http://localhost" + port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}