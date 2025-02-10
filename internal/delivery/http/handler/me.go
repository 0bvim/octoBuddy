package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	accessToken := GetAccessToken(c)
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No access token provided"})
		return
	}

	user, err := GetGitHubUser(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user info: %v", err)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListFollowers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on list followers",
	})
}

func ListFollowing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on list following.",
	})
}

func GetAllowList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on get allow list",
	})
}

func AddAllowList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on add allow list",
	})
}

func DeleteAllowList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on delete allow list",
	})
}

func GetDenyList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on get deny list",
	})
}

func AddDenyList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on add deny list",
	})
}

func DeleteDenyList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success on delete deny list",
	})
}
