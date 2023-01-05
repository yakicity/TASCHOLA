package controllers

import (
	"net/http"
	"strconv"
	"taschola/db"
	"taschola/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//newuser
//POST"v1/user/new"
//
// @Param user [database.User]
//
// @Success 200 {"user": user}
// @Failure 400 {"error": "Bad Request (invalid user)"}
// @Failure 409 {"error": "Conflict (duplicate user name)"}
// @Failure 500 {"error": "Internal Server Error"}
func CreateUser(ctx *gin.Context){
	// get user
	var user db.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, gin.H{"error": "Bad Request (invalid user)"})
		return
	}

	// check duplicate user name
	switch models.CheckDuplicateUserName(user.Name) {
	case true: // duplicate
		ctx.JSON(http.StatusConflict, gin.H{"error": "Conflict (duplicate user name)"})
		
		return
	default: // not duplicate
		// create user
		err = models.CreateUser(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}


//GET"v1/user/:id"
//
// @Param id [uint64]
// @Param user [database.User]
//
// @Success 200 { "user": database.User}
// @Failure 404 
func ShowUser(ctx *gin.Context){
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Bad Request (invalid user id)"})
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


//PUT"v1/user/:id"
//
// @Param id [uint64]
//
// @Success 200 { "user_id": uint64}
// @Failure 400 { "error": "invalid user id"}
// @Failure 401 { "error": "Internal Server Error"} 
// @Failure 409 {"error": "Conflict (duplicate user name)"}
func EditUser(ctx *gin.Context){
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Bad Request (invalid user id)"})
		return
	}
	// get user
	var user db.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, gin.H{"error": "Bad Request (invalid user)"})
		return
	}

	//重複判定
	if (models.CheckDuplicateUserName(user.Name)){
		ctx.JSON(http.StatusConflict, gin.H{"error": "Conflict (duplicate user name)"})
	}

	//update user information
	err = models.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_id": userID})
}

//userdelete
//DELETE"v1/user/:id"
//
// @Param id [uint64]
//
// @Success 200 { "user_id": uint64}
// @Failure 400 { "error": "invalid user id"}
// @Failure 401 { "error": "Internal Server Error"}
func DeleteUser(ctx *gin.Context){
	// get user_id
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Bad Request (invalid user id)"})
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
