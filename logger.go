package logging

import (
	"context"
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultLevel             = LevelInfo
	defaultAddSource         = true
	defaultIsJSON            = true
	defaultSetDefault        = true
	defaultLogFile           = ""
	defaultLogFileMaxSizeMB  = 10
	defaultLogFileMaxBackups = 3
	defaultLogFileMaxAgeDays = 14
)

func NewLogger(opts ...LoggerOption) *Logger {
	config := &LoggerOptions{
		Level:             defaultLevel,
		AddSource:         defaultAddSource,
		IsJSON:            defaultIsJSON,
		SetDefault:        defaultSetDefault,
		LogFilePath:       defaultLogFile,
		LogFileMaxSizeMB:  defaultLogFileMaxSizeMB,
		LogFileMaxBackups: defaultLogFileMaxBackups,
		LogFileMaxAgeDays: defaultLogFileMaxAgeDays,
	}

	for _, opt := range opts {
		opt(config)
	}

	options := &HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
	}

	// by default we write to stdout.
	var w io.Writer = os.Stdout

	// file or stdout.
	if config.LogFilePath != "" {
		w = &lumberjack.Logger{
			Filename:   config.LogFilePath,
			MaxSize:    config.LogFileMaxSizeMB,
			MaxBackups: config.LogFileMaxBackups,
			MaxAge:     config.LogFileMaxAgeDays,
			Compress:   config.LogFileCompress,
		}
	}

	var h Handler = NewTextHandler(w, options)

	if config.IsJSON {
		h = NewJSONHandler(w, options)
	}

	logger := New(h)

	if config.SetDefault {
		SetDefault(logger)
	}

	return logger
}

type LoggerOptions struct {
	Level             Level
	AddSource         bool
	IsJSON            bool
	SetDefault        bool
	LogFilePath       string
	LogFileMaxSizeMB  int
	LogFileMaxBackups int
	LogFileMaxAgeDays int
	LogFileCompress   bool
}

type LoggerOption func(*LoggerOptions)

// WithLevel logger option sets the log level, if not set, the default level is Info.
func WithLevel(level string) LoggerOption {
	return func(o *LoggerOptions) {
		var l Level
		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = LevelInfo
		}

		o.Level = l
	}
}

// WithAddSource logger option sets the add source option, which will add source file and line number to the log record.
func WithAddSource(addSource bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.AddSource = addSource
	}
}

// WithIsJSON logger option sets the is json option, which will set JSON format for the log record.
func WithIsJSON(isJSON bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.IsJSON = isJSON
	}
}

// WithSetDefault logger option sets the set default option, which will set the created logger as default logger.
func WithSetDefault(setDefault bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.SetDefault = setDefault
	}
}

// WithLogFilePath logger option sets the file where logs will be written.
func WithLogFilePath(logFilePath string) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFilePath = logFilePath
	}
}

// WithLogFileMaxSizeMB logger option sets the maximum file size for rotation.
func WithLogFileMaxSizeMB(maxSize int) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileMaxSizeMB = maxSize
	}
}

// WithLogFileMaxBackups logger option sets the number of backup files to retain.
func WithLogFileMaxBackups(maxBackups int) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileMaxBackups = maxBackups
	}
}

// WithLogFileMaxAgeDays logger option sets the maximum age of the log files.
func WithLogFileMaxAgeDays(maxAge int) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileMaxAgeDays = maxAge
	}
}

// WithLogFileCompress logger options set needs compression.
func WithLogFileCompress(compression bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileCompress = compression
	}
}

// WithAttrs returns logger with attributes.
func WithAttrs(ctx context.Context, attrs ...Attr) *Logger {
	logger := L(ctx)
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

// WithDefaultAttrs returns logger with default attributes.
func WithDefaultAttrs(logger *Logger, attrs ...Attr) *Logger {
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

func L(ctx context.Context) *Logger {
	return loggerFromContext(ctx)
}

func Default() *Logger {
	return slog.Default()
}
