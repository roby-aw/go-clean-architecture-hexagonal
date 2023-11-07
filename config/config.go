package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// AppConfig Application configuration structure
type AppConfig struct {
	App struct {
		Port string `toml:"port"`
	} `toml:"app"`
	Database struct {
		DBURL  string
		DBNAME string
		DBPASS string
		DBUSER string
		DBETC  string
		DBPORT string
	} `toml:"database"`
	JWT struct {
		Secret string `toml:"token"`
	} `toml:"secrettoken"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	_ = godotenv.Load()
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var finalConfig AppConfig

	finalConfig.App.Port = "8080"
	finalConfig.Database.DBNAME = os.Getenv("DBNAME")
	finalConfig.Database.DBURL = os.Getenv("DBURL")
	finalConfig.Database.DBPORT = os.Getenv("DBPORT")
	finalConfig.Database.DBPASS = os.Getenv("DBPASS")
	finalConfig.Database.DBUSER = os.Getenv("DBUSER")
	finalConfig.Database.DBETC = os.Getenv("DBETC")

	return &finalConfig
}
