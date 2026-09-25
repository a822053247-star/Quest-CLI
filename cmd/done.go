package cmd

import (
	"fmt"
	"quest/internal/player"
	"quest/internal/quest"
	"quest/internal/storage"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Complete a quest",
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
		q, found := quest.FindQuestByID(quests, id)
		if !found {
			fmt.Println("Quest not found")
			return nil
		}
		quests, _, completed := quest.CompleteQuest(quests, id)
		//if !found {
		//	fmt.Println("Quest not found")
		//	return nil
		//}
		if !completed {
			fmt.Println("Quest already completed")
			return nil
		}
		p, err := storage.LoadPlayer()
		if err != nil {
			return err
		}
		player.AddXP(&p, q.XP)
		player.UpdateStreak(&p, time.Now())
		err = storage.SaveQuests(quests)
		if err != nil {
			return err
		}
		err = storage.SavePlayer(p)
		if err != nil {
			return err
		}
		err = storage.SaveQuests(quests)
		if err != nil {
			return err
		}
		fmt.Println("Quest completed! +%d XP\n", q.XP)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
