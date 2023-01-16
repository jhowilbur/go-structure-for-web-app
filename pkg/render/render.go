package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	error := parsedTemplate.Execute(w, nil)
	if error != nil {
		fmt.Println("Error executing template :", error)
		return
	}
}
