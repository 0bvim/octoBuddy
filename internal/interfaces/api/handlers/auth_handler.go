package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/0bvim/octoBuddy/internal/application/dto"
	"github.com/0bvim/octoBuddy/internal/application/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) GithubLogin(c *gin.Context) {
	authURL := h.authService.GetAuthURL()
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (h *AuthHandler) GithubCallback(c *gin.Context) {
	code := c.Query("code")
	tokenPair, user, err := h.authService.HandleCallback(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Construct the redirect URL with query parameters
	redirectURL := fmt.Sprintf(
		os.Getenv("REDIRECT_URL"),
		url.QueryEscape(tokenPair.AccessToken),
		url.QueryEscape(tokenPair.RefreshToken),
		url.QueryEscape(strconv.Itoa(user.ID)),
		url.QueryEscape(user.Login),
		url.QueryEscape(user.Name),
		url.QueryEscape(user.Email),
		url.QueryEscape(user.AvatarURL),
		url.QueryEscape(strconv.Itoa(user.Followers)),
		url.QueryEscape(strconv.Itoa(user.Following)),
		url.QueryEscape(user.Email),
	)

	// Redirect to the new URL
	c.Redirect(http.StatusFound, redirectURL)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	tokenPair, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, dto.TokenResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
