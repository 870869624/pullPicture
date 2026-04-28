package main

import (
	"log"

	"pullpicture/backend/conf"
	"pullpicture/backend/models"
	"pullpicture/backend/routes"
)

func main() {
	if err := conf.InitConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	defer models.DB.Close()

	models.DB.AutoMigrate(&models.Picture{})

	r := routes.SetupRouter()

	port := conf.AppConfig.Server.Port
	log.Printf("Server is running on http://%s:%s", conf.AppConfig.Server.Host, port)
	r.Run(":" + port)
}