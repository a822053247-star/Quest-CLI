package cmd

import (
	"fmt"
	"quest/internal/player"
	"quest/internal/storage"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show player stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := storage.LoadPlayer()
		if err != nil {
			return err
		}
		level := player.Level(p.XP)
		fmt.Printf("Lv: %d\nXP: %d\n🔥 Streak: %d\n🏆 Best Streak: %d\n", level, p.XP, p.CurrentStreak, p.BestStreak)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
