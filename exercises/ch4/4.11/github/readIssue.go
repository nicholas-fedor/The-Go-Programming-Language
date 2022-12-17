package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const URL = ""

// ReadIssue requests a specific issue and
// prints out more detailed information.
func ReadIssue(terms []string) (*IssueReadResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(URL + "?q=" + q)
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
