package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveBook(t *testing.T) {
	RemoveBookRequest = func(bookID int) (*http.Response, error) {
		assert.Equal(t, 123, bookID)
		mockResp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader([]byte{})),
		}
		return mockResp, nil
	}

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	os.Stdout = w

	cmd := removeCmd()
	cmd.SetArgs([]string{"123"})

	err := cmd.Execute()
	assert.NoError(t, err)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Book removed successfully!")
}

func TestRemoveBookInvalidID(t *testing.T) {
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	os.Stdout = w

	cmd := removeCmd()
	cmd.SetArgs([]string{"invalid-id"})

	err := cmd.Execute()
	assert.NoError(t, err)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Invalid book ID")
}

func TestRemoveBookAPIError(t *testing.T) {
	RemoveBookRequest = func(bookID int) (*http.Response, error) {
		assert.Equal(t, 456, bookID)
		apiError := map[string]string{"error": "Book not found"}
		mockData, _ := json.Marshal(apiError)
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(bytes.NewReader(mockData)),
		}, nil
	}

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	os.Stdout = w

	cmd := removeCmd()
	cmd.SetArgs([]string{"456"})

	err := cmd.Execute()
	assert.NoError(t, err)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Error removing book: Book not found")
}
