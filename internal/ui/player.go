package ui

import (
	"fmt"
	"quest/internal/player"
)

func PrintStats(p player.Player) {
	level := player.Level(p.XP)
	fmt.Printf("Lv: %d\nXP: %d\n🔥 Streak: %d\n🏆 Best Streak: %d\n", level, p.XP, p.CurrentStreak, p.BestStreak)
}
