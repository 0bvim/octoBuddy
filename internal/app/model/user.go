package model

type User struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}
