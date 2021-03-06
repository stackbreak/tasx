package repositories

import (
	"github.com/stackbreak/tasx/internal/pkg/models"
)

type RepositoryManager struct {
	Person
	TaskList
	Task
}

type Person interface {
	CreateOne(person *models.Person) (int, error)
	GetOne(username string) (*models.Person, error)
}

type TaskList interface { //
	CreateOne(personId int, list *models.TaskList) (int, error)
	GetAll(personId int) ([]models.TaskList, error)
	GetOne(personId, taskListId int) (*models.TaskList, error)
}

type Task interface { //
}
