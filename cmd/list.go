package cmd

import (
	"books-cli/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list all books",
		Run:   listBooks,
	}
	return listCmd
}

func listBooks(cmd *cobra.Command, args []string) {
	resp, err := GetBooksRequest()
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		var books []api.Book
		json.NewDecoder(resp.Body).Decode(&books)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "Title", "Author", "Published Date", "Edition", "Genre(s)", "Description"})

		for _, book := range books {
			t.AppendRow(table.Row{
				book.ID,
				book.Title,
				book.Author,
				book.PublishedDate,
				book.Edition,
				strings.Join(book.Genre, ", "),
				book.Description,
			})
		}

		t.Render()
	} else {
		fmt.Println("Error getting books:", resp.Status)
	}
}
