package logger

import "go.uber.org/zap"

// NewLogger provides a reference to a zap 'sugared' logger.
func NewLogger(inDevelopment bool) *zap.SugaredLogger {
	var logger *zap.Logger
	if inDevelopment {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync()
	return logger.Sugar()
}
