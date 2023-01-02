package models

import (
	database "taschola/db"
)

func GetTasksByUserID(userID int) ([]database.Task, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var tasks []database.Task
	err = db.Select(&tasks, "SELECT * FROM tasks INNER JOIN ownerships ON task.id = ownerships.task_id WHERE ownerships.user_id = ?ORDER BY tasks.due_date ASC, tasks.priority ASC", userID) // 締切が近く、優先度の高い順に並べる
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksByUserIDAndKeywordAndStatus(userID int, keyword string, status []string) ([]database.Task, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var tasks []database.Task
	err = db.Select(&tasks, "SELECT * FROM tasks INNER JOIN ownerships ON task.id = ownerships.task_id WHERE ownerships.user_id = ? AND (tasks.title LIKE ? OR tasks.description LIKE ?) AND tasks.status IN (?) ORDER BY tasks.due_date ASC, tasks.priority ASC", userID, "%"+keyword+"%", "%"+keyword+"%", status) // 締切が近く、優先度の高い順に並べる
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByTaskID(taskID int) (database.Task, error) {
	db, err := database.GetConnection()
	if err != nil {
		return database.Task{}, err
	}

	var task database.Task
	err = db.Get(&task, "SELECT * FROM tasks WHERE id = ?", taskID)
	if err != nil {
		return database.Task{}, err
	}

	return task, nil
}

func CreateTask(task database.Task, userID uint64) (int, error) {
	db, err := database.GetConnection()
	if err != nil {
		return 0, err
	}

	tx, err := db.Beginx()
	if err != nil {
		return 0, err
	}

	err = tx.Get(&task, "INSERT INTO tasks (title, description, status, priority, due_date) VALUES (?, ?, ?, ?, ?)", task.Title, task.Description, task.Status, task.Priority, task.DueDate)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var taskID int
	err = tx.Get(&taskID, "SELECT LAST_INSERT_ID()")
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec("INSERT INTO ownerships (user_id, task_id) VALUES (?, ?)", userID, taskID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return taskID, nil
}

func UpdateTask(task database.Task) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, priority = ?, due_date = ? WHERE id = ?", task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(taskID int) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	// delete task from tasks table and ownerships table
	_, err = db.Exec("DELETE tasks, ownerships FROM tasks INNER JOIN ownerships ON tasks.id = ownerships.task_id WHERE tasks.id = ?", taskID)
	if err != nil {
		return err
	}

	return nil
}
