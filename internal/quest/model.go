package quest

import "time"

// Quest 代表一个任务
// 一个任务包含以下字段：例如ID是任务编号
// json:“id” 是在Quest变成json后，给json字段定义tag标签是id，而不是ID
type Quest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	XP        int       `json:"xp"`
	Boss      bool      `json:"boss"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"create_at"`
}
