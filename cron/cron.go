package cron

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"todo-list/entities"
	"todo-list/repository"
	"todo-list/service"
	"todo-list/utils"

	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

type TaskCron struct {
	MainRepo repository.TaskRepository
	BackupRepo repository.TaskRepository
}

func NewTaskCron(Main repository.TaskRepository, Backup repository.TaskRepository) *TaskCron {
	return &TaskCron{
		MainRepo: Main,
		BackupRepo: Backup,
	}
}
func (tc *TaskCron) Start(t time.Duration) {
	c := cron.New()
	// c.AddFunc("0 9 * * *", tc.SendDailyRemainders)
	c.AddFunc("@every 1m", tc.SendDailyRemainders)
	// c.AddFunc("0 0 * * *", tc.CleanExpiredTasks)
	c.AddFunc("@every 2m", tc.CleanExpiredTasks)
	c.AddFunc("0 2 * * 0", tc.BackupDatabase)
	// c.AddFunc("@every 10m", tc.RefreshCaches)
	c.AddFunc("@every 15m", tc.RefreshCaches)
	c.Start()
}
func (c *TaskCron) SendDailyRemainders() {
	log.Println("SendDailyRemainders triggered...")
	tasks, err := c.MainRepo.GetTaskDueBefore(time.Now().Add(24 * time.Hour))
	if err != nil {
		log.Printf("cron daily remainder error: %v", err)
		return
	}
	if len(tasks) == 0 {
		log.Println("No tasks due in next 24s")
	}
	for _, t := range tasks {
		log.Printf("cron Remainder: Task %d: %s (due: %v)", t.ID, t.Title, t.DueDate)
		Subject := fmt.Sprintf("Remainder, Task %s is due", t.Title)
		body := fmt.Sprintf("Hello\n\nYour task \"%s\" is due on %v. \n\nDescription: %s", t.Title, t.DueDate, t.Description)
		if err := utils.SendOtp(t.UserEmail, Subject, body); err != nil {
			log.Printf("failed to send mail to ID:%d, %v", t.ID, err)
		} else {
			log.Printf("Email sent to user ID: %d", t.ID)
		}
	}
}

func (c *TaskCron) CleanExpiredTasks() {
	cutoff := time.Now().Add(-30 * 24 * time.Hour)
	count, err := c.BackupRepo.DeleteTaskOlderThan(cutoff)
	if err != nil {
		log.Printf("cron clean expired time error: %v", err)
		return
	}
	log.Printf("cron cleared %d expired token", count)
}

func (c *TaskCron) BackupDatabase() {
	log.Println("cron database backup started")
	tasks, err := c.MainRepo.GetAllRepo()
	if err != nil {
		fmt.Printf("error fetching tasks for backup:%v", err)
		return
	}
	tasksJson, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("error marshalling tasks to JSON %v", err)
		return
	}
	backup := entities.TaskBackup{
		Tasks:      string(tasksJson),
		BackupTime: time.Now(),
	}
	if _, err := c.BackupRepo.CreateBackupRepo(backup); err != nil {
		log.Printf("err saving backup :%v", err)
		return
	}
	log.Println("backup saved inside postgres database")
}

func (c *TaskCron) RefreshCaches() {
	task, err := c.MainRepo.GetAllRepo()
	if err != nil {
		log.Printf("cron refresh cache error: %v", err)
		return
	}
	service.TodoCache.Set("allTasks", task, cache.DefaultExpiration)
	log.Println("log cache refreshed for allTasks")
}
