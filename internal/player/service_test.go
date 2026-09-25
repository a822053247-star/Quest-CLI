package player

import (
	"testing"
	"time"
)

func TestUpdateStreakFirstTime(t *testing.T) {
	player := Player{}
	now := time.Date(2026, 9, 24, 10, 0, 0, 0, time.UTC)
	UpdateStreak(&player, now)
	if player.CurrentStreak != 1 {
		t.Errorf("Expected streak 1,got %d", player.CurrentStreak)
	}
	if player.BestStreak != 1 {
		t.Errorf("Expected best streak 1,got %d", player.BestStreak)
	}
}

func TestUpdateStreakSameDay(t *testing.T) {
	player := Player{}
	morning := time.Date(2026, 9, 24, 10, 0, 0, 0, time.UTC)
	evening := time.Date(2026, 9, 24, 18, 0, 0, 0, time.UTC)
	UpdateStreak(&player, morning)
	UpdateStreak(&player, evening)
	if player.CurrentStreak != 1 {
		t.Errorf("Expected streak 1,got %d", player.CurrentStreak)
	}
	if player.BestStreak != 1 {
		t.Errorf("Expected best streak 1,got %d", player.BestStreak)
	}
}

func TestUpdateStreakNextDay(t *testing.T) {
	player := Player{}
	yesterday := time.Date(2026, 9, 23, 10, 0, 0, 0, time.UTC)
	now := time.Date(2026, 9, 24, 10, 0, 0, 0, time.UTC)
	UpdateStreak(&player, yesterday)
	UpdateStreak(&player, now)
	if player.CurrentStreak != 2 {
		t.Errorf("Expected streak 2,got %d", player.CurrentStreak)
	}
	if player.BestStreak != 2 {
		t.Errorf("Expected best streak 2,got %d", player.BestStreak)
	}
}

func TestUpdateStreakResetAfterGap(t *testing.T) {
	player := Player{}
	time1 := time.Date(2026, 9, 20, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2026, 9, 21, 10, 0, 0, 0, time.UTC)
	time3 := time.Date(2026, 9, 24, 10, 0, 0, 0, time.UTC)
	UpdateStreak(&player, time1)
	UpdateStreak(&player, time2)
	UpdateStreak(&player, time3)
	if player.CurrentStreak != 1 {
		t.Errorf("Expected streak 1,got %d", player.CurrentStreak)
	}
	if player.BestStreak != 2 {
		t.Errorf("Expected best streak 2,got %d", player.BestStreak)
	}
}
