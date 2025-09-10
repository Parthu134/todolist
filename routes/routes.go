package routes

import (
	"todo-list/controller"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, taskController *controller.TaskController){
	app.Post("/tasks", taskController.CreateTaskHandler)
	app.Get("/tasks", taskController.GetAllTasksHandler)
	app.Get("/tasks/:id", taskController.GetTaskHandler)
	app.Put("/tasks/:id", taskController.UpdateTaskHandler)
	app.Delete("/tasks/:id", taskController.DeleteTaskHandler)
}

	