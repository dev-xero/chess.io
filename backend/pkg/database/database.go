package database

import (
	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/jmoiron/sqlx"
)

// InitializeDatabase constructs a
func NewDatabaseConnection(c *config.Config) sqlx.DB {
	// datasourceName := fmt.Sprintf(
	// 	"host=%s",
	// )
}
