package main

import (
	"fmt"
	"github.com/0bvim/octoBuddy/internal/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

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
	router.Initialize()
}
