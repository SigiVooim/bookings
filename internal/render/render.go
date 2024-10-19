package render

import (
	"bytes"
	"github.com/sigivooim/bookings/internal/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewRenderer sets the config for the render package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// Template renders templates using html/templates
func Template(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("Error creating template cache:", err)
			return
		}
	}

	// get request template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache: ", tmpl)
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Fatal("Error executing:", err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing page:", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	// get all of the files named *.page.gothml from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return tmplCache, err
	}

	// range through all files ending with *.page.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		// get all of the files ending with *.layout.gohtml
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return tmplCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return tmplCache, err
			}
		}
		tmplCache[name] = ts
	}

	return tmplCache, nil
}
