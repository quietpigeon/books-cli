package cmd

import (
	"books-cli/api"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddBookRequest(newBook api.Book) (*http.Response, error) {
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

func UpdateBookRequest(book api.Book) (*http.Response, error) {
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

func GetBooksRequest() (*http.Response, error) {
	// Create the HTTP request
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/books", nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
