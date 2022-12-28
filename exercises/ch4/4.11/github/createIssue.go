package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// CreateIssue is a POST request to create a new GitHub issue.
// https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#create-an-issue
func CreateIssue(request []string) (*IssueCreateResult, error) {
	// POST Request
	// Path Parameters: owner and repo
	// Body Parameters: title

	// URL: URL + owner/{owner}/repo/{repo}
	// Body: {title}
	q := url.QueryEscape(strings.Join(request, " "))
	fmt.Println("request: " + q)
	
	resp, err := http.Post(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	var result IssueCreateResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}
