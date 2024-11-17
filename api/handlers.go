package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleGetBooks(c *gin.Context) {
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

func HandleAddBooks(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequestBody})
		return
	}

	// Check if the book already exists (using title and author)
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM books WHERE title = ? AND author = ?", newBook.Title, newBook.Author).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check for duplicate book"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": ErrDuplicateBook})
		return
	}

	err = SaveBook(&newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newBook.ID})
}

func HandleGetBooksByID(c *gin.Context) {
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

func HandleUpdateBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequestBody})
		return
	}

	// Ensure the ID is not changed
	updatedBook.ID = id
	err = updateBook(&updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update book"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func HandleDeleteBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}

	err = deleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book"})
		return
	}

	c.Status(http.StatusOK)
}
