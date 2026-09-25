package player

import "time"

// Player 表示玩家的等级与连续完成任务状态。
// Player XP              总经验
// CurrentStreak   当前连续天数
// BestStreak      历史最高连续天数
// LastActiveAt    上一次完成任务的时间

type Player struct {
	XP            int       `json:"xp"`
	CurrentStreak int       `json:"current_streak"`
	BestStreak    int       `json:"best_streak"`
	LastActiveAt  time.Time `json:"last_active_at"`
}
