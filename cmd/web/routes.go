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
	//mux.Use(NoSurf)
	//mux.Use(SessionLoad)

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/search", handlers.Repo.Search)

	mux.Post("/home", handlers.Repo.PostHome)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}