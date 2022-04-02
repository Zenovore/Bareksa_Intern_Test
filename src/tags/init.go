package tags

import (
	"log"
)

var (
	agent Agent
)

func init() {
	log.SetFlags(log.Llongfile)

	log.Println("Tags Initializing")

	agent = NewService()

	log.Println("Tags Initialiazed")
}
