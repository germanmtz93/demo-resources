FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -s -w' -o api .

# Create minimal image using scratch
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/api /api

# Expose port
EXPOSE 3000

# Run the application
CMD ["/api"]