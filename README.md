# Golang Boilerplate

A production-ready Go backend boilerplate with enterprise-grade features, optimized for companies like Uber. Built with Gin framework, PostgreSQL, Redis, JWT authentication, and comprehensive deployment configurations.

## Features

### Core Features
- **Framework**: Gin web framework with middleware chain
- **Database**: PostgreSQL with connection pooling and migrations
- **Cache**: Redis with circuit breaker protection
- **Authentication**: JWT-based auth with bcrypt password hashing
- **API Versioning**: v1 API with backward compatibility
- **Rate Limiting**: Per-user rate limiting with configurable limits

### Observability & Monitoring
- **Metrics**: Prometheus metrics collection
- **Tracing**: Distributed tracing with Jaeger
- **Logging**: Structured logging with Zap and request tracing
- **Health Checks**: Comprehensive service health monitoring

### Reliability & Resilience
- **Circuit Breakers**: Fault tolerance for database and Redis connections
- **Retry Logic**: Exponential backoff retry utility
- **Graceful Shutdown**: Proper server shutdown with signal handling
- **Async Processing**: Worker pool for background tasks

### Security
- **CORS**: Configurable CORS middleware
- **Security Headers**: OWASP recommended security headers
- **Input Validation**: Request validation with Gin binding

### Deployment & DevOps
- **Docker**: Multi-stage optimized builds
- **Kubernetes**: Production-ready K8s manifests with health checks
- **Load Balancing**: Nginx-based load balancing with health checks
- **CI/CD**: GitHub Actions with testing and building
- **Configuration**: Environment-based config with Viper

## Tech Stack

- **Go**: 1.24+ with latest features
- **Web Framework**: Gin with extensive middleware
- **Database**: PostgreSQL with golang-migrate
- **Cache**: Redis with go-redis client
- **Logging**: Zap structured logging
- **Config**: Viper with YAML/ENV support
- **Metrics**: Prometheus client
- **Tracing**: Jaeger OpenTracing
- **Testing**: Go testing with testify
- **Container**: Docker with multi-stage builds
- **Orchestration**: Kubernetes with Helm-ready manifests

## Architecture

```
├── cmd/                    # Application entrypoints
├── main/                   # Main application code
│   ├── config/            # Configuration management
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # Gin middlewares (auth, logging, etc.)
│   ├── models/            # Data models
│   ├── repo/              # Repository layer (data access)
│   ├── routes/            # Route definitions
│   └── server/            # HTTP server setup
├── pkg/                   # Shared packages
│   ├── async/            # Async processing utilities
│   ├── metrics/          # Prometheus metrics
│   └── utils/            # Common utilities
├── docker/                # Docker configurations
├── migrations/            # Database migrations
├── scripts/               # Database initialization
└── tests/                 # Integration tests
```

## Quick Start

### Prerequisites
- Go 1.24+
- Docker & Docker Compose
- PostgreSQL (for local development)
- Redis (for local development)

### Development Setup

1. **Install development tools**:
   ```bash
   ./scripts/setup-dev.sh
   ```

2. **Clone and setup**:
   ```bash
   git clone <repository>
   cd golang-boilerplate
   cp config.yaml.example config.yaml
   ```

3. **Start dependencies**:
   ```bash
   make dev-deps
   ```

4. **Run the application**:
   ```bash
   make run
   ```

5. **Run tests**:
   ```bash
   make test
   ```

### Docker Development

```bash
# Start full stack with Docker Compose
make docker-dev

# View logs
make logs

# Stop services
make docker-stop
```

### Production Deployment

```bash
# Build for production
make build

# Deploy with Docker Compose
make deploy

# Or deploy to Kubernetes
kubectl apply -f docker/k8s-deployment.yaml
```

## API Documentation

### API v1 Endpoints

#### Public Endpoints
- `GET /api/v1/health` - Comprehensive health check
- `GET /api/v1/ping` - Simple ping response
- `POST /api/v1/register` - User registration
- `POST /api/v1/login` - User authentication

