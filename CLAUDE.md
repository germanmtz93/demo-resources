# CLAUDE.md

This file documents the codebase structure, conventions, and development workflows for AI assistants working in this repository.

## Project Overview

This is a minimal Go REST API used as a demonstration target for ArgoCD and Kubernetes deployments. It is intentionally simple — three endpoints, no database, no auth — to keep the focus on GitOps and container delivery workflows.

- **Module**: `github.com/organization/go-api-template`
- **Go version**: 1.21
- **Web framework**: [Gin](https://github.com/gin-gonic/gin) v1.9.1
- **Default port**: 3000 (overridden by `PORT` env var)

## Repository Layout

```
.
├── main.go                    # Entry point: loads env, starts Gin server
├── config/
│   └── env.go                 # godotenv loader; reads .env and .env.{GO_ENV}
├── router/
│   └── router.go              # Route registration (3 GET routes)
├── handlers/
│   └── handlers.go            # HTTP handlers: Welcome, HealthCheck, GetDocs
├── tests/
│   ├── handlers_test.go       # Unit tests for handlers (see known issues below)
│   └── router_test.go         # Integration tests for router setup
├── k8s/
│   ├── base/                  # Kustomize base manifests (Deployment + Service)
│   └── overlays/dev/          # Dev overlay: 1 replica, dev namespace
├── argocd/
│   ├── application.yaml       # ArgoCD Application CR (targets k8s/overlays/dev)
│   ├── argocd-rbac-cm.yaml    # RBAC: demouser gets read-only role
│   ├── integration-values.yml # Port.io integration for ArgoCD event listener
│   └── new-user.yaml          # Demo user ConfigMap
├── .github/workflows/
│   └── ci-cd.yml              # GitHub Actions: build + push Docker image on dev push
├── Dockerfile                 # Multi-stage build → scratch final image
├── docker-compose.yml         # Local dev container (port 3000)
├── Makefile                   # Common development tasks
├── deploy.sh                  # Build image, apply ArgoCD app, wait for sync
├── test-deploy.sh             # Build image, apply Kustomize, wait for pod readiness
├── cleanup.sh                 # Delete ArgoCD app and namespace
├── go.mod / go.sum            # Go module files
└── .env.example               # Template for local environment variables
```

## API Endpoints

| Method | Path      | Handler      | Response                                    |
|--------|-----------|--------------|---------------------------------------------|
| GET    | `/`       | `Welcome`    | `{"message": "Welcome to Demo API", "version": "2.0.0"}` |
| GET    | `/health` | `HealthCheck`| `{"status": "healthy", "service": "demo-api v2"}` |
| GET    | `/docs`   | `GetDocs`    | `{"api": "Demo API", "endpoints": [...], "description": "..."}` |

## Development Workflows

### Local Setup

```bash
cp .env.example .env
go mod tidy
go run main.go          # server starts on :3000
```

### Common Make Targets

| Target       | What it does                                       |
|--------------|----------------------------------------------------|
| `make build` | Compiles binary to `./go-api-template`             |
| `make run`   | `go run main.go`                                   |
| `make test`  | `go test -v ./...`                                 |
| `make clean` | Removes binary and generated Swagger docs          |
| `make swagger`| Runs `swag init` to generate Swagger docs         |
| `make docker`| Builds Docker image tagged `go-api-template`       |
| `make docker-run` | Runs container, maps host 8080 → container 8080 |
| `make dev`   | `swagger` → `build` → `run` (all-in-one)          |

### Docker

```bash
docker build -t demo-api .
docker run -p 3000:3000 demo-api
```

The multi-stage `Dockerfile` compiles with `CGO_ENABLED=0` and uses a `scratch` final image — the resulting binary has no OS dependencies.

### Environment Variables

Defined in `.env` (copy from `.env.example`):

| Variable      | Default       | Purpose                        |
|---------------|---------------|--------------------------------|
| `PORT`        | `3000`        | HTTP listen port               |
| `GO_ENV`      | `development` | Controls which .env file loads |
| `API_VERSION` | `v1`          | Informational version tag      |

`config.LoadEnv()` loads `.env` and, if `GO_ENV` is set, also loads `.env.{GO_ENV}` (latter takes precedence).

## Testing

Run tests:
```bash
make test
# or
go test -v ./...
```

Tests live in `tests/` and use `github.com/stretchr/testify/assert` with Gin's `TestMode`.

### Known Test/Implementation Mismatches

The current test files reference types and handlers that **do not exist** in `handlers/handlers.go`:

- `handlers_test.go` references `handlers.HealthResponse`, `handlers.HostnameResponse`, and `handlers.GetHostname`
- `router_test.go` tests a `/swagger/index.html` endpoint that is not registered

Running `go test -v ./...` will fail until these are resolved. When adding new handlers, define exported response structs in `handlers/handlers.go` and register routes in `router/router.go`.

## CI/CD Pipeline

**File**: `.github/workflows/ci-cd.yml`

Triggers:
- **Push to `dev` branch** → runs `build-and-push` job
- Pull requests to `dev` → workflow runs but build job is skipped (condition: `github.event_name == 'push'`)
- Can also be called as a reusable workflow (`workflow_call`) with an optional `service_name` input

The job:
1. Checks out code
2. Sets up Docker Buildx
3. Logs in to DockerHub using `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` secrets
4. Builds and pushes multi-platform images (`linux/amd64`, `linux/arm64`)
5. Tags: `latest` and the commit SHA
6. Uses registry-based layer caching

**Development branch is `dev`** — merging to `dev` triggers a new Docker image build and push.

## Kubernetes & ArgoCD

### Kustomize Structure

```
k8s/base/            → 2 replicas, ClusterIP service (port 80 → 3000), liveness/readiness probes
k8s/overlays/dev/    → overrides to 1 replica, dev namespace, development image
```

Apply manually:
```bash
kubectl apply -k k8s/overlays/dev/
```

### ArgoCD

`argocd/application.yaml` defines an ArgoCD `Application` CR pointing at `k8s/overlays/dev` with **auto-sync enabled**. Any push to the tracked branch that modifies manifests triggers an automatic sync.

Helper scripts:
```bash
./deploy.sh        # Build image → apply ArgoCD app → wait for sync
./test-deploy.sh   # Build image → apply Kustomize → wait for pod ready
./cleanup.sh       # Delete ArgoCD app and namespace
```

## Code Conventions

- **Package layout**: one package per directory; package names match the directory name (`config`, `router`, `handlers`, `tests`)
- **Handler signature**: always `func Name(c *gin.Context)` — no middleware wrapping
- **JSON responses**: use `gin.H{}` for ad-hoc maps; define exported structs in `handlers/` for structured responses that tests need to deserialize
- **No global state**: configuration is loaded once in `main()` and passed via environment variables
- **Error handling**: fatal errors in `main()` use `log.Fatalf`; handlers return appropriate HTTP status codes via `c.JSON`
- **No comments on obvious code** — only add comments when the reason is non-obvious

## Adding a New Endpoint

1. Add a handler function to `handlers/handlers.go`
2. If the handler returns a structured response, define an exported struct in the same file so tests can deserialize it
3. Register the route in `router/router.go` inside `SetupRouter()`
4. Add a test in `tests/handlers_test.go` and/or `tests/router_test.go`
5. Update the endpoint list in `handlers.GetDocs` if relevant
