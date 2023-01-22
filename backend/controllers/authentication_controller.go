package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"taschola/models"
)

// Login
//
//	@Summary		Login
//	@Description	Set cookie "user_id" if login success
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			userForm	body	controllers.UserForm	true	"user"
//	@Success		200
//	@Failure		400	{object}	models.HTTPError
//	@Failure		401	{object}	models.HTTPError
//	@Failure		404	{object}	models.HTTPError
//	@Router			/v1/login [post]
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
	if string(user.Password) != string(hash(userForm.Password)) {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized (invalid password)",
			Place: "controllers.Login",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	cookie, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.SetCookie("user_id", strconv.FormatUint(user.ID, 10), 3600, "/", "localhost", false, true)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login success" + cookie})
}

// Logout
//
//	@Summary		Logout
//	@Description	Delete cookie "user_id" if logout success
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Cookie			user_id	string	true	"user id"
//	@Success		200
//	@Router			/v1/logout [post]
func Logout(ctx *gin.Context) {
	if _, err := ctx.Cookie("user_id"); err == nil {
		ctx.SetCookie("user_id", "", -1, "/", "localhost", false, true)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "logout failed (cookie does not exist)"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
