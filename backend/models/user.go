package models

import (
	"log"
	database "taschola/db"
)

// user information whose password is plaintext
// HTTP Request Body From Frontend
type EncryptedUserForm struct {
	Name     string `json:"name" binding:"required"`
	Password []byte `json:"password" binding:"required"`
}

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

func GetUserByName(name string) (database.User, error) {
	db, err := database.GetConnection()
	if err != nil {
		return database.User{}, err
	}

	var user database.User
	err = db.Get(&user, "SELECT * FROM users WHERE name = ?", name) // unique constraint により、name は一意である
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func CreateUser(user EncryptedUserForm) error {
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

func UpdateUser(userID uint64, user EncryptedUserForm) error {
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
		log.Println("Check Duplicate User Name: DB connection error: " + err.Error())
		return true // return true if error occurs
	}

	var count int64
	// count number of users with the same user name
	err = db.Get(&count, "SELECT count(*) FROM users WHERE name = ?", userName)
	if err != nil {
		log.Println("Check Duplicate User Name: DB query error: " + err.Error())
		return true // return true if error occurs
	}
	// SELECT count(*) FROM users WHERE name = 'userName';
	return count > 0
	// return true if user name is already taken
}
