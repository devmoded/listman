package cmd

import (
	"github.com/devmoded/listman/internal/db"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "",
}

var newEntryCmd = &cobra.Command{
	Use:   "entry [list-name] [entry-name] [data]",
	Short: "",
	Args:  cobra.ExactArgs(3),
	RunE:  newEntry,
}

var newListCmd = &cobra.Command{
	Use:   "list [name]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE:  newList,
}

func newEntry(cmd *cobra.Command, args []string) error {
	s, err := db.NewStore(dbPath)
	if err != nil {
		return err
	}
	defer s.Close()

	list_id, err := s.GetListID(args[0])
	if err != nil {
		return err
	}

	s.AddEntry(list_id, args[1], args[2])
	return nil
}

func newList(cmd *cobra.Command, args []string) error {
	s, err := db.NewStore(dbPath)
	if err != nil {
		return err
	}
	defer s.Close()
	s.AddList(args[0])
	return nil
}
