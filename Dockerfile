# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

# Install necessary tools
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o app ./cmd/server

# Stage 2: Minimal image for running
FROM alpine:3.18

# Install CA certificates (needed if your app calls HTTPS endpoints)
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the compiled binary from builder
COPY --from=builder /app/app .

COPY .env .


# Expose port (change to your app’s port)
EXPOSE 8080

# Run the binary
CMD ["./app"]
