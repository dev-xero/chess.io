package server

import (
	"net/http"
	"time"

	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewServer creates a new webserver. It accepts high-level
// server configuration.
func NewServer(config *config.Config, logger *zap.SugaredLogger) {
	gin.SetMode(config.ServerMode)
	mux := gin.Default()
	mux.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Info("http webserver listening at http://localhost:" + config.Port)
	s.ListenAndServe()
}
