package main

import (
	"fmt"
	//"html/template"
	"net/http"
	"strconv"
	"errors"
	"snippetbox-webapp/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, gist := range gists {
		fmt.Fprintf(w, "%+v\n", gist)
	}

	/*
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	*/
}

func (app *application) gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	gist, err := app.gists.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r) // 404 NotFound error 
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Written as a plain text HTTP response body.
	fmt.Fprintf(w, "%v", gist)
}

func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {
	
}

func (app *application) gistCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// Pass the data to the GistModel.insert() method, receiving the ID of the new record back.
	id , err := app.gists.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page to view the gist.
	http.Redirect(w, r, fmt.Sprintf("/gist/view/%d", id), http.StatusSeeOther)
}
