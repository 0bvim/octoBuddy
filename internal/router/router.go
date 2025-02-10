package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func Initialize(ctx context.Context) *Server {
	router := gin.Default()

	// TODO: setup trusted proxies
	initializeRoutes(router)

	srv := &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}

	go func() {
		if err := srv.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("error starting server: %v\n", err)
		}
	}()

	// Listen for the context cancellation (shutdown signal)
	go func() {
		<-ctx.Done()
		if err := srv.httpServer.Shutdown(context.Background()); err != nil {
			fmt.Printf("error shutting down server: %v\n", err)
		}
	}()

	return srv
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
