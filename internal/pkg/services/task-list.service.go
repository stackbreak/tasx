package services

import (
	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (s *Services) TaskListServiceCreate(personId int, list *models.TaskList) (int, error) {
	return s.repo.TaskList.CreateOne(personId, list)
}

func (s *Services) TaskListServiceGetAll(personId int) ([]models.TaskList, error) {
	return s.repo.TaskList.GetAll(personId)
}

func (s *Services) TaskListServiceGetOne(personId, taskListId int) (*models.TaskList, error) {
	return s.repo.TaskList.GetOne(personId, taskListId)
}

func (s *Services) TaskListServiceDeleteOne(personId, taskListId int) error {
	return s.repo.TaskList.DeleteOne(personId, taskListId)
}
