package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func Init(addr, pass string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
}

type Job struct {
	Type   string `json:"type"`
	TaskID string `json:"task_id omitempty"`
}

const (
	queueName = "todo_jobs"
)

func Enqueue(job Job) error { // enque means push a job
	b, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return rdb.RPush(ctx, queueName, b).Err()
}

func Dequeue(timeout time.Duration) (*Job, error) { //this will get a job from the queue and give it to worker for processing
	res, err := rdb.BLPop(ctx, timeout, queueName).Result()
	if err != nil {
		return nil, err
	}
	if len(res) < 2 {
		return nil, fmt.Errorf("invalid job payload")
	}
	var job Job
	if err := json.Unmarshal([]byte(res[1]), &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func Startworker() {
	for {
		job, err := Dequeue(5 * time.Second)
		if err != nil {
			continue
		}
		handlejob(job)
	}
}

func handlejob(job *Job) {
	switch job.Type {
	case "task_created":
		processTaskCreated(job.TaskID)
	case "task_updated":
		processTaskUpdated(job.TaskID)
	case "task_completed":
		processTaskCompleted(job.TaskID)
	}
}

func processTaskCreated(taskID string) {
	fmt.Printf("Task Created:%s\n", taskID)
	fmt.Printf("sending notification for task: %s\n", taskID)
	err := rdb.Del(ctx, "allTasks").Err()
	if err != nil {
		fmt.Printf("Cache invalidation failed %v\n", err)

	}
}

func processTaskUpdated(taskID string) {
	fmt.Printf("Task Updated: %s\n", taskID)
	fmt.Printf("sending updated notification for task: %s\n", taskID)
	err := rdb.Del(ctx, "allTasks").Err()
	if err != nil {
		fmt.Printf("cache invalidation failed: %v\n", err)
	}
}

func processTaskCompleted(taskID string) {
	fmt.Printf("Task completed: %s\n", taskID)
	fmt.Printf("Sending completion notification for task %s\n", taskID)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Printf("Archiving completed task: %s\n", taskID)
	}()
}

func MonitorQueues() {
	for {
		length, err := rdb.LLen(ctx, "todo_jobs").Result() //return the length of the list stored in queues
		if err != nil {
			log.Printf("failed to get queue length: %v\n", err)
		} else {
			log.Printf("jobs in queue: %d\n", length)
		}
		time.Sleep(10 * time.Second)
	}
}
