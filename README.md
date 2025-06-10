# Go API Template

A standard template for creating REST APIs in Go.

## Features

- Standard REST API structure
- Environment management with .env files
- API documentation with Swagger
- Docker support
- CI/CD with GitHub Actions
- Test setup

## Endpoints

- `/` - Returns the hostname of the machine
- `/health` - Returns the API health status
- `/swagger/*any` - API documentation

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker (optional)

### Installation

1. Clone the repository
```bash
git clone https://github.com/organization/go-api-template.git
cd go-api-template
```

2. Install dependencies
```bash
go mod download
```

3. Create a .env file
```bash
cp .env.example .env
```

4. Generate Swagger documentation
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

5. Run the application
```bash
go run main.go
```

## Docker

Build the Docker image:
```bash
docker build -t go-api-template .
```

Run the container with environment variables:
```bash
docker run -p 8080:8080 --env-file .env go-api-template
```

## Testing

Run tests:
```bash
go test -v ./...
```

## CI/CD

The repository includes GitHub Actions workflows for:
- Running tests
- Building and pushing Docker images to Docker Hub

Required secrets:
- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN`

### Reusing the workflow

This workflow can be reused in other repositories by calling it with a custom service name:

```yaml
jobs:
  build:
    uses: organization/go-api-template/.github/workflows/ci-cd.yml@main
    with:
      service_name: my-custom-service
    secrets:
      DOCKERHUB_USERNAME: ${{ secrets.DOCKERHUB_USERNAME }}
      DOCKERHUB_TOKEN: ${{ secrets.DOCKERHUB_TOKEN }}
```