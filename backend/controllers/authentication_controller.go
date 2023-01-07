package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"taschola/models"
)

// POST /v1/login
//
// @Param userForm [UserForm]
//
// @Success 200 { "message": "login success" }
// @Failure 400 { "error": "Bad Request (userForm is invalid)" }
// @Failure 404 { "error": "Not Found (such a user name does not exist)" }
// @Failure 401 { "error": "Unauthorized (invalid password)" }
func Login(ctx *gin.Context) {
	var userForm UserForm
	err := ctx.BindJSON(&userForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (userForm is invalid)"})
		return
	}

	user, err := models.GetUserByName(userForm.Name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found (such a user name does not exist)"})
		return
	}
	// check password is correct
	if string(user.Password) != userForm.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized (invalid password)"})
		return
	}

	session := sessions.Default(ctx)
	session.Set(userKey, user.ID)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "login success"})
}

// POST /v1/logout
//
// # No parameters
//
// @Success 200 { "message": "logout success" }
func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
