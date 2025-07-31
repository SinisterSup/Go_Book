// Package githubissues provides a GO API for Github issues tracker.
// See https://docs.github.com/en/rest/issues/issues
package githubissues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const IssuesURL = "https://api.github.com/search/issues"

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
	CreatedAt string `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// // SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch issues: %w", err)
	}

	// We must close resp.Body when we're done with it, on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	resp.Body.Close()
	return &result, nil
}

// // SearchIssues queries the Github issue tracker.
// // Copilot suggested code!!
// func SearchIssues(terms []string) (*IssuesSearchResult, error) {
// 	// Construct the search URL with the provided terms.
// 	url := IssuesURL + "?q=" + url.QueryEscape(strings.Join(terms, " "))
//
// 	// Make the HTTP GET request.
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch issues: %w", err)
// 	}
// 	defer resp.Body.Close()
//
// 	// Check for a successful response.
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("error: %s", resp.Status)
// 	}
//
// 	// Decode the JSON response into IssuesSearchResult.
// 	var result IssuesSearchResult
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return nil, fmt.Errorf("failed to decode JSON: %w", err)
// 	}
//
// 	return &result, nil
// }
