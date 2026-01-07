# Go Demo API _test

A minimal REST API in Go for ArgoCD deployment demonstrations.

## Features

- Simple REST API with 3 endpoints
- Environment management with .env files
- Docker support
- CI/CD with GitHub Actions

## Endpoints

- `/` - Returns welcome message for demo API
- `/health` - Returns API health status
- `/docs` - Returns dummy documentation data

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
go mod tidy
```

3. Create a .env file
```bash
cp .env.example .env
```

4. Run the application
```bash
go run main.go
```

## Docker

Build the Docker image:
```bash
docker build -t demo-api .
```

Run the container:
```bash
docker run -p 8080:8080 demo-api
```

## Testing

Test the endpoints:
```bash
curl http://localhost:8080/
curl http://localhost:8080/health
curl http://localhost:8080/docs
```

## ArgoCD Deployment

This API is designed for demonstration purposes with ArgoCD running on OrbStack.
