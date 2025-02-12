package routes

import (
	"github.com/0bvim/octoBuddy/internal/interfaces/api/handlers"
	"github.com/0bvim/octoBuddy/internal/interfaces/api/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine     *gin.Engine
	auth       *handlers.AuthHandler
	user       *handlers.UserHandler
	authMiddle *middleware.AuthMiddleware
	security   *middleware.SecurityMiddleware
}

func NewRouter(
	engine *gin.Engine,
	auth *handlers.AuthHandler,
	user *handlers.UserHandler,
	authMiddle *middleware.AuthMiddleware,
) *Router {
	return &Router{
		engine:     engine,
		auth:       auth,
		user:       user,
		authMiddle: authMiddle,
		security:   middleware.NewSecurityMiddleware(),
	}
}

func (r *Router) Setup() {
	// Apply CORS middleware to all routes
	r.engine.Use(r.security.CORS())

	// Public routes
	r.engine.GET("/auth/github", r.auth.GithubLogin)
	r.engine.GET("/callback", r.auth.GithubCallback)
	r.engine.POST("/auth/refresh", r.auth.RefreshToken)

	// Protected routes
	api := r.engine.Group("/api")
	api.Use(r.authMiddle.AuthRequired())
	{
		api.GET("/user", r.user.GetUser)
	}
}
