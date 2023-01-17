package controllers

import (
	"fmt"
	"net/http"

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
