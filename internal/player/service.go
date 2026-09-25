package player

import "time"

// 在 Go 里，小写开头的函数只能在当前 package 内使用；大写开头才可以被其他 package 调用。
func Level(xp int) int {
	return xp/100 + 1
}

// 给玩家增加经验
// 用指针可以直接修改传入的 Player 对象
// 如果是p Player，那么修改XP只是修改副本，没办法改原对象
func AddXP(p *Player, xp int) {
	p.XP += xp
}

// 判断两个时间是不是同日
func SameDay(a, b time.Time) bool {
	same := a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
	return same
}

func IsYesterday(a, b time.Time) bool {
	// 不能用Day()-1，因为会遇到月底、年底
	lastday := b.AddDate(0, 0, -1)
	return SameDay(a, lastday)
}

// 更新持续时间的状态
func UpdateStreak(p *Player, now time.Time) {
	// 判断是否为零时间（没有设置过时间）
	if p.LastActiveAt.IsZero() {
		p.CurrentStreak = 1
	} else if SameDay(p.LastActiveAt, now) {
	} else if IsYesterday(p.LastActiveAt, now) {
		p.CurrentStreak++
	} else {
		p.CurrentStreak = 1
	}
	if p.CurrentStreak > p.BestStreak {
		p.BestStreak = p.CurrentStreak
	}
	p.LastActiveAt = now
}
