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
	
	handler := &api.Handler{
		Manager: manager,
	}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Post("/api/connections/test", handler.TestConnection)

	port := ":8181"
	log.Println("server running at " + port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}