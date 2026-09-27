package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"quest/internal/player"
	"quest/internal/quest"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// OpenMySQL 从环境变量 QUEST_MYSQL_DSN 读取 DSN 并建立 MySQL 连接。
// 连接成功后返回可用的 *sql.DB，失败时返回对应错误。
func OpenMySQL() (*sql.DB, error) {
	// 读取 DSN 配置，未配置则直接报错
	dsn := os.Getenv("QUEST_MYSQL_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("QUEST_MYSQL_DSN is not set")
	}

	// sql.Open 仅校验参数格式，不会真正建立网络连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping 触发真正的 TCP 握手与认证，确保数据库可达
	if err := db.Ping(); err != nil {
		db.Close() // 连接失败时释放资源，避免泄漏
		return nil, err
	}

	return db, nil
}

func CreateQuestMySQL(db *sql.DB, q quest.Quest) (int64, error) {
	result, err := db.Exec("INSERT INTO quests (title, xp, boss, completed, created_at) VALUES (?, ?, ?, ?, ?)", q.Title, q.XP, q.Boss, q.Completed, q.CreatedAt)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func ListQuestsMySQL(db *sql.DB) ([]quest.Quest, error) {
	rows, err := db.Query("SELECT id, title, xp, boss, completed, created_at FROM quests ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	quests := []quest.Quest{}
	for rows.Next() {
		q := quest.Quest{}
		err := rows.Scan(&q.ID, &q.Title, &q.XP, &q.Boss, &q.Completed, &q.CreatedAt)
		if err != nil {
			return nil, err
		}
		quests = append(quests, q)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return quests, nil
}

func DeleteQuestMySQL(db *sql.DB, id int) (bool, error) {
	result, err := db.Exec("DELETE FROM quests WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	// result.RowsAffected() 会告诉影响了多少行
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func FindQuestByIDMySQL(db *sql.DB, id int) (quest.Quest, bool, error) {
	row := db.QueryRow("SELECT id, title, xp, boss, completed, created_at FROM quests WHERE id = ?", id)
	q := quest.Quest{}
	err := row.Scan(&q.ID, &q.Title, &q.XP, &q.Boss, &q.Completed, &q.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return q, false, nil
		}
		return q, false, err
	}
	return q, true, nil
}

func CompleteQuestMySQL(db *sql.DB, id int) (bool, error) {
	result, err := db.Exec("UPDATE quests SET completed = true WHERE id = ? AND completed = false", true, id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// LoadPlayerMySQL 从 players 表加载 id=1 的玩家数据。
// last_active_at 字段允许为 NULL，通过 sql.NullTime 安全扫描。
func LoadPlayerMySQL(db *sql.DB) (player.Player, error) {
	row := db.QueryRow("SELECT xp, best_streak, current_streak, last_active_at FROM players WHERE id = 1")
	p := player.Player{}

	// 使用 sql.NullTime 处理 last_active_at 可能为 NULL 的情况
	var lastActive sql.NullTime
	err := row.Scan(&p.XP, &p.BestStreak, &p.CurrentStreak, &lastActive)
	if err != nil {
		return p, err
	}

	// 仅当 last_active_at 非 NULL 时才赋值
	if lastActive.Valid {
		p.LastActiveAt = lastActive.Time
	}

	return p, nil
}

func SavePlayerMySQL(db *sql.DB, p player.Player) error {
	//lastactive可以装两种不同类型变量
	//如果直接用 := 声明，则 lastactive 的类型会根据赋值自动推断为time
	var lastactive any
	if p.LastActiveAt.IsZero() {
		lastactive = nil
	} else {
		lastactive = p.LastActiveAt
	}
	_, err := db.Exec("UPDATE players SET xp = ?, best_streak = ?, current_streak = ?, last_active_at = ? WHERE id = 1", p.XP, p.BestStreak, p.CurrentStreak, lastactive)
	if err != nil {
		return err
	}
	return nil
}

func EnsurePlayerMySQL(db *sql.DB) error {
	// ON DUPLICATE KEY UPDATE id = id;
	// 表示当出现主键重复时，执行 UPDATE id = id;
	_, err := db.Exec("INSERT INTO players (id, xp, best_streak, current_streak, last_active_at) VALUES (1, 0, 0, 0, NULL) ON DUPLICATE KEY UPDATE id = id")
	if err != nil {
		return err
	}
	return nil
}

// 事务处理
func CompleteQuestWithReward(db *sql.DB, id int) (quest.Quest, error) {
	// tx是一次连接conn，conn执行BEGIN
	// *sql.DB是连接池，最后执行完会自动归还连接池
	// *sql.Tx是单个事务会话，代表这条连接被tx占用，不会归还连接池
	tx, err := db.Begin()
	if err != nil {
		return quest.Quest{}, err
	}
	// 事务回滚.条件：如果return err就触发。如果成功tx.Commit()，回滚也不会把已经提交的数据撤销
	// `sql.Tx` 对象内部有一个事务状态标记，一旦 `Commit()` 成功，状态被置为 `done`。后续任何 `Rollback/Query/Exec` 都会先检查这个状态，直接在 Go 代码层面返回 `sql.ErrTxDone`，根本不会往数据库发送 ROLLBACK 语句
	defer tx.Rollback()
	var q quest.Quest

	err = tx.QueryRow(`SELECT id, title, xp, boss, completed, created_at
	FROM quests WHERE id = ?`, id).Scan(
		&q.ID,
		&q.Title,
		&q.XP,
		&q.Boss,
		&q.Completed,
		&q.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return quest.Quest{}, fmt.Errorf("quest %d not found", id)
	}
	if err != nil {
		return quest.Quest{}, err
	}
	if q.Completed {
		return quest.Quest{}, fmt.Errorf("quest %d already completed", id)
	}
	result, err := tx.Exec(
		"UPDATE quests SET completed = true WHERE id = ? AND completed = false",
		id,
	)
	if err != nil {
		return quest.Quest{}, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return quest.Quest{}, err
	}
	if affected == 0 {
		return quest.Quest{}, fmt.Errorf("quest %d not found", id)
	}

	// 同一个事务里查玩家
	var p player.Player
	var lastActive sql.NullTime
	err = tx.QueryRow("SELECT xp, best_streak, current_streak, last_active_at FROM players WHERE id = 1").Scan(&p.XP, &p.BestStreak, &p.CurrentStreak, &lastActive)
	if err != nil {
		return quest.Quest{}, err
	}
	if lastActive.Valid {
		p.LastActiveAt = lastActive.Time
	}
	player.AddXP(&p, q.XP)
	player.UpdateStreak(&p, time.Now())
	_, err = tx.Exec(
		"UPDATE players SET xp = ?, best_streak = ?, current_streak = ?, last_active_at = ? WHERE id = 1",
		p.XP,
		p.BestStreak,
		p.CurrentStreak,
		p.LastActiveAt,
	)
	if err != nil {
		return quest.Quest{}, err
	}
	if err := tx.Commit(); err != nil {
		return quest.Quest{}, err
	}
	return q, nil
}
