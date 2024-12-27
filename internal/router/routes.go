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
		user.GET("/followings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong2",
			})
		})
	}

	group := router.Group("/me")
	{
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
	}
}
