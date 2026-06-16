package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/evandroad/gomyadm/internal/api"
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/router"
	"github.com/evandroad/gomyadm/internal/services/column"
	"github.com/evandroad/gomyadm/internal/services/connection"
	"github.com/evandroad/gomyadm/internal/services/session"
	. "github.com/evandroad/gomyadm/internal/respond"

	_ "github.com/evandroad/gomyadm/docs"
	"github.com/swaggo/http-swagger"
)

// @title Database Manager API
// @version 1.0
// @description API para gerenciamento de bancos de dados.
// @host localhost:8181
// @BasePath /api/
func main() {
	err := connectionService.Init()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(router.CORS)
	r.Use(router.Recovery)
	r.Use(router.Logger)

	manager := db.NewConnectionManager()

	sessionService := sessionService.NewSessionService(manager)
	connectionService := connectionService.GetStore()
	columnService := columnService.NewColumnService(manager)
	
	sessionHandler := api.NewSessionHandler(sessionService)
	connectionHandler := api.NewConnectionHandler(connectionService)
	databaseHandler := &api.DatabaseHandler{ Connection: manager }
	tableHandler := &api.TableHandler{ Connection: manager }
	itemHandler := &api.ItemHandler{ Connection: manager }
	columnHandler := api.NewColumnHandler(columnService)
	queryHandler := &api.QueryHandler{ Connection: manager }

	r.Get("/health", health)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/session", func(r chi.Router) {
			r.Post("/", sessionHandler.Connect)
			r.Delete("/", sessionHandler.Disconnect)
			r.Get("/", sessionHandler.Active)
		})

		r.Route("/connection", func(r chi.Router) {
			r.Get("/", connectionHandler.GetAll)
			r.Post("/", connectionHandler.Create)
			r.Put("/", connectionHandler.Update)
			r.Delete("/{id}", connectionHandler.Delete)	
		})

		r.Post("/database/select", databaseHandler.SelectDatabase)

		r.Route("/table", func(r chi.Router) {
			r.Get("/", tableHandler.GetAll)
			r.Post("/", tableHandler.Create)
			
			r.Route("/item", func(r chi.Router) {
				r.Get("/{table}", itemHandler.GetAll)
				r.Post("/", itemHandler.Create)
				r.Put("/", itemHandler.Update)
				r.Delete("/", itemHandler.Delete)
			})
			
			r.Route("/column", func(r chi.Router) {
				r.Get("/{table}", columnHandler.GetAll)
				r.Post("/", columnHandler.Create)
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