package utils

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	database *sqlx.DB
)

func init() {
	log.SetFlags(log.Llongfile)
	log.Println("Connection Pool Initializing")

	InitDB()

	log.Println("Connection Pool Initialiazed")
}

func InitDB() {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		GetEnvVar("POSTGRES_USER"),
		GetEnvVar("POSTGRES_PASSWORD"),
		GetEnvVar("POSTGRES_HOST"),
		GetEnvVar("POSTGRES_PORT"),
		GetEnvVar("POSTGRES_DB"))

	db, err := sqlx.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	database = db
}

func GetDB() *sqlx.DB {
	return database
}
