package quest

import "testing"

func TestNextID(t *testing.T) {
	// 空任务列表
	//→ NextID = 1
	quests := []Quest{}
	nextID := NextID(quests)
	if nextID != 1 {
		t.Errorf("Expected next ID to be 1, but got %d", nextID)
	}
}

func TestNextIDWithQuests(t *testing.T) {
	quests := []Quest{
		{ID: 1},
		{ID: 3},
		{ID: 5},
	}
	nextID := NextID(quests)
	if nextID != 6 {
		t.Errorf("Expected next ID to be 6, but got %d", nextID)
	}
}

func TestDeleteQuest(t *testing.T) {
	quests := []Quest{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	deletedQuests, success := DeleteQuest(quests, 2)
	if !success {
		t.Errorf("Expected deletion to be successful")
	}
	if len(deletedQuests) != 2 {
		t.Errorf("Expected 2 quests to remain, but got %d", len(deletedQuests))
	}
	for _, q := range deletedQuests {
		if q.ID == 2 {
			t.Errorf("quest 2 should have been deleted")
		}
	}
}

func TestCompleteQuest(t *testing.T) {
	quests := []Quest{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	quests, found, completed := CompleteQuest(quests, 2)
	if !found {
		t.Errorf("Expected completion to be successful")
	}
	if !completed {
		t.Errorf("Expected completion to be successful")
	}
}

func TestCompleteQuestAlreadyCompleted(t *testing.T) {
	quests := []Quest{
		{ID: 1, Completed: true},
	}

	quests, found, completed := CompleteQuest(quests, 1)

	if !found {
		t.Errorf("expected quest to be found")
	}

	if completed {
		t.Errorf("expected already completed quest not to complete again")
	}

	if !quests[0].Completed {
		t.Errorf("expected quest to remain completed")
	}
}
