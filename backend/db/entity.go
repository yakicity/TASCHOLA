package db

import (
	"time"
)

type Task struct {
	ID          uint64    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	IsDone      bool      `db:"is_done"`
	Status      string    `db:"status"`
	DueDate     time.Time `db:"due_date"`
	Priority    int       `db:"priority"`
}

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Password []byte `db:"password"`
}

type OwnerShip struct {
	ID     uint64 `db:"id"`
	UserID uint64 `db:"user_id"`
	TaskID uint64 `db:"task_id"`
}

type Student struct {
	ID        uint64 `db:"id"`
	UserID		uint64 `db:"user_id"`
	StudentID string `db:"student_id"` // 学籍番号
	Password	[]byte `db:"password"` // パスワード(portal)
	MatrixCode string `db:"matrix_code"` // マトリックスコード
}
