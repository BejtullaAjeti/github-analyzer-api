package github

import (
	"encoding/json"
	"net/http"
	"time"
)

var githubToken string

const baseURL = "https://api.github.com"

type GitHubError struct {
	StatusCode int
	Message    string
}

func SetToken(token string) {
	githubToken = token
}

func (e *GitHubError) Error() string { return e.Message }

func get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "github-analyzer-app")
	req.Header.Set("Authorization", "Bearer "+githubToken)
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FetchUser(username string) (*User, error) {
	resp, err := get(baseURL + "/users/" + username)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FetchRepos(username string) ([]Repository, error) {
	resp, err := get(baseURL + "/users/" + username + "/repos?per_page=100&sort=pushed")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var repos []Repository
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func FetchEvents(username string) ([]Event, error) {
	resp, err := get(baseURL + "/users/" + username + "/events?per_page=100")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func FetchRepoDetail(owner, repo string) (*RepoDetail, error) {
	resp, err := get(baseURL + "/repos/" + owner + "/" + repo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var detail RepoDetail
	err = json.NewDecoder(resp.Body).Decode(&detail)
	if err != nil {
		return nil, err
	}

	return &detail, nil
}

func FetchCommits(owner, repo string) ([]Commit, error) {
	resp, err := get(baseURL + "/repos/" + owner + "/" + repo + "/commits?per_page=100")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var commits []Commit
	err = json.NewDecoder(resp.Body).Decode(&commits)
	if err != nil {
		return nil, err
	}

	return commits, nil
}

func FetchContributors(owner, repo string) ([]Contributor, error) {
	resp, err := get(baseURL + "/repos/" + owner + "/" + repo + "/contributors?per_page=100")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var contributors []Contributor
	err = json.NewDecoder(resp.Body).Decode(&contributors)
	if err != nil {
		return nil, err
	}

	return contributors, nil
}

func FetchCommitActivity(owner, repo string) ([]WeeklyCommitActivity, error) {
	resp, err := get(baseURL + "/repos/" + owner + "/" + repo + "/stats/commit_activity")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusAccepted {
		return nil, &GitHubError{StatusCode: 202, Message: "stats computing"}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &GitHubError{StatusCode: resp.StatusCode, Message: "github request failed"}
	}
	var commitActivity []WeeklyCommitActivity
	err = json.NewDecoder(resp.Body).Decode(&commitActivity)
	if err != nil {
		return nil, err
	}

	return commitActivity, nil
}
