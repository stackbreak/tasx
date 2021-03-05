package services

import (
	"github.com/stackbreak/tasx/internal/pkg/repositories"
	"github.com/stackbreak/tasx/internal/pkg/tokens"
)

type Services struct {
	repo   *repositories.RepositoryManager
	tokens *tokens.TokenManagerHS
}

func NewServices(repo *repositories.RepositoryManager, tokenManager *tokens.TokenManagerHS) *Services {
	return &Services{repo, tokenManager}
}
