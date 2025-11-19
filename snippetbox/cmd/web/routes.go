package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.ServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HanldeFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
