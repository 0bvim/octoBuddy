package router

import (
	"fmt"
	"github.com/0bvim/octoBuddy/configs"
	"github.com/gin-gonic/gin"
	"os"
)

func Initialize() {
	// Setup OAuth provider TODO: fix the login with github
	configs.SetupOAuth(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "http://localhost:8080/auth/github/callback")

	router := gin.Default()

	// TODO: setup trusted proxies
	initializeRoutes(router)

	err := router.Run(":8042")
	if err != nil {
		fmt.Printf("error starting server: %v\n", err)
		return
	}
}
