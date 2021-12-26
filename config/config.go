package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Function making for read configuration from env file by parameter key
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error while loading .env file")
	}
	return os.Getenv(key)
}

//Function making for initial jwt secret for authorize
var JWTSecret string

func InitJWT() {
	JWTSecret = Config("APP_SECRET")
}
