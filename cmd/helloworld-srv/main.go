package main

import (
	"flag"
	helloworld_srv "github.com/worldiety/hg-example/internal/app/helloworld-srv"
	"github.com/worldiety/hg-example/internal/sys/logging"
	"golang.org/x/exp/slog"
)

func main() {
	host := flag.String("host", "localhost:8080", "the host to bind on")
	flag.Parse()

	if err := helloworld_srv.Run(*host); err != nil {
		logging.NewLoggerFromEnv().Error("failed to start helloword service", slog.Any("err", err))
	}
}
