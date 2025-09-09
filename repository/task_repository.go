package repository

import (
	"todo-list/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateReppo(task entities.Task) (entities.Task, error)
	GetAll() ([]entities.Task, error)
	GetById(id uint) (entities.Task, error)
	Update(task entities.Task) (entities.Task, error)
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateReppo(task entities.Task) (entities.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAll() ([]entities.Task, error) {
	var tasks []entities.Task
	err := r.db.Find(&tasks).Error
		return tasks,err
}
func (r *taskRepository) GetById(id uint) (entities.Task, error) {
	var task entities.Task
	err := r.db.First(&task, id).Error
	return task, err
}
func (r *taskRepository) Update(task entities.Task) (entities.Task, error) {
	err := r.db.Save(&task).Error
	return task, err
}
func (r *taskRepository) Delete(id uint) error {
	var task entities.Task
	err := r.db.Delete(task, id).Error
	return err
}
