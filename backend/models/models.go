package models

import (
	"github.com/jinzhu/gorm"
	"pullpicture/backend/conf"
)

var DB *gorm.DB

func InitDB() error {
	config := conf.AppConfig.Database
	dsn := "host=" + config.Host + 
		" port=" + config.Port + 
		" user=" + config.User + 
		" dbname=" + config.Name + 
		" password=" + config.Password + 
		" sslmode=" + config.SSLMode

	db, err := gorm.Open("postgres", dsn)
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