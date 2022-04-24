package main

import (
	"net/http"
	"web-application-template/pkg/config"

	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := -chi.NewRouter()

	return mux
}
