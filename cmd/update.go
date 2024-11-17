package cmd

import (
	"books-cli/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

func updateCmd() *cobra.Command {
	var updateCmd = &cobra.Command{
		Use:   "update <book_id>",
		Short: "update a book by ID",
		Args:  cobra.ExactArgs(1),
		Run:   updateBook,
	}
	updateCmd.Flags().StringP("title", "t", "", "New title of the book (optional)")
	updateCmd.Flags().StringP("author", "a", "", "New author of the book (optional)")
	updateCmd.Flags().StringP("published-date", "p", "", "New published date (YYYY-MM-DD) (optional)")
	updateCmd.Flags().StringP("edition", "e", "", "New edition of the book (optional)")
	updateCmd.Flags().StringSliceP("genre", "g", []string{}, "New genre(s) of the book (optional)")
	updateCmd.Flags().StringP("description", "d", "", "New description of the book (optional)")
	return updateCmd
}

func updateBook(cmd *cobra.Command, args []string) {
	bookID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid book ID:", err)
		return
	}

	title, _ := cmd.Flags().GetString("title")
	author, _ := cmd.Flags().GetString("author")
	publishedDate, _ := cmd.Flags().GetString("published-date")
	edition, _ := cmd.Flags().GetInt("edition")
	genres, _ := cmd.Flags().GetStringSlice("genre")
	description, _ := cmd.Flags().GetString("description")

	updatedBook := api.Book{
		ID:            bookID,
		Title:         title,
		Author:        author,
		PublishedDate: publishedDate,
		Edition:       edition,
		Genre:         genres,
		Description:   description,
	}

	resp, err := UpdateBookRequest(updatedBook)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Book updated successfully!")
	} else {
		var apiError struct {
			Error string `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiError)
		fmt.Println("Error updating book:", apiError.Error)
	}
}
