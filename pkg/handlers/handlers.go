package handlers

import (
	"github.com/jhowilbur/golang-web-app/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateWithCache(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateWithCache(w, "about.page.tmpl")
}
