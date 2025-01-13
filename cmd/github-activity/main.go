package main

import (
	"log"

	"github.com/natnael-alemayehu/github-activity-cli/internal"
)

func main() {
	err := internal.ConsumeBody()
	if err != nil {
		log.Fatal(err)
	}

}
