package main

import (
	"fmt"
	"github.com/jhowilbur/golang-web-app/pkg/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const portNUmber = ":8080"

func main() {

	http.Handle("/metrics", promhttp.Handler()) // Prometheus metrics
	handlers.RecordMetrics()                    // Prometheus custom metrics

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("Server starting on port %s", portNUmber))
	_ = http.ListenAndServe(portNUmber, nil)

}
