package models

import (
	"pullpicture/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() error {
	config := conf.AppConfig.Database
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.LogMode(true)
	DB = db
	return nil
}

type Picture struct {
	gorm.Model
	URL      string `json:"url"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email"`
	Role     string `json:"role" gorm:"default:'user'"`
}
