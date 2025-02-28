
# Students API

A Go-based API for managing student data.

## Overview
This project provides a basic API framework for managing student information. It includes configuration management, HTTP server setup, and environment handling.

## Prerequisites
- Go 1.23 or higher
- Basic understanding of Go programming and HTTP servers

## Getting Started

1. Clone the repository:
```bash
git clone github.com/Piyu-Pika/students-api.git
```

2. Install dependencies:
```bash
go mod tidy
```

3. Copy and update the configuration file:
```bash
cp config/local.yaml config/your_env.yaml
```
Modify the values as needed in your YAML file.

4. Run the server:
```bash
go run cmd/students-api/main.go
```

## Configuration
The application uses YAML configuration files. The default configuration is loaded from `config/local.yaml` if no other config file is specified via the `CONFIG` environment variable.

### Example Configuration (config/local.yaml):
```yaml
env: "local"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8080"
  read_timeout: 10000000000
  write_timeout: 10000000000
  # Other server settings...
```

## Environment Variables
- `CONFIG`: Path to the YAML configuration file
- `ENV`: Environment name (Defaults to "prod" if not specified)

## Project Structure
```
.
├── cmd/                    # Command executables
│   └── students-api/      # Students API main package
│       └── main.go        # Main entry point
├── config/                 # Configuration files
│   └── local.yaml         # Local environment configuration
└── internal/               # Internal packages
    └── config/            # Configuration handling
        └── config.go      # Configuration structures and loading
```

## Future Work
- Add database integration
- Implement proper route handlers
- Add logging middleware
- Include authentication
- Add unit tests
- Implement proper error handling
- Add more comprehensive configuration validation
- Add Docker support for easier deployment

This is a basic structure that you can expand upon as you add more features to your API.
