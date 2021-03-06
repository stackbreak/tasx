package services

import "github.com/stackbreak/tasx/internal/pkg/models"

func (s *Services) TaskListServiceCreate(personId int, list *models.TaskList) (int, error) {
	taskListId, err := s.repo.TaskList.CreateOne(personId, list)
	if err != nil {
		return 0, err
	}

	return taskListId, nil
}
