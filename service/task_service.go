package service

import (
	"fmt"
	"strconv"
	"time"
	"todo-list/entities"
	"todo-list/repository"

	"github.com/patrickmn/go-cache"
)

type TaskService interface {
	CreateTaskService(task entities.Task) (entities.Task, error)
	GetAllTasksService() ([]entities.Task, error)
	GetTaskService(id uint) (entities.Task, error)
	UpdateTaskService(task entities.Task) (entities.Task, error)
	DeleteTaskService(id uint) error
}

var todoCache = cache.New(5*time.Minute, 10*time.Minute)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}
func (s *taskService) CreateTaskService(task entities.Task) (entities.Task, error) {
	createdTask, err := s.repo.CreateRepo(task)
	if err != nil {
		return entities.Task{}, err
	}
	todoCache.Delete("allTasks")
	return createdTask, nil
}
func (s *taskService) GetAllTasksService() ([]entities.Task, error) {
	if cachedTasks, found := todoCache.Get("allTasks"); found {
        fmt.Println("serving all tasks from cache:",found)
        return cachedTasks.([]entities.Task), nil
    }
	tasks, err := s.repo.GetAllRepo()
	if err != nil {
		return nil, err
	}
	todoCache.Set("allTasks", tasks, cache.DefaultExpiration)
	fmt.Println("All tasks stored in cache")
	return tasks, nil
}
func (s *taskService) GetTaskService(id uint) (entities.Task, error) {
	if cachedTask, found := todoCache.Get(strconv.FormatUint(uint64(id), 10)); found {
        fmt.Println("Serving task from cache:", found)
        return cachedTask.(entities.Task), nil
    }

	task, err := s.repo.GetByIdRepo(id)
	if err != nil {
		return entities.Task{}, err
	}
	todoCache.Set(strconv.FormatUint(uint64(id), 10), task, cache.DefaultExpiration)
	fmt.Println("task stored in cache:", id)
	return task, nil
}
func (s *taskService) UpdateTaskService(task entities.Task) (entities.Task, error) {
	updateTasks, err := s.repo.UpdateRepo(task)
	if err != nil {
		return entities.Task{}, err
	}
	idStr := strconv.FormatUint(uint64(task.ID), 10)
	todoCache.Set(idStr, updateTasks, cache.DefaultExpiration)
	todoCache.Delete("allTasks")
	fmt.Println("Updated Task")
	return updateTasks, nil
}
func (s *taskService) DeleteTaskService(id uint) error {
	err := s.repo.DeleteRepo(id)
	if err != nil {
		return err
	}
	idStr := strconv.FormatUint(uint64(id), 10)
	todoCache.Delete(idStr)
	todoCache.Delete("allTasks")
	fmt.Println("Deleted Task removed from the cache")
	return nil
}
