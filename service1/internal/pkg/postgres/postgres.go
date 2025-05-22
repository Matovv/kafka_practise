package postgres

import (
	"local/service1/internal/pkg/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driverName = "postgres"

func NewDatabase(config *configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, config.DatabaseURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}