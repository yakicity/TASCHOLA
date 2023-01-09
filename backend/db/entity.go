package db

import (
	"time"
)

type Task struct {
	ID          uint64    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`   // [TODO, DOING, DONE]
	Priority    int       `db:"priority"` // [1, 2, 3, 4, 5] (1: highest, 5: lowest)
	CreatedAt   time.Time `db:"created_at"`
	DueDate     time.Time `db:"due_date"`
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
	UserID    uint64 `db:"user_id"`
	StudentID string `db:"student_id"` // 学籍番号
	Password  string `db:"password"`   // パスワード(portal) 平文

	MatrixCodeAColumn string `db:"matrix_code_a_column"` // マトリックスコードA列
	MatrixCodeBColumn string `db:"matrix_code_b_column"` // マトリックスコードB列
	MatrixCodeCColumn string `db:"matrix_code_c_column"` // マトリックスコードC列
	MatrixCodeDColumn string `db:"matrix_code_d_column"` // マトリックスコードD列
	MatrixCodeEColumn string `db:"matrix_code_e_column"` // マトリックスコードE列
	MatrixCodeFColumn string `db:"matrix_code_f_column"` // マトリックスコードF列
	MatrixCodeGColumn string `db:"matrix_code_g_column"` // マトリックスコードG列
	MatrixCodeHColumn string `db:"matrix_code_h_column"` // マトリックスコードH列
	MatrixCodeIColumn string `db:"matrix_code_i_column"` // マトリックスコードI列
	MatrixCodeJColumn string `db:"matrix_code_j_column"` // マトリックスコードJ列
}
