package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PgConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

func NewPgDB(cfg *PgConfig) (*sqlx.DB, error) {
	dbURI := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Pass,
		cfg.SSLMode,
	)

	db, err := sqlx.Open("pgx", dbURI)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
