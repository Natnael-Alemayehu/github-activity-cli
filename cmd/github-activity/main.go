package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/natnael-alemayehu/github-activity-cli/internal"
)

func main() {
	username := os.Args[1]
	fmt.Println(username)

	url := fmt.Sprintf(`https://api.github.com/users/%v/events`, username)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Printf("Content length: %v\n", res.ContentLength)

	if res.StatusCode == http.StatusOK {
		bodyByte, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyByte)

		file, err := internal.ReadFile("respBody.json")
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(bodyString)
	}

}