#### Protected Endpoints
- `GET /api/v1/protected` - Example protected route (requires JWT)

#### Monitoring Endpoints
- `GET /metrics` - Prometheus metrics
- `GET /health` - Service health status

### Authentication

Use JWT tokens obtained from `/api/v1/login`:

```bash
curl -X POST http://localhost/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"pass"}'
```

Use the returned token in subsequent requests:

```bash
curl -H "Authorization: Bearer <token>" \
  http://localhost/api/v1/protected
```

## Configuration

Configuration is managed via `config.yaml` and environment variables:

```yaml
server:
  port: 8080
  host: "0.0.0.0"

database:
  url: "postgres://user:pass@localhost/db"

redis:
  url: "redis://localhost:6379"

jwt:
  secret: "your-secret-key"
```

## Monitoring & Observability

### Metrics
Access Prometheus metrics at `/metrics`

### Tracing
Distributed tracing is configured for Jaeger. Set `JAEGER_AGENT_HOST` for production.

### Logging
Structured logs include request IDs, latency, and user context.

## Development

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run locally
make watch         # Live reload on file changes (like nodemon)
make test          # Run tests
make docker-dev    # Start Docker development environment
make deploy        # Deploy to production
make clean         # Clean build artifacts
make lint          # Run linter
make fmt           # Format code
```

### Live Reloading

For development with automatic restarts on file changes (similar to nodemon):

**Prerequisites:**
```bash
# Install reflex for live reloading
go install github.com/cespare/reflex@latest
```

**Usage:**
```bash
# Using make command (recommended)
make watch

# Or directly with reflex
reflex -s -r '\.go$$' -- go run ./cmd/main.go
```

This will automatically restart your Go application whenever you save changes to `.go` files.

## Deployment

### Docker Compose (Development)
```bash
docker-compose -f docker/docker-compose.yml up -d
```

### Docker Compose (Production)
```bash
docker-compose -f docker/docker-compose.prod.yml up -d
```

### Kubernetes
```bash
kubectl apply -f docker/k8s-deployment.yaml
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.
- `GET /api/protected` - Protected route (requires JWT)

### Authentication

Register a user:
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username": "test", "password": "password"}'
```

Login:
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "test", "password": "password"}'
```

Use the returned token for protected routes:
```bash
curl -H "Authorization: Bearer <token>" http://localhost:8080/api/protected
```

## Deployment

### Docker

```bash
make docker-build
make docker-run
```

### Docker Compose (Production)

```bash
make prod
```

### Kubernetes

```bash
kubectl apply -f docker/k8s-deployment.yaml
```

## Configuration

Edit `config.yaml` or set environment variables:

- `SERVER_PORT` - Server port (default: 8080)
- `SERVER_HOST` - Server host (default: 0.0.0.0)
- `DATABASE_URL` - PostgreSQL connection string
- `REDIS_URL` - Redis connection string
- `JWT_SECRET` - JWT signing secret

## Development

- `make build` - Build the application
- `make test` - Run tests
- `make lint` - Run linter
- `make fmt` - Format code
- `make tidy` - Tidy modules

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── main/
│   ├── common/                 # Shared clients (redis, etc.)
│   ├── config/                 # Configuration management
│   ├── handlers/               # HTTP handlers
│   ├── middleware/             # Gin middlewares
│   ├── models/                 # Data models
│   ├── repo/                   # Data access layer
│   ├── routes/                 # Route definitions
│   ├── server/                 # Server setup
│   └── service/                # External services (DB, Redis)
├── pkg/                        # Shared packages
├── docker/                     # Deployment configurations
│   ├── Dockerfile
│   ├── docker-compose.prod.yml
│   └── k8s-deployment.yaml
├── scripts/                    # Database scripts
├── tests/                      # Integration tests
├── .github/workflows/          # CI/CD
├── config.yaml                 # Configuration file
├── Makefile                    # Build automation
└── README.md
```