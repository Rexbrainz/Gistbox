package main

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"snippetbox-webapp/internal/models"
	"snippetbox-webapp/internal/validator"
)


type gistCreateForm struct {
	Title								string `form:"title"`
	Content							string `form:"content"`
	Expires							int		 `form:"expires"`
	validator.Validator		`form:"-"` // Embeded (means gistCreateForm inherits all the fields and methods
																	// of our Validator struct).
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Gists = gists

	app.render(w, r, http.StatusOK, "home.tmpl", data)
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

	data := app.newTemplateData(r)
	data.Gist = gist

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = gistCreateForm{
		Expires: 365,
	}
	app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) gistCreatePost(w http.ResponseWriter, r *http.Request) {
	var form gistCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")
	
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}
	// Pass the data to the GistModel.insert() method, receiving the ID of the new record back.
	id, err := app.gists.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Use the Put() method to add a string value and the corresponding key
	// to the session data.
	app.sessionManager.Put(r.Context(), "flash", "Gist successfully created!")

	// Redirect the user to the relevant page to view the gist.
	http.Redirect(w, r, fmt.Sprintf("/gist/view/%d", id), http.StatusSeeOther)
}
