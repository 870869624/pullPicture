package routes

import (
	"pullpicture/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	pictures := r.Group("/pictures")
	{
		pictures.GET("", controllers.GetPictures)
		pictures.GET("/:id", controllers.GetPicture)
		pictures.POST("", controllers.CreatePicture)
		pictures.PUT("/:id", controllers.UpdatePicture)
		pictures.DELETE("/:id", controllers.DeletePicture)
	}

	return r
}
