package repository

import "github.com/jmoiron/sqlx"

type Authorization interface { //
}

type TaskList interface { //
}

type TaskListItem interface { //
}

type Repository struct {
	Authorization
	TaskList
	TaskListItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
