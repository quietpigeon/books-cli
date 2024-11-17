package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Cli() {
	var cli = &cobra.Command{
		Use:   "books-cli",
		Short: "CLI for Book Management",
	}
	cli.AddCommand(addCmd())

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
