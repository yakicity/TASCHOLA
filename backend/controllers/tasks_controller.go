package controllers

import (
	"net/http"
	"strconv"
	"taschola/db"
	"taschola/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GET /v1/tasks
//
// # No parameters
//
// @Success 200 { "tasks": []database.Task }
// @Failure 401 { "error": "unauthorized" }
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

// GET /v1/tasks?keyword=foo&status=TODO&status=DOING&status=DONE
//
// @Param keyword [string]
// @Param status [[]string]
//
// @Success 200 { "tasks": []database.Task }
// @Failure 500 { "error": "Internal Server Error" }
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

// GET /v1/task/:task_id
//
// @Param task_id [uint64]
//
// @Success 200 { "task": database.Task }
// @Failure 400 { "error": "invalid task ID" }
// @Failure 401 { "error": "unauthorized" }
func GetTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get task ID from request
	taskID := ctx.Param("task_id")
	taskIDUint64, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	// Get task from database
	task, err := models.GetTaskByUserIDAndTaskID(userID.(uint64), taskIDUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// POST /v1/task/new
//
// @Param task [database.Task]
//
// @Success 200 { "task_id": uint64 }
// @Failure 400 { "error": "invalid task" }
// @Failure 401 { "error": "unauthorized" }
func CreateTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get task from request
	var task db.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task"})
		return
	}

	// Create task in database
	taskID, err := models.CreateTask(task, userID.(uint64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task_id": taskID})
}

// PUT /v1/task/:task_id
//
// @Param task_id [uint64]
// @Param task [database.Task]
//
// @Success 200 { "task_id": uint64 }
// @Failure 400 { "error": "invalid task ID" }
// @Failure 401 { "error": "unauthorized" }
func UpdateTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get task ID from request
	taskID := ctx.Param("task_id")
	taskIDUint64, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	// Get task from request
	var task db.Task
	err = ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task"})
		return
	}

	// Update task in database
	err = models.UpdateTaskByUserIDAndTaskID(task, userID.(uint64), taskIDUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task_id": taskID})
}

// DELETE /v1/task/:task_id
//
// @Param task_id [uint64]
//
// @Success 200 { "task_id": uint64 }
// @Failure 400 { "error": "invalid task ID" }
// @Failure 401 { "error": "unauthorized" }
func DeleteTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get task ID from request
	taskID := ctx.Param("task_id")
	taskIDUint64, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	// Delete task from database
	err = models.DeleteTaskByUserIDAndTaskID(userID.(uint64), taskIDUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task_id": taskID})
}
