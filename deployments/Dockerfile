# Use official Golang image as a builder
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project source
COPY . .

# Ensure we are in the correct directory before building
WORKDIR /app/cmd/server

# Force Linux AMD64 Build (x86_64)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/iam-server

# Build the migration binary
WORKDIR /app/cmd/migrate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/migrate

# Use a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Install necessary dependencies
RUN apk --no-cache add ca-certificates curl

# Copy the built binaries from the builder stage
COPY --from=builder /app/iam-server /root/iam-server
COPY --from=builder /app/migrate /root/migrate

# Ensure the binaries are executable
RUN chmod +x /root/iam-server /root/migrate

# ✅ Copy the database migration files into the final container
COPY --from=builder /app/db/migrations /root/migrations

# Expose port
EXPOSE 8080

# Docker Health Check
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:8080/health || exit 1

# Default command to run the IAM server
CMD ["/root/iam-server"]
