package cmd

import (
	"quest/internal/storage"
	"quest/internal/ui"

	"github.com/spf13/cobra"
)

// 打印任务
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all quests",
	RunE: func(cmd *cobra.Command, args []string) error {
		quests, err := storage.LoadQuests()
		if err != nil {
			return err
		}
		ui.PrintQuests(quests)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
