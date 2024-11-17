package api

import (
	"encoding/json"
	"errors"
	"os"
)

type Book struct {
	ID            int
	Title         string
	Author        string
	PublishedDate string
	Edition       int
	Genre         []string
	Description   string
}

var books []Book

func LoadBooks() {
	data, err := os.ReadFile("data/books.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			books = []Book{}
			return
		}
	}
	err = json.Unmarshal(data, &books)
	if err != nil {
		panic(err)
	}
}

func SaveBooks() {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("data/books.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
