package main

import (
	"log"

	"pullpicture/backend/conf"
	"pullpicture/backend/models"
	"pullpicture/backend/routes"

	"golang.org/x/crypto/bcrypt"
)

func initAdmin() {
	var admin models.User
	if err := models.DB.Where("username = ?", "admin").First(&admin).Error; err == nil {
		log.Println("Admin user already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash admin password: %v", err)
	}

	admin = models.User{
		Username: "admin",
		Password: string(hashedPassword),
		Email:    "admin@example.com",
		Role:     "admin",
	}

	if err := models.DB.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Println("Admin user created successfully")
}

func main() {
	if err := conf.InitConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	defer models.DB.Close()

	models.DB.AutoMigrate(&models.Picture{}, &models.User{})

	initAdmin()

	r := routes.SetupRouter()

	port := conf.AppConfig.Server.Port
	log.Printf("Server is running on http://%s:%s", conf.AppConfig.Server.Host, port)
	r.Run(":" + port)
}