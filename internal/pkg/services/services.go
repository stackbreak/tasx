package services

import (
	"github.com/stackbreak/tasx/internal/pkg/repositories"
)

type Services struct {
	repo *repositories.RepositoryManager
}

func NewServices(repo *repositories.RepositoryManager) *Services {
	return &Services{repo}
}
