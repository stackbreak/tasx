package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/stackbreak/tasx/internal/pkg/repositories"
)

const (
	tablePerson     = "person"
	tableTask       = "task"
	tableTaskList   = "task_list"
	tableTaskToList = "task_to_list"
)

type PgConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

func NewDB(cfg *PgConfig) (*sqlx.DB, error) {
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

func NewRepositoryManager(db *sqlx.DB) *repositories.RepositoryManager {
	return &repositories.RepositoryManager{
		Person:   &PersonRepo{db},
		TaskList: &TaskListRepo{db},
		Task:     &TaskRepo{db},
	}
}
