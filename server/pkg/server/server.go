package server

import (
	"net/http"
	"time"

	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// NewServer creates a new webserver. It accepts high-level
// server configuration.
func NewServer(
	logger *zap.SugaredLogger,
	config *config.Config,
	database *sqlx.DB,
) *http.Server {
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
	return s
}
