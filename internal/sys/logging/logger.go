package logging

import (
	"context"
	"github.com/worldiety/hg-example/internal/sys/version"
	"os"
	"strconv"
	"sync"

	"golang.org/x/exp/slog"
)

// contextKey is a private string type to prevent collisions in the context map.
type contextKey string

// loggerKey points to the value in the context where the logger is stored.
const loggerKey = contextKey("logger")

var (
	// defaultLogger is the default logger. It is initialized once per package
	// include upon calling defLogger.
	// This is just an optimization detail.
	defaultLogger     *slog.Logger
	defaultLoggerOnce sync.Once
)

// NewLoggerFromEnv allocates a new base logger based on the environment settings.
func NewLoggerFromEnv() *slog.Logger {
	_, isDevMode := os.LookupEnv("XPC_SERVICE_NAME") // some MacOS specific IPC stuff
	level := os.Getenv("LOG_LEVEL")

	logger := slog.Default()

	if !isDevMode {
		opts := &slog.HandlerOptions{
			AddSource: isDevMode,
			Level:     getLogLevel(isDevMode, level),
		}

		logger = slog.New(slog.NewJSONHandler(os.Stderr, opts)).With(
			slog.String("build_sha", version.BuildSha),
			slog.String("build_tag", version.BuildTag),
		)
	}

	return logger
}

func getLogLevel(devMode bool, levelText string) slog.Level {
	if devMode {
		return slog.LevelDebug
	}

	intLevel, _ := strconv.Atoi(levelText)

	return slog.Level(intLevel)
}

// WithLogger creates a new context with the provided logger attached.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// defLogger returns the default logger for the package and is intentionally package private.
// Use always FromContext instead.
func defLogger() *slog.Logger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLoggerFromEnv()
	})

	return defaultLogger
}

// FromContext returns the logger stored in the context. If no such logger
// exists, a default logger is returned.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}

	return defLogger()
}

const Incident = "incident"
