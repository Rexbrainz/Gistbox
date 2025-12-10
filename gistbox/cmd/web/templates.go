package main

import (
	"html/template"
	"path/filepath"
	"time"

 	"snippetbox-webapp/internal/models"
 )

// Define a templateData type to act as the holding structure for 
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll andd more 
// to it as the build progresses.
type templateData struct {
	CurrentYear int
	Gist 				models.Gist
	Gists 			[]models.Gist
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize  a new map to act as the cache
		cache := map[string]*template.Template{}

		// Use filepath.Glob() function to get a slice of all filepaths that match the pattern.
		pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
		if err != nil {
			return nil, err
		}

		for _, page := range pages {
			// Extract the file name from the filepath
			name := filepath.Base(page)

			// Parse the base template file into a template set.
			ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
			if err != nil {
				return nil, err
			}

			// Call ParseGlob() *on this template set* to add any partials.
			ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
			if err != nil {
				return nil, err
			}

		// parse the files into a template set.
			ts, err = ts.ParseFiles(page)
			if err != nil {
				return nil, err
			}

			// Add the template set to the map, using the name of the page.
			cache[name] = ts
		}
		return cache, nil
}


