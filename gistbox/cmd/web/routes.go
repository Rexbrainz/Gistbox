package main

import (
	"net/http"

	"github.com/justinas/alice"
)


func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /gist/view/{id}", app.gistView)
	mux.HandleFunc("GET /gist/create", app.gistCreate)
	mux.HandleFunc("POST /gist/create", app.gistCreatePost)

	// Create a middleware chain containing our standard middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(mux)
}
