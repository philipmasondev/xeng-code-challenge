package main

import (
	"net/http"
	"web-application-template/pkg/config"
	"web-application-template/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/", handlers.Repo.Search)

	return mux
}
