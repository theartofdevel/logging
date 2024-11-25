package logging

import (
	"context"
	"log/slog"
	"os"
	"path"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	ctx := context.Background()

	logger := NewLogger()

	if logger == nil {
		t.Errorf("logger should not be nil")
	}

	if Default() != logger {
		t.Errorf("logger should be default logger")
	}

	if slog.Default() != Default() {
		t.Errorf("logger should be default logger")
	}

	if logger != Default() {
		t.Errorf("logger should be default logger")
	}

	logger = NewLogger(WithIsJSON(true))
	if slog.Default() != logger {
		t.Errorf("logger should be default logger")
	}

	_ = NewLogger(WithIsJSON(false))
	if slog.Default() == logger {
		t.Errorf("logger should NOT be default logger")
	}

	logger = NewLogger(WithLevel("info"))
	if logger.Handler().Enabled(ctx, LevelDebug) {
		t.Errorf("logger should NOT be enabled for debug level")
	}

	logger = NewLogger(WithLevel("warn"))
	if logger.Handler().Enabled(ctx, LevelDebug) {
		t.Errorf("logger should NOT be enabled for debug level")
	}

	logger = NewLogger(WithLevel("debug"))
	enabled := []Level{LevelDebug, LevelInfo, LevelWarn, LevelError}
	for _, level := range enabled {
		if !logger.Handler().Enabled(ctx, level) {
			t.Errorf("logger should be enabled for all levels")
		}
	}

	logger = NewLogger(WithLevel("abcdef"))
	if !logger.Handler().Enabled(ctx, LevelInfo) {
		t.Errorf("logger should be enabled for info level")
	}

	if logger.Handler().Enabled(ctx, LevelDebug) {
		t.Errorf("logger should NOT be enabled for info level")
	}

	logger = NewLogger(WithAddSource(true))
	if slog.Default() != logger {
		t.Errorf("logger should be default logger")
	}

	logger = NewLogger(WithSetDefault(false))
	if slog.Default() == logger {
		t.Errorf("logger should NOT be default logger")
	}

	if logger == Default() {
		t.Errorf("logger should NOT be default logger")
	}

	if L(ctx) == logger {
		t.Errorf("logger should NOT be from context logger")
	}

	logger = NewLogger(WithAddSource(false), WithIsJSON(false))
	ctx = ContextWithLogger(ctx, logger)

	if L(ctx) != logger {
		t.Errorf("logger should be from context logger")
	}
}

func TestLoggerWithFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "logfile_*.log")
	if err != nil {
		t.Fatalf("Failed to create temporary log file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	logger := NewLogger(WithLogFilePath(tempFile.Name()), WithIsJSON(false))
	if logger == nil {
		t.Fatalf("Failed to create logger with file output")
	}

	testMessage := "This is a test log message"
	logger.Info(testMessage)

	// Check file content
	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), testMessage) {
		t.Errorf("Log file does not contain expected message. Got: %s", string(content))
	}
}

func TestLoggerRotationRetention(t *testing.T) {
	tempDir := "./" + "logs"

	err := os.Mkdir(tempDir, 0755) // 0755 defines the permissions
	if err != nil {
		t.Errorf("Failed to create temporary directory: %v", err)
	}

	defer os.RemoveAll(tempDir)

	fileName := "all-log"

	logFilePath := path.Join(tempDir, fileName+".log")
	maxSizeMB := 1
	maxBackups := 3

	NewLogger(
		WithLogFilePath(logFilePath),
		WithLogFileMaxSizeMB(maxSizeMB),   // 1 MB for testing
		WithLogFileMaxBackups(maxBackups), // Keep only 1 backup
		WithLogFileCompress(false),
		WithIsJSON(false),
		WithSetDefault(true),
	)

	if Default() == nil {
		t.Fatalf("Failed to create logger")
	}

	logMessage := "This is a sample log entry for testing purposes."

	messageSize := len(logMessage)

	targetSize := maxSizeMB*1024*1024 + 5000

	for currentSize := 0; currentSize < targetSize; currentSize += messageSize {
		Default().Info(logMessage, IntAttr("currentSize", currentSize))
	}

	// Отладочное сообщение для проверки состояния
	t.Logf("Log file path: %s", logFilePath)

	// Проверяем количество файлов, созданных логгером
	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read temporary directory: %v", err)
	}

	// Фильтруем файлы, связанные с логгером
	var matchingFiles []os.DirEntry
	for _, file := range files {
		if strings.HasPrefix(file.Name(), fileName) {
			matchingFiles = append(matchingFiles, file)
		}
	}

	if len(matchingFiles) == 0 {
		t.Errorf("No log files were created. Check lumberjack configuration or logger output.")
	} else if len(matchingFiles) != maxBackups+1 { // 1 current file + 1 backup
		t.Errorf("Expected exactly 2 log files (1 current + 1 backup), found %d", len(matchingFiles))
	}

	// Отладочная информация о найденных файлах
	for _, file := range matchingFiles {
		t.Logf("Found log file: %s", file.Name())
	}
}
