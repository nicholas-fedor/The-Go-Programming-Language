package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UpdateIssue updates a GitHub issue.
// https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#update-an-issue
func UpdateIssue() (*IssueUpdateResult, error) {
	// Request Type: Patch
	// URL schema: /repos/{owner}/{repo}/issues/{issue_number}
	// 
	// Body parameters enable updating of the issue.
	resp, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	var result IssueUpdateResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}