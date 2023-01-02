package models

import (
	database "taschola/db"
)

func GetUserByID(userID uint64) (database.User, error) {
	db, err := database.GetConnection()
	if err != nil {
		return database.User{}, err
	}

	var user database.User
	err = db.Get(&user, "SELECT * FROM users WHERE id = ?", userID)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func UpdateUser(user database.User) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	// userIDが一致するユーザーのnameとpasswordを更新
	_, err = db.Exec("UPDATE users SET name = ?, password = ? WHERE id = ?", user.Name, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(userID uint64) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	// delete user from users table
	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}

	// delete task from tasks table and ownerships table
	_, err = db.Exec("DELETE tasks, ownerships FROM tasks INNER JOIN ownerships ON tasks.id = ownerships.task_id WHERE ownerships.user_id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}
