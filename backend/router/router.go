package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"taschola/controllers"
)

func Init() *gin.Engine {
	// initialize engine
	engine := gin.Default()

	// CORS settings
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// health check
	engine.GET("/health", controllers.HealthCheck)

	// endpoints version 1
	v1 := engine.Group("v1")
	{
		// home
		v1.GET("/", controllers.NotImplemented)

		// tasks
		tasks := v1.Group("/tasks")
		tasks.Use(controllers.LoginCheck)
		{
			tasks.GET("/", controllers.GetTasks)
			tasks.GET("/t2schola_sync", controllers.NotImplemented)
		}
		task := v1.Group("/task")
		task.Use(controllers.LoginCheck)
		{
			task.GET("/:id", controllers.GetTask)
			task.PUT("/:id", controllers.UpdateTask)
			task.DELETE("/:id", controllers.DeleteTask)
			task.POST("/new", controllers.CreateTask)
		}

		// auth
		v1.POST("/login", controllers.Login)
		v1.POST("/logout", controllers.Logout)

		// user
		v1.POST("/user/new", controllers.CreateUser)
		user := v1.Group("/user")
		user.Use(controllers.LoginCheck)
		{
			v1.GET("/:id", controllers.GetUser)
			v1.PUT("/:id", controllers.UpdateUser)    // edit
			v1.DELETE("/:id", controllers.DeleteUser) // delete
		}
	}

	return engine
}
