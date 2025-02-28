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

# Use a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/iam-server /root/iam-server

# Ensure the binary is executable
RUN chmod +x /root/iam-server

# Check binary architecture
RUN ldd /root/iam-server || true

# Expose port
EXPOSE 8080

# Correct CMD
CMD ["/root/iam-server"]

# Docker Health Check
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:8080/health || exit 1

