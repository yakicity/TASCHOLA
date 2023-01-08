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

// GET v1/user/:id
//
// @Param id [uint64]
// @Param user [database.User]
//
// @Success 200 { "user": database.User}
// @Failure 404
func GetUser(ctx *gin.Context) {
	userID, err := strconv.ParseUint((ctx.Param("id")), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (invalid user id)"})
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// POST v1/user/new
//
// @Param user [database.User]
//
// @Success 200 {"user": user}
// @Failure 400 {"error": "Bad Request (invalid user)"}
// @Failure 409 {"error": "Conflict (duplicate user name)"}
// @Failure 500 {"error": "Internal Server Error"}

func CreateUser(ctx *gin.Context) {
	// get user
	var userInfo UserForm
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (userForm is invalid)"})
		return
	}

	user := encryptUserInfo(userInfo)
	// check duplicate user name
	switch models.CheckDuplicateUserName(user.Name) {
	case true: // duplicate
		ctx.JSON(http.StatusConflict, gin.H{"error": "Conflict (duplicate user name)"})
		return

	default: // not duplicate
		// create user
		err := models.CreateUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// PUT v1/user/:id
//
// @Param id [uint64]
//
// @Success 200 { "user_id": uint64}
// @Failure 400 { "error": "invalid user id"}
// @Failure 401 { "error": "Internal Server Error"}
// @Failure 409 {"error": "Conflict (duplicate user name)"}
func UpdateUser(ctx *gin.Context) {
	userID, err := strconv.ParseUint((ctx.Param("id")), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (invalid user id)"})
		return
	}

	var userForm UpdateUserForm
	err = ctx.BindJSON(&userForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (updateUserForm is invalid)"})
		return
	}

	// check old password
	oldUserInfo, err := models.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "The user is not found"})
		return
	}
	if oldUserInfo.Password == nil || string(oldUserInfo.Password) != string(hash(userForm.OldPassword)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (old password is wrong)"})
		return
	}

	// check duplicate user name
	switch models.CheckDuplicateUserName(userForm.User.Name) {
	case true: // duplicate
		ctx.JSON(http.StatusConflict, gin.H{"error": "Conflict (duplicate user name)"})
		return

	default: // not duplicate
		// update user
		user := encryptUserInfo(userForm.User)
		err := models.UpdateUser(userID, user)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user_id": userID})
	}
}

// DELETE v1/user/:id
//
// @Param id [uint64]
//
// @Success 200 { "user_id": uint64}
// @Failure 400 { "error": "invalid user id"}
// @Failure 401 { "error": "Internal Server Error"}
func DeleteUser(ctx *gin.Context) {
	// get user_id
	userID, err := strconv.ParseUint((ctx.Param("id")), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request (invalid user id)"})
		return
	}

	// delete user
	err = models.DeleteUser(userID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_id": userID})
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
