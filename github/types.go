package github

type User struct {
	Login       string  `json:"login"`
	Name        *string `json:"name,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	AvatarURL   string  `json:"avatar_url"`
	HTMLURL     string  `json:"html_url"`
	Followers   int     `json:"followers"`
	Following   int     `json:"following"`
	PublicRepos int     `json:"public_repos"`
	Location    *string `json:"location,omitempty"`
	Blog        *string `json:"blog,omitempty"`
	Company     *string `json:"company,omitempty"`
	CreatedAt   string  `json:"created_at"`
}

type Repository struct {
	Name            string  `json:"name"`
	HTMLURL         string  `json:"html_url"`
	Description     *string `json:"description,omitempty"`
	StargazersCount int     `json:"stargazers_count"`
	ForksCount      int     `json:"forks_count"`
	Language        *string `json:"language,omitempty"`
	Fork            bool    `json:"fork"`
}

type Event struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

type RepoDetail struct {
	Name            string  `json:"name"`
	FullName        string  `json:"full_name"`
	HTMLURL         string  `json:"html_url"`
	Description     *string `json:"description"`
	StargazersCount int     `json:"stargazers_count"`
	ForksCount      int     `json:"forks_count"`
	OpenIssuesCount int     `json:"open_issues_count"`
	WatchersCount   int     `json:"watchers_count"`
	Language        *string `json:"language"`
	PushedAt        string  `json:"pushed_at"`
	CreatedAt       string  `json:"created_at"`
}

type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Name string `json:"name"`
			Date string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
}

type Contributor struct {
	Login         string `json:"login"`
	AvatarURL     string `json:"avatar_url"`
	Contributions int    `json:"contributions"`
}

type WeeklyCommitActivity struct {
	Total int    `json:"total"`
	Week  int64  `json:"week"`
	Days  [7]int `json:"days"`
}
