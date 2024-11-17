package cmd

import (
	"books-cli/api"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var AddBookRequest = func(newBook api.Book) (*http.Response, error) {
	data, err := json.Marshal(newBook)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/books", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

var UpdateBookRequest = func(book api.Book) (*http.Response, error) {
	jsonData, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://localhost:8080/v1/books/%d", book.ID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

var GetBooksRequest = func() (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/v1/books", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

func RemoveBookRequest(bookID int) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/books/%d", bookID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
