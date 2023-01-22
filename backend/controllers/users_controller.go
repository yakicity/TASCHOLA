package controllers

import (
	"crypto/sha256"
	"net/http"
	"strconv"
	"taschola/models"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserForm struct {
	OldPassword string   `json:"old_password"`
	User        UserForm `json:"user"`
}

// GetUser
//
//	@Summary	Get User
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		id	path		uint64	true	"user id"
//	@Success	200	{object}	db.User
//	@Failure	400	{object}	models.HTTPError
//	@Failure	404	{object}	models.HTTPError
//	@Router		/v1/user/{id} [get]
func GetUser(ctx *gin.Context) {
	userID, err := strconv.ParseUint((ctx.Param("id")), 10, 64)

	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (invalid user id)",
			Place: "controllers.GetUser",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusNotFound,
			Error: "Not Found (such a user id does not exist)",
			Place: "controllers.GetUser",
		}
		ctx.JSON(http.StatusNotFound, httpError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// CreateUser
//
//	@Summary	Create User
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		userForm	body	UserForm	true	"user"
//	@Success	200
//	@Failure	400	{object}	models.HTTPError
//	@Failure	409	{object}	models.HTTPError
//	@Failure	500	{object}	models.HTTPError
//	@Router		/v1/user/new [post]
func CreateUser(ctx *gin.Context) {
	// get user
	var userInfo UserForm
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (userForm is invalid)",
			Place: "controllers.CreateUser",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	user := encryptUserInfo(userInfo)
	// check duplicate user name
	switch models.CheckDuplicateUserName(user.Name) {
	case true: // duplicate
		httpError := models.HTTPError{
			Code:  http.StatusConflict,
			Error: "Conflict (duplicate user name)",
			Place: "controllers.CreateUser",
		}
		ctx.JSON(http.StatusConflict, httpError)
		return

	default: // not duplicate
		// create user
		err := models.CreateUser(user)
		if err != nil {
			httpError := models.HTTPError{
				Code:  http.StatusInternalServerError,
				Error: "Internal Server Error: " + err.Error(),
				Place: "controllers.CreateUser",
			}
			ctx.JSON(http.StatusInternalServerError, httpError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// UpdateUser
//
//	@Summary	Update User
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		id		path		uint64			true	"user id"
//	@Param		user	body		UpdateUserForm	true	"user"
//	@Success	200		integer		user			ID
//	@Failure	400		{object}	models.HTTPError
//	@Failure	401		{object}	models.HTTPError
//	@Failure	409		{object}	models.HTTPError
//	@Failure	500		{object}	models.HTTPError
//	@Router		/v1/user/{id} [put]
func UpdateUser(ctx *gin.Context) {
	userID, err := strconv.ParseUint((ctx.Param("id")), 10, 64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (invalid user id)",
			Place: "controllers.UpdateUser",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	var userForm UpdateUserForm
	err = ctx.BindJSON(&userForm)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (updateUserForm is invalid)",
			Place: "controllers.UpdateUser",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// check old password
	oldUserInfo, err := models.GetUserByID(userID)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Not Found (such a user id does not exist)",
			Place: "controllers.UpdateUser",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}
	if oldUserInfo.Password == nil || string(oldUserInfo.Password) != string(hash(userForm.OldPassword)) {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (old password is wrong)",
			Place: "controllers.UpdateUser",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// check duplicate user name
	switch models.CheckDuplicateUserName(userForm.User.Name) {
	case true: // duplicate
		httpError := models.HTTPError{
			Code:  http.StatusConflict,
			Error: "Conflict (duplicate user name)",
			Place: "controllers.UpdateUser",
		}
		ctx.JSON(http.StatusConflict, httpError)
		return

	default: // not duplicate
		// update user
		user := encryptUserInfo(userForm.User)
		err := models.UpdateUser(userID, user)
		if err != nil {
			httpError := models.HTTPError{
				Code:  http.StatusInternalServerError,
				Error: "Internal Server Error: " + err.Error(),
				Place: "controllers.UpdateUser",
			}
			ctx.JSON(http.StatusInternalServerError, httpError)
			return
		}

		ctx.JSON(http.StatusOK, userID)
	}
}

// DeleteUser
//
//	@Summary	Delete User
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		id	path		uint64	true	"user id"
//	@Success	200	integer		user	ID
//	@Failure	400	{object}	models.HTTPError
//	@Failure	500	{object}	models.HTTPError
//	@Router		/v1/user/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	// get user_id
	userID, err := strconv.ParseUint((ctx.Param("id")), 10, 64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Bad Request (invalid user id)",
			Place: "controllers.DeleteUser",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// delete user
	err = models.DeleteUser(userID)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error: " + err.Error(),
			Place: "controllers.DeleteUser",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, userID)
}

// private functions

func hash(pw string) []byte {
	const salt = "blue_taschola"
	h := sha256.New()
	h.Write([]byte(salt))
	h.Write([]byte(pw))
	return h.Sum(nil)
}

func encryptUserInfo(userInfo UserForm) models.EncryptedUserForm {
	var user models.EncryptedUserForm
	user.Name = userInfo.Name
	user.Password = hash(userInfo.Password)

	return user
}
