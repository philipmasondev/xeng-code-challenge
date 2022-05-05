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
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)         // home page. This has a field to take a url and update db with the response
	mux.Get("/list", handlers.Repo.List)     // returns all records in db as json
	mux.Get("/search", handlers.Repo.Search) // serves the search page

	mux.Post("/", handlers.Repo.PostHome) // processes post of home page when url is provided and submit button is clicked.

	// sub router
	mux.Route("/{id}", func(mux chi.Router) {

		mux.Get("/", handlers.Repo.Get) // GET /search/{name} - return json of recipe by name.
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
