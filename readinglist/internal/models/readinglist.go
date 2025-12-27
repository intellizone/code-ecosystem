package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float32  `json:"rating"`
}

type BookResponse struct {
	Book *Book `json:"book"`
}

type BooksResponse struct {
	Books *[]Book `json:"books"`
}

type ReadinglistModel struct {
	Endpoint string
}

func (m *ReadinglistModel) GetAll() (*[]Book, error) {
	res, err := http.Get(m.Endpoint)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var booksResponse BooksResponse
	err = json.Unmarshal(data, &booksResponse)
	if err != nil {
		return nil, err
	}
	return booksResponse.Books, nil
}

func (m *ReadinglistModel) Get(id int) (*Book, error) {
	url := fmt.Sprintf("%s/%d", m.Endpoint, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var bookResp BookResponse
	err = json.Unmarshal(data, &bookResp)
	if err != nil {
		return nil, err
	}

	return bookResp.Book, nil
}

func (m *ReadinglistModel) Create(book Book) error {
	data, err := json.Marshal(book)
	if err != nil {
		return err
	}
	res, err := http.Post(m.Endpoint, "", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %v", res.Status)
	}
	return nil
}
