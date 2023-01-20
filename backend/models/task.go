package models

import (
	"strings"
	database "taschola/db"
	"time"
)

type TaskForm struct {
	Title       string
	Description string
	Status      string // [TODO, DOING, DONE]
	Priority    int    // [1, 2, 3, 4, 5] (1: highest, 5: lowest)
	DueDate     time.Time
}

func GetTasksByUserID(userID int64) ([]database.Task, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var tasks []database.Task
	err = db.Select(&tasks,
		"SELECT tasks.id, tasks.title, tasks.description, tasks.status, tasks.priority, tasks.created_at, tasks.due_date FROM tasks INNER JOIN ownerships ON tasks.id = ownerships.task_id WHERE ownerships.user_id = ? ORDER BY tasks.due_date ASC, tasks.priority ASC", userID) // 締切が近く、優先度の高い順に並べる
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksByUserIDAndKeywordAndStatus(userID int64, keyword string, status []string) ([]database.Task, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}
	if len(status) == 0 {
		status = []string{"TODO", "DOING", "DONE"}
	}
	s := strings.Join(status, ",")

	var tasks []database.Task
	err = db.Select(&tasks, "SELECT tasks.id, tasks.title, tasks.description, tasks.status, tasks.priority, tasks.created_at, tasks.due_date FROM tasks INNER JOIN ownerships ON tasks.id = ownerships.task_id WHERE ownerships.user_id = ? AND (tasks.title LIKE ? OR tasks.description LIKE ?) AND tasks.status IN (?) ORDER BY tasks.due_date ASC, tasks.priority ASC", userID, "%"+keyword+"%", "%"+keyword+"%", s) // 締切が近く、優先度の高い順に並べる
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByUserIDAndTaskID(userID int64, taskID int64) (database.Task, error) {
	db, err := database.GetConnection()
	if err != nil {
		return database.Task{}, err
	}

	var task database.Task
	err = db.Get(&task, "SELECT tasks.id, tasks.title, tasks.description, tasks.status, tasks.priority, tasks.created_at, tasks.due_date FROM tasks INNER JOIN ownerships ON tasks.id = ownerships.task_id WHERE ownerships.user_id = ? AND tasks.id = ?", userID, taskID)
	if err != nil {
		return database.Task{}, err
	}

	return task, nil
}

func CreateTask(task TaskForm, userID int64) (int, error) {
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

func UpdateTaskByUserIDAndTaskID(task TaskForm, userID int64, taskID int64) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, priority = ?, due_date = ? WHERE id = ? AND id IN (SELECT task_id FROM ownerships WHERE user_id = ?)", task.Title, task.Description, task.Status, task.Priority, task.DueDate, taskID, userID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTaskByUserIDAndTaskID(userID int64, taskID int64) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM tasks WHERE id = ? AND id IN (SELECT task_id FROM ownerships WHERE user_id = ?)", taskID, userID)
	if err != nil {
		return err
	}

	return nil
}
