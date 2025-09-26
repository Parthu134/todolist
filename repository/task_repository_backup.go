package repository

import (
	"time"
	"todo-list/entities"

	"gorm.io/gorm"
)

type TaskBackupRepository interface{
	GetTaskDueBefore(deadline time.Time) ([]entities.Task, error)
	DeleteTaskOlderThan(cutoff time.Time) (int64, error)

	CreateBackupRepo(backup entities.TaskBackup) (entities.TaskBackup, error)
}
type taskBackupRepository struct {
	DB *gorm.DB
}

func NewTaskBackupRepository(db *gorm.DB) TaskBackupRepository {
	return &taskBackupRepository{DB: db}
}

func (r *taskBackupRepository) CreateBackupRepo(backup entities.TaskBackup) (entities.TaskBackup, error) {
	err := r.DB.Create(&backup).Error
	return backup, err
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