# 🚀 Go Redis Implementation

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)
![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=for-the-badge)

*A lightweight, high-performance Redis-compatible key-value store built from scratch in Go*

</div>

---

## ✨ Features

- 🔥 **High Performance** - Built with Go's concurrency primitives
- 🌐 **Network Ready** - TCP server with RESP protocol support
- 🔐 **Thread Safe** - Concurrent read/write operations with mutex protection
- 📦 **Simple Client** - Easy-to-use Go client library included
- 🧪 **Well Tested** - Comprehensive test suite for reliability
- ⚡ **Fast Setup** - Get running in seconds

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│                 │    │                 │    │                 │
│     Client      │◄──►│     Server      │◄──►│   Key-Value     │
│                 │    │                 │    │     Store       │
│  - TCP Client   │    │  - TCP Server   │    │  - Thread Safe  │
│  - RESP Proto   │    │  - RESP Parser  │    │  - In Memory    │
│                 │    │  - Peer Manager │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- Go 1.19 or higher
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/Adityaadpandey/Redis.git
cd Redis

# Install dependencies
go mod tidy

# Build and Run the server
make
```

The server will start on `localhost:5832` by default.

## 🔧 Usage

### Server Commands

```bash
# Start server on default port (5832)
./bin/redis

# Start server on custom port
./bin/redis --listenAddr :6379
```

### Client Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/adityaadpandey/Redis/client"
)

func main() {
    // Create a new client
    c, err := client.New("localhost:5832")
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // Set a key-value pair
    err = c.Set(context.TODO(), "mykey", "myvalue")
    if err != nil {
        log.Fatal(err)
    }

    // Get the value
    val, err := c.Get(context.TODO(), "mykey")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Value: %s\n", val)
}
```

## 📋 Supported Commands

| Command | Description | Example |
|---------|-------------|---------|
| `SET` | Set a key to hold a string value | `SET mykey "hello"` |
| `GET` | Get the value of a key | `GET mykey` |

## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./...

# Run client tests specifically
make test

# Run with verbose output
go test -v ./...
```

### Concurrent Client Test

The project includes a stress test that spawns multiple concurrent clients:

```bash
# Test with 10 concurrent clients
go test -timeout 30s -run ^TestNewClients$ ./client -v -count=1
```

## 🏗️ Project Structure

```
.
├── client/
│   ├── client.go          # Redis client implementation
│   └── client_test.go     # Client tests
├── src/
│   ├── main.go           # Server entry point
│   ├── server.go         # Core server logic
│   ├── peer.go           # Connection handling
│   ├── kv.go             # Key-value store
│   ├── protocol.go       # RESP protocol parser
│   └── *_test.go         # Test files
├── bin/                  # Compiled binaries
├── Makefile             # Build automation
└── README.md
```

## 🔄 Protocol

This implementation uses the **RESP (Redis Serialization Protocol)** for client-server communication, making it compatible with standard Redis clients.

### Message Format Example

```
*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n
```

This represents: `SET foo bar`

## ⚡ Performance

- **Concurrent Operations**: Supports multiple simultaneous client connections
- **Memory Efficient**: In-memory storage with minimal overhead
- **Fast Networking**: Built on Go's efficient TCP stack
- **Lock Optimization**: Read-write mutex for optimal concurrent access

## 🛠️ Development

### Building from Source

```bash
# Install dependencies
go mod tidy

# Build the project
make build

# Run development server
make run
```

### Adding New Commands

1. Define the command in `protocol.go`
2. Add parsing logic in `parseCommand()`
3. Handle the command in `server.go`'s `handleMessage()`
4. Add tests for the new functionality

## 📝 Learning Journey

This project was built as a **first-time Go learning experience**, implementing:

- **Concurrency patterns** with goroutines and channels
- **Network programming** with TCP sockets
- **Protocol parsing** with RESP format
- **Memory management** with proper resource cleanup
- **Testing strategies** for concurrent systems

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by the original [Redis](https://redis.io/) project
- Built while learning Go programming language
- Uses [tidwall/resp](https://github.com/tidwall/resp) for RESP protocol parsing

---

<div align="center">

**⭐ Star this repo if you found it helpful!**

Made with ❤️ while learning Go

</div>
