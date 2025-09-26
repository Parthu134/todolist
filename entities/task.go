package entities

import "time"

type Task struct {
	ID          uint      `gorm:"primarykey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"`
	UserEmail   string    `json:"user_email"`
}
