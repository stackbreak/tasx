package services

import "github.com/stackbreak/tasx/internal/pkg/models"

func (s *Services) TaskServiceCreateOne(personId, listId int, task *models.Task) (int, error) {
	_, err := s.repo.TaskList.GetOne(personId, listId)
	if err != nil {
		return -1, err
	}

	return s.repo.Task.CreateOne(listId, task)
}

func (s *Services) TaskServiceGetAll(personId, listId int) ([]models.Task, error) {
	return s.repo.Task.GetAll(personId, listId)
}
