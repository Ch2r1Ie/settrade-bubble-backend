package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	apiKey string
}

func Init() error {

	env()
	location(os.Getenv(timezone))
	initAppCfg(os.Getenv(app_prefix))

	return nil
}

func env() {

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
	log.Printf("Local time zone %v", time.Now().In(ict))
}

func initAppCfg(prefix string) AppConfig {
	return AppConfig{
		apiKey: connected(prefix, os.Getenv(apiKey)),
	}
}

func connected(prefix, config string) string {
	return prefix + interconnected + config
}
