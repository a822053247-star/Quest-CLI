# Quest CLI

一个把 Todo 变成 RPG 任务的命令行工具。

你可以创建任务、完成任务获得 XP、升级等级、积累连续完成天数，并把重要任务标记为 Boss。

## Features

- 创建 Quest
- 查看 Quest 列表
- 完成 Quest 并获得 XP
- 删除 Quest
- Player 等级系统
- Streak 连续完成系统
- Boss Quest
- 本地 JSON 持久化
- CLI RPG 风格输出

## Example

```text
⚔️ Today's Quests

1  [ ] Read paper       +30 XP
2  [√] Run 5km          +50 XP
3  [ ] Finish project   +100 XP 👹

Lv.3  ██████░░░░ 260 / 300 XP
🔥 Streak: 5 days
```

## Installation

确保已经安装 Go。

```bash
go version
```

克隆项目后进入项目目录：

```bash
cd quest
```

下载依赖：

```bash
go mod tidy
```

直接运行：

```bash
go run . --help
```

也可以编译：

```bash
go build -o quest .
```

Windows：

```powershell
go build -o quest.exe .
```

## Commands

### Add Quest

创建一个新任务：

```bash
quest add "Read paper"
```

指定 XP：

```bash
quest add "Read paper" --xp 30
```

默认 XP 为：

```text
10 XP
```

创建 Boss Quest：

```bash
quest add "Finish project" --xp 100 --boss
```

### List Quests

查看任务列表：

```bash
quest list
```

示例：

```text
⚔️ Today's Quests

1 [ ] Read paper +30 XP
2 [ ] Run 5km +50 XP
3 [ ] Finish project +100 XP 👹
```

其中：

```text
[ ]   未完成

[√]   已完成

👹    Boss Quest
```

### Complete Quest

完成指定任务：

```bash
quest done 1
```

完成任务后会：

```text
Quest Completed
      ↓
获得 XP
      ↓
更新 Level
      ↓
更新 Streak
      ↓
保存数据
```

示例：

```text
⚔️ Quest completed: Read paper (+30 XP)
```

已经完成的任务不能重复完成，因此不会重复获得 XP。

### Delete Quest

删除指定任务：

```bash
quest delete 1
```

如果任务存在：

```text
🗑️ Quest 1 deleted.
```

如果任务不存在：

```text
Quest with id 1 not found.
```

### Player Stats

查看玩家状态：

```bash
quest stats
```

示例：

```text
Lv.3
XP: 260
🔥 Streak: 5
🏆 Best Streak: 8
```

## XP & Level

玩家通过完成 Quest 获得 XP。

等级计算公式：

```go
level := xp/100 + 1
```

例如：

```text
0 XP    → Lv.1
99 XP   → Lv.1
100 XP  → Lv.2
260 XP  → Lv.3
300 XP  → Lv.4
```

每获得 100 XP 提升一级。

## Streak

每天至少完成一个 Quest，就可以维持连续完成记录。

规则：

```text
第一次完成任务
→ CurrentStreak = 1

同一天再次完成任务
→ CurrentStreak 不变

连续第二天完成任务
→ CurrentStreak + 1

中间断了一天或以上
→ CurrentStreak = 1
```

同时会记录：

```text
BestStreak
```

表示历史最高连续完成天数。

## Boss Quest

使用：

```bash
quest add "Finish project" --xp 100 --boss
```

创建 Boss Quest。

Boss Quest 会在列表中显示：

```text
👹
```

例如：

```text
3 [ ] Finish project +100 XP 👹
```

Boss 当前只是任务标记，可以在后续版本中扩展更多 RPG 机制。

## Data Storage

Quest CLI 不使用数据库。

所有数据保存在本地：

```text
~/.quest/
├── quests.json
└── player.json
```

`quests.json` 保存任务：

```json
[
  {
    "id": 1,
    "title": "Read paper",
    "xp": 30,
    "boss": false,
    "completed": false,
    "created_at": "2026-09-25T10:00:00Z"
  }
]
```

`player.json` 保存玩家状态：

```json
{
  "xp": 260,
  "current_streak": 5,
  "best_streak": 8,
  "last_active_at": "2026-09-25T10:00:00Z"
}
```

因此程序关闭后，Quest、XP 和 Streak 都不会丢失。

## Project Structure

```text
quest/
├── cmd/
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── done.go
│   ├── delete.go
│   └── stats.go
│
├── internal/
│   ├── quest/
│   │   ├── model.go
│   │   └── service.go
│   │
│   ├── player/
│   │   ├── model.go
│   │   └── service.go
│   │
│   ├── storage/
│   └── ui/
│
├── main.go
├── go.mod
└── README.md
```

各层职责：

```text
cmd
→ Cobra 命令、参数解析和业务流程编排

quest
→ Quest 数据模型和任务业务逻辑

player
→ XP、Level、Streak

storage
→ JSON 文件读取和保存

ui
→ CLI 输出和 RPG 风格展示
```

## Architecture

例如：

```bash
quest add "Read paper" --xp 30
```

内部流程：

```text
Cobra
  ↓
cmd/add.go
  ↓
LoadQuests
  ↓
CreateQuest
  ↓
AddQuest
  ↓
SaveQuests
  ↓
CLI UI
```

而：

```bash
quest done 1
```

内部流程：

```text
LoadQuests
    ↓
FindQuestByID
    ↓
CompleteQuest
    ↓
LoadPlayer
    ↓
AddXP
    ↓
UpdateStreak
    ↓
SaveQuests
    ↓
SavePlayer
```

## Development

格式化整个项目：

```bash
gofmt -w .
```

运行所有测试：

```bash
go test ./...
```

查看详细测试结果：

```bash
go test -v ./...
```

运行 CLI：

```bash
go run . --help
```

## MVP

`v0.1.0` 的目标是保证以下命令正常工作：

```bash
quest add "Read paper"
quest list
quest done 1
quest delete 1
quest stats
```

并满足：

- Quest 可以持久化
- Player 可以持久化
- XP 计算正确
- Level 计算正确
- Streak 计算正确
- 已完成任务不能重复获得 XP
- Boss Quest 可以正常显示
- `go test ./...` 通过

## Future Ideas

后续版本可以考虑：

```text
Daily Quest
Quest Difficulty
Achievements
Boss HP
Gold / Coins
Equipment
Quest Categories
Deadlines
Random Rewards
Colored CLI UI
Configuration
Cloud Sync
```

## Tech Stack

- Go
- Cobra
- JSON
- Go standard library

## Version

Current MVP:

```text
v0.1.0
```