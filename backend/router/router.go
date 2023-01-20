package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "taschola/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"taschola/controllers"
)

func Init() *gin.Engine {
	// initialize engine
	engine := gin.Default()

	// CORS settings
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// jwt
	authMiddleware, _ := controllers.JWTInit()

	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health check
	engine.GET("/health", controllers.HealthCheck)
	engine.GET("/healthz", controllers.CheckDBConnection)

	// endpoints version 1
	v1 := engine.Group("v1")
	{
		// tasks
		v1.GET("/t2schola_sync", controllers.NotImplemented)

		task := v1.Group("/tasks")
		{
			task.GET("/", controllers.GetTasks)
			task.GET("/:id", controllers.GetTask)
			task.PUT("/:id", controllers.UpdateTask)
			task.DELETE("/:id", controllers.DeleteTask)
			task.POST("/new", controllers.CreateTask)
		}

		// auth
		v1.POST("/login", authMiddleware.LoginHandler)
		v1.POST("/logout", authMiddleware.LogoutHandler)

		auth := v1.Group("/auth")
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.Use(authMiddleware.MiddlewareFunc())
		{
			auth.GET("/hello", controllers.HelloHandler)
		}

		// user
		v1.POST("/user/new", controllers.CreateUser)

		user := v1.Group("/user")
		// user.Use(controllers.LoginCheck)
		{
			user.GET("/:id", controllers.GetUser)
			user.PUT("/:id", controllers.UpdateUser)    // edit
			user.DELETE("/:id", controllers.DeleteUser) // delete
		}
	}

	engine.NoRoute(controllers.NotFound)

	return engine
}
