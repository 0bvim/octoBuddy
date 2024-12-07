package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	godotenv.Load()
}

func initLogger() *log.Logger {
	file, err := os.OpenFile("tmp/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Configure the logger
	return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := initLogger()

	println(ctx)
	println(logger)
	// connect db

	// run migrations

	// seed if exists

	// redis or mongo

	// health  check

}
