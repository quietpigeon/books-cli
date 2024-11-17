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

func TestListBooks(t *testing.T) {
	GetBooksRequest = func() (*http.Response, error) {
		mockBooks := []api.Book{
			{ID: 1, Title: "Book One", Author: "Author One"},
			{ID: 2, Title: "Book Two", Author: "Author Two"},
		}
		mockData, _ := json.Marshal(mockBooks)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(mockData)),
		}, nil
	}

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	defer func() {
		os.Stdout = oldStdout
	}()

	cmd := listCmd()
	listBooks(cmd, []string{})

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Book One")
	assert.Contains(t, output, "Author One")
	assert.Contains(t, output, "Book Two")
	assert.Contains(t, output, "Author Two")
}
