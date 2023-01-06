package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userkey = "user"
const frontEndURL string = "http://localhost:3000"

// Logincheck
func LoginCheck(ctx *gin.Context) {
	userID := sessions.Default(ctx).Get(userkey)
	if userID == nil {
		ctx.Redirect(http.StatusFound, frontEndURL+"/login")
		ctx.Abort()
	} else {
		ctx.Next()
	}
}
