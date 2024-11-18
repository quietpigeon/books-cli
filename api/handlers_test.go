package api_test

import (
	"books-cli/api"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGetBooks(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	mock.ExpectQuery("SELECT \\* FROM books").WillReturnRows(
		sqlmock.NewRows([]string{"id", "title", "author", "published_date", "edition", "genre", "description"}).
			AddRow(1, "Book One", "Author One", "2024-01-01", 1, "Fiction,Adventure", "A great book"),
	)

	router := gin.Default()
	router.GET("/books", func(c *gin.Context) {
		api.HandleGetBooks(c, mockDB)
	})

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book One")
}

func TestHandleAddBooks(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM books WHERE title = \\? AND author = \\?\\)").
		WithArgs("Book One", "Author One").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("INSERT INTO books \\(title, author, published_date, edition, genre, description\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)").
		WithArgs("Book One", "Author One", "2024-01-01", 1, "Fiction,Adventure", "A great book").
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := gin.Default()
	router.POST("/books", func(c *gin.Context) {
		api.HandleAddBooks(c, mockDB)
	})

	body := `{
		"title": "Book One",
		"author": "Author One",
		"published_date": "2024-01-01",
		"edition": 1,
		"genre": ["Fiction", "Adventure"],
		"description": "A great book"
	}`
	req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestHandleGetBooksByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT \\* FROM books WHERE id = \\?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "published_date", "edition", "genre", "description"}).
			AddRow(1, "Book One", "Author One", "2024-01-01", 1, "Fiction,Adventure", "A great book"))

	router := gin.Default()
	router.GET("/books/:id", func(c *gin.Context) {
		api.HandleGetBooksByID(c, mockDB)
	})

	req, _ := http.NewRequest(http.MethodGet, "/books/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"id": 1,
		"title": "Book One",
		"author": "Author One",
		"published_date": "2024-01-01",
		"edition": 1,
		"genre": ["Fiction", "Adventure"],
		"description": "A great book"
	}`, w.Body.String())

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestHandleUpdateBookByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	// Simulate book exists.
	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM books WHERE id = \\?\\)").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectExec("UPDATE books SET title=\\?, author=\\?, published_date=\\?, edition=\\?, genre=\\?, description=\\? WHERE id=\\?").
		WithArgs("Updated Title", "Updated Author", "2024-12-01", 2, "Fiction,Drama", "An updated description", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	router := gin.Default()
	router.PUT("/books/:id", func(c *gin.Context) {
		api.HandleUpdateBookByID(c, mockDB)
	})

	body := `{
		"title": "Updated Title",
		"author": "Updated Author",
		"published_date": "2024-12-01",
		"edition": 2,
		"genre": ["Fiction", "Drama"],
		"description": "An updated description"
	}`
	req, _ := http.NewRequest(http.MethodPut, "/books/1", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"id": 1,
		"title": "Updated Title",
		"author": "Updated Author",
		"published_date": "2024-12-01",
		"edition": 2,
		"genre": ["Fiction", "Drama"],
		"description": "An updated description"
	}`, w.Body.String())

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestHandleDeleteBookByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectExec(`DELETE FROM books WHERE id\s*=\s*\?`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	router := gin.Default()
	router.DELETE("/books/:id", func(c *gin.Context) {
		api.HandleDeleteBookByID(c, mockDB)
	})

	req, _ := http.NewRequest(http.MethodDelete, "/books/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestHandleDeleteBookByID_InvalidID(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	router := gin.Default()
	router.DELETE("/books/:id", func(c *gin.Context) {
		api.HandleDeleteBookByID(c, mockDB)
	})

	req, err := http.NewRequest(http.MethodDelete, "/books/abc", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":"invalid book ID"}`, w.Body.String())
}
