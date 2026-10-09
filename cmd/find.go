package cmd

import (
	"fmt"

	"github.com/devmoded/listman/internal/db"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "",
}

var findListCmd = &cobra.Command{
	Use:   "list [name]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE:  findList,
}

func findList(cmd *cobra.Command, args []string) error {
	s, err := db.NewStore(dbPath)
	if err != nil {
		return err
	}
	defer s.Close()
	id, err := s.GetListID(args[0])
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
