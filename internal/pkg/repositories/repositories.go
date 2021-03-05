package repositories

import (
	"github.com/stackbreak/tasx/internal/pkg/models"
)

type RepositoryManager struct {
	Person   models.PersonInterface
	TaskList models.TaskListInterface
	Task     models.TaskInterface
}
