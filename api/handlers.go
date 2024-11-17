package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func HandleAddBooks(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequestBody})
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	SaveBooks()
	c.JSON(http.StatusOK, gin.H{"id": newBook.ID})
}

func HandleGetBooksByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": ErrInvalidRequestBody})
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
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			books[i].ID = id
			SaveBooks()
			c.JSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": ErrInvalidRequestBody})
}

func HandleDeleteBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBookID})
		return
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			SaveBooks()
			c.Status(http.StatusOK)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": ErrInvalidRequestBody})
}
