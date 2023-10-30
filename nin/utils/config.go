package utils

import (
	// "log"
	"os"

	// "github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	MongoURI     string
}

func LoadEvn() *Config {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("error loading env file >> %s", err.Error())
	// }
	cfg := &Config{
		AppPort:      os.Getenv("APP_PORT"),
		MongoURI:     os.Getenv("MONGO_URI"),
	}
	return cfg
}
