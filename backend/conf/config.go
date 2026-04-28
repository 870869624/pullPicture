package conf

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	Server struct {
		Port string `default:"8080"`
		Host string `default:"localhost"`
	}
	Database struct {
		Type     string `default:"postgres"`
		Host     string `default:"localhost"`
		Port     string `default:"5432"`
		Name     string `default:"postgres"`
		User     string `default:"postgres"`
		Password string `default:"password"`
		SSLMode  string `default:"disable"`
	}
}

var AppConfig Config

func InitConfig() error {
	return configor.Load(&AppConfig, "conf/config.yml")
}