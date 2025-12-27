package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/api/v1/books", app.getCreateBooks)
	mux.HandleFunc("/api/v1/books/", app.updateDeleteBooks)
	return mux
}
