package internal

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestConsumeBody(t *testing.T) {
	// Save original args and restore them after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Setup test cases
	tests := []struct {
		name           string
		username       string
		mockResponse   string
		expectedStatus int
		wantPush       []string
		wantTypes      map[string]int
		wantRepos      map[string]int
		wantErr        bool
	}{
		{
			name:     "successful response",
			username: "testuser",
			mockResponse: `[
				{
					"id": "45492759947",
					"type": "CreateEvent",
					"actor": {
					"id": 61546827,
					"login": "Natnael-Alemayehu",
					"display_login": "Natnael-Alemayehu",
					"gravatar_id": "",
					"url": "https://api.github.com/users/Natnael-Alemayehu",
					"avatar_url": "https://avatars.githubusercontent.com/u/61546827?"
					},
					"repo": {
					"id": 915949224,
					"name": "Natnael-Alemayehu/github-activity-cli",
					"url": "https://api.github.com/repos/Natnael-Alemayehu/github-activity-cli"
					},
					"payload": {
					"ref": null,
					"ref_type": "repository",
					"master_branch": "main",
					"description": null,
					"pusher_type": "user"
					},
					"public": true,
					"created_at": "2025-01-13T06:56:17Z"
				},
				{
					"id": "45492759888",
					"type": "CreateEvent",
					"actor": {
					"id": 61546827,
					"login": "Natnael-Alemayehu",
					"display_login": "Natnael-Alemayehu",
					"gravatar_id": "",
					"url": "https://api.github.com/users/Natnael-Alemayehu",
					"avatar_url": "https://avatars.githubusercontent.com/u/61546827?"
					},
					"repo": {
					"id": 915949224,
					"name": "Natnael-Alemayehu/github-activity-cli",
					"url": "https://api.github.com/repos/Natnael-Alemayehu/github-activity-cli"
					},
					"payload": {
					"ref": "main",
					"ref_type": "branch",
					"master_branch": "main",
					"description": null,
					"pusher_type": "user"
					},
					"public": true,
					"created_at": "2025-01-13T06:56:17Z"
				}]`,
			expectedStatus: http.StatusOK,
			wantPush:       []string{"Pushed 2 to testuser/repo1 \n"},
			wantTypes:      map[string]int{"PushEvent": 1, "WatchEvent": 1},
			wantRepos:      map[string]int{"testuser/repo1": 1, "testuser/repo2": 1},
			wantErr:        false,
		},
		{
			name:           "HTTP error",
			username:       "testuser",
			mockResponse:   `{}`,
			expectedStatus: http.StatusNotFound,
			wantPush:       nil,
			wantTypes:      nil,
			wantRepos:      nil,
			wantErr:        true,
		},
		{
			name:     "invalid JSON response",
			username: "testuser",
			mockResponse: `{
				invalid json
			}`,
			expectedStatus: http.StatusOK,
			wantPush:       nil,
			wantTypes:      nil,
			wantRepos:      nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server and patch the default HTTP client
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.expectedStatus)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			// Store the original HTTP client and restore it after the test
			originalClient := http.DefaultClient
			defer func() { http.DefaultClient = originalClient }()

			// Set command line argument
			os.Args = []string{"cmd", tt.username}

			// Call ConsumeBody
			gotPush, gotTypes, gotRepos, err := ConsumeBody()

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("ConsumeBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If expecting error, no need to check other values
			if tt.wantErr {
				return
			}

			// Compare results
			if !reflect.DeepEqual(gotPush, tt.wantPush) {
				t.Errorf("ConsumeBody() gotPush = %v, want %v", gotPush, tt.wantPush)
			}

			if !reflect.DeepEqual(gotTypes, tt.wantTypes) {
				t.Errorf("ConsumeBody() gotTypes = %v, want %v", gotTypes, tt.wantTypes)
			}

			if !reflect.DeepEqual(gotRepos, tt.wantRepos) {
				t.Errorf("ConsumeBody() gotRepos = %v, want %v", gotRepos, tt.wantRepos)
			}
		})
	}
}

func TestReadUsername(t *testing.T) {
	// Save original args and restore them after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct {
		name      string
		args      []string
		want      string
		wantPanic bool
	}{
		{
			name:      "valid username",
			args:      []string{"cmd", "testuser"},
			want:      "testuser",
			wantPanic: false,
		},
		{
			name:      "missing username",
			args:      []string{"cmd"},
			want:      "",
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("ReadUsername() should have panicked")
					}
				}()
			}

			if got := ReadUsername(); got != tt.want {
				t.Errorf("ReadUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
