package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/0bvim/octoBuddy/internal/app/model"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOAuthConfig *oauth2.Config

func InitOAuthConfig() {
	githubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		Scopes:       []string{"user:follow", "read:user"},
		Endpoint:     github.Endpoint,
	}
}

// LoginHandler starts the GitHub login process
func LoginHandler(c *gin.Context) {
	// redirect to the github oauth page
	url := githubOAuthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "state" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	token, err := githubOAuthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch user information from GitHub API
	client := githubOAuthConfig.Client(c, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user info: %v", err)})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read response: %v", err)})
		return
	}

	var user model.GitHubUser
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to parse user info: %v", err)})
		return
	}

	// Store user info in session/context
	c.Set("user", user)
	
	// Set access token in cookie
	c.SetCookie("access_token", token.AccessToken, 3600, "/", "", false, true)

	// Redirect to the me page
	c.Redirect(http.StatusTemporaryRedirect, "/me")
}

// GetAccessToken retrieves the access token from context or cookie
func GetAccessToken(c *gin.Context) string {
	// First try to get from context
	if token, exists := c.Get("access_token"); exists {
		return token.(string)
	}

	// Then try to get from Authorization header
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}

	// Then try to get from query parameter
	if token := c.Query("access_token"); token != "" {
		return token
	}

	// Finally try to get from cookie
	token, err := c.Cookie("access_token")
	if err == nil {
		return token
	}

	return ""
}

// LogoutHandler logs the user out
func LogoutHandler(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// GetGitHubUser is a helper function to fetch GitHub user info using an access token
func GetGitHubUser(accessToken string) (*model.GitHubUser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status: %d", resp.StatusCode)
	}

	var user model.GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
