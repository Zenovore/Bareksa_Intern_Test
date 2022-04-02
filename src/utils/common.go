package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func GetTimeNow() (timeNow time.Time) {
	location, err := time.LoadLocation(GetEnvVar("TIMEZONE"))
	if err != nil {
		log.Fatal(err)
	}
	timeNow = time.Now().In(location)
	return
}
