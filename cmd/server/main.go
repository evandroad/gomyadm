package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"gomyadm/internal/api"
	"gomyadm/internal/db"
	. "gomyadm/internal/respond"
	"gomyadm/internal/router"
	"gomyadm/internal/services"

	_ "gomyadm/docs"
)

//go:embed all:web
var webFS embed.FS

// @title Database Manager API
// @version 1.0
// @description API para gerenciamento de bancos de dados.
// @host localhost:8181
// @BasePath /api/
func main() {
	err := services.LoadConnections()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(router.CORS)
	r.Use(router.Recovery)
	r.Use(router.Logger)

	manager := db.NewConnectionManager()

	sessionService    := services.NewSessionService(manager)
	connectionService := services.NewConnectionStore()
	databaseService   := services.NewDatabaseService(manager)
	columnService     := services.NewColumnService(manager)
	queryService      := services.NewQueryService(manager)
	itemService       := services.NewItemService(manager)
	tableService      := services.NewTableService(manager)
	appService        := services.NewAppService()
	
	sessionHandler    := api.NewSessionHandler(sessionService)
	connectionHandler := api.NewConnectionHandler(connectionService)
	databaseHandler   := api.NewDatabaseHandler(databaseService)
	tableHandler      := api.NewTableHandler(tableService)
	itemHandler       := api.NewItemHandler(itemService)
	columnHandler     := api.NewColumnHandler(columnService)
	queryHandler      := api.NewQueryHandler(queryService)
	appHandler        := api.NewAppHandler(appService)

	r.Get("/health", health)
	setupSwagger(r)
	
	r.Route("/api", func(r chi.Router) {
		r.Get("/version", appHandler.Version)

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

		r.Route("/database", func(r chi.Router) {
			r.Get("/", databaseHandler.GetAll)
			r.Post("/select", databaseHandler.Select)
			r.Post("/", databaseHandler.Create)
			r.Put("/", databaseHandler.Update)
			r.Delete("/{name}", databaseHandler.Delete)
		})

		r.Route("/table", func(r chi.Router) {
			r.Get("/", tableHandler.GetAll)
			r.Post("/", tableHandler.Create)
			r.Put("/{oldName}/{newName}", tableHandler.Update)
			r.Delete("/{table}", tableHandler.Delete)
			
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

		r.NotFound(func(w http.ResponseWriter, r *http.Request) { notFound(w, r) })
	})

	sub, _ := fs.Sub(webFS, "web")
	r.Get("/*", spaHandler(sub))

	port := ":8181"
	log.Println("server running at http://localhost" + port)
	log.Println("swagger at http://localhost" + port + "/swagger/index.html")

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

func spaHandler(sub fs.FS) http.HandlerFunc {
	fileServer := http.FileServer(http.FS(sub))

	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if f, err := sub.Open(path); err == nil {
			f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		r.URL.Path = "/"
		fileServer.ServeHTTP(w, r)
	}
}