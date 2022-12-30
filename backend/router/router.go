package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *gin.Engine {
	// initialize engine
	engine := gin.Default()

	return engine
}
