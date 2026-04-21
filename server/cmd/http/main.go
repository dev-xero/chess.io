package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/dev-xero/chess.io/internal/server"
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
	return server.InitAndServe()
}
