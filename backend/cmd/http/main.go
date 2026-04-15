package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/dev-xero/chess.io/pkg/logger"
	"github.com/dev-xero/chess.io/pkg/server"
)

// main is the entry point to the webserver.
func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run is where the magic happens, it essentially bootstraps
// our server.
func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to configure server: %v", err)
	}

	inDevelopment := os.Getenv("SERVER_MODE") == "debug"
	logger := logger.NewLogger(inDevelopment)

	server.NewServer(config, logger)

	return nil
}
