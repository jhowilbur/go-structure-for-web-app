package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplateWithCache(w http.ResponseWriter, tmpl string) {
	// create template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	templateWanted, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)               // execute values get from map and execute that directly
	err = templateWanted.Execute(buffer, nil) // to help identify if it has an error
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// it will execute and show the template again if uncomment
	// this way don't cache the tmpl

	//parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//error := parsedTemplate.Execute(w, nil)
	//if error != nil {
	//	fmt.Println("Error executing template :", error)
	//	return
	//}
}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{} // same functionality like above.

	// get all the files named *.page.tmpl from ./tempaltes
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)                             // to get the last element in the path (file tmpl)
		templateSet, err := template.New(name).ParseFiles(page) // populate name with HTML

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
