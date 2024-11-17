package cmd

import (
	"books-cli/api"
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	AddBookRequest = func(newBook api.Book) (*http.Response, error) {
		assert.Equal(t, "Test Title", newBook.Title)
		assert.Equal(t, "Test Author", newBook.Author)

		mockResp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"id": 123}`)),
		}
		return mockResp, nil
	}

	input := bytes.NewBufferString("y\nA great adventure book.\n")

	cmd := &cobra.Command{}
	cmd.Flags().StringP("title", "t", "Test Title", "")
	cmd.Flags().StringP("author", "a", "Test Author", "")
	cmd.Flags().StringSliceP("genre", "g", []string{"Fiction"}, "")

	addBook(cmd, []string{}, input)
}
