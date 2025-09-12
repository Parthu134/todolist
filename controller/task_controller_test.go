package controller_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-list/controller"
	"todo-list/entities"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
)

type FakeRepo struct {
	tasks []entities.Task
}
func (f *FakeRepo) CreateRepo(task entities.Task) (entities.Task, error){
	task.ID=uint(len(f.tasks)+1)
	f.tasks=append(f.tasks, task)
	return task,nil
}
func (f *FakeRepo) GetAllRepo() ([]entities.Task, error){
	return f.tasks,nil
}
func (f *FakeRepo) GetByIdRepo(id uint) (entities.Task, error){
	for _,t:=range f.tasks{
		if t.ID==id{
			return t,nil
		}
	}
	return entities.Task{},nil
}
func (f *FakeRepo) UpdateRepo(task entities.Task) (entities.Task, error){
	for i,t:=range f.tasks{
		if t.ID==task.ID{
			f.tasks[i]=task
			return task,nil
		}
	}
	return task,nil
}
func (f *FakeRepo) DeleteRepo(id uint) error{
	for i,t:=range f.tasks{
		if t.ID==id{
			f.tasks=append(f.tasks[:i], f.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}

func seedFakeRepo() *FakeRepo{
	return &FakeRepo{
		tasks:[]entities.Task{
			{ID: 1,Title: "Seed", Description: "Seed Task"},
		},
	}
}
func TestCreateTaskHandler(t *testing.T){
	app:=fiber.New()
	repo:=&FakeRepo{}
	srv:=service.NewTaskService(repo)
	control:=controller.NewTaskController(srv)
	app.Post("/tasks", control.CreateTaskHandler)
	body:=`{"title":"New Task", "description":"Testing"}`
	req:=httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp,_:=app.Test(req)
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("expected status 200 got status %d", resp.StatusCode)
	}
	if len(repo.tasks)!=1{
		t.Errorf("expected 1 task in repo but found %d tasks", len(repo.tasks))
	}
}

func TestGetAllHandler(t *testing.T){
	app:=fiber.New()
	repo:=seedFakeRepo()
	srv:=service.NewTaskService(repo)
	control:=controller.NewTaskController(srv)
	app.Get("/all", control.GetAllTasksHandler)
	req:=httptest.NewRequest(http.MethodGet, "/all", nil)
	resp,_:=app.Test(req)
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("expected status 200 got status %d",resp.StatusCode)
	}
	bodybytes,_:=io.ReadAll(resp.Body)
	if !strings.Contains(string(bodybytes), "Seed Task"){
		t.Errorf("response does not contain expected task, but it contains %s",string(bodybytes))
	}
}

func TestGetByIdHandler(t *testing.T){
	app:=fiber.New()
	repo:=seedFakeRepo()
	srv:=service.NewTaskService(repo)
	control:=controller.NewTaskController(srv)
	app.Get("/task/:id", control.GetTaskHandler)
	req:=httptest.NewRequest(http.MethodGet, "/task/1", nil)
	resp,_:=app.Test(req)
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("expected status 200, got status %d",resp.StatusCode)
	}
}

func TestUpdateHandler(t *testing.T){
	app:=fiber.New()
	repo:=seedFakeRepo()
	srv:=service.NewTaskService(repo)
	control:=controller.NewTaskController(srv)
	app.Put("/task/:id",control.UpdateTaskHandler)
	body:=`{"title":"Harry Potter", "description":"A Comic book"}`
	req:=httptest.NewRequest(http.MethodPut, "/task/1", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp,_:=app.Test(req)
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("expected status code 200, got status %d", resp.StatusCode)
	}
}

func TestDeleteHandler(t *testing.T){
	app:=fiber.New()
	repo:=seedFakeRepo()
	srv:=service.NewTaskService(repo)
	control:=controller.NewTaskController(srv)
	app.Delete("/task/:id", control.DeleteTaskHandler)
	req:=httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	resp,_:=app.Test(req)
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("expected status 200 got status %d",resp.StatusCode)
	}
}

