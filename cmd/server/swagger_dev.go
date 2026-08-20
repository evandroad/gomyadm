//go:build dev

package main

import (
	"github.com/go-chi/chi/v5"
	_ "gomyadm/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func setupSwagger(r chi.Router) {
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}