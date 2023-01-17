package db

import (
	"time"
)

type Task struct {
	ID          uint64    `db:"id" example:"1" json:"id"`
	Title       string    `db:"title" example:"title" json:"title"`
	Description string    `db:"description" example:"description" json:"description"`
	Status      string    `db:"status" example:"TODO" json:"status"`  // [TODO, DOING, DONE]
	Priority    int       `db:"priority" example:"2" json:"priority"` // [1, 2, 3, 4, 5] (1: highest, 5: lowest)
	CreatedAt   time.Time `db:"created_at" example:"2020-01-01 00:00:00" json:"created_at"`
	DueDate     time.Time `db:"due_date" example:"2020-01-01 00:00:00" json:"due_date"`
}

type User struct {
	ID       uint64 `db:"id" example:"1" json:"id"`
	Name     string `db:"name" example:"user-name" json:"name"`
	Password []byte `db:"password" json:"password"`
}

type OwnerShip struct {
	ID     uint64 `db:"id"`
	UserID uint64 `db:"user_id"`
	TaskID uint64 `db:"task_id"`
}

type Student struct {
	ID        uint64 `db:"id" example:"1" json:"id"`
	UserID    uint64 `db:"user_id" example:"1" json:"user_id"`
	StudentID string `db:"student_id" example:"20B30790" json:"student_id"` // 学籍番号
	Password  string `db:"password" example:"password" json:"password"`     // パスワード(portal) 平文

	MatrixCodeAColumn string `db:"matrix_code_a_column" example:"ABCDEFG" json:"matrix_code_a_column"` // マトリックスコードA列
	MatrixCodeBColumn string `db:"matrix_code_b_column" example:"ABCDEFG" json:"matrix_code_b_column"` // マトリックスコードB列
	MatrixCodeCColumn string `db:"matrix_code_c_column" example:"ABCDEFG" json:"matrix_code_c_column"` // マトリックスコードC列
	MatrixCodeDColumn string `db:"matrix_code_d_column" example:"ABCDEFG" json:"matrix_code_d_column"` // マトリックスコードD列
	MatrixCodeEColumn string `db:"matrix_code_e_column" example:"ABCDEFG" json:"matrix_code_e_column"` // マトリックスコードE列
	MatrixCodeFColumn string `db:"matrix_code_f_column" example:"ABCDEFG" json:"matrix_code_f_column"` // マトリックスコードF列
	MatrixCodeGColumn string `db:"matrix_code_g_column" example:"ABCDEFG" json:"matrix_code_g_column"` // マトリックスコードG列
	MatrixCodeHColumn string `db:"matrix_code_h_column" example:"ABCDEFG" json:"matrix_code_h_column"` // マトリックスコードH列
	MatrixCodeIColumn string `db:"matrix_code_i_column" example:"ABCDEFG" json:"matrix_code_i_column"` // マトリックスコードI列
	MatrixCodeJColumn string `db:"matrix_code_j_column" example:"ABCDEFG" json:"matrix_code_j_column"` // マトリックスコードJ列
}
