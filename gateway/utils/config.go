package utils

import (
	// "log"
	"os"

	// "github.com/joho/godotenv"
)

type microServices struct {
	IncidentService string
	NIN             string
}
type Config struct {
	JWTSecret     string
	AppPort       string
	MongoURI      string
	JwtEpiration  string
	MicroServices microServices
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
		MicroServices: microServices{
			IncidentService: os.Getenv("INCIDENT_SERVICE"),
			NIN:             os.Getenv("NIN_SERVICE"),
		},
	}
	return cfg
}
