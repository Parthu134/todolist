package service

import (
	"todo-list/entities"
	"todo-list/repository"
)

type TaskService interface {
	CreateTaskService(task entities.Task) (entities.Task, error)
	GetAllTasks() ([]entities.Task, error)
	GetTask(id uint) (entities.Task, error)
	UpdateTask(task entities.Task) (entities.Task, error)
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}
func (s *taskService) CreateTaskService(task entities.Task) (entities.Task, error) {
	return s.repo.CreateReppo(task)
}
func (s *taskService) GetAllTasks() ([]entities.Task, error) {
	return s.repo.GetAll()
}
func (s *taskService) GetTask(id uint) (entities.Task, error) {
	return s.repo.GetById(id)
}
func (s *taskService) UpdateTask(task entities.Task) (entities.Task, error) {
	return s.repo.Update(task)
}
func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
