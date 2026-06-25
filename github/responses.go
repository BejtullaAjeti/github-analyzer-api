package github

type ActivityPoint struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type ReposResponse struct {
	TopRepos  []Repository   `json:"top_repos"`
	Languages map[string]int `json:"languages"`
}

type RepoHealth struct {
	LastCommitAt   string `json:"last_commit_at"`
	OpenIssues     int    `json:"open_issues"`
	StarsTrendNote string `json:"stars_trend_note"`
	Stars          int    `json:"stars"`
	Watchers       int    `json:"watchers"`
}

type CommitFrequencyPoint struct {
	WeekStart string `json:"week_start"`
	Count     int    `json:"count"`
}

type ContributorSummary struct {
	Login         string `json:"login"`
	AvatarURL     string `json:"avatar_url"`
	Contributions int    `json:"contributions"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
