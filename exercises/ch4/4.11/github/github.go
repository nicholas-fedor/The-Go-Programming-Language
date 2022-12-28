package github

import "time"

const URL = "https://api.github.com/"

type IssueReadRequest struct {
	Owner       string
	Repo        string
	IssueNumber int
}

type IssueCreateResult struct {
	Owner   string
	Repo    string
	Title   string
	HTMLURL string `json:"html_url"`
}

type IssueReadResult struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	Title     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssueUpdateResult struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	Title     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
