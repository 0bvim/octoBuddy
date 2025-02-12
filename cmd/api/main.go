package main

import (
	"log"

	"github.com/0bvim/octoBuddy/config"
	"github.com/0bvim/octoBuddy/internal/application/service"
	"github.com/0bvim/octoBuddy/internal/infrastructure/auth/github"
	"github.com/0bvim/octoBuddy/internal/infrastructure/auth/jwt"
	"github.com/0bvim/octoBuddy/internal/infrastructure/persistence/memory"
	"github.com/0bvim/octoBuddy/internal/interfaces/api/handlers"
	"github.com/0bvim/octoBuddy/internal/interfaces/api/middleware"
	"github.com/0bvim/octoBuddy/internal/interfaces/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Initialize infrastructure
	githubClient := github.NewGithubClient(config)
	tokenService := jwt.NewTokenService(config.JWTSecret)
	userRepo := memory.NewUserRepository()

	// Initialize services
	authService := service.NewAuthService(githubClient, tokenService, userRepo)

	// Initialize HTTP server and middleware
	engine := gin.Default()
	authMiddleware := middleware.NewAuthMiddleware(tokenService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler()

	// Setup routes
	router := routes.NewRouter(engine, authHandler, userHandler, authMiddleware)
	router.Setup()

	// Start server
	log.Fatal(engine.Run(":8080"))
}
