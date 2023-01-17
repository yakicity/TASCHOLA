package controllers

import (
	"fmt"
	"net/http"
	database "taschola/db"

	"github.com/gin-gonic/gin"
)

// NotImplemented
//	@Summary		Not implemented
//	@Description	for not implemented methods
//	@Tags			default
//	@Accept			json
//	@Produce		json
//	@Success		501	{object}	models.Error
//	@Router			/ [get]
func NotImplemented(ctx *gin.Context) {
	msg := fmt.Sprintf("%s access to %s is not implemented yet", ctx.Request.Method, ctx.Request.URL)
	ctx.Header("Cache-Control", "no-cache")
	ctx.String(http.StatusNotImplemented, msg)
}

// HealthCheck
//	@Summary	Health check
// 	@Description	Health check (for Docker healthcheck)
//	@Tags		default
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	string	"OK"
//	@Router		/health [get]
func HealthCheck(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache")
	ctx.String(http.StatusOK, "OK")
}

// CheckDBConnection
//	@Summary	Check DB connection
// 	@Description	Check DB connection
//	@Tags		default
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	string	"OK"
//	@Failure	500	{object}	models.Error
//	@Router		/heathz [get]
func CheckDBConnection(ctx *gin.Context) {
	db, err := database.GetConnection()
	if err != nil {
		ctx.Header("Cache-Control", "no-cache")
		ctx.String(http.StatusInternalServerError, "DB connection error")
		return
	}
	err = db.Ping()
	if err != nil {
		ctx.Header("Cache-Control", "no-cache")
		ctx.String(http.StatusInternalServerError, "DB connection error")
		return
	}

	var tableNames []string
	err = db.Select(&tableNames, "SHOW TABLES")
	if err != nil {
		ctx.Header("Cache-Control", "no-cache")
		ctx.String(http.StatusInternalServerError, "DB connection error")
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.JSON(http.StatusOK, tableNames)
}
