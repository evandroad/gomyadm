package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/evandroad/gomyadm/internal/api"
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/storage"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/router"
)

func main() {
	store := storage.GetConnectionsStore()
	err := store.Init()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(router.CORS)
	r.Use(router.Recovery)
	r.Use(router.Logger)

	manager := db.NewConnectionManager()
	
	connectionHandler := &api.ConnectionHandler{ Connection: manager }
	schemaHandler := &api.SchemaHandler{ Connection: manager }
	queryHandler := &api.QueryHandler{ Connection: manager }

	r.Get("/health", health)

	r.Route("/api", func(r chi.Router) {
		r.Route("/connection", func(r chi.Router) {
			r.Post("/connect", connectionHandler.Connect)
			r.Post("/disconnect", connectionHandler.Disconnect)
			r.Get("/active", connectionHandler.Active)

			r.Get("/", connectionHandler.GetAllConnections)
			r.Post("/", connectionHandler.InsertConnection)
			r.Put("/", connectionHandler.UpdateConnection)
			r.Delete("/{id}", connectionHandler.DeleteConnection)	
		})

		r.Post("/database/select", connectionHandler.SelectDatabase)

		r.Route("/tables", func(r chi.Router) {
			r.Get("/", schemaHandler.ListTables)
			r.Get("/{table}", schemaHandler.SelectTable)
			r.Get("/struct/{table}", schemaHandler.DescribeTable)
			r.Post("/", schemaHandler.InsertValue)
			r.Put("/", schemaHandler.UpdateValue)
			r.Delete("/", schemaHandler.DeleteValue)
		})

		r.Post("/query", queryHandler.ExecuteQuery)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { notFound(w, r) })

	port := ":8181"
	log.Println("server running at http://localhost" + port)

	err = http.ListenAndServe(port, r)
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