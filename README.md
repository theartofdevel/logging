# Logging Library for Go

A robust and extensible logging library for Go applications, with support for structured logging, JSON output, log file rotation, and configurable options. This library is built on top of Go's standard logging interface (`slog`) and extends it with advanced features like log file rotation using `[lumberjack](https://github.com/natefinch/lumberjack)`.

---

## Features

- **JSON and Text Logging**: Choose between structured JSON output or human-readable text format.
- **Log Levels**: Configurable levels (`debug`, `info`, `warn`, `error`) for fine-grained control.
- **File Logging with Rotation**: Automatically rotate log files based on size, age, and backups.
- **Extensible Options**: Easily configure attributes like log level, file path, and formatting.

---

## Installation

```bash
go get github.com/theartofdevel/logging
```

---

## Usage

### Basic Example

```go
package main

import (
	"context"
	"github.com/theartofdevel/logging"
)

func main() {
	logger := logging.NewLogger(
		logging.WithLevel("info"),
		logging.WithIsJSON(false),
		logging.WithLogFilePath("app.log"),
		logging.WithLogFileMaxSizeMB(10),
	)

	logger.Info("Application started")
	logger.With(logging.IntAttr("user_id", 123)).Error("User not found")
}
```

### Advanced Example with Context

```go
package main

import (
	"context"
	"github.com/theartofdevel/logging"
)

func main() {
	logger := logging.NewLogger(
		logging.WithAddSource(true),
		logging.WithSetDefault(true),
	)

	ctx := logging.ContextWithLogger(context.Background(), logger)
	logging.L(ctx).Info("This log includes context and source")
}
```

---

## Configuration

### Options

| Option                 | Description                                              | Default          |
|------------------------|----------------------------------------------------------|------------------|
| `WithLevel`            | Sets the log level (`debug`, `info`, `warn`, `error`).   | `info`           |
| `WithIsJSON`           | Enables JSON output.                                     | `true`           |
| `WithLogFilePath`      | Specifies the log file path.                             | Writes to stdout |
| `WithLogFileMaxSizeMB` | Maximum log file size in MB before rotation.             | `10`             |
| `WithLogFileMaxBackups`| Number of rotated log files to retain.                   | `3`              |
| `WithLogFileMaxAgeDays`| Maximum number of days to retain old log files.          | `30`             |
| `WithAddSource`        | Includes source file and line number in logs.            | `true`           |
| `WithSetDefault`       | Sets this logger as the default logger for `slog`.       | `true`           |

---

## Testing

Run the tests using the following command:

```bash
go test ./...
```

---

## Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m "Add new feature"`).
4. Push to the branch (`git push origin feature-name`).
5. Open a Pull Request.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Authors

Developed by Artur Karapetov for Youtube