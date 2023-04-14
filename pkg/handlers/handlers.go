package handlers

import (
	"github.com/jhowilbur/golang-web-app/pkg/config"
	"github.com/jhowilbur/golang-web-app/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(applicationConfig *config.AppConfig) *Repository {
	return &Repository{
		App: applicationConfig,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateWithCache(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateWithCache(w, "about.page.tmpl")
}
