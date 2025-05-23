package dto

type UserResponse struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}
