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

func (s *Services) TaskListServiceGetOneById(personId, taskListId int) (*models.TaskList, error) {
	return s.repo.TaskList.GetOne(personId, taskListId)
}
