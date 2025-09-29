package repository

import (
	"todo-list/entities"

	"gorm.io/gorm"
)

type TaskBackupRepository interface{
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