package service

import "github.com/stackbreak/tasx/internal/pkg/repository"

type Authorization interface { //
}

type TaskList interface { //
}

type TaskListItem interface { //
}

type Service struct {
	Authorization
	TaskList
	TaskListItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
