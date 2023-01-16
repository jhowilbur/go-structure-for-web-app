package render

import (
	"fmt"
	"html/template"
	"log"
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

var templateCache = make(map[string]*template.Template)

func RenderTemplateWithCache(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if template is already in cache
	_, isInMap := templateCache[t]
	if !isInMap {
		log.Println("Need to create template - Creating template cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Template already in cache")
	}

	// execute template
	tmpl = templateCache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse template
	tmpl, err := template.ParseFiles(templates...) // ... is a variadic parameter, and will take a look to all the parameters
	if err != nil {
		return err
	}

	// add template to cache
	templateCache[t] = tmpl
	return nil
}
