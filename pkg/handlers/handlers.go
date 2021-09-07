package handlers

import (
	"net/http"

	"github.com/tatyanayagolnikov/bookings/pkg/config"
	"github.com/tatyanayagolnikov/bookings/pkg/models"
	"github.com/tatyanayagolnikov/bookings/pkg/render"
)

// Repo - the repository used by the handlers
var Repo *Repository 

// Repository - is the Repository type
type Repository struct {
	App *config.AppConfig
}
 
// NewRepo - creates a new repository 
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers - sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r 
}

// ----- HANDLER func ----- 
// In order for a function to respond to a request
// from a web browser, it has to handle 2 parameters,
// first- a response writer 
// second- has to handle a request



// Home - is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

// About - is the about page HANDLER  
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again, Tanya programmer"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Person - is the person page handler
func (m *Repository) Person(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "person.page.tmpl.html", &models.TemplateData{})
}
