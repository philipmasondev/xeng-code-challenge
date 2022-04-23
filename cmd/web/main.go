package main

import (
	"fmt"
	"log"
	"net/http"
	"web-application-template/pkg/config"
	"web-application-template/pkg/driver"
	"web-application-template/pkg/handlers"
	"web-application-template/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// connect to db
	log.Println("connecting to db")
	db, err := driver.ConnectSQL("host=127.0.0.1 port=5432 dbname=postgres user=postgres password=0O9i8u7y6t5r4e3w@q1")
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}

	defer db.SQL.Close()

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
}
