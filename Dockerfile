FROM golang:1.21-alpine AS builder

# Install git and certificates for downloads
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies with retry logic
RUN go env -w GOPROXY=direct && \
    for i in $(seq 1 3); do go mod download && break || sleep 5; done

# Copy source code
COPY . .

# Generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

# Create a minimal image
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/api .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./api"]