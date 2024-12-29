package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	user := router.Group("/:user") // ':name' notation are to indicates a variable
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "General user status like followers and following and so forth",
			})
		})
		user.GET("/followers", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong1",
			})
		})
		user.GET("/following", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong2",
			})
		})
		user.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong3",
			})
		})
	}

	group := router.Group("/me")
	{
		group.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "General information about logged user",
			})
		})
		group.GET("/followers", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong4",
			})
		})
		group.GET("/following", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong5",
			})
		})
		group.GET("/allow-list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "list of users that are allowed",
			})
		})
		group.POST("/allow-list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "include of users that are allowed",
			})
		})
		group.GET("/deny-list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get list of users that are denied",
			})
		})
		group.POST("/deny-list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Post Deny List",
			})
		})
	}
}
