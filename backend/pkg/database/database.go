package database

import (
	"fmt"

	"github.com/dev-xero/chess.io/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitializeDatabase constructs a postgres datasource name and attempts
// to connect to the database.
func NewDatabaseConnection(c *config.Config) (*sqlx.DB, error) {
	datasourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.DbHost,
		c.DbUser,
		c.DbPass,
		c.DbName,
		c.DbPort,
		c.DbSslMode,
	)
	return sqlx.Connect("postgres", datasourceName)
}
