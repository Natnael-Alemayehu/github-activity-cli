package data

import "time"

type Body struct {
	Id           string    `json:"id"`
	Type         string    `json:"type"`
	ActorField   Actor     `json:"actor"`
	RepoField    Repo      `json:"repo"`
	PayloadField Payload   `json:"payload"`
	Public       bool      `json:"public"`
	CreatedAt    time.Time `json:"created_time"`
}

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	Ref          string        `json:"ref"`
	RefType      string        `json:"ref_type"`
	MasterBranch string        `json:"master_branch"`
	Description  string        `json:"description"`
	Pushertype   string        `json:"pusher_type"`
	Commits      []CommitField `json:"commits"`
	Action       string        `json:"action"`
}

type CommitField struct {
	Sha      string      `json:"sha"`
	Author   AuthorField `json:"author"`
	Message  string      `json:"message"`
	Distinct bool        `json:"distinct"`
}

type AuthorField struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
