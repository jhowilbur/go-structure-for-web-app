package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNUmber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	error := parsedTemplate.Execute(w, nil)
	if error != nil {
		fmt.Println("Error executing template :", error)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server starting on port %s", portNUmber))
	_ = http.ListenAndServe(portNUmber, nil)

}
