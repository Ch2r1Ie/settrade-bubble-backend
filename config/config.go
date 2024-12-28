package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server          Server
	AccessControl   AccessControl
	Database        Database
	Header          Header
	YahooFinanceURL string
}

func Init() *Config {

	config(localhost_path, secret_path)
	location(os.Getenv(timezone))
	enviroment(os.Getenv(app_prefix))

	appCfg := initAppCfg(os.Getenv(app_prefix))

	return appCfg
}

func config(localhost_path, secret_path string) {

	if err := godotenv.Load(localhost_path); err != nil {
		panic("Error loading localhost.configmap.env file")
	}

	if err := godotenv.Load(secret_path); err != nil {
		panic("Error loading localhost.secret.env file")
	}
}

func location(timezone string) {
	ict, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("error loading location '%s': %v\n", timezone, err)
	}
	time.Local = ict
}

func initAppCfg(prefix string) *Config {
	return &Config{
		Server: Server{
			Hostname: os.Getenv(connected(prefix, hostname)),
			Port:     os.Getenv(connected(prefix, port)),
		},
		YahooFinanceURL: os.Getenv(connected(prefix, yahoo_finance)),
	}
}

func connected(prefix, config string) string {
	return prefix + interconnected + config
}
