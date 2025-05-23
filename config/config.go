package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const (
	FOLLOWING_URL = "https://api.github.com/users/%s/following?per_page=100"
	FOLLOWERS_URL = "https://api.github.com/users/%s/followers?per_page=100"
)

type Config struct {
	GithubClientID     string
	GithubClientSecret string
	GithubRedirectURL  string
	JWTSecret          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		GithubRedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
	}

	// Validate required environment variables
	if config.GithubClientID == "" {
		return nil, errors.New("GITHUB_CLIENT_ID is required")
	}
	if config.GithubClientSecret == "" {
		return nil, errors.New("GITHUB_CLIENT_SECRET is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return config, nil
}
