package main

import (
	"log"
	"todo-list/controller"
	"todo-list/entities"
	"todo-list/repository"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost dbname=list user=postgres password=Parthu732 sslmode=disable port=5432"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&entities.Task{})
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)
	app := fiber.New()
	app.Post("/tasks", taskController.CreateTaskHandler)
	app.Get("/tasks", taskController.GetAllTasks)
	app.Get("/tasks/:id", taskController.GetTask)
	app.Put("/tasks/:id", taskController.UpdateTask)
	app.Delete("/tasks/:id", taskController.DeleteTask)

	app.Listen(":5000")
}
