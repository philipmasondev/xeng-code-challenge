package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"web-application-template/pkg/config"
	"web-application-template/pkg/driver"
	"web-application-template/pkg/handlers"
	"web-application-template/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	db, err := run()
	if err != nil {
		fmt.Errorf("run: failed to execute run in main: %w", err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal("Error at srv.ListenAndServe.", err)
	fmt.Println("Error at srv.ListenAndServe.", err)
}

func run() (*driver.DB, error) {
	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=xeng-coding-challenge user=postgres password=" + DBPassword)

	if err != nil {
		fmt.Errorf("ConnectSQL: failed to connect to database. %v", err)
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Errorf("CreateTemplateCache: Can not create template cache. %v", err)
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	return db, nil
}
