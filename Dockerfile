# Use Go 1.24 base image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum .env ./

# Download all dependencies
RUN go mod download

# Copy the source code excluding web directory
COPY api ./api
COPY cmd ./cmd
COPY config ./config
COPY docs ./docs
COPY internal ./internal
COPY migrations ./migrations
COPY pkg ./pkg
COPY test ./test

# Build the application
RUN go build -o bin/server ./cmd/api/main.go

# Expose port (adjust if needed)
EXPOSE 8080

# Command to run the executable
CMD ["./bin/server"]
