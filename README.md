# URL Shortener

![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-in%20development-yellow)

A lightweight and efficient URL shortening service built with Go. This service allows you to convert long URLs into short, manageable links.

## Features

- Generate short URLs from long links
- SQLite-based storage for simplicity
- Clean and minimal configuration
- Structured logging with `log/slog`
- Environment-based configuration

## Project Stack

- **Language**: Go 1.25
- **Configuration**: [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- **Logging**: `log/slog` (standard library)
- **Storage**: SQLite
- **Router**: chi, chi/render
- **Environment**: godotenv

## Getting Started

### Prerequisites

- Go 1.25 or higher
- SQLite3

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd url-shortener
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
```

The `.env` file should contain:
```
CONFIG_PATH=./config/local.yaml
```

### Configuration

The application uses a YAML configuration file located at `config/local.yaml`:

```yaml
env: "local"
storage_path: "./storage/storage.db"
http_server:
    address: "localhost:8080"
    timeout: 4s
    idle_timeout: 60s
```

### Running the Application

```bash
go run cmd/url-shortener/main.go
```

Or build and run:
```bash
go build -o url-shortener cmd/url-shortener/main.go
./url-shortener
```

## Project Structure

```
.
├── cmd/
│   └── url-shortener/       # Application entrypoint
├── config/                   # Configuration files
│   └── local.yaml
├── internal/
│   ├── config/              # Configuration loading
│   └── storage/             # Database layer
├── .env                     # Environment variables
└── go.mod
```

## Future Roadmap

- [ ] **PostgreSQL Support**: Replace SQLite with PostgreSQL for production-ready scalability
- [ ] **Basic Authentication**: Add Basic Auth for protected endpoints
- [ ] **Custom Short URLs**: Allow users to specify custom short codes
- [ ] **Analytics**: Track click statistics and analytics
- [ ] **Rate Limiting**: Implement rate limiting for API endpoints
- [ ] **REST API**: Full REST API with comprehensive documentation
- [ ] **Docker Support**: Containerize the application

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Linting

```bash
golangci-lint run
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
