package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/natnael-alemayehu/github-activity-cli/internal/data"
)

func ReadUsername() string {
	username := os.Args[1]
	return username
}

// ConsumeBody returns the
// 1. pushMessage as a []string
// 2. typeCount as map[string]int
// 3. activeRepos as map[string]int
func ConsumeBody() ([]string, map[string]int, map[string]int, error) {
	usr := ReadUsername()

	url := fmt.Sprintf(`https://api.github.com/users/%v/events`, usr)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, nil, nil, fmt.Errorf("HTTP method failed")
	}

	var items []data.Body

	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		return nil, nil, nil, err
	}

	var pushMessages []string
	typeCount := make(map[string]int)
	activeRepos := make(map[string]int)
	// var starMessages []string

	for _, item := range items {
		typeCount[item.Type]++
		activeRepos[item.RepoField.Name]++

		if len(item.PayloadField.Commits) > 0 {
			commit := len(item.PayloadField.Commits)
			name := item.RepoField.Name

			newPushMessage := fmt.Sprintf("Pushed %d to %v \n", commit, name)
			pushMessages = append(pushMessages, newPushMessage)
		}

	}

	return pushMessages, typeCount, activeRepos, nil
}
