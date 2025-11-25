package main

import (
	"log"
	"os"
	"time"
	"todo-list/controller"
	"todo-list/entities"
	"todo-list/repository"
	"todo-list/routes"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// queues
// crons

// repeating check if today is the day of the task completion
// send a email to the user

func initDB(dsn string) *gorm.DB{
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect database")
	}
	return db
}
func main() {
	maindsn := os.Getenv("DATABASE_URL")
	if maindsn==""{
		maindsn="host=localhost dbname=list1 user=postgres password=Parthu732 sslmode=disable port=5432"
	}
	backupdsn:=os.Getenv("BACKUP_URL")
	if backupdsn==""{
		backupdsn= "host=localhost dbname=list_db1 user=postgres password=Parthu732 sslmode=disable port=5432"
	}
	mainDB:=initDB(maindsn)
	backupDB:=initDB(backupdsn)
	if err:=mainDB.AutoMigrate(&entities.Task{}); err!=nil{
		log.Fatalln("auto-migrate for mainDB failed", err)
	}
	if err:=backupDB.AutoMigrate(&entities.TaskBackup{}); err!=nil{
		log.Fatalln("auto-migrate for backupdb filed", err)
	}
	mainRepo := repository.NewTaskRepository(mainDB)
	backupRepo:=repository.NewTaskBackupRepository(backupDB)
	taskService := service.NewTaskService(mainRepo)
	taskController := controller.NewTaskController(taskService)
	cronService:=service.NewTaskCron(mainRepo, backupRepo)
	go cronService.Start(3*time.Hour)

	service.Init("localhost:6379", "", 0)
	go service.Startworker()
	go service.MonitorQueues()

	app := fiber.New()
	routes.TaskRoutes(app, taskController)
	if err:=app.Listen(":5000"); err!=nil{
		log.Fatalf("fiber failed: %v",err)
	}
}


