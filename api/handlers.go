package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleGetBooks(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrFetchBooks})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		var genre string
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Edition, &genre, &book.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan for books"})
			return
		}
		book.Genre = strings.Split(genre, ",")
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func HandleAddBooks(c *gin.Context, db *sql.DB) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequestBody})
		return
	}

	// Check if the book already exists (using title and author)
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE title = ? AND author = ?)", newBook.Title, newBook.Author).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check duplicate book", "details": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": ErrDuplicateBook})
		return
	}

	result, err := db.Exec(
		"INSERT INTO books (title, author, published_date, edition, genre, description) VALUES (?, ?, ?, ?, ?, ?)",
		newBook.Title, newBook.Author, newBook.PublishedDate, newBook.Edition, strings.Join(newBook.Genre, ","), newBook.Description,
	)
	if err != nil {
		fmt.Printf("Exec error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add book"})
		return
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve inserted book ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func HandleGetBooksByID(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}

	var book Book
	var genre string
	err = db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Edition, &genre, &book.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrBookNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrFetchBooks})
		return
	}
	book.Genre = strings.Split(genre, ",")

	c.JSON(http.StatusOK, book)
}

func HandleUpdateBookByID(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}

	// Check if the book exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		log.Printf("Error checking for book existence: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check book existence"})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": ErrBookNotFound})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequestBody})
		return
	}

	// Ensure the ID is not changed
	updatedBook.ID = id
	err = updateBook(db, &updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update book"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func HandleDeleteBookByID(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}

	err = deleteBook(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book"})
		return
	}

	c.Status(http.StatusOK)
}
