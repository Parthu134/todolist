package service

import (
	"todo-list/entities"
	"todo-list/repository"
)

type TaskService interface {
	CreateTaskService(task entities.Task) (entities.Task, error)
	GetAllTasksService() ([]entities.Task, error)
	GetTaskService(id uint) (entities.Task, error)
	UpdateTaskService(task entities.Task) (entities.Task, error)
	DeleteTaskService(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}
func (s *taskService) CreateTaskService(task entities.Task) (entities.Task, error) {
	return s.repo.CreateRepo(task)
}
func (s *taskService) GetAllTasksService() ([]entities.Task, error) {
	return s.repo.GetAllRepo()
}
func (s *taskService) GetTaskService(id uint) (entities.Task, error) {
	return s.repo.GetByIdRepo(id)
}
func (s *taskService) UpdateTaskService(task entities.Task) (entities.Task, error) {
	return s.repo.UpdateRepo(task)
}
func (s *taskService) DeleteTaskService(id uint) error {
	return s.repo.DeleteRepo(id)
}

