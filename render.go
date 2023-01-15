package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	error := parsedTemplate.Execute(w, nil)
	if error != nil {
		fmt.Println("Error executing template :", error)
		return
	}
}
