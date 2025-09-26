package entities

import (
	"time"
)

type TaskBackup struct {
	ID         uint      `gorm:"primarykey"`
	Tasks      string   `gorm:"type:jsonb"`
	BackupTime time.Time `json:"backup_time"`
}
