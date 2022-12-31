package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"taschola/controllers"
)

func Init() *gin.Engine {
	// initialize engine
	engine := gin.Default()

	// endpoints version 1
	v1 := engine.Group("/api/v1")
	{
		// home
		v1.GET("/", controllers.NotImplemented)

		// tasks
		v1.GET("/tasks", controllers.NotImplemented)
		v1.GET("/tasks/t2schola_sync", controllers.NotImplemented)

		v1.GET("/task/:id", controllers.NotImplemented)
		v1.PUT("/task/:id", controllers.NotImplemented)
		v1.DELETE("/task/:id", controllers.NotImplemented)
		v1.POST("/task/new", controllers.NotImplemented)

		// auth
		v1.POST("/login", controllers.NotImplemented)
		v1.POST("/logout", controllers.NotImplemented)

		// user
		v1.POST("/user/new", controllers.NotImplemented)
		v1.GET("/user/:id", controllers.NotImplemented)
		v1.PUT("/user/:id", controllers.NotImplemented)    // edit
		v1.DELETE("/user/:id", controllers.NotImplemented) // delete
	}

	return engine
}
