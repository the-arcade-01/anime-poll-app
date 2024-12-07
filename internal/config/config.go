package config

import (
	"log"
	"os"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var appConfig *AppConfig

type AppConfig struct {
	ApiUrl   string
	DbClient *gorm.DB
}

func NewAppConfig() *AppConfig {
	once.Do(func() {
		db, err := newDBClient()
		if err != nil {
			log.Fatalln(err)
		}
		appConfig = &AppConfig{
			ApiUrl:   os.Getenv("API_URL"),
			DbClient: db,
		}
	})
	return appConfig
}
