package services

import (
	"agenda-escolar/internal/domain"
	"agenda-escolar/internal/storage/repository"
)

var taskRepository repository.TaskRepository

type TaskService struct {
}

func (*TaskService) Create(task *domain.Task) (*domain.Task, error) {
	return taskRepository.Create(task)
}

func (*TaskService) Update(task *domain.Task) (*domain.Task, error) {
	return taskRepository.Update(task)
}

func (*TaskService) List(userID int, dateFrom, dateTo string) ([]domain.Task, error) {
	return taskRepository.List(userID, dateFrom, dateTo)
}

func (*TaskService) GetByID(id int) (*domain.Task, error) {
	return taskRepository.GetByID(id)
}