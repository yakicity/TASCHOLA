package controllers

import (
	"net/http"
	"taschola/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get tasks from database
	tasks, err := models.GetTasksByUserID(userID.(uint64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func GetTasksByKeywordAndStatus(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	keyword := ctx.Query("keyword")
	status := ctx.QueryArray("status")
	if len(status) == 0 {
		status = []string{"TODO", "DOING", "DONE"}
	}

	// Get tasks from database
	tasks, err := models.GetTasksByUserIDAndKeywordAndStatus(userID.(uint64), keyword, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
