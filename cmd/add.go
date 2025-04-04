package cmd

import (
	"books-cli/api"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	var add = &cobra.Command{
		Use:   "add",
		Short: "add a new book",
		Run: func(cmd *cobra.Command, args []string) {
			addBook(cmd, args, os.Stdin)
		},
	}
	add.Flags().StringP("title", "t", "", "Title of the book")
	add.Flags().StringP("author", "a", "", "Author of the book")
	add.Flags().StringP("published-date", "p", "", "Published date (YYYY-MM-DD)")
	add.Flags().StringP("edition", "e", "", "Edition of the book (optional)")
	add.Flags().StringSliceP("genre", "g", []string{}, "Genre(s) of the book")
	add.Flags().StringP("description", "d", "", "Description of the book (optional)")
	return add
}

func addBook(cmd *cobra.Command, _ []string, input io.Reader) {
	title, _ := cmd.Flags().GetString("title")
	author, _ := cmd.Flags().GetString("author")
	publishedDate, _ := cmd.Flags().GetString("published-date")
	edition, _ := cmd.Flags().GetInt("edition")
	genres, _ := cmd.Flags().GetStringSlice("genre")

	newBook := api.Book{
		Title:         title,
		Author:        author,
		PublishedDate: publishedDate,
		Edition:       edition,
		Genre:         genres,
	}

	resp, err := AddBookRequest(newBook)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Book added successfully!")

		// Ask for book description.
		reader := bufio.NewReader(input)
		fmt.Print("Would you like to add a short description? [y/n] ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// Get ID of the book added.
		var responseBody struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			log.Fatal("Error decoding response:", err)
		}

		if answer == "y" || answer == "Y" {
			fmt.Print("Enter description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			newBook.Description = description
			newBook.ID = responseBody.ID
			updateResp, err := UpdateBookRequest(newBook)
			if err != nil {
				log.Fatal(err)
			}
			if updateResp.StatusCode == http.StatusOK {
				fmt.Println("Description added.")
			} else {
				fmt.Println("Error adding description:", updateResp.Status)
			}
		}
	} else {
		var apiError struct {
			Error string `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiError)
		fmt.Println("Error adding book:", apiError.Error)
	}
}
