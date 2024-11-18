package api

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	PublishedDate string   `json:"published_date"`
	Edition       int      `json:"edition"`
	Genre         []string `json:"genre"`
	Description   string   `json:"description"`
}

var db *sql.DB

func InitializeDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./db/books.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the books table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		published_date TEXT,
		edition TEXT,
		genre TEXT, -- Store genres as comma-separated string
		description TEXT
	);
	`
	if _, err = db.Exec(createTableSQL); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create books table: %w", err)
	}

	return db, nil
}

func SaveBook(book *Book) error {
	// Convert genre slice to comma-separated string
	genre := strings.Join(book.Genre, ",")
	result, err := db.Exec("INSERT INTO books(title, author, published_date, edition, genre, description) values(?, ?, ?, ?, ?, ?)",
		book.Title, book.Author, book.PublishedDate, book.Edition, genre, book.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(id)
	return nil
}

func updateBook(db *sql.DB, book *Book) error {
	log.Printf("Updating book: %+v", book)
	_, err := db.Exec(
		"UPDATE books SET title=?, author=?, published_date=?, edition=?, genre=?, description=? WHERE id=?",
		book.Title, book.Author, book.PublishedDate, book.Edition, strings.Join(book.Genre, ","), book.Description, book.ID,
	)
	return err
}

func deleteBook(db *sql.DB, bookID int) error {
	_, err := db.Exec("DELETE FROM books WHERE id=?", bookID)
	if err != nil {
		log.Printf("Error executing delete query: %v", err)
	}
	return err
}
