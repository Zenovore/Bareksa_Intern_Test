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

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
