package main

import (
	"net/http"

	"github.com/justinas/alice"
)


func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Create a new middleware chain containing the middleware specific to our
	// dynamic application routes.
	dynamic := alice.New(app.sessionManager.LoadAndSave)


	// update routes to use the new dynamic middleware chain followed by
	// the appropriate handler function
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /gist/view/{id}", dynamic.ThenFunc(app.gistView))
	mux.Handle("GET /gist/create", dynamic.ThenFunc(app.gistCreate))
	mux.Handle("POST /gist/create", dynamic.ThenFunc(app.gistCreatePost))

	// Create a middleware chain containing our standard middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(mux)
}
