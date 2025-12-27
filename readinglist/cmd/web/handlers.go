package main

import (
	"fmt"
	"net/http"
	"readinglist/internal/models"
	"strconv"
)

type envelope map[string]any

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status": "ok",
	}
	err := app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	books, err := app.readinglist.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "<html><head><title>Reading List</title></head><body><h1>Reading List</h1><ul>")
	for _, book := range *books {
		fmt.Fprintf(w, "<li>%s (%d)</li>", book.Title, book.Pages)
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	book, err := app.readinglist.Get(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s (%d)", book.Title, book.Pages)
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.bookCreateForm(w, r)
	case http.MethodPost:
		app.bookCreateProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	book := models.Book{
		Title:     "Stranger Things",
		Published: 2015,
		Genres:    []string{"Fantasy"},
		Pages:     175,
		Rating:    5,
	}

	err := app.readinglist.Create(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}

func (app *application) bookCreateForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Create Book</title></head></html>")
}

func (app *application) bookCreateProcess(w http.ResponseWriter, r *http.Request) {

}
