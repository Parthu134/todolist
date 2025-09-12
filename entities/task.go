package entities

type Task struct {
	ID          uint `gorm:"primarykey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

