package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Cli() {
	var cli = &cobra.Command{
		Use:   "book",
		Short: "cli for bool management",
	}
	cli.AddCommand(addCmd(), listCmd(), removeCmd(), updateCmd())

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
