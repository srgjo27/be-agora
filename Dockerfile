# ----- Step 1: Build -----
FROM golang:1.25.3-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies
RUN go mod download

# Copy the remaining source code
COPY . .

# Build the Go app
# CGO_ENABLED=0 and GOOS=linux can be set static binary
# Change 'username-github' with your actual GitHub username
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /agora-api ./cmd/api/main.go

# ----- Step 2: Run -----
FROM alpine:3.18

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /agora-api /agora-api

# (IMPORTANT) Copy your SQL migration files if you have them
# COPY ./migrations /migrations

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/agora-api"]