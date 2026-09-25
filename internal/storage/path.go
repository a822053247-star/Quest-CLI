package storage

// 确定数据文件保存到什么地方

import (
	"encoding/json"
	"os"
	"path/filepath"
	"quest/internal/player"
	"quest/internal/quest"
)

func DataDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".quest"), nil
}

func EnsureDataDir() (string, error) {
	dir, err := DataDir()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func QuestFilePath() (string, error) {
	dir, err := EnsureDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "quests.json"), nil
}

func LoadQuests() ([]quest.Quest, error) {
	path, err := QuestFilePath()
	if err != nil {
		return nil, err
	}
	var quests []quest.Quest
	data, err := os.ReadFile(path)
	if err != nil {
		// 文件不存在
		if os.IsNotExist(err) {
			return []quest.Quest{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &quests)
	if err != nil {
		return nil, err
	}
	return quests, nil
}

func SaveQuests(quests []quest.Quest) error {
	path, err := QuestFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(quests, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func PlayerFilePath() (string, error) {
	dir, err := EnsureDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "player.json"), nil
}

func LoadPlayer() (player.Player, error) {
	p := player.Player{}
	path, err := PlayerFilePath()
	if err != nil {
		return p, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return p, nil
		}
		return p, err
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SavePlayer(p player.Player) error {
	path, err := PlayerFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
