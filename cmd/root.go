package cmd

import (
	"github.com/spf13/cobra"
)

var dbPath string

var rootCmd = &cobra.Command{
	Use:           "listman [command]",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbPath, "database", "lists.db", "define database root")

	findCmd.AddCommand(findListCmd)

	newCmd.AddCommand(newEntryCmd)
	newCmd.AddCommand(newListCmd)

	printCmd.AddCommand(printEntriesCmd)
	printCmd.AddCommand(printListsCmd)

	rootCmd.AddCommand(findCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(printCmd)
}

func Run() error {
	err := rootCmd.Execute()
	return err
}
