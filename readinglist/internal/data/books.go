package data

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Version   int32     `json:"-"`
}

type BookModel struct {
	DB *sql.DB
}

func (b BookModel) Insert(book *Book) error {
	query := `
	INSERT INTO books(title, published, pages, genere, rating)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, version`
	args := []interface{}{book.Title, book.Published, book.Pages, pq.StringArray(book.Genres), book.Rating}

	return b.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

func (b BookModel) Get(id int64) (*Book, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT id, created_at, title, published, pages, genere, rating, version
	FROM books
	WHERE id=$1`

	var book Book

	err := b.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Published,
		&book.Pages,
		pq.Array(&book.Genres),
		&book.Rating,
		&book.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &book, nil
}

func (b *BookModel) DeleteRow(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
	DELETE FROM books
	WHERE id=$1`
	result, err := b.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (b *BookModel) UpdateRow(book *Book) error {

	query := `
		UPDATE books
		SET (title, published, pages, genere, rating, version) = ($2, $3, $4, $5, $6,version + 1)
		WHERE id=$1 AND version = $7
		RETURNING version`
	args := []interface{}{book.ID, book.Title, book.Published, book.Pages, pq.StringArray(book.Genres), book.Rating, book.Version}
	return b.DB.QueryRow(query, args...).Scan(&book.Version)
}

func (b *BookModel) GetAll() ([]Book, error) {
	query := `
		SELECT id, created_at, title, published, pages, genere, rating, version
		FROM books
		ORDER BY id`

	rows, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		if err := rows.Scan(
			&book.ID,
			&book.CreatedAt,
			&book.Title,
			&book.Published,
			&book.Pages,
			pq.Array(&book.Genres),
			&book.Rating,
			&book.Version); err != nil {
			log.Println(err)
			return nil, err
		}
		books = append(books, book)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err

	}
	return books, nil

}
