package routes

import (
	"github.com/gin-gonic/gin"
	"pullpicture/backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

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