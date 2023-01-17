package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"taschola/models"
)

// Login
//
//	@Summary	Login
//	@Tags		authentication
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Failure	400	{object}	models.HTTPError
//	@Failure	401	{object}	models.HTTPError
//	@Failure	404	{object}	models.HTTPError
//	@Router		/v1/login [post]
func Login(ctx *gin.Context) {
	var userForm UserForm
	err := ctx.BindJSON(&userForm)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (userForm is invalid)",
			Place: "controllers.Login",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	user, err := models.GetUserByName(userForm.Name)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusNotFound,
			Error: "Not Found (such a user name does not exist)",
			Place: "controllers.Login",
		}
		ctx.JSON(http.StatusNotFound, httpError)
		return
	}
	// check password is correct
	if string(user.Password) != userForm.Password {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized (invalid password)",
			Place: "controllers.Login",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	session := sessions.Default(ctx)
	session.Set(userKey, user.ID)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "login success"})
}

// Logout
//
//	@Summary	Logout
//	@Tags		authentication
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Router		/v1/logout [post]
func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
