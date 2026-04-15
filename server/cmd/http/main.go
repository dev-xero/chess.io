package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/dev-xero/chess.io/pkg/database"
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

	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to configure server: %v", err)
	}

	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize a database connection: %v", err)
	}

	server.NewServer(
		logger.NewLogger(cfg.ServerMode),
		cfg,
		db,
	)

	return nil
}
