package entity

import "time"

// GitHubUser represents the structure of the GitHub user API response.
type User struct {
	Follower          []Follower
	Followed          []Followed
	Login             string    `json:"login"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	ID                int       `json:"id"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	SiteAdmin         bool      `json:"site_admin"`
	Hireable          bool      `json:"hireable"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Follower struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	ID        int    `json:"id"`
	HTMLURL   string `json:"html_url"`
}

type Followed struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	ID        int    `json:"id"`
	HTMLURL   string `json:"html_url"`
}
