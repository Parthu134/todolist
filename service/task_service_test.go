package service_test

import (
	"testing"
	"todo-list/entities"
	"todo-list/service"
)

type FakeRepo struct {
	tasks []entities.Task
}

func (f *FakeRepo) CreateRepo(task entities.Task) (entities.Task, error) {
	task.ID = uint(len(f.tasks) + 1)
	f.tasks = append(f.tasks, task)
	return task, nil
}
func (f *FakeRepo) GetAllRepo() ([]entities.Task, error) {
	return f.tasks, nil
}
func (f *FakeRepo) GetByIdRepo(id uint) (entities.Task, error) {
	for _, t := range f.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return entities.Task{}, nil
}
func (f *FakeRepo) UpdateRepo(task entities.Task) (entities.Task, error) {
	for i, t := range f.tasks {
		if t.ID == task.ID {
			f.tasks[i] = task
			return task, nil
		}
	}
	return task, nil
}
func (f *FakeRepo) DeleteRepo(id uint) error {
	for i, t := range f.tasks {
		if t.ID == id {
			f.tasks = append(f.tasks[:i], f.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestCreateTaskSevice(t *testing.T) {
	repo := &FakeRepo{}
	svc := service.NewTaskService(repo)
	task := entities.Task{Title: "Learn Go"}
	created, err := svc.CreateTaskService(task)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}
	if created.Title != "Learn Go" {
		t.Errorf("expected title 'Learn Go', got %s", created.Title)
	}
}

func TestGetAllTaskService(t *testing.T) {
	repo := &FakeRepo{}
	service := service.NewTaskService(repo)
	service.CreateTaskService(entities.Task{Title: "Task A"})
	service.CreateTaskService(entities.Task{Title: "Task B"})
	tasks, _ := service.GetAllTasksService()
	if len(tasks) != 2 {
		t.Errorf("expectd 2 tasks got %d", len(tasks))
	}
}

func TestGetByIdService(t *testing.T) {
	repo := &FakeRepo{}
	service := service.NewTaskService(repo)
	task, _ := service.CreateTaskService(entities.Task{Title: "Learn Go"})
	g, _ := service.GetTaskService(task.ID)
	if g.Title != "Learn Go" {
		t.Errorf("Expected Learn Go, got %s", g.Title)
	}
}

func TestUpdateService(t *testing.T) {
	repo := &FakeRepo{}
	service := service.NewTaskService(repo)
	task, _ := service.CreateTaskService(entities.Task{Title: "Old Title"})
	task.Title = "New Title"
	updated, _ := service.UpdateTaskService(task)
	if updated.Title != "New Title" {
		t.Errorf("expected New Title, got %s", updated.Title)
	}
}

func TestDeleteService(t *testing.T) {
	repo := &FakeRepo{}
	service := service.NewTaskService(repo)
	task, _ := service.CreateTaskService(entities.Task{Title: "Learn"})
	_ = service.DeleteTaskService(task.ID)
	tasks, _ := service.GetAllTasksService()
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}
}
