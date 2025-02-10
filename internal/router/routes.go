package router

import (
	"github.com/0bvim/octoBuddy/internal/delivery/http/handler"
	"github.com/0bvim/octoBuddy/internal/router/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Protected routes
	protected := router.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handler.ProfileHandler)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Home page",
		})
	})

	router.GET("/login", handler.LoginHandler)
	router.GET("/callback", handler.CallbackHandler)
	router.GET("/logout", handler.LogoutHandler)

	user := router.Group("/:user") // ':name' notation are to indicates a variable
	{
		user.GET("/", handler.GetUser)
		user.GET("/followers", handler.GetUserFollowers)
		user.GET("/following", handler.GetUserFollowing)
		user.GET("/status", handler.GetStatus)
	}

	group := router.Group("/me")
	{
		group.GET("/", handler.GetMe)
		group.GET("/followers", handler.ListFollowers)
		group.GET("/following", handler.ListFollowing)
		group.GET("/allow-list", handler.GetAllowList)
		group.POST("/allow-list", handler.AddAllowList)
		group.DELETE("/allow-list", handler.DeleteAllowList)
		group.GET("/deny-list", handler.GetDenyList)
		group.POST("/deny-list", handler.AddDenyList)
		group.DELETE("/deny-list", handler.DeleteDenyList)
	}
}
