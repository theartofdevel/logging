package logging

import (
	"context"
	"log/slog"
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
