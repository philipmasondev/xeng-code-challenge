package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"web-application-template/pkg/config"
	"web-application-template/pkg/driver"
	"web-application-template/pkg/models"
	"web-application-template/pkg/render"
	"web-application-template/pkg/repository"
	"web-application-template/pkg/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// PostHome is the handler for the home page
func (m *Repository) PostHome(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	apiJsonString := r.FormValue("url_of_api")

	// check url format and if http:// is missing, add http:// to beginning
	_, err := url.ParseRequestURI(apiJsonString)
	if err != nil {
		apiJsonString = "http://" + apiJsonString
	}

	urlString := apiResponse(apiJsonString)

	fmt.Println(apiJsonString)

	if json.Valid([]byte(urlString)) {
		errInsert := m.DB.InsertRecipe(urlString)
		if errInsert != nil {
			fmt.Println("Error at InsertRecipe with ", errInsert)
		}
	}

	render.Template(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// PostSearch is the handler to serve the search page
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "search.page.tmpl", &models.TemplateData{
		//
	})
}

// Search is the handler to serve the search page after POST
func (m *Repository) PostSearch(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "search.page.tmpl", &models.TemplateData{
		//
	})
}

// apiResponse return api response from provided url
func apiResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error at apiResponse. Error - ", err)
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error at apiResponse. Error - ", err)
	}
	return string(text)
}
