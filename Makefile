APP_NAME := octoBuddy
BIN_DIR := bin

.PHONY: all build run clean

all: clean build run

build:
	go build -o $(BIN_DIR)/$(APP_NAME) cmd/application/main.go

run:
	@./bin/$(APP_NAME)

clean:
	@rm -f $(BIN_DIR)/$(APP_NAME)
