package controllers

import (
	"net/http"
	database "taschola/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// logins
func Login(ctx *gin.Context) {
	var userForm UserForm
	err := ctx.BindJSON(&userForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (invalid user)"})
		return
	}
	db, err := database.GetConnection()
	if err != nil {
		return
	}
	var user database.User
	err = db.Get(&user, "SELECT id, name, password FROM users WHERE name = ? AND password = ?", userForm.UserInfo.Name, hash(userForm.UserInfo.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (incorrect password)"})
		return
	}

	session := sessions.Default(ctx)
	session.Set(userkey, user.ID)
	session.Save()
	ctx.Redirect(http.StatusFound, frontEndURL+"/tasks")
}

// logout
func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	ctx.Redirect(http.StatusFound, frontEndURL+"/")
}
