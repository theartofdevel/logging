package logging

import (
	"context"
)

func simple() {
	logger := NewLogger()
	logger.Info("Hello, World!")

	// {"time":"2024-01-27T00:37:27.452291+03:00","level":"INFO","source":{"function":"logging.simple","file":"logging/example.go","line":5},"msg":"Hello, World!"}

	logger = NewLogger(WithLevel("debug"), WithAddSource(false))
	logger.Debug("Hello, World!")

	// {"time":"2024-01-27T00:38:29.168834+03:00","level":"DEBUG","msg":"Hello, World!"}

	logger = NewLogger(WithLevel("debug"), WithAddSource(false), WithIsJSON(false))
	logger.Debug("Hello, World!")

	// time=2024-01-27T00:39:04.497+03:00 level=DEBUG msg="Hello, World!"

	logger = NewLogger(WithAddSource(true))
	Default().Info("Hello, World!")

	// {"time":"2024-01-27T00:39:43.120848+03:00","level":"INFO","msg":"Hello, World!"}

	ctx := context.Background()
	logger = WithAttrs(ctx, StringAttr("hello", "world"))
	logger.Info("OK")

	// {"time":"2024-01-27T00:44:55.083891+03:00","level":"INFO","source":{"function":"logging.simple","file":"theartofdevelopment/libs/logging/example.go","line":30},"msg":"OK","hello":"world"}

	ctx = ContextWithLogger(ctx, WithAttrs(ctx, StringAttr("where?", "from context!")))
	fromContext(ctx)
}

func fromContext(ctx context.Context) {
	L(ctx).Info("me")

	// {"time":"2024-01-27T00:46:33.235805+03:00","level":"INFO","source":{"function":"logging.fromContext","file":"theartofdevelopment/libs/logging/example.go","line":39},"msg":"me","where?":"from context!"}
}

func byDefault() {
	NewLogger(WithIsJSON(false), WithAddSource(false), WithSetDefault(false))
	Default().Info("Hello, World!")

	// {"time":"2024-01-27T00:41:49.6227+03:00","level":"INFO","source":{"function":"logging.byDefault","file":"theartofdevelopment/libs/logging/example.go","line":27},"msg":"Hello, World!"}
}
