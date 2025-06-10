.PHONY: build run test clean swagger docker docker-run

# Build the application
build:
	go build -o go-api-template .

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f go-api-template
	rm -rf docs/docs.go docs/swagger.json docs/swagger.yaml

# Generate Swagger documentation
swagger:
	swag init

# Build Docker image
docker:
	docker build -t go-api-template .

# Run Docker container
docker-run:
	docker run -p 8080:8080 go-api-template

# All-in-one command for development
dev: swagger build run