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
		Type     string `default:"mysql"`
		Host     string `default:"113.44.79.143"`
		Port     string `default:"3306"`
		Name     string `default:"pullpicture"`
		User     string `default:"pullpicture"`
		Password string `default:"EHWWxrhJthEHy5Ez"`
		SSLMode  string `default:"disable"`
	}
}

var AppConfig Config

func InitConfig() error {
	return configor.Load(&AppConfig, "conf/config.yml")
}
