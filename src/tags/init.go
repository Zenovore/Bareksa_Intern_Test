package tags

import (
	"log"
)

var (
	agent Agent
)

func init() {
	log.SetFlags(log.Llongfile)

	log.Println("Account Initializing")

	agent = NewService()

	log.Println("Account Initialiazed")
}
