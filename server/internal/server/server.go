package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/dev-xero/chess.io/pkg/database"
	"github.com/dev-xero/chess.io/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Server struct {
	config *config.Config
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

// InitAndServe creates a new webserver. It accepts high-level server
// configuration.
func InitAndServe() error {
	server, err := bootstrap()
	if err != nil {
		return fmt.Errorf("failed to bootstrap server: %v", err)
	}

	gin.SetMode(server.config.ServerMode)

	mux := gin.Default()
	mux.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srv := http.Server{
		Addr:           ":" + server.config.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.logger.Info("http webserver listening at http://localhost:" + server.config.Port)
	return srv.ListenAndServe()
}

// bootstrap configures and instantiates the server's external
// dependencies
func bootstrap() (*Server, error) {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		return nil, err
	}

	logger := logger.NewLogger(cfg.ServerMode)

	return &Server{cfg, db, logger}, nil
}
