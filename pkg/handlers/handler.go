package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"web-application-template/pkg/config"
	"web-application-template/pkg/models"
	"web-application-template/pkg/render"
	"web-application-template/pkg/repository"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
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

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// PostHome is the handler for the home page
func (m *Repository) PostHome(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	apiJsonString := apiResponse("http://www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata")

	data["apiResponse"] = apiJsonString
	data["test"] = "test"

	apiForDB := models.RecipiesJSON{
		RecipeJsonString: apiJsonString,
	}

	err := m.DB.InsertRecipiesToDB(apiForDB)
	if err != nil {
		fmt.Println("issue with db insert in PostHome handler", err)
	}

	m.App.Session.Put(r.Context(), "apiForDb", apiForDB)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// PostSearch is the handler to serve the search page
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search.page.tmpl", &models.TemplateData{
		//
	})
}

// Search is the handler to serve the search page after POST
func (m *Repository) PostSearch(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search.page.tmpl", &models.TemplateData{
		//
	})
}

// apiResponse return api response from provided url
func apiResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(text)
}
