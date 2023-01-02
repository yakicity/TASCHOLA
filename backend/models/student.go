package models

import (
	database "taschola/db"
)

func GetStudentByUserID(userID uint64) (database.Student, error) {
	db, err := database.GetConnection()
	if err != nil {
		return database.Student{}, err
	}

	var student database.Student
	err = db.Get(&student, "SELECT * FROM students WHERE user_id = ?", userID)
	if err != nil {
		return database.Student{}, err
	}

	return student, nil
}

func GetStudentByStudentID(studentID uint64) (database.Student, error) {
	db, err := database.GetConnection()
	if err != nil {
		return database.Student{}, err
	}

	var student database.Student
	err = db.Get(&student, "SELECT * FROM students WHERE id = ?", studentID)
	if err != nil {
		return database.Student{}, err
	}

	return student, nil
}

func CreateStudent(student database.Student) (uint64, error) {
	db, err := database.GetConnection()
	if err != nil {
		return 0, err
	}

	tx, err := db.Beginx()
	if err != nil {
		return 0, err
	}

	var studentID uint64
	err = tx.QueryRowx("INSERT INTO students (user_id, student_id, password, matrix_code_a_column, matrix_code_b_column, matrix_code_c_column, matrix_code_d_column, matrix_code_e_column, matrix_code_f_column, matrix_code_g_column, matrix_code_h_column, matrix_code_i_column, matrix_code_j_column) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", student.UserID, student.StudentID, student.Password, student.MatrixCodeAColumn, student.MatrixCodeBColumn, student.MatrixCodeCColumn, student.MatrixCodeDColumn, student.MatrixCodeEColumn, student.MatrixCodeFColumn, student.MatrixCodeGColumn, student.MatrixCodeHColumn, student.MatrixCodeIColumn, student.MatrixCodeJColumn).Scan(&studentID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return studentID, nil
}

func UpdateStudent(student database.Student) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	// userIDが一致するユーザーのnameとpasswordを更新
	_, err = db.Exec("UPDATE students SET user_id = ?, student_id = ?, password = ?, matrix_code_a_column = ?, matrix_code_b_column = ?, matrix_code_c_column = ?, matrix_code_d_column = ?, matrix_code_e_column = ?, matrix_code_f_column = ?, matrix_code_g_column = ?, matrix_code_h_column = ?, matrix_code_i_column = ?, matrix_code_j_column = ? WHERE id = ?", student.UserID, student.StudentID, student.Password, student.MatrixCodeAColumn, student.MatrixCodeBColumn, student.MatrixCodeCColumn, student.MatrixCodeDColumn, student.MatrixCodeEColumn, student.MatrixCodeFColumn, student.MatrixCodeGColumn, student.MatrixCodeHColumn, student.MatrixCodeIColumn, student.MatrixCodeJColumn, student.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStudent(studentID uint64) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM students WHERE id = ?", studentID)
	if err != nil {
		return err
	}

	return nil
}
