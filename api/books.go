package api

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID            int
	Title         string
	Author        string
	PublishedDate string
	Edition       int
	Genre         []string
	Description   string
}

var db *sql.DB

func InitializeDB() {
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
		published_date   
 TEXT,
		edition TEXT,
		genre TEXT, -- Store genres as comma-separated string
		description TEXT
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
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

func updateBook(book *Book) error {
	genre := strings.Join(book.Genre, ",")
	_, err := db.Exec("UPDATE books SET title=?, author=?, published_date=?, edition=?, genre=?, description=? WHERE id=?",
		book.Title, book.Author, book.PublishedDate, book.Edition, genre, book.Description, book.ID)
	return err
}

func deleteBook(bookID int) error {
	_, err := db.Exec("DELETE FROM books WHERE id=?", bookID)
	return err
}
