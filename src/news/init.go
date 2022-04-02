package news

import (
	"bareksaIntern/src/tags"
	"log"
)

var (
	agent     Agent
	tagsAgent tags.Agent
)

func init() {
	log.SetFlags(log.Llongfile)

	log.Println("News Initializing")

	agent = NewService()
	tagsAgent = tags.NewService()

	log.Println("News Initialiazed")
}
