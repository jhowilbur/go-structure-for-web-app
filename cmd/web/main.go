package main

import (
	"fmt"
	"github.com/jhowilbur/golang-web-app/pkg/handlers"
	"net/http"
)

const portNUmber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server starting on port %s", portNUmber))
	_ = http.ListenAndServe(portNUmber, nil)

}
