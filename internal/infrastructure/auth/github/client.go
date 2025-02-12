package github

import (
	"context"
	"encoding/json"

	"github.com/0bvim/octoBuddy/config"
	"github.com/0bvim/octoBuddy/internal/domain/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Client struct {
	config      *config.Config
	oauthConfig *oauth2.Config
}

func NewGithubClient(config *config.Config) *Client {
	oauthConfig := &oauth2.Config{
		ClientID:     config.GithubClientID,
		ClientSecret: config.GithubClientSecret,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"user:follow",
			"read:user",
		},
		Endpoint: github.Endpoint,
	}

	return &Client{
		config:      config,
		oauthConfig: oauthConfig,
	}
}

func (c *Client) ExchangeCode(code string) (*oauth2.Token, error) {
	return c.oauthConfig.Exchange(context.Background(), code)
}

func (c *Client) GetUserInfo(token *oauth2.Token) (*entity.User, error) {
	client := c.oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user entity.User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Client) GetAuthCodeURL(state string) string {
	return c.oauthConfig.AuthCodeURL(state)
}
