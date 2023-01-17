package controllers

import (
	"fmt"
	"net/http"
	database "taschola/db"
	"taschola/models"

	"github.com/gin-gonic/gin"
)

// NotImplemented
//
//	@Summary		Not implemented
//	@Description	for not implemented methods
//	@Tags			default
//	@Accept			json
//	@Produce		json
//	@Success		501	{object}	models.HTTPError
//	@Router			/ [get]
func NotImplemented(ctx *gin.Context) {
	msg := fmt.Sprintf("%s access to %s is not implemented yet", ctx.Request.Method, ctx.Request.URL)
	err := models.HTTPError{
		Code:  http.StatusNotImplemented,
		Error: msg,
		Place: "NotImplemented",
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.JSON(http.StatusNotImplemented, err)
}

// HealthCheck
//
//	@Summary		Health check
//	@Description	Health check (for Docker health check)
//	@Tags			default
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	OK	"Health check OK"
//	@Router			/health [get]
func HealthCheck(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache")
	ctx.String(http.StatusOK, "OK")
}

// CheckDBConnection
//
//	@Summary		Check DB connection
//	@Description	Check DB connection
//	@Tags			default
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		string	"OK"
//	@Failure		500	{object}	models.HTTPError
//	@Router			/healthz [get]
func CheckDBConnection(ctx *gin.Context) {
	db, err := database.GetConnection()
	if err != nil {
		ctx.Header("Cache-Control", "no-cache")
		errResponse := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "DB connection error: " + err.Error(),
			Place: "CheckDBConnection",
		}
		ctx.JSON(http.StatusInternalServerError, errResponse.Error)
		return
	}
	err = db.Ping()
	if err != nil {
		errResponse := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "DB connection error: " + err.Error(),
			Place: "CheckDBConnection",
		}
		ctx.Header("Cache-Control", "no-cache")
		ctx.JSON(http.StatusInternalServerError, errResponse.Error)
		return
	}

	var tableNames []string
	err = db.Select(&tableNames, "SHOW TABLES")
	if err != nil {
		errResponse := models.HTTPError{
			Code:  http.StatusInternalServerError,
			Error: "DB connection error: " + err.Error(),
			Place: "CheckDBConnection",
		}
		ctx.Header("Cache-Control", "no-cache")
		ctx.JSON(http.StatusInternalServerError, errResponse.Error)
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.JSON(http.StatusOK, tableNames)
}
