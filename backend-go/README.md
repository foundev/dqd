# DQD Go Backend

Go REST API server for Dremio Query Doctor (DQD).

## Overview

This is a modern rewrite of the DQD backend using Go. It provides a REST API for analyzing Dremio query profiles, system metrics, and generating reproduction scripts.

## Project Structure

```
backend-go/
├── cmd/
│   └── server/          # Main application entry point
│       └── main.go
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware (CORS, logging, etc.)
│   ├── models/          # Data models and types
│   └── services/        # Business logic services
├── bin/                 # Compiled binaries (generated)
├── go.mod              # Go module definition
├── go.sum              # Go dependencies checksums
└── Makefile            # Build automation
```

## Prerequisites

- Go 1.21 or higher
- Make (optional, for convenience)

## Quick Start

### 1. Install Dependencies

```bash
go mod download
# or
make install-deps
```

### 2. Run the Server

```bash
# Direct run
go run ./cmd/server/main.go

# Or using Make
make run

# Or build and run binary
make build
./bin/dqd-server
```

The server will start on port `8080` by default.

### 3. Test the API

```bash
# Check version
curl http://localhost:8080/api/about.json

# Health check
curl http://localhost:8080/health
```

## Configuration

Create a `.env` file in the root directory (see `.env.example`):

```env
PORT=8080
LOG_LEVEL=info
```

## API Endpoints

### Implemented

- `GET /api/about.json` - Get version information
- `GET /health` - Health check endpoint

### Coming Soon (Return 501 Not Implemented)

- `POST /api/profile` - Detailed profile analysis
- `POST /api/simple-profile` - Simple profile analysis
- `POST /api/profiles` - Compare two profiles
- `POST /api/queriesjson` - Queries.json analysis
- `POST /api/reproduction` - Generate schema reproduction scripts
- `POST /api/iostat` - IOStat analysis
- `POST /api/ttop` - Thread top analysis

## Development

### Available Make Commands

```bash
make help              # Show all available commands
make install-deps      # Install Go dependencies
make build            # Build the server binary
make run              # Run development server
make test             # Run tests
make test-coverage    # Run tests with coverage report
make clean            # Clean build artifacts
make fmt              # Format code
make lint             # Run linter (requires golangci-lint)
```

### Running Tests

```bash
go test -v ./...
# or
make test
```

### Code Formatting

```bash
go fmt ./...
# or
make fmt
```

## Migration Strategy

This Go backend is being developed in parallel with the existing Java backend. The migration follows this approach:

1. ✅ **Phase 1**: Basic server setup with routing and middleware
2. ✅ **Phase 2**: Simple endpoints (`/about.json`)
3. 🔄 **Phase 3**: File upload handling and processing
4. 📋 **Phase 4**: Profile JSON parsing and analysis
5. 📋 **Phase 5**: Queries JSON analysis
6. 📋 **Phase 6**: Schema generation (repro)
7. 📋 **Phase 7**: System analysis (iostat, top)

## Tech Stack

- **Router**: [Chi](https://github.com/go-chi/chi) - Lightweight, composable HTTP router
- **Standards**: Standard library for most functionality
- **Future**: May add for complex JSON parsing, compression, etc.

## Architecture

The application follows a clean architecture pattern:

- **Handlers**: HTTP request/response handling
- **Services**: Business logic (to be implemented)
- **Models**: Data structures and types
- **Middleware**: Cross-cutting concerns (CORS, logging, auth)

## CORS

CORS is enabled by default to allow frontend development. In production, configure `CORS_ALLOWED_ORIGINS` in your environment.

## Contributing

1. Write clean, idiomatic Go code
2. Add tests for new functionality
3. Run `make fmt` and `make lint` before committing
4. Update documentation as needed

## License

Apache License 2.0 - See LICENSE file for details
