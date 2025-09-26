package repository

import (
	"time"
	"todo-list/entities"

	"gorm.io/gorm"
)

type TaskBackupRepository interface{
	DeleteOldBackups(cutoff time.Time) (int64, error)
	GetAllBackups() ([]entities.TaskBackup, error)
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

func (r *taskBackupRepository) DeleteOldBackups(cutoff time.Time) (int64, error) {
	result := r.DB.Where("backup_time <= ?", cutoff).Delete(&entities.TaskBackup{})
	return result.RowsAffected, result.Error
}

func (r *taskBackupRepository) GetAllBackups() ([]entities.TaskBackup, error) {
	var backups []entities.TaskBackup
	err := r.DB.Find(&backups).Error
	return backups, err
}