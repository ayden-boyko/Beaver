# Beaver - Logging Service

## Overview
Beaver is a lightweight and efficient logging service written in Go. It provides structured logging capabilities, including middleware for HTTP routes, to enhance observability and debugging for your applications.

## Features
- Structured logging with JSON output.
- Middleware for logging HTTP requests and responses.
- Configurable log levels (debug, info, warn, error).
- Support for logging to different outputs (console, file, remote services).
- Easy integration with existing Go applications.

## Installation
```sh
# Clone the repository
git clone https://github.com/yourusername/beaver.git
cd beaver

# Build the service
go build -o beaver

# Run the service
./beaver
```

## Usage
### Importing Beaver in Your Go Application
```go
import (
    "github.com/yourusername/beaver"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    // Attach Beaver logging middleware
    r.Use(beaver.LoggingMiddleware())
    
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    
    r.Run(":8080")
}
```

### Logging Example
```go
package main

import (
    "github.com/yourusername/beaver"
    "log"
)

func main() {
    logger := beaver.NewLogger(beaver.Config{Level: "info"})
    
    logger.Info("Beaver logging initialized successfully")
    logger.Warn("This is a warning message")
    logger.Error("An error occurred")
}
```

## Configuration
Beaver can be configured using environment variables or a config file:

| Environment Variable | Description | Default |
|----------------------|-------------|---------|
| `BEAVER_LOG_LEVEL` | Log level (`debug`, `info`, `warn`, `error`) | `info` |
| `BEAVER_LOG_OUTPUT` | Log output (`console`, `file`, `remote`) | `console` |
| `BEAVER_LOG_FILE` | File path if `file` output is selected | `logs/app.log` |


## Authors
- [Ayden](https://github.com/ayden-boyko)

