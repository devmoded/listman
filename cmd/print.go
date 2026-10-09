package cmd

import (
	"fmt"

	"github.com/devmoded/listman/internal/db"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "",
}

var printEntriesCmd = &cobra.Command{
	Use:   "entries [list-name]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE:  printEntries,
}

var printListsCmd = &cobra.Command{
	Use:   "lists [list-name]",
	Short: "",
	Args:  cobra.NoArgs,
	RunE:  printLists,
}

func printEntries(cmd *cobra.Command, args []string) error {
	s, err := db.NewStore(dbPath)
	if err != nil {
		return err
	}
	defer s.Close()

	id, err := s.GetListID(args[0])
	if err != nil {
		return err
	}

	c, err := s.GetListEntries(id)
	if err != nil {
		return err
	}

	for _, line := range c {
		fmt.Println(line)
	}
	return nil
}

func printLists(cmd *cobra.Command, args []string) error {
	s, err := db.NewStore(dbPath)
	if err != nil {
		return err
	}
	defer s.Close()

	c, err := s.GetLists()
	if err != nil {
		return err
	}

	for _, line := range c {
		fmt.Println(line)
	}
	return nil
}
