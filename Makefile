APP_NAME := octoBuddy
BIN_DIR := bin

all: clean build run

build:
	go build -o $(BIN_DIR)/$(APP_NAME) cmd/api/main.go

run:
	@./bin/$(APP_NAME)

clean:
	@rm -f $(BIN_DIR)/$(APP_NAME)

.PHONY: all build run clean
