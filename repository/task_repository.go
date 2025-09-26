package repository

import (
	"time"
	"todo-list/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateRepo(task entities.Task) (entities.Task, error)
	GetAllRepo() ([]entities.Task, error)
	GetByIdRepo(id uint) (entities.Task, error)
	UpdateRepo(task entities.Task) (entities.Task, error)
	DeleteRepo(id uint) error
	GetTaskDueBefore(deadline time.Time) ([]entities.Task, error)
	DeleteTaskOlderThan(cutoff time.Time) (int64, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateRepo(task entities.Task) (entities.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllRepo() ([]entities.Task, error) {
	var tasks []entities.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) GetByIdRepo(id uint) (entities.Task, error) {
	var task entities.Task
	err := r.db.First(&task, id).Error
	return task, err
}
func (r *taskRepository) UpdateRepo(task entities.Task) (entities.Task, error) {
	err := r.db.Save(&task).Error
	return task, err
}
func (r *taskRepository) DeleteRepo(id uint) error {
	err := r.db.Delete(&entities.Task{}, id).Error
	return err
}

func (r *taskRepository) GetTaskDueBefore(deadline time.Time) ([]entities.Task, error) {
	var tasks []entities.Task
	err := r.db.Where("due_date <= ?", deadline).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) DeleteTaskOlderThan(cutoff time.Time) (int64, error) {
	result := r.db.Where("due_date <= ?", cutoff).Delete(&entities.Task{})
	return result.RowsAffected, result.Error
}
