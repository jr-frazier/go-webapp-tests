package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// registar middleware
	mux.Use(middleware.Recoverer)

	// register routes
	mux.Get("/", app.Home)

	// static assets

	return mux
}
