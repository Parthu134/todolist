package service

import (
	"strings"
	"time"
	"todo-list/entities"

	"github.com/gofiber/fiber/v2"
)

var Tasks = []entities.Tasksearch{
	{
		ID:      "1",
		Title:   "Buy Groceries",
		Status:  "pending",
		Tags:    []string{"shopping"},
		DueDate: time.Now().Add(24 * time.Hour),
	},
	{
		ID:      "2",
		Title:   "Buy Fruits",
		Status:  "Pending",
		Tags:    []string{"work"},
		DueDate: time.Now().Add(-1 * time.Hour),
	},
}

func SearchFilter(c *fiber.Ctx) error {
	title := strings.ToLower(c.Query("title"))
	status := strings.ToLower(c.Query("status"))
	tag := strings.ToLower(c.Query("tag"))
	dueDate := c.Query("due_date")

	var filtered []entities.Tasksearch

	for _, task := range Tasks {
		if title != "" && !strings.Contains(strings.ToLower(task.Title), title) {
			continue
		}
		if status != "" && strings.ToLower(task.Status) != status {
			continue
		}
		if tag != "" {
			found := false
			for _, t := range task.Tags {
				if strings.ToLower(t) == tag {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		if dueDate != "" {
			d, err := time.Parse("2006-01-02", dueDate)
			if err == nil && task.DueDate.Format("2006-01-02") != d.Format("2006-01-02") {
				continue
			}
		}
		filtered = append(filtered, task)
	}
	return c.JSON(filtered)
}
