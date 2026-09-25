package ui

import (
	"fmt"
	"quest/internal/player"
	"quest/internal/quest"
)

func PrintQuests(quests []quest.Quest, p player.Player) error {
	fmt.Println("⚔️ Today's Quests")
	fmt.Println()
	if len(quests) == 0 {
		fmt.Println("No quests found.")
		return nil
	}
	for _, q := range quests {
		status := "[ ]"
		bossMark := ""
		if q.Completed {
			status = "[√]"
		}
		if q.Boss {
			bossMark = "👹"
		}
		fmt.Printf("%d %s %s +%d XP %s\n", q.ID, status, q.Title, q.XP, bossMark)
	}
	level := player.Level(p.XP)
	nextLevelXP := level * 100
	currentLevelXP := p.XP % 100
	bar := ProgressBar(currentLevelXP)
	fmt.Printf("\nLv.%d %s %d / %d XP\n", level, bar, p.XP, nextLevelXP)
	fmt.Printf("🔥 Streak: %d days\n", p.CurrentStreak)
	return nil
}

func PrintQuestAdded(q quest.Quest) {
	fmt.Printf("✅ Quest added: %s (+%d XP)\n", q.Title, q.XP)
}

func PrintQuestNotFound(id int) {
	fmt.Printf("Quest with ID %d not found.\n", id)
}

func PrintQuestDeleted(id int) {
	fmt.Printf("Quest %d deleted.\n", id)
}

func PrintQuestAlreadyCompleted(id int) {
	fmt.Printf("Quest %d is already completed.\n", id)
}

func PrintQuestCompleted(q quest.Quest) {
	fmt.Printf("⚔️ Quest completed: %s (+%d XP)\n", q.Title, q.XP)
}

func ProgressBar(xp int) string {
	bar := ""
	filled := xp / 10
	for i := 0; i < 10; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}
