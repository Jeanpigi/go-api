package loadenv

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carga todas la variables de entorno del archivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}