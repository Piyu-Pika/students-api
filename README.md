

# 🎓 Students API

A powerful Go-based API for managing student data with elegance and efficiency. 🚀

## 📋 Overview
This project provides a robust and scalable API framework for managing student information. It includes advanced configuration management, high-performance HTTP server setup, and flexible environment handling for different deployment scenarios. The API currently supports basic CRUD operations for student records.

## 🛡️ Features
- RESTful API endpoints for student management
- SQLite database integration
- Structured configuration management
- Custom error handling
- HTTP server with configurable timeouts
- Environment-specific configurations

## 🔧 Prerequisites
- Go 1.23 or higher installed on your system
- Basic understanding of Go programming and RESTful APIs
- Text editor or IDE (VS Code recommended)
- SQLite installed on your system

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone github.com/Piyu-Pika/students-api.git
cd students-api
```

2. Install dependencies:
```bash
go mod tidy
go mod verify
```

3. Set up configuration:
```bash
cp config/local.yaml config/your_env.yaml
```
💡 Customize the configuration values in your YAML file based on your environment needs.

4. Create storage directory and initialize database (if not exists):
```bash
mkdir -p storage
touch storage/storage.db
```

5. Launch the server:
```bash
go run cmd/students-api/main.go
```

## ⚙️ Configuration
The application uses YAML configuration files for flexible setup. The default configuration loads from `config/local.yaml` unless specified otherwise via the `CONFIG` environment variable.

### 📝 Example Configuration (config/local.yaml):
```yaml
env: "local"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8080"
  read_timeout: 10s
  write_timeout: 10s
  max_header_size: 1048576
  idle_timeout: 30s
```

### Configuration Parameters
| Parameter            | Description                              | Default Value       |
|----------------------|----------------------------------------|---------------------|
| `env`                | Environment name                       | "local"             |
| `storage_path`       | Path to SQLite database file          | "storage/storage.db"|
| `http_server.address`| Server address to listen on           | "localhost:8080"   |
| `http_server.read_timeout` | HTTP read timeout                  | 10 seconds          |
| `http_server.write_timeout` | HTTP write timeout                | 10 seconds          |
| `http_server.max_header_size` | Maximum HTTP header size          | 1MB                |
| `http_server.idle_timeout`  | HTTP idle connection timeout      | 30 seconds          |

## 🌍 Environment Variables
The following environment variables can be used to configure the application:

### General
- `CONFIG`: Path to YAML configuration file
- `ENV`: Environment name (defaults to "prod" if not specified)

### Database
- `STORAGE_PATH`: Path to SQLite database file

### Server
- `ADDR`: Override server address
- `READ_TIMEOUT`: HTTP read timeout
- `WRITE_TIMEOUT`: HTTP write timeout
- `MAX_HEADER_SIZE`: Maximum HTTP header size
- `IDLE_TIMEOUT`: HTTP idle timeout

## 📁 Project Structure
```
.
├── 📂 cmd/                    # Command executables
│   └── 📂 students-api/      # Main application package
│       └── 📄 main.go        # Entry point
├── 📂 config/                # Configuration directory
│   └── 📄 local.yaml         # Local environment config
├── 📂 internal/              # Internal packages
│   ├── 📂 config/           # Configuration handling
│   │   └── 📄 config.go     # Config structures
│   ├── 📂 models/            # Data models
│   │   └── 📄 student.go     # Student data model
│   ├── 📂 handlers/          # HTTP handlers
│   │   └── 📄 student.go     # Student handler
│   ├── 📂 storage/           # Storage implementations
│   │   └── 📂 sqlite/       # SQLite storage
│   │       └── 📄 sqlite.go # SQLite implementation
│   └── 📂 utils/             # Utility functions
│       └── 📂 response/      # HTTP response utilities
│           └── 📄 response.go # Response helper functions
```

## 🌐 API Documentation

### Endpoints

#### 1. Add Student
`POST /api/students`

Example request body:
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "1234567890",
    "age": 20
}
```

#### 2. Get Student by ID
`GET /api/students/{id}`

Example response:
```json
{
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "1234567890",
    "age": 20
}
```

#### 3. List All Students
`GET /api/students`

Example response:
```json
[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "phone": "1234567890",
        "age": 20
    }
]
```

## 🧮 Error Handling

The API returns JSON-formatted error responses with a consistent structure:

```json
{
    "status": "error",
    "message": "Error description here"
}
```

### Error Response Example
```json
{
    "status": "error",
    "message": "student with id 1 not found"
}
```

## 🛡️ Security Considerations

- Input validation is performed using the [go-playground/validator](https://github.com/go-playground/validator) package
- Request processing includes proper error handling and validation
- Configuration validation is performed at startup
- Secure HTTP headers are configured by default

## 🎯 Future Enhancements
- 💾 Implement robust database integration (PostgreSQL/MongoDB)
- 🛣️ Create RESTful route handlers with proper documentation
- 📝 Add comprehensive logging system with rotation
- 🔐 Implement JWT-based authentication
- 🧪 Add extensive unit and integration tests
- ⚠️ Implement proper error handling with custom errors
- ✅ Add thorough configuration validation
- 🐳 Include Docker support with multi-stage builds
- 📊 Add performance monitoring and metrics
- 📚 Generate API documentation with Swagger/OpenAPI

## 🤝 Contributing

### Code of Conduct
Please read our [Code of Conduct](CODE_OF_CONDUCT.md) before participating in the project.

### How to Contribute

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Commit your changes: `git commit -m "Add your feature name"`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Open a Pull Request against the `main` branch

### Pull Request Guidelines
- Make sure your code is properly formatted and includes necessary tests
- Ensure all required imports are included
- Provide clear documentation for any new functionality

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Support

Please contact the maintainers for any questions or support requests:
- piyushbhardwaj1603@gmail.com
- Piyu-Pika

--- 

⭐️ If you find this project useful, please consider starring it to show your support! ⭐
