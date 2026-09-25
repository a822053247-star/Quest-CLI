package cmd

import (
	"quest/internal/storage"
	"quest/internal/ui"

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
		ui.PrintStats(p)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
