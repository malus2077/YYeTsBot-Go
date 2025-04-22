# YYeTsBot-Go

[![Go Report Card](https://goreportcard.com/badge/github.com/malus2077/YYeTsBot-Go)](https://goreportcard.com/report/github.com/yourusername/YYeTsBot-Go)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`YYeTsBot-Go` is the Go implementation of [YYeTsBot](https://github.com/tgbot-collection/YYeTsBot), a project for backing up resources from YYeTs and user shares. Built with the [Kratos](https://go-kratos.dev) framework, it provides a modern backend service.

## ✨ Features

* **HTTP + gRPC** APIs generated from Protobuf.
* **JWT** middleware for authentication.
* **Redis** for caching.
* **MongoDB** for data storage.
* **Wire** for dependency injection.
* **Colly** for web scraping.

## 🚀 Getting Started

### ✅ Prerequisites

* Go **1.24+**
* MongoDB **≥ 7.0**
* Redis **≥ 7.0**
* (Optional) Docker + Docker Compose

### Run Locally

1.  **Clone:**
    ```bash
    git clone https://github.com/malus2077/YYeTsBot-Go.git
    cd YYeTsBot-Go
    ```

2.  **Initialize:** (Installs deps, generates code)
    ```bash
    make init
    ```

3.  **Configure:**
    * Edit `configs/config.yaml` with your DB, Redis, and JWT settings.

4.  **Run:** (Starts dev server)
    ```bash
    kratos run
    ```

Server runs on:
* **HTTP:** `http://localhost:8000`
* **gRPC:** `localhost:9000`

## 🐳 Docker

1.  **Build Image:**
    ```bash
    docker build -t yyetsbot-go:latest .
    ```

2.  **Run Container:**
    ```bash
    docker run --rm \
           -p 8000:8000 \
           -p 9000:9000 \
           -v $(pwd)/configs:/data/conf \
           --name yyetsbot-go-app \
           yyetsbot-go:latest
    ```
    * Mount your local `configs` directory to `/data/conf` in the container. Use an absolute path for `$(pwd)/configs`.

## ⚙️ Configuration

* Configure the application via YAML files in the `configs/` directory.
* Key settings: `database`, `redis`, `server` ports, `auth.jwt.secret`.
* Environment variables might override YAML values (check Kratos setup).

## 🏗️ Project Structure

```
.
├── api/            # Protobuf definitions (.proto)
├── cmd/            # Main application entrypoint
├── configs/        # Configuration files
├── internal/       # Core application code
│   ├── biz/        # Business logic
│   ├── data/       # Data access (DB, cache)
│   └── service/    # API service implementations
├── Makefile        # Build commands
├── go.mod          # Go modules
└── Dockerfile      # Docker build file
```

## 🤝 Contributing

Contributions are welcome!

* Report bugs or suggest features via [Issues](https://github.com/malus2077/YYeTsBot-Go/issues).
* Submit code changes via [Pull Requests](https://github.com/malus2077/YYeTsBot-Go/pulls).

## 📄 License

Licensed under the **Apache License 2.0**. See the [LICENSE](LICENSE) file for details.