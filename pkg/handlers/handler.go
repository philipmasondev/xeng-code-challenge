package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"web-application-template/pkg/config"
	"web-application-template/pkg/models"
	"web-application-template/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	resp, err := http.Get("http://www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata")
	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	data["apiReponse"] = string(text)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
