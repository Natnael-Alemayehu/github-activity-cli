package main

import (
	"fmt"
	"log"
	"os"

	"github.com/natnael-alemayehu/github-activity-cli/internal"
)

func main() {
	if os.Args[0] != "github-activity" {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}
	pushMessage, typeCount, activeRepos, err := internal.ConsumeBody()
	if err != nil {
		log.Fatal(err)
	}

	// psuhMessage printed
	fmt.Println("Push Message")
	for _, i := range pushMessage {
		fmt.Printf("%v", i)
	}

	// typeCount printed
	fmt.Println("\nEvent types")
	for key, val := range typeCount {
		fmt.Printf("%v - %v \n", key, val)
	}

	// activeRepos
	fmt.Println("\nRepositories user interacted with")
	for key, val := range activeRepos {
		fmt.Printf("%v - %v \n", key, val)
	}

}
