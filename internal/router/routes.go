package router

import (
	"github.com/0bvim/octoBuddy/internal/delivery/http/handler"
	"github.com/0bvim/octoBuddy/internal/router/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Home page",
		})
	})

	router.GET("/login", handler.LoginHandler)
	router.GET("/callback", handler.CallbackHandler)
	router.GET("/logout", handler.LogoutHandler)

	// Protected user routes
	user := router.Group("/:user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", handler.GetUser)
		user.GET("/followers", handler.GetUserFollowers)
		user.GET("/following", handler.GetUserFollowing)
		user.GET("/status", handler.GetStatus)
	}

	// Protected me routes
	me := router.Group("/me")
	me.Use(middleware.AuthMiddleware())
	{
		me.GET("/", handler.GetMe)
		me.GET("/followers", handler.ListFollowers)
		me.GET("/following", handler.ListFollowing)
		me.GET("/allow-list", handler.GetAllowList)
		me.POST("/allow-list", handler.AddAllowList)
		me.DELETE("/allow-list", handler.DeleteAllowList)
		me.GET("/deny-list", handler.GetDenyList)
		me.POST("/deny-list", handler.AddDenyList)
		me.DELETE("/deny-list", handler.DeleteDenyList)
	}
}
