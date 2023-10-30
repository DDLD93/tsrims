package utils

import (
	// "log"
	"os"

	// "github.com/joho/godotenv"
)

type Config struct {
	JWTSecret    string
	AppPort      string
	MongoURI     string
	JwtEpiration string
}

func LoadEvn() *Config {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("error loading env file >> %s", err.Error())
	// }
	cfg := &Config{
		JWTSecret:    os.Getenv("JWT_SECRET"),
		AppPort:      os.Getenv("APP_PORT"),
		MongoURI:     os.Getenv("MONGO_URI"),
		JwtEpiration: os.Getenv("JWT_EXPIRATION"),
	}
	return cfg
}
