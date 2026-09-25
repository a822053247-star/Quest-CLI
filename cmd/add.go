package cmd

import (
	"quest/internal/quest"
	"quest/internal/storage"
	"quest/internal/ui"

	"github.com/spf13/cobra"
)

// 保存 --xp 和 --boss flag 的值
var addXP int
var addBoss bool

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new quest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// title 是args[0]
		title := args[0]
		quests, err := storage.LoadQuests()
		if err != nil {
			return err
		}
		q := quest.CreateQuest(quests, title, addXP, addBoss)
		quests = quest.AddQuest(quests, q)
		err = storage.SaveQuests(quests)
		if err != nil {
			return err
		}
		ui.PrintQuestAdded(q)
		return nil
	},
}

func init() {
	// 注册 --xp
	addCmd.Flags().IntVar(&addXP, "xp", 10, "XP for completing the quest")
	addCmd.Flags().BoolVar(&addBoss, "boss", false, "Is this a boss quest?")
	rootCmd.AddCommand(addCmd)
}
