package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadIssue requests a specific issue and
// prints out more detailed information.
// https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#get-an-issue
func ReadIssue(request string) (*IssueReadResult, error) {
	// Request /repos/{owner}/{repo}/issues/{issue_number}
	resp, err := http.Get(URL + "repos/" + request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	var result IssueReadResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}
