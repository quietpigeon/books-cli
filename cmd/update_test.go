package cmd

import (
	"books-cli/api"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateBook(t *testing.T) {
	originalUpdateBookRequest := UpdateBookRequest
	defer func() { UpdateBookRequest = originalUpdateBookRequest }()

	UpdateBookRequest = func(updatedBook api.Book) (*http.Response, error) {
		assert.Equal(t, 123, updatedBook.ID)
		assert.Equal(t, "Updated Title", updatedBook.Title)
		assert.Equal(t, "Updated Author", updatedBook.Author)
		assert.ElementsMatch(t, []string{"Fiction", "Drama"}, updatedBook.Genre)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader([]byte{})),
		}, nil
	}

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	cmd := updateCmd()
	cmd.SetArgs([]string{
		"123",
		"--title", "Updated Title",
		"--author", "Updated Author",
		"--genre", "Fiction,Drama",
	})

	err := cmd.Execute()
	assert.NoError(t, err)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Book updated successfully!")
}

func TestUpdateBookAPIError(t *testing.T) {
	originalUpdateBookRequest := UpdateBookRequest
	defer func() { UpdateBookRequest = originalUpdateBookRequest }()

	UpdateBookRequest = func(updatedBook api.Book) (*http.Response, error) {
		assert.Equal(t, 456, updatedBook.ID)
		apiError := map[string]string{"error": "Book not found"}
		mockData, _ := json.Marshal(apiError)
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(bytes.NewReader(mockData)),
		}, nil
	}

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	cmd := updateCmd()
	cmd.SetArgs([]string{
		"456",
		"--title", "Nonexistent Book",
	})

	err := cmd.Execute()
	assert.NoError(t, err)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Error updating book: Book not found")
}

func TestUpdateBookInvalidID(t *testing.T) {
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	cmd := updateCmd()
	cmd.SetArgs([]string{"invalid-id"})

	err := cmd.Execute()
	assert.NoError(t, err)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Invalid book ID")
}
