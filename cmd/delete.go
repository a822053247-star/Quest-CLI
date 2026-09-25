package cmd

import (
	"fmt"
	"quest/internal/quest"
	"quest/internal/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a quest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		quests, err := storage.LoadQuests()
		if err != nil {
			return err
		}
		quests, isFound := quest.DeleteQuest(quests, id)
		if !isFound {
			fmt.Printf("Quest with id %d not found.\n", id)
			return nil
		}
		err = storage.SaveQuests(quests)
		if err != nil {
			return err
		}
		fmt.Printf("Quest %d deleted.\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
