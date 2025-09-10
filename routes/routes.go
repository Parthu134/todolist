package routes

import (
	"todo-list/controller"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, taskController *controller.TaskController){
	app.Post("/tasks", taskController.CreateTaskHandler)
	app.Get("/tasks", taskController.GetAllTasks)
	app.Get("/tasks/:id", taskController.GetTask)
	app.Put("/tasks/:id", taskController.UpdateTask)
	app.Delete("/tasks/:id", taskController.DeleteTask)
}

	