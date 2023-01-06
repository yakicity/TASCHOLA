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

func CreateUser(user database.User) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	// insert user into users table
	_, err = db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", user.Name, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(userID uint64, user database.User) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	// userIDが一致するユーザーのnameとpasswordを更新
	_, err = db.Exec("UPDATE users SET name = ?, password = ? WHERE id = ?", user.Name, user.Password, userID)
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

func CheckDuplicateUserName(userName string) bool {
	// connect to database
	db, err := database.GetConnection()
	if err != nil {
		return true // return true if error occurs
	}

	var count int64
	// count number of users with the same user name
	err = db.Get(&count, "SELECT count(*) FROM users WHERE user_name = ?", userName)
	if err != nil {
		return true // return true if error occurs
	}
	// SELECT count(*) FROM users WHERE user_name = 'userName';
	return count > 0
	// return true if user name is already taken
}
