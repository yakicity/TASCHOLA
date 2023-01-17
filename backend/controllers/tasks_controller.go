package controllers

import (
	"net/http"
	"strconv"
	"taschola/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetTasks
//
//	@Summary	Get tasks
//	@Tags		task
//	@Accept		json
//	@Produce	json
//	@Param		keyword	query		string		false	"keyword"
//	@Param		status	query		[]string	false	"status"
//	@Success	200		{object}	[]db.Task
//	@Failure	401		{object}	models.HTTPError
//	@Failure	500		{object}	models.HTTPError
//	@Router		/v1/tasks [get]
func GetTasks(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized",
			Place: "controllers.GetTasks",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	// Get tasks from database
	tasks, err := models.GetTasksByUserID(userID.(uint64))
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error (models error)" + err.Error(),
			Place: "controllers.GetTasks",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// GetTasksByKeywordAndStatus
//
//	@Summary	Get tasks by keyword and status
//	@Tags		task
//	@Accept		json
//	@Produce	json
//	@Param		keyword	query		string		false	"keyword"
//	@Param		status	query		[]string	false	"status"
//	@Success	200		{object}	[]db.Task
//	@Failure	401		{object}	models.HTTPError
//	@Failure	500		{object}	models.HTTPError
//	@Router		/v1/tasks [get]
func GetTasksByKeywordAndStatus(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized",
			Place: "controllers.GetTasksByKeywordAndStatus",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
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
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error (models error)" + err.Error(),
			Place: "controllers.GetTasksByKeywordAndStatus",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// GetTask
//
//	@Summary	Get task
//	@Tags		task
//	@Accept		json
//	@Produce	json
//	@Param		task_id	path		uint64	true	"task ID"
//	@Success	200		{object}	db.Task
//	@Failure	401		{object}	models.HTTPError
//	@Failure	400		{object}	models.HTTPError
//	@Failure	500		{object}	models.HTTPError
//	@Router		/v1/tasks/{task_id} [get]
func GetTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized",
			Place: "controllers.GetTask",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	// Get task ID from request
	taskID := ctx.Param("task_id")
	taskIDUint64, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Invalid task ID",
			Place: "controllers.GetTask",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// Get task from database
	task, err := models.GetTaskByUserIDAndTaskID(userID.(uint64), taskIDUint64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error (models error)" + err.Error(),
			Place: "controllers.GetTask",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// CreateTask
//
//	@Summary	Create task
//	@Tags		task
//	@Accept		json
//	@Produce	json
//	@Success	200	integer		task	ID
//	@Failure	400	{object}	models.HTTPError
//	@Failure	401	{object}	models.HTTPError
//	@Failure	500	{object}	models.HTTPError
//	@Router		/v1/tasks [post]
func CreateTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized",
			Place: "controllers.CreateTask",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	// Get task from request
	var task models.TaskForm
	err := ctx.BindJSON(&task)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Invalid task form",
			Place: "controllers.CreateTask",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// Create task in database
	taskID, err := models.CreateTask(task, userID.(uint64))
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error (models error)" + err.Error(),
			Place: "controllers.CreateTask",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, taskID)
}

// UpdateTask
//
//	@Summary	Update task
//	@Tags		task
//	@Accept		json
//	@Produce	json
//	@Param		task_id	path		uint64	true	"task ID"
//	@Success	200		integer		task	ID
//	@Failure	400		{object}	models.HTTPError
//	@Failure	401		{object}	models.HTTPError
//	@Failure	500		{object}	models.HTTPError
//	@Router		/v1/tasks/{task_id} [put]
func UpdateTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized",
			Place: "controllers.UpdateTask",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	// Get task ID from request
	taskID := ctx.Param("task_id")
	taskIDUint64, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Invalid task ID",
			Place: "controllers.UpdateTask",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// Get task from request
	var task models.TaskForm
	err = ctx.BindJSON(&task)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Invalid task form",
			Place: "controllers.UpdateTask",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// Update task in database
	err = models.UpdateTaskByUserIDAndTaskID(task, userID.(uint64), taskIDUint64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error (models error)" + err.Error(),
			Place: "controllers.UpdateTask",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, taskID)
}

// DeleteTask
//
//	@Summary	Delete task
//	@Tags		task
//	@Accept		json
//	@Produce	json
//	@Param		task_id	path		uint64	true	"task ID"
//	@Success	200		integer		task	ID
//	@Failure	400		{object}	models.HTTPError
//	@Failure	401		{object}	models.HTTPError
//	@Failure	500		{object}	models.HTTPError
//	@Router		/v1/tasks/{task_id} [delete]
func DeleteTask(ctx *gin.Context) {
	// Get user ID from sessions
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		httpError := models.HTTPError{
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized",
			Place: "controllers.DeleteTask",
		}
		ctx.JSON(http.StatusUnauthorized, httpError)
		return
	}

	// Get task ID from request
	taskID := ctx.Param("task_id")
	taskIDUint64, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusBadRequest,
			Error: "Invalid task ID",
			Place: "controllers.DeleteTask",
		}
		ctx.JSON(http.StatusBadRequest, httpError)
		return
	}

	// Delete task from database
	err = models.DeleteTaskByUserIDAndTaskID(userID.(uint64), taskIDUint64)
	if err != nil {
		httpError := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "Internal Server Error (models error)" + err.Error(),
			Place: "controllers.DeleteTask",
		}
		ctx.JSON(http.StatusInternalServerError, httpError)
		return
	}

	ctx.JSON(http.StatusOK, taskID)
}
