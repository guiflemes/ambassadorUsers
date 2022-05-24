package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func GETENV(key string) string {
	env, isPresent := os.LookupEnv(key)

	if !isPresent {
		log.Fatalf("%s is not present in .env", key)
	}
	return env
}
