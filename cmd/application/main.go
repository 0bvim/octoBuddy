package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0bvim/octoBuddy/internal/delivery/http/handler"
	"github.com/0bvim/octoBuddy/internal/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	handler.InitOAuthConfig()
}

// TODO: implement logger with context
func initLogger() *log.Logger {
	file, err := os.OpenFile("tmp/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("failed to close log file: %v", err)
		}
	}(file)

	// Configure the logger
	return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Create a context that we'll use to gracefully shutdown our application
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialize the router with our context
	server := router.Initialize(ctx)

	// Wait for shutdown signal
	<-ctx.Done()

	// Create a timeout context for graceful shutdown
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
