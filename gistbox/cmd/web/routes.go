package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /gist/view/{id}", app.gistView)
	mux.HandleFunc("GET /gist/create", app.gistCreate)
	mux.HandleFunc("POST /gist/create", app.gistCreatePost)

	return mux
}
