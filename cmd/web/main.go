package main

import (
	"fmt"
	"github.com/jhowilbur/golang-web-app/pkg/config"
	"github.com/jhowilbur/golang-web-app/pkg/handlers"
	"github.com/jhowilbur/golang-web-app/pkg/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const portNUmber = ":8080"

func main() {

	// calling template cache from config
	var app config.AppConfig
	templateCache, err := render.CreateTemplate()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// handle endpoints
	http.Handle("/metrics", promhttp.Handler()) // Prometheus metrics
	handlers.RecordMetrics()                    // Prometheus custom metrics

	//http.HandleFunc("/", repo.Home)
	//http.HandleFunc("/about", repo.About)

	log.Println(fmt.Sprintf("Server starting on port %s", portNUmber))
	//_ = http.ListenAndServe(portNUmber, nil)

	serve := &http.Server{
		Addr:    portNUmber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
