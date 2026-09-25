package cmd

import (
	"fmt"
	"quest/internal/storage"

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
		for _, quest := range quests {
			status := "[ ]"
			bossMark := ""
			if quest.Completed {
				status = "[√]"
			}
			if quest.Boss {
				bossMark = "👹"
			}
			fmt.Printf("%d %s %s +%d XP %s\n", quest.ID, status, quest.Title, quest.XP, bossMark)
		}
		if len(quests) == 0 {
			fmt.Println("No quests found.")
			return nil
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
