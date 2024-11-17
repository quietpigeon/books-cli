package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

func removeCmd() *cobra.Command {
	var removeCmd = &cobra.Command{
		Use:   "remove <book_id>",
		Short: "remove a book by ID",
		Args:  cobra.ExactArgs(1),
		Run:   removeBook,
	}
	return removeCmd
}

func removeBook(cmd *cobra.Command, args []string) {
	bookID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid book ID:", err)
		return
	}

	resp, err := RemoveBookRequest(bookID)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Book removed successfully!")
	} else {
		var apiError struct {
			Error string `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiError)
		fmt.Println("Error removing book:", apiError.Error)
	}
}
