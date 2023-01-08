package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userKey string = "user_id"

func LoginCheck(ctx *gin.Context) {
	userID := sessions.Default(ctx).Get(userKey)
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		ctx.Abort()
	} else {
		ctx.Next()
	}
}
