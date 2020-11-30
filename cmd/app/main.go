package main

import (
	"context"
	"io"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/schigh-ntwrk/ent-poc/internal/bootstrap"
)

// create top-level application activity logger
func mkLogger() *zap.SugaredLogger {
	lg, lgErr := zap.NewProduction(zap.OnFatal(zapcore.WriteThenFatal))
	if lgErr != nil {
		panic(lgErr)
	}
	return lg.Sugar()
}

// this seeds the uuid generator with a
// node identifier and clock sequence
func seedUUID() {
	uuid.SetNodeID([]byte("NTWRK_"))
	uuid.SetClockSequence(math.MaxInt32)
}

func main() {

	logger := mkLogger()
	defer logger.Sync()

	seedUUID()

	// run the wire bootstrapper
	deps, depsErr := bootstrap.Up(context.Background())
	if depsErr != nil {
		logger.Fatalw("bootstrapping application failed", "error", depsErr)
	}

	// run application server
	go deps.Server.Start()

	// populate closers from closable dependencies
	closers := []io.Closer{
		deps.EntClient,
	}

	// listen for SIGTERM/SIGINT
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	logger.Infow("pet hospital application running", "event", "UP")
	<-stopChan

	logger.Infow("shutdown signal received.  cleaning up", "event", "TERM")
	if srvErr := deps.Server.Shutdown(context.Background()); srvErr != nil {
		logger.Errorw("server shutdown error", "error", srvErr)
	}

	// clean up
	for _, c := range closers {
		_ = c.Close()
	}
	logger.Infow("pet hospital application shutdown complete", "event", "DOWN")
}
