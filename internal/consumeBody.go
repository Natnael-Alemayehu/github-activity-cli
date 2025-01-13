package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/natnael-alemayehu/github-activity-cli/internal/data"
)

func ReadUsername() string {
	username := os.Args[1]
	return username
}
func ConsumeBody() error {
	usr := ReadUsername()

	url := fmt.Sprintf(`https://api.github.com/users/%v/events`, usr)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP method failed")
	}

	var items []data.Body

	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		return err
	}

	var pushMessages []string
	typeCount := make(map[string]int)
	activeRepos := make(map[string]int)
	activeRepos2 := make(map[string]int)
	// var starMessages []string

	for _, item := range items {
		typeCount[item.Type]++
		activeRepos2[item.RepoField.Name]++
		if strings.EqualFold(item.ActorField.Login, ReadUsername()) {
			activeRepos[item.RepoField.Name]++
		}

		if len(item.PayloadField.Commits) > 0 {
			commit := len(item.PayloadField.Commits)
			name := item.RepoField.Name

			newPushMessage := fmt.Sprintf("Pushed %d to %v \n", commit, name)
			pushMessages = append(pushMessages, newPushMessage)
		}

	}

	// for _, i := range pushMessages {
	// 	fmt.Printf("%v", i)
	// }

	// for key, val := range typeCount {
	// 	fmt.Printf("%v - %v \n", key, val)
	// }

	fmt.Println("Before")
	for key, val := range activeRepos {
		fmt.Printf("%v - %v \n", key, val)
	}

	fmt.Println("After")
	for key, val := range activeRepos2 {
		fmt.Printf("%v - %v \n", key, val)
	}

	return nil
}
