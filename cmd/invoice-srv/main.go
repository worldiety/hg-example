package main

import (
	"context"
	"fmt"
	"github.com/worldiety/hg-example/internal/sys/logging"
	"os/signal"
	"syscall"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	logger := logging.NewLoggerFromEnv()
	ctx = logging.WithLogger(ctx, logger)

	defer func() {
		done()

		if r := recover(); r != nil {
			logger.Error("application panic", fmt.Errorf("panic: %v", r))
		}
	}()

	err := realMain(ctx)

	done()

	if err != nil {
		logger.Error("application error", err)
	}

	logger.Info("successful shutdown")
}

func realMain(ctx context.Context) error {
	return fmt.Errorf("TODO")
}
