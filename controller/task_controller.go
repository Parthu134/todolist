package controller

import (
	"strconv"
	"todo-list/entities"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	service service.TaskService
}

func NewTaskController(service service.TaskService) *TaskController {
	return &TaskController{service: service}
}
func (c *TaskController) CreateTaskHandler(ctx *fiber.Ctx) error {
	var task entities.Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	createdTask, err := c.service.CreateTaskService(task)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(createdTask)
}
func (c *TaskController) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := c.service.GetAllTasks()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(tasks)
}
func (c *TaskController) GetTask(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	task,err:=c.service.GetTask(uint(id))
	if err!=nil{
		return ctx.Status(404).JSON(fiber.Map{
			"error":"task not found",
		})
	}
	return ctx.JSON(task)
}
func (c *TaskController) UpdateTask(ctx *fiber.Ctx) error{
	id, _ := strconv.Atoi(ctx.Params("id"))
	var task entities.Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	task.ID=uint(id)
	updatedTask, err:=c.service.UpdateTask(task)
	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	return ctx.JSON(updatedTask)
}
func (c *TaskController) DeleteTask(ctx *fiber.Ctx) error{
	id,_:=strconv.Atoi(ctx.Params("id"))
	err:=c.service.DeleteTask(uint(id))
	if err!=nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"message":"data deleted successfully",
	})
}

