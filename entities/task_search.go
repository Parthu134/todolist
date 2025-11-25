package entities

import "time"

type Tasksearch struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Tags      []string  `json:"tags"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
}
