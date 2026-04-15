package logger

import (
	"go.uber.org/zap"
)

// NewLogger provides a reference to a zap 'sugared' logger.
func NewLogger(serverMode string) *zap.SugaredLogger {
	var logger *zap.Logger
	if serverMode == "debug" {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync()
	return logger.Sugar()
}
