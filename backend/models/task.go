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
