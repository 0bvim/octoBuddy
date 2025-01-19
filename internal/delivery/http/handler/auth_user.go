package handler

import (
	"github.com/gin-gonic/gin"
)

// ProfileHandler returns the authenticated user's profile
func ProfileHandler(c *gin.Context) {
	user := c.MustGet("user") // Get user info from context
	c.JSON(200, gin.H{
		"message": "User profile",
		"user":    user,
	})
}
