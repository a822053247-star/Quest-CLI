package quest

import "time"

func NextID(quests []Quest) int {
	maxID := 0
	for _, quest := range quests {
		if quest.ID > maxID {
			maxID = quest.ID
		}
	}
	return maxID + 1
}

func CreateQuest(quests []Quest, title string, xp int, boss bool) Quest {
	q := Quest{
		ID:        NextID(quests),
		Title:     title,
		XP:        xp,
		Boss:      boss,
		Completed: false,
		CreatedAt: time.Now(),
	}
	return q
}

func AddQuest(quests []Quest, q Quest) []Quest {
	// 把q加到quests的尾部
	quests = append(quests, q)
	return quests
}

// 第二个bool代表是否找到
func DeleteQuest(quests []Quest, id int) ([]Quest, bool) {
	result := []Quest{}
	deleted := false
	for _, quest := range quests {
		if quest.ID != id {
			result = append(result, quest)
		} else {
			deleted = true
		}
	}
	return result, deleted
}

// 标记任务已经完成,第一个bool代表found，第二个bool代表completed
func CompleteQuest(quests []Quest, id int) ([]Quest, bool, bool) {
	found := false
	completed := false
	// 不要用下面这个，因为quest只是修改了副本，最后返回的quests并没有修改completed
	//for _, quest := range quests {
	//	if quest.ID == id {
	//		quest.Completed = true
	//		completed = true
	//	}
	//}
	for i := range quests {
		if quests[i].ID == id {
			found = true
			if !quests[i].Completed {
				quests[i].Completed = true
				completed = true
			}
			break
		}
	}
	return quests, found, completed
}

func FindQuestByID(quests []Quest, id int) (Quest, bool) {
	for _, quest := range quests {
		if quest.ID == id {
			return quest, true
		}
	}
	return Quest{}, false
}
