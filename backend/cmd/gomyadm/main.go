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
	_ "github.com/evandroad/gomyadm/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Database Manager API
// @version 1.0
// @description API para gerenciamento de bancos de dados.
// @host localhost:8181
// @BasePath /api/
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
	itemHandler := &api.ItemHandler{ Connection: manager }
	columnHandler := &api.ColumnHandler{ Connection: manager }
	queryHandler := &api.QueryHandler{ Connection: manager }

	r.Get("/health", health)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/connection", func(r chi.Router) {
			r.Post("/connect", connectionHandler.Connect)
			r.Post("/disconnect", connectionHandler.Disconnect)
			r.Get("/active", connectionHandler.Active)

			r.Get("/", connectionHandler.GetAll)
			r.Post("/", connectionHandler.Insert)
			r.Put("/", connectionHandler.Update)
			r.Delete("/{id}", connectionHandler.Delete)	
		})

		r.Post("/database/select", connectionHandler.SelectDatabase)

		r.Route("/tables", func(r chi.Router) {
			r.Get("/", schemaHandler.ListTables)
			
			r.Route("/item", func(r chi.Router) {
				r.Get("/{table}", itemHandler.GetAll)
				r.Post("/", itemHandler.Insert)
				r.Put("/", itemHandler.Update)
				r.Delete("/", itemHandler.Delete)
			})
			
			r.Route("/column", func(r chi.Router) {
				r.Get("/{table}", columnHandler.GetAll)
				r.Post("/", columnHandler.Insert)
				r.Put("/", columnHandler.Update)
				r.Delete("/{table}/{column}", columnHandler.Delete)
			})
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
	Success(w, http.StatusOK, "API ok.", nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	Error(w, http.StatusNotFound, "rota não encontrada", H{ "path": r.URL.Path })
}