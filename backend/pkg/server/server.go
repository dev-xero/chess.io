package server

import (
	"github.com/dev-xero/chess.io/pkg/config"
	"go.uber.org/zap"
)

// NewServer creates a new webserver. It accepts high-level
// server configuration.
func NewServer(
	config *config.Config,
	logger *zap.SugaredLogger,
) {
	logger.Infof("config: %+v\n", config)
}
